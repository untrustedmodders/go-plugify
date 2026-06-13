package generator

import (
	"fmt"
	"go/ast"
	"go/token"
	"go/types"
	"os"
	"slices"
	"strings"

	"golang.org/x/tools/go/packages"
)

type Package struct {
	currentFile string
	verbose     bool

	mainPkg *Package

	exportsRenameIndex int
	Exports            []ExportedFunction

	importsRenameIndex int
	imports            []packageImport

	Fset *token.FileSet
	Pkg  *packages.Package
}

func (g *Package) GenerateMain() ([]ExportedFunction, error) {
	_, err := g.Generate(nil)
	if err != nil {
		return nil, err
	}

	pkgDir := g.Pkg.Dir

	err = g.generateAutoExports(pkgDir)
	if err != nil {
		return nil, fmt.Errorf("error generating autoexports: %w", err)
	}

	err = GenerateAutoExportsHeader(pkgDir)
	if err != nil {
		return nil, fmt.Errorf("error generating autoexports header: %w", err)
	}

	fmt.Println("Generated autoexports.go and autoexports.h")

	return g.Exports, nil
}

func (g *Package) Generate(mainGen *Package) ([]ExportedFunction, error) {
	g.mainPkg = mainGen

	err := g.extractExportedFunctions()
	if err != nil {
		return nil, err
	}

	if len(g.Exports) == 0 {
		if g.verbose {
			fmt.Fprintf(os.Stderr, "Warning: No functions with //plugify:export comment found for %s\n", g.Pkg.Dir)
		}

		return nil, nil
	}

	if mainGen != nil {
		for _, exp := range g.Exports {
			idx := slices.IndexFunc(mainGen.Exports, func(e ExportedFunction) bool {
				return e.ExportName == exp.ExportName
			})
			if idx != -1 {

				exp.FuncName = fmt.Sprintf("%s_%d", exp.FuncName, g.exportsRenameIndex)

				g.exportsRenameIndex++

				if g.verbose {
					existExport := mainGen.Exports[idx]

					fmt.Fprintf(os.Stderr, "Warning: Name collision:\n\t%s.%s and %s.%s (solved by indexing)\n",
						existExport.packageImport.path, existExport.ExportName,
						g.Pkg.PkgPath, exp.ExportName,
					)
				}
			}

			mainGen.Exports = append(mainGen.Exports, exp)
		}
	}

	return g.Exports, nil
}

type packageImport struct {
	name string
	path string
}

// Returns the import name for the given path, adding it if necessary
func (g *Package) AddImport(name string, path string) string {
	if path == "" {
		return ""
	}

	var renameNeeded bool
	for _, imp := range g.imports {
		if imp.path == path {
			return imp.name
		}

		if imp.name == name {
			renameNeeded = true
		}
	}

	if renameNeeded {
		g.importsRenameIndex++
		name = fmt.Sprintf("%s_%d", name, g.importsRenameIndex)
	}

	g.imports = append(g.imports, packageImport{name: name, path: path})
	return name
}

func (g *Package) GetFileLinePosNumber(pos token.Pos) string {
	fpos := g.Fset.Position(pos)
	return fmt.Sprintf("%d:%d", fpos.Line, fpos.Column)
}

func (g *Package) extractExportedFunctions() (retErr error) {
	for _, file := range g.Pkg.Syntax {
		ast.Inspect(file, func(n ast.Node) bool {
			funcDecl, ok := n.(*ast.FuncDecl)
			if !ok {
				return true
			}

			// Skip methods (functions with receivers)
			if funcDecl.Recv != nil {
				return true
			}

			// Look for //plugify:export comment
			exportName := ""
			if funcDecl.Doc != nil {
				for _, comment := range funcDecl.Doc.List {
					text := strings.TrimSpace(strings.TrimPrefix(comment.Text, "//"))
					if strings.HasPrefix(text, "plugify:export") {
						parts := strings.SplitN(text, " ", 2)
						if len(parts) == 2 {
							exportName = strings.TrimSpace(parts[1])
						} else {
							exportName = funcDecl.Name.Name
						}
						break
					}
				}
			}

			if exportName == "" {
				return true
			}

			// Skip unexported functions
			if !funcDecl.Name.IsExported() {
				retErr = fmt.Errorf("Function is not exported but marked to export %s:%s: %s\n", file.Name.Name, g.GetFileLinePosNumber(funcDecl.Name.Pos()), funcDecl.Name.Name)
				return false
			}

			// Parse documentation comments
			docComment := parseDocComment(funcDecl.Doc)

			// Get function signature from type info
			obj := g.Pkg.TypesInfo.ObjectOf(funcDecl.Name)
			if obj == nil {
				return true
			}

			sig, ok := obj.Type().(*types.Signature)
			if !ok {
				return true
			}

			params, err := g.extractParams(sig, g.Pkg.TypesInfo, g.Pkg)
			if err != nil {
				retErr = fmt.Errorf("%s:%s: %s %w", file.Name.Name, g.GetFileLinePosNumber(funcDecl.Name.Pos()), funcDecl.Name.Name, err)
				return false
			}

			// Add parameter descriptions from doc comments
			for i := range params {
				if desc, ok := docComment.ParamDescs[params[i].Name]; ok {
					params[i].Description = desc
				}
			}

			// Extract return type
			retType, _, err := g.extractReturnType(sig.Results(), g.Pkg.TypesInfo, g.Pkg)
			if err != nil {
				retErr = err
				return false
			}

			// Add return description from doc comments
			retType.Description = docComment.ReturnDesc

			var pkgImport packageImport
			if g.mainPkg != nil {
				importName := g.mainPkg.AddImport(g.Pkg.Name, g.Pkg.PkgPath)
				pkgImport = packageImport{name: importName, path: g.Pkg.PkgPath}
			}

			g.Exports = append(g.Exports, ExportedFunction{
				ExportName:  exportName,
				FuncName:    "__" + funcDecl.Name.Name,
				Description: docComment.Description,
				Params:      params,
				ReturnType:  retType,

				originalFuncName: funcDecl.Name.Name,
				packageImport:    pkgImport,
			})

			return true
		})
	}

	return
}

func (g *Package) prepareImport(param *types.TypeName) packageImport {
	pkg := param.Pkg()
	if pkg == nil {
		return packageImport{}
	}

	pkgName := pkg.Name()

	if pkgName == "main" {
		return packageImport{}
	}

	pkgPath := pkg.Path()

	genPkg := g
	if g.mainPkg != nil {
		genPkg = g.mainPkg
	}

	importName := genPkg.AddImport(pkgName, pkgPath)

	return packageImport{name: importName, path: pkgPath}
}
