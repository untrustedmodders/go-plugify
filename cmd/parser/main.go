package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/untrustedmodders/go-plugify"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/packages"
	"os"
	"strings"
)

type ExportedFunction struct {
	ExportName string
	FuncName   string
	Params     []ParamInfo
	ReturnType TypeInfo
}

type ParamInfo struct {
	Name string
	Type TypeInfo
}

type TypeInfo struct {
	TypeString   string
	IsFunc       bool
	FuncSig      *FuncSignature
	IsEnum       bool
	EnumTypeName string
	EnumValues   []plugify.EnumValue
	IsArray      bool
	ElemType     *TypeInfo
}

type FuncSignature struct {
	Name   string
	Params []ParamInfo
	Return TypeInfo
}

func main() {
	var (
		pkgPath     = flag.String("package", ".", "Package path to analyze")
		outputFile  = flag.String("output", "", "Output manifest file (default: <packagename>.pplugin)")
		name        = flag.String("name", "", "Plugin name (default: package name)")
		version     = flag.String("version", "1.0.0", "Plugin version")
		description = flag.String("description", "", "Plugin description")
		author      = flag.String("author", "", "Plugin author")
		website     = flag.String("website", "", "Plugin website")
		license     = flag.String("license", "", "Plugin license")
		entry       = flag.String("entry", "", "Plugin entry point (default: <packagename>)")
	)

	flag.Parse()

	// Load package with type information
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedSyntax |
			packages.NeedTypes | packages.NeedTypesInfo | packages.NeedImports,
	}

	pkgs, err := packages.Load(cfg, *pkgPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading package: %v\n", err)
		os.Exit(1)
	}

	if len(pkgs) == 0 {
		fmt.Fprintf(os.Stderr, "No packages found\n")
		os.Exit(1)
	}

	pkg := pkgs[0]
	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1)
	}

	// Extract exported functions
	exportedFuncs := extractExportedFunctions(pkg)

	if len(exportedFuncs) == 0 {
		fmt.Fprintf(os.Stderr, "Warning: No functions with //plugify:export comment found\n")
	}

	// Build manifest
	pluginName := *name
	if pluginName == "" {
		pluginName = pkg.Name
	}

	pluginEntry := *entry
	if pluginEntry == "" {
		pluginEntry = pluginName
	}

	outputFileName := *outputFile
	if outputFileName == "" {
		outputFileName = pluginName + ".pplugin"
	}

	manifest := plugify.Manifest{
		Schema:      "https://raw.githubusercontent.com/untrustedmodders/plugify/refs/heads/main/schemas/plugin.schema.json",
		Name:        pluginName,
		Version:     *version,
		Description: *description,
		Author:      *author,
		Website:     *website,
		License:     *license,
		Entry:       pluginEntry,
		Language:    "golang",
		Methods:     convertToManifestMethods(exportedFuncs),
	}

	// Write JSON
	data, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error marshaling JSON: %v\n", err)
		os.Exit(1)
	}

	err = os.WriteFile(outputFileName, data, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing manifest file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated manifest: %s (%d methods)\n", outputFileName, len(exportedFuncs))
}

func extractExportedFunctions(pkg *packages.Package) []ExportedFunction {
	var exports []ExportedFunction

	for _, file := range pkg.Syntax {
		ast.Inspect(file, func(n ast.Node) bool {
			funcDecl, ok := n.(*ast.FuncDecl)
			if !ok {
				return true
			}

			// Skip methods (functions with receivers)
			if funcDecl.Recv != nil {
				return true
			}

			// Skip unexported functions
			if !funcDecl.Name.IsExported() {
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

			// Get function signature from type info
			obj := pkg.TypesInfo.ObjectOf(funcDecl.Name)
			if obj == nil {
				return true
			}

			sig, ok := obj.Type().(*types.Signature)
			if !ok {
				return true
			}

			// Extract parameters
			params := extractParams(sig.Params(), pkg.TypesInfo)

			// Extract return type
			retType := extractReturnType(sig.Results(), pkg.TypesInfo)

			exports = append(exports, ExportedFunction{
				ExportName: exportName,
				FuncName:   funcDecl.Name.Name,
				Params:     params,
				ReturnType: retType,
			})

			return true
		})
	}

	return exports
}

func extractParams(params *types.Tuple, info *types.Info) []ParamInfo {
	if params == nil {
		return nil
	}

	result := make([]ParamInfo, params.Len())
	for i := 0; i < params.Len(); i++ {
		param := params.At(i)
		result[i] = ParamInfo{
			Name: param.Name(),
			Type: mapTypeInfo(param.Type(), info),
		}
	}
	return result
}

func extractReturnType(results *types.Tuple, info *types.Info) TypeInfo {
	if results == nil || results.Len() == 0 {
		return TypeInfo{TypeString: "void"}
	}

	// For now, handle single return value
	// You could extend this for multiple returns
	return mapTypeInfo(results.At(0).Type(), info)
}

func mapTypeInfo(t types.Type, info *types.Info) TypeInfo {
	// Handle pointers (treat as references)
	if ptr, ok := t.(*types.Pointer); ok {
		// For now, just get the underlying type
		t = ptr.Elem()
	}

	// Handle slices/arrays
	if slice, ok := t.(*types.Slice); ok {
		elemType := mapTypeInfo(slice.Elem(), info)
		return TypeInfo{
			TypeString: elemType.TypeString + "[]",
			IsArray:    true,
			ElemType:   &elemType,
		}
	}

	// Handle function types
	if sig, ok := t.(*types.Signature); ok {
		params := extractParams(sig.Params(), info)
		retType := extractReturnType(sig.Results(), info)

		return TypeInfo{
			TypeString: "function",
			IsFunc:     true,
			FuncSig: &FuncSignature{
				Name:   "callback",
				Params: params,
				Return: retType,
			},
		}
	}

	// Handle named types (check for enums - const groups)
	if named, ok := t.(*types.Named); ok {
		typeName := named.Obj().Name()

		// Check if it's an enum-like type (underlying basic type with const values)
		if basic, ok := named.Underlying().(*types.Basic); ok {
			// Try to find enum values (this is heuristic)
			enumValues := findEnumValues(named, info)
			if len(enumValues) > 0 {
				underlyingType := mapBasicType(basic)
				return TypeInfo{
					TypeString:   underlyingType,
					IsEnum:       true,
					EnumTypeName: typeName,
					EnumValues:   enumValues,
				}
			}
		}
	}

	// Handle basic types
	if basic, ok := t.(*types.Basic); ok {
		return TypeInfo{TypeString: mapBasicType(basic)}
	}

	// Default
	return TypeInfo{TypeString: t.String()}
}

func mapBasicType(basic *types.Basic) string {
	switch basic.Kind() {
	case types.Bool:
		return "bool"
	case types.Int8:
		return "int8"
	case types.Int16:
		return "int16"
	case types.Int32:
		return "int32"
	case types.Int64, types.Int:
		return "int64"
	case types.Uint8:
		return "uint8"
	case types.Uint16:
		return "uint16"
	case types.Uint32:
		return "uint32"
	case types.Uint64, types.Uint:
		return "uint64"
	case types.Float32:
		return "float"
	case types.Float64:
		return "double"
	case types.String:
		return "string"
	case types.UntypedNil:
		return "ptr64" // or ptr32 depending on architecture
	default:
		return basic.String()
	}
}

func findEnumValues(named *types.Named, info *types.Info) []plugify.EnumValue {
	// This is a heuristic approach
	// In Go, enums are typically const groups with the same type
	// We'd need to scan the package scope for const declarations
	// For now, return empty - you could extend this
	return nil
}

func convertToManifestMethods(funcs []ExportedFunction) []plugify.Method {
	methods := make([]plugify.Method, len(funcs))
	for i, f := range funcs {
		methods[i] = plugify.Method{
			Name:       f.ExportName,
			FuncName:   f.FuncName,
			ParamTypes: convertParams(f.Params),
			RetType:    convertReturnType(f.ReturnType),
		}
	}
	return methods
}

func convertParams(params []ParamInfo) []plugify.Property {
	result := make([]plugify.Property, len(params))
	for i, p := range params {
		if p.Type.IsFunc {
			// Function parameter with prototype
			result[i] = plugify.Property{
				Type: "function",
				Name: p.Name,
				Prototype: &plugify.Method{
					Name:       p.Type.FuncSig.Name,
					ParamTypes: convertParams(p.Type.FuncSig.Params),
					RetType:    convertReturnType(p.Type.FuncSig.Return),
				},
			}
		} else if p.Type.IsEnum {
			// Enum parameter
			result[i] = plugify.Property{
				Type: p.Type.TypeString,
				Name: p.Name,
				Enumerator: &plugify.EnumObject{
					Name:   p.Type.EnumTypeName,
					Values: p.Type.EnumValues,
				},
			}
		} else {
			// Regular parameter
			result[i] = plugify.Property{
				Type: p.Type.TypeString,
				Name: p.Name,
			}
		}
	}
	return result
}

func convertReturnType(t TypeInfo) plugify.Property {
	if t.IsEnum {
		return plugify.Property{
			Type: t.TypeString,
			Enumerator: &plugify.EnumObject{
				Name:   t.EnumTypeName,
				Values: t.EnumValues,
			},
		}
	}

	return plugify.Property{
		Type: t.TypeString,
	}
}
