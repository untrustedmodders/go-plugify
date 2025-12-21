package plugify

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/constant"
	"go/types"
	"os"
	"sort"
	"strings"
	"unsafe"

	"golang.org/x/tools/go/packages"
)

type ExportedFunction struct {
	ExportName  string
	FuncName    string
	Params      []ParamInfo
	ReturnType  TypeInfo
	Description string
}

type ParamInfo struct {
	Name        string
	Type        TypeInfo
	Description string
}

type TypeInfo struct {
	TypeString string
	IsRef      bool
	IsFunc     bool
	IsEnum     bool
	IsArray    bool

	EnumTypeName string
	EnumValues   []EnumValue
	ElemType     *TypeInfo
	FuncSig      *FuncSignature
	Description  string
}

type FuncSignature struct {
	Name        string
	Params      []ParamInfo
	Return      TypeInfo
	Description string
}

// DocComment represents parsed documentation from comments
type DocComment struct {
	Description  string
	ParamDescs   map[string]string // param name -> description
	ReturnDesc   string
	EnumValueMap map[string]string // enum value name -> description
}

// parseDocComment parses doxygen-style comments and extracts @param, @return, @brief, etc.
func parseDocComment(commentGroup *ast.CommentGroup) DocComment {
	doc := DocComment{
		ParamDescs:   make(map[string]string),
		EnumValueMap: make(map[string]string),
	}

	if commentGroup == nil {
		return doc
	}

	var descriptionLines []string
	var briefDesc string
	inDescription := true

	for _, comment := range commentGroup.List {
		text := strings.TrimSpace(strings.TrimPrefix(comment.Text, "//"))
		text = strings.TrimSpace(strings.TrimPrefix(text, "/*"))
		text = strings.TrimSpace(strings.TrimSuffix(text, "*/"))
		text = strings.TrimSpace(strings.TrimPrefix(text, "*"))

		// Skip plugify:export directives
		if strings.HasPrefix(text, "plugify:export") {
			continue
		}

		// Parse @brief tag
		if strings.HasPrefix(text, "@brief") {
			inDescription = false
			parts := strings.SplitN(text, "@brief", 2)
			if len(parts) == 2 {
				briefDesc = strings.TrimSpace(parts[1])
			}
			continue
		}

		// Parse @param tag
		if strings.HasPrefix(text, "@param") {
			inDescription = false
			parts := strings.Fields(text)
			if len(parts) >= 3 {
				paramName := parts[1]
				paramDesc := strings.Join(parts[2:], " ")
				doc.ParamDescs[paramName] = paramDesc
			}
			continue
		}

		// Parse @return tag
		if strings.HasPrefix(text, "@return") {
			inDescription = false
			parts := strings.SplitN(text, "@return", 2)
			if len(parts) == 2 {
				doc.ReturnDesc = strings.TrimSpace(parts[1])
			}
			continue
		}

		// Collect description lines
		if inDescription && text != "" {
			descriptionLines = append(descriptionLines, text)
		}
	}

	// Use @brief if provided, otherwise use the collected description lines
	if briefDesc != "" {
		doc.Description = briefDesc
	} else {
		doc.Description = strings.Join(descriptionLines, " ")
	}

	return doc
}

// invalidNames contains Go's reserved keywords and predeclared identifiers
var invalidNames = map[string]struct{}{
	"break": {}, "case": {}, "chan": {}, "const": {}, "continue": {},
	"default": {}, "defer": {}, "else": {}, "fallthrough": {}, "for": {},
	"func": {}, "go": {}, "goto": {}, "if": {}, "import": {}, "interface": {},
	"map": {}, "package": {}, "range": {}, "return": {}, "select": {}, "struct": {},
	"switch": {}, "type": {}, "var": {}, "append": {}, "bool": {}, "byte": {},
	"cap": {}, "close": {}, "complex": {}, "complex64": {}, "complex128": {},
	"copy": {}, "delete": {}, "error": {}, "false": {}, "float32": {}, "float64": {},
	"imag": {}, "int": {}, "int8": {}, "int16": {}, "int32": {}, "int64": {},
	"iota": {}, "len": {}, "make": {}, "new": {}, "nil": {}, "panic": {},
	"print": {}, "println": {}, "real": {}, "recover": {}, "rune": {}, "string": {},
	"true": {}, "uint": {}, "uint8": {}, "uint16": {}, "uint32": {}, "uint64": {},
	"uintptr": {},
}

func generateName(name string) string {
	if _, exists := invalidNames[name]; exists {
		return name + "_"
	}
	return name
}

func Generate(
	patterns string,
	output string,
	name string,
	version string,
	description string,
	author string,
	website string,
	license string,
	platforms []string,
	dependencies []string,
	conflicts []string,
	entry string,
	target string,
) error {
	// Load package with type information
	pkg, err := loadPackage(patterns, target)
	if err != nil {
		return err
	}

	// Extract exported functions]
	exportedFuncs := extractExportedFunctions(pkg)

	if len(exportedFuncs) == 0 {
		fmt.Fprintf(os.Stderr, "Warning: No functions with //plugify:export comment found\n")
	}

	// Build manifest
	pluginName := name
	if pluginName == "" {
		pluginName = pkg.Name
	}

	pluginEntry := entry
	if pluginEntry == "" {
		pluginEntry = pluginName
	}

	var pluginDependencies []Dependency
	for _, dependency := range dependencies {
		pluginDependencies = append(pluginDependencies, Dependency{Name: dependency})
	}

	var pluginConflicts []Conflict
	for _, conflict := range conflicts {
		pluginConflicts = append(pluginConflicts, Conflict{Name: conflict})
	}

	outputFile := output
	if outputFile == "" {
		outputFile = pluginName + ".pplugin"
	}

	manifest := Manifest{
		Schema:       "https://raw.githubusercontent.com/untrustedmodders/plugify/refs/heads/main/schemas/plugin.schema.json",
		Name:         pluginName,
		Version:      version,
		Description:  description,
		Author:       author,
		Website:      website,
		License:      license,
		Platforms:    platforms,
		Entry:        pluginEntry,
		Dependencies: pluginDependencies,
		Conflicts:    pluginConflicts,
		Language:     "golang",
		Methods:      convertToManifestMethods(exportedFuncs),
	}

	// Write JSON
	data, err := json.MarshalIndent(manifest, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %w", err)
	}

	err = os.WriteFile(outputFile, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing manifest file %s: %w", outputFile, err)
	}

	fmt.Printf("Generated manifest: %s (%d methods)\n", outputFile, len(exportedFuncs))

	if err := generateAutoExports(exportedFuncs, target); err != nil {
		return fmt.Errorf("error generating autoexports: %w", err)
	}

	if err := generateAutoExportsHeader(); err != nil {
		return fmt.Errorf("error generating autoexports header: %w", err)
	}

	fmt.Println("Generated autoexports.go and autoexports.h")

	return nil
}

func loadPackage(patterns string, target string) (*packages.Package, error) {
	cfg := &packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedSyntax |
			packages.NeedTypes | packages.NeedTypesInfo,
		Tests:      false,
		Env:        append(os.Environ(), "CGO_ENABLED=1"),
		BuildFlags: []string{"-tags", "cgo"},
	}

	pkgs, err := packages.Load(cfg, patterns)

	if err != nil {
		return nil, fmt.Errorf("error loading package: %w", err)
	}

	if len(pkgs) == 0 {
		return nil, fmt.Errorf("no packages found in patterns: %s", patterns)
	}

	var targetPkg *packages.Package
	for _, pkg := range pkgs {
		if pkg.Name == target {
			targetPkg = pkg
		}
	}

	if targetPkg == nil {
		return nil, fmt.Errorf("no target package found in patterns: %s", patterns)
	}

	return targetPkg, nil
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

			// Parse documentation comments
			docComment := parseDocComment(funcDecl.Doc)

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
			params := extractParams(sig.Params(), pkg.TypesInfo, pkg)

			// Add parameter descriptions from doc comments
			for i := range params {
				if desc, ok := docComment.ParamDescs[params[i].Name]; ok {
					params[i].Description = desc
				}
			}

			// Extract return type
			retType := extractReturnType(sig.Results(), pkg.TypesInfo, pkg)

			// Add return description from doc comments
			retType.Description = docComment.ReturnDesc

			exports = append(exports, ExportedFunction{
				ExportName:  exportName,
				FuncName:    "__" + funcDecl.Name.Name,
				Description: docComment.Description,
				Params:      params,
				ReturnType:  retType,
			})

			return true
		})
	}

	return exports
}

func extractParams(params *types.Tuple, info *types.Info, pkg *packages.Package) []ParamInfo {
	if params == nil {
		return nil
	}

	result := make([]ParamInfo, params.Len())
	for i := 0; i < params.Len(); i++ {
		param := params.At(i)
		result[i] = ParamInfo{
			Name: param.Name(),
			Type: mapTypeInfo(param.Type(), info, pkg),
		}
	}
	return result
}

func extractReturnType(results *types.Tuple, info *types.Info, pkg *packages.Package) TypeInfo {
	if results == nil || results.Len() == 0 {
		return TypeInfo{TypeString: "void"}
	}

	// For now, handle single return value
	// You could extend this for multiple returns
	return mapTypeInfo(results.At(0).Type(), info, pkg)
}

func mapTypeInfo(t types.Type, info *types.Info, pkg *packages.Package) TypeInfo {
	return mapTypeInfoWithRef(t, info, pkg, false)
}

func mapTypeInfoWithRef(t types.Type, info *types.Info, pkg *packages.Package, isRef bool) TypeInfo {
	// Handle pointer types - set isRef and unwrap
	if ptr, ok := t.(*types.Pointer); ok {
		return mapTypeInfoWithRef(ptr.Elem(), info, pkg, true)
	}

	// Handle type aliases (Go 1.22+)
	if alias, ok := t.(*types.Alias); ok {
		aliasName := alias.Obj().Name()
		rhs := alias.Rhs()

		// Special case: 'any' is a built-in alias for interface{}, unwrap it
		if aliasName == "any" {
			return mapTypeInfoWithRef(rhs, info, pkg, isRef)
		}

		// Check if alias is to a function type
		if sig, ok := rhs.(*types.Signature); ok {
			params := extractParams(sig.Params(), info, pkg)
			retType := extractReturnType(sig.Results(), info, pkg)

			// Get type-level documentation for delegate with doxygen parsing
			var delegateDoc DocComment
			if pkg != nil {
				delegateDoc = findTypeDelegateDoc(pkg, aliasName)
			}

			// Apply parameter descriptions from delegate documentation
			for i := range params {
				if desc, ok := delegateDoc.ParamDescs[params[i].Name]; ok {
					params[i].Description = desc
				}
			}

			// Apply return description from delegate documentation
			if delegateDoc.ReturnDesc != "" {
				retType.Description = delegateDoc.ReturnDesc
			}

			return TypeInfo{
				TypeString: "function",
				IsRef:      isRef,
				IsFunc:     true,
				FuncSig: &FuncSignature{
					Name:        aliasName, // Use the alias name
					Params:      params,
					Return:      retType,
					Description: delegateDoc.Description,
				},
			}
		}

		// For other aliases (potential enums), check if underlying type is basic
		if basic, ok := rhs.(*types.Basic); ok {
			underlyingType := mapBasicType(basic)
			// This is a type alias to a basic type (likely an enum)
			// Try to extract enum values with descriptions
			enumValues := findEnumValues(pkg, alias.Obj())

			// Get type-level documentation for enum
			var typeDesc string
			if pkg != nil {
				typeDesc = findTypeComment(pkg, aliasName)
			}

			return TypeInfo{
				TypeString:   underlyingType,
				IsRef:        isRef,
				IsEnum:       true,
				EnumTypeName: aliasName,
				EnumValues:   enumValues,
				Description:  typeDesc,
			}
		}

		// For other aliases, unwrap them
		return mapTypeInfoWithRef(rhs, info, pkg, isRef)
	}

	// Handle interface{} (any) - check this before slices since []any needs to detect any
	if iface, ok := t.(*types.Interface); ok {
		// Check if it's the empty interface (any)
		if iface.NumMethods() == 0 {
			return TypeInfo{
				TypeString: "any",
				IsRef:      isRef,
			}
		}
		// Non-empty interfaces are treated as pointers
		return TypeInfo{
			TypeString: getPtrType(),
			IsRef:      isRef,
		}
	}

	// Handle slices/arrays
	if slice, ok := t.(*types.Slice); ok {
		elemType := mapTypeInfoWithRef(slice.Elem(), info, pkg, false)
		return TypeInfo{
			TypeString: elemType.TypeString + "[]",
			IsRef:      isRef,
			IsArray:    true,
			ElemType:   &elemType,
		}
	}

	// Handle named types (check for plugify structs and enums)
	if named, ok := t.(*types.Named); ok {
		obj := named.Obj()
		typeName := obj.Name()
		pkgPath := ""
		if obj.Pkg() != nil {
			pkgPath = obj.Pkg().Path()
		}

		// Check for plugify struct types (Vector2, Vector3, Vector4, Matrix4x4)
		if pkgPath == "github.com/untrustedmodders/go-plugify" {
			switch typeName {
			case "Vector2":
				return TypeInfo{TypeString: "vec2", IsRef: isRef}
			case "Vector3":
				return TypeInfo{TypeString: "vec3", IsRef: isRef}
			case "Vector4":
				return TypeInfo{TypeString: "vec4", IsRef: isRef}
			case "Matrix4x4":
				return TypeInfo{TypeString: "mat4x4", IsRef: isRef}
			}
		}

		// Check underlying type for enums, functions, and type aliases
		underlying := named.Underlying()

		// Handle named function types (e.g., type Func1 func())
		if sig, ok := underlying.(*types.Signature); ok {
			params := extractParams(sig.Params(), info, pkg)
			retType := extractReturnType(sig.Results(), info, pkg)

			// Get type-level documentation for delegate with doxygen parsing
			var delegateDoc DocComment
			if pkg != nil {
				delegateDoc = findTypeDelegateDoc(pkg, typeName)
			}

			// Apply parameter descriptions from delegate documentation
			for i := range params {
				if desc, ok := delegateDoc.ParamDescs[params[i].Name]; ok {
					params[i].Description = desc
				}
			}

			// Apply return description from delegate documentation
			if delegateDoc.ReturnDesc != "" {
				retType.Description = delegateDoc.ReturnDesc
			}

			return TypeInfo{
				TypeString: "function",
				IsRef:      isRef,
				IsFunc:     true,
				FuncSig: &FuncSignature{
					Name:        typeName, // Use the name from the named type
					Params:      params,
					Return:      retType,
					Description: delegateDoc.Description,
				},
			}
		}

		// Handle type aliases to basic types (enums)
		if basic, ok := underlying.(*types.Basic); ok {
			underlyingType := mapBasicType(basic)

			// Check if this is a type alias in the same package (likely an enum)
			// Type aliases have the same underlying type but different name
			if typeName != "" && typeName != underlyingType {
				// Try to extract enum values with descriptions
				enumValues := findEnumValues(pkg, obj)

				// Get type-level documentation for enum
				var typeDesc string
				if pkg != nil {
					typeDesc = findTypeComment(pkg, typeName)
				}

				return TypeInfo{
					TypeString:   underlyingType,
					IsRef:        isRef,
					IsEnum:       true,
					EnumTypeName: typeName,
					EnumValues:   enumValues,
					Description:  typeDesc,
				}
			}

			return TypeInfo{
				TypeString: underlyingType,
				IsRef:      isRef,
			}
		}

		// Handle type aliases to structs
		if structType, ok := underlying.(*types.Struct); ok {
			_ = structType // Use the variable to avoid unused warning
			// For now, treat unknown structs as unsafe.Pointer
			return TypeInfo{
				TypeString: getPtrType(),
				IsRef:      isRef,
			}
		}
	}

	// Handle function types
	if sig, ok := t.(*types.Signature); ok {
		params := extractParams(sig.Params(), info, pkg)
		retType := extractReturnType(sig.Results(), info, pkg)

		funcName := "callback"
		// Try to get function type name from context if available
		return TypeInfo{
			TypeString: "function",
			IsRef:      isRef,
			IsFunc:     true,
			FuncSig: &FuncSignature{
				Name:   funcName,
				Params: params,
				Return: retType,
			},
		}
	}

	// Handle basic types
	if basic, ok := t.(*types.Basic); ok {
		return TypeInfo{
			TypeString: mapBasicType(basic),
			IsRef:      isRef,
		}
	}

	// Default - unknown type
	return TypeInfo{
		TypeString: getPtrType(),
		IsRef:      isRef,
	}
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
	case types.Int64:
		return "int64"
	case types.Uint8:
		return "uint8"
	case types.Uint16:
		return "uint16"
	case types.Uint32:
		return "uint32"
	case types.Uint64:
		return "uint64"
	case types.Float32:
		return "float"
	case types.Float64:
		return "double"
	case types.String:
		return "string"
	case types.UnsafePointer:
		return getPtrType()
	case types.Uintptr:
		return getPtrType()
	case types.UntypedNil:
		return getPtrType()
	case types.Uint:
		return getUIntType()
	case types.Int:
		return getIntType()
	default:
		return basic.String()
	}
}

// findConstComment finds the comment for a constant declaration in AST
func findConstComment(pkg *packages.Package, constName string) string {
	for _, file := range pkg.Syntax {
		for _, decl := range file.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}

			for _, spec := range genDecl.Specs {
				valueSpec, ok := spec.(*ast.ValueSpec)
				if !ok {
					continue
				}

				for _, name := range valueSpec.Names {
					if name.Name == constName {
						// Check for comment on the same line or doc comment
						if valueSpec.Doc != nil {
							docComment := parseDocComment(valueSpec.Doc)
							return docComment.Description
						}
						if valueSpec.Comment != nil {
							comment := valueSpec.Comment.Text()
							comment = strings.TrimSpace(strings.TrimPrefix(comment, "//"))
							comment = strings.TrimSpace(strings.TrimPrefix(comment, "/*"))
							comment = strings.TrimSpace(strings.TrimSuffix(comment, "*/"))
							return comment
						}
					}
				}
			}
		}
	}
	return ""
}

// findTypeComment finds the comment for a type declaration in AST
func findTypeComment(pkg *packages.Package, typeName string) string {
	for _, file := range pkg.Syntax {
		for _, decl := range file.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}

			for _, spec := range genDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				if typeSpec.Name.Name == typeName {
					// First check for doc comment above the type
					if genDecl.Doc != nil {
						docComment := parseDocComment(genDecl.Doc)
						if docComment.Description != "" {
							return docComment.Description
						}
					}

					// Then check for inline comment on the same line
					if typeSpec.Comment != nil {
						comment := typeSpec.Comment.Text()
						comment = strings.TrimSpace(strings.TrimPrefix(comment, "//"))
						comment = strings.TrimSpace(strings.TrimPrefix(comment, "/*"))
						comment = strings.TrimSpace(strings.TrimSuffix(comment, "*/"))
						return comment
					}
				}
			}
		}
	}
	return ""
}

// findTypeDelegateDoc finds and parses doxygen-style documentation for a delegate type
func findTypeDelegateDoc(pkg *packages.Package, typeName string) DocComment {
	for _, file := range pkg.Syntax {
		for _, decl := range file.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}

			for _, spec := range genDecl.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}

				if typeSpec.Name.Name == typeName {
					// Check for doc comment above the type
					if genDecl.Doc != nil {
						return parseDocComment(genDecl.Doc)
					}
				}
			}
		}
	}
	return DocComment{
		ParamDescs:   make(map[string]string),
		EnumValueMap: make(map[string]string),
	}
}

func findEnumValues(pkg *packages.Package, typeObj types.Object) []EnumValue {
	// Get the package where this type is defined
	typePkg := typeObj.Pkg()
	if typePkg == nil {
		return nil
	}

	// Get the type we're looking for
	enumType := typeObj.Type()
	enumTypeName := typeObj.Name()

	// If we have AST access, use it to find constants with explicit type annotations
	if pkg != nil {
		enumValues := findEnumValuesFromAST(pkg, enumTypeName)
		if len(enumValues) > 0 {
			return enumValues
		}
	}

	return findEnumValuesInScope(pkg, typePkg, enumType)
}

// findEnumValuesFromAST extracts enum values by examining the AST to find constants
// with explicit type annotations matching the enum type name
func findEnumValuesFromAST(pkg *packages.Package, enumTypeName string) []EnumValue {
	var enumValues []EnumValue

	for _, file := range pkg.Syntax {
		for _, decl := range file.Decls {
			genDecl, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}

			for _, spec := range genDecl.Specs {
				valueSpec, ok := spec.(*ast.ValueSpec)
				if !ok {
					continue
				}

				// Check if this constant has an explicit type annotation
				if valueSpec.Type == nil {
					continue
				}

				// Get the type name from the AST
				var typeName string
				switch t := valueSpec.Type.(type) {
				case *ast.Ident:
					typeName = t.Name
				default:
					continue
				}

				// Only include constants explicitly declared with this enum type
				if typeName != enumTypeName {
					continue
				}

				// Extract all constants in this spec
				for _, name := range valueSpec.Names {
					// Get the constant object from type info
					obj := pkg.TypesInfo.ObjectOf(name)
					if obj == nil {
						continue
					}

					constObj, ok := obj.(*types.Const)
					if !ok {
						continue
					}

					// Extract the constant value
					val := constObj.Val()
					if val == nil {
						continue
					}

					// Convert to int64
					var intValue int64
					switch val.Kind() {
					case constant.Int:
						if i, ok := constant.Int64Val(val); ok {
							intValue = i
						} else {
							continue
						}
					default:
						continue
					}

					// Get description
					var description string
					if valueSpec.Doc != nil {
						docComment := parseDocComment(valueSpec.Doc)
						description = docComment.Description
					}
					if description == "" && valueSpec.Comment != nil {
						comment := valueSpec.Comment.Text()
						comment = strings.TrimSpace(strings.TrimPrefix(comment, "//"))
						comment = strings.TrimSpace(strings.TrimPrefix(comment, "/*"))
						comment = strings.TrimSpace(strings.TrimSuffix(comment, "*/"))
						description = comment
					}

					enumValues = append(enumValues, EnumValue{
						Name:        name.Name,
						Value:       intValue,
						Description: description,
					})
				}
			}
		}
	}

	// Sort by value to ensure consistent ordering
	sort.Slice(enumValues, func(i, j int) bool {
		return enumValues[i].Value < enumValues[j].Value
	})

	return enumValues
}

func findEnumValuesInScope(pkg *packages.Package, typePkg *types.Package, enumType types.Type) []EnumValue {
	var enumValues []EnumValue

	// Fallback: Iterate through all objects in the package scope
	scope := typePkg.Scope()
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)

		// Check if it's a constant
		constObj, ok := obj.(*types.Const)
		if !ok {
			continue
		}

		// Check if the constant's type matches our enum type
		if !types.Identical(constObj.Type(), enumType) {
			continue
		}

		// Extract the constant value
		val := constObj.Val()
		if val == nil {
			continue
		}

		// Convert to int64
		var intValue int64
		switch val.Kind() {
		case constant.Int:
			// Get the int64 value
			if i, ok := constant.Int64Val(val); ok {
				intValue = i
			} else {
				// Value too large for int64, skip
				continue
			}
		default:
			// Not an integer constant, skip
			continue
		}

		// Try to get comment for this constant
		var description string
		if pkg != nil {
			description = findConstComment(pkg, constObj.Name())
		}

		enumValues = append(enumValues, EnumValue{
			Name:        constObj.Name(),
			Value:       intValue,
			Description: description,
		})
	}

	// Sort by value to ensure consistent ordering
	sort.Slice(enumValues, func(i, j int) bool {
		return enumValues[i].Value < enumValues[j].Value
	})

	return enumValues
}

func convertToManifestMethods(funcs []ExportedFunction) []Method {
	methods := make([]Method, len(funcs))
	for i, f := range funcs {
		methods[i] = Method{
			Name:        f.ExportName,
			FuncName:    f.FuncName,
			Description: f.Description,
			ParamTypes:  convertParams(f.Params),
			RetType:     convertReturnType(f.ReturnType),
		}
	}
	return methods
}

func convertParams(params []ParamInfo) []Property {
	result := make([]Property, len(params))
	for i, p := range params {
		if p.Type.IsFunc {
			// Function parameter with prototype
			result[i] = Property{
				Type:        "function",
				Name:        p.Name,
				Description: p.Description,
				Ref:         p.Type.IsRef,
				Prototype: &Method{
					Name:        p.Type.FuncSig.Name,
					FuncName:    "_",
					Description: p.Type.FuncSig.Description,
					ParamTypes:  convertParams(p.Type.FuncSig.Params),
					RetType:     convertReturnType(p.Type.FuncSig.Return),
				},
			}
		} else if p.Type.IsArray && p.Type.ElemType != nil && p.Type.ElemType.IsEnum {
			// Array of enum parameter
			result[i] = Property{
				Type:        p.Type.TypeString,
				Name:        p.Name,
				Description: p.Description,
				Ref:         p.Type.IsRef,
				Enumerator: &EnumObject{
					Name:        p.Type.ElemType.EnumTypeName,
					Description: p.Type.ElemType.Description,
					Values:      p.Type.ElemType.EnumValues,
				},
			}
		} else if p.Type.IsEnum {
			// Enum parameter
			result[i] = Property{
				Type:        p.Type.TypeString,
				Name:        p.Name,
				Description: p.Description,
				Ref:         p.Type.IsRef,
				Enumerator: &EnumObject{
					Name:        p.Type.EnumTypeName,
					Description: p.Type.Description,
					Values:      p.Type.EnumValues,
				},
			}
		} else {
			// Regular parameter
			result[i] = Property{
				Type:        p.Type.TypeString,
				Ref:         p.Type.IsRef,
				Name:        p.Name,
				Description: p.Description,
			}
		}
	}
	return result
}

func convertReturnType(t TypeInfo) Property {
	if t.IsFunc {
		// Function return type with prototype
		return Property{
			Type:        "function",
			Description: t.Description,
			Prototype: &Method{
				Name:        t.FuncSig.Name,
				FuncName:    "_",
				Description: t.FuncSig.Description,
				ParamTypes:  convertParams(t.FuncSig.Params),
				RetType:     convertReturnType(t.FuncSig.Return),
			},
		}
	}

	if t.IsArray && t.ElemType != nil && t.ElemType.IsEnum {
		// Array of enum return type
		return Property{
			Type:        t.TypeString,
			Description: t.Description,
			Enumerator: &EnumObject{
				Name:        t.ElemType.EnumTypeName,
				Description: t.ElemType.Description,
				Values:      t.ElemType.EnumValues,
			},
		}
	}

	if t.IsEnum {
		return Property{
			Type:        t.TypeString,
			Description: t.Description,
			Enumerator: &EnumObject{
				Name:        t.EnumTypeName,
				Description: t.Description,
				Values:      t.EnumValues,
			},
		}
	}

	return Property{
		Type:        t.TypeString,
		Description: t.Description,
	}
}

func getPtrType() string {
	if unsafe.Sizeof(uintptr(0)) == 4 {
		return "ptr32"
	} else {
		return "ptr64"
	}
}

func getIntType() string {
	if unsafe.Sizeof(uintptr(0)) == 4 {
		return "int32"
	} else {
		return "int64"
	}
}

func getUIntType() string {
	if unsafe.Sizeof(uintptr(0)) == 4 {
		return "uint32"
	} else {
		return "uint64"
	}
}

// Generate autoexports.go
func generateAutoExports(funcs []ExportedFunction, target string) error {
	code := []string{fmt.Sprintf(`package %s

// #include "autoexports.h"
import "C"
import (
	"reflect"
	"unsafe"
	"github.com/untrustedmodders/go-plugify"
)

var _ = reflect.TypeOf(0)
var _ = unsafe.Sizeof(0)
var _ = plugify.Plugin.Loaded

// Exported methods
`, target)}

	for _, fn := range funcs {
		code = append(code, generateWrapper(fn))
	}

	return os.WriteFile("autoexports.go", []byte(strings.Join(code, "")), 0644)
}

func generateWrapper(fn ExportedFunction) string {
	var params []string
	var callParams []string
	var tempVars []string
	var assignBack []string

	for _, param := range fn.Params {
		cType := mapToCType(param.Type)
		goType := mapToGoType(param.Type)
		varType := mapToFunc(param.Type)
		name := generateName(param.Name)

		// Determine if this is a ref parameter (passed by pointer)
		if param.Type.IsRef || strings.HasPrefix(cType, "C.") {
			cType = "*" + cType
		}

		params = append(params, fmt.Sprintf("%s %s", name, cType))

		castName := name
		if strings.HasPrefix(cType, "*C.") {
			castName = fmt.Sprintf("(*plugify.Plg%s)(unsafe.Pointer(%s))", cType[3:], name)
		}

		// Handle ref parameters with temporary variables
		if param.Type.IsRef && (cType == "*C.String" || cType == "*C.Vector" || cType == "*C.Variant") {
			tempVar := fmt.Sprintf("_%s", name)

			// Create temporary Go variable and convert from C
			switch param.Type.TypeString {
			case "string", "any":
				tempVars = append(tempVars, fmt.Sprintf("\t%s := plugify.Get%sData(%s)\n", tempVar, varType, castName))
				assignBack = append(assignBack, fmt.Sprintf("\tplugify.Assign%s(%s, %s)\n", varType, castName, tempVar))
			default:
				if param.Type.IsArray {
					typeName := ""
					if param.Type.ElemType != nil && param.Type.ElemType.IsEnum {
						typeName = fmt.Sprintf("T[%s]", param.Type.ElemType.EnumTypeName)
					}
					tempVars = append(tempVars, fmt.Sprintf("\t%s := plugify.GetVectorData%s%s(%s)\n", tempVar, varType, typeName, castName))
					assignBack = append(assignBack, fmt.Sprintf("\tplugify.AssignVector%s(%s, %s)\n", varType, castName, tempVar))
				}
			}

			// Pass address of temporary variable
			callParams = append(callParams, fmt.Sprintf("&%s", tempVar))
		} else {
			// Handle non-ref parameters
			switch {
			case param.Type.IsFunc:
				funcName := generateName(param.Type.FuncSig.Name)
				callParams = append(callParams, fmt.Sprintf("plugify.GetDelegateForFunctionPointer(%s, reflect.TypeOf(%s(nil))).(%s)", name, funcName, funcName))
			case param.Type.TypeString == "string" || param.Type.TypeString == "any":
				callParams = append(callParams, fmt.Sprintf("plugify.Get%sData(%s)", varType, castName))
			case param.Type.TypeString == "vec2" || param.Type.TypeString == "vec3" || param.Type.TypeString == "vec4" || param.Type.TypeString == "mat4x4":
				deref := "*"
				if param.Type.IsRef {
					deref = ""
				}
				callParams = append(callParams, fmt.Sprintf("%s(*%s)(unsafe.Pointer(%s))", deref, goType, name))
			case param.Type.IsArray:
				typeName := ""
				if param.Type.ElemType != nil && param.Type.ElemType.IsEnum {
					typeName = fmt.Sprintf("T[%s]", param.Type.ElemType.EnumTypeName)
				}
				callParams = append(callParams, fmt.Sprintf("plugify.GetVectorData%s%s(%s)", varType, typeName, castName))
			case param.Type.IsEnum:
				if param.Type.IsRef {
					callParams = append(callParams, fmt.Sprintf("(*%s)(%s)", param.Type.EnumTypeName, name))
				} else {
					callParams = append(callParams, fmt.Sprintf("%s(%s)", param.Type.EnumTypeName, name))
				}
			default:
				callParams = append(callParams, name)
			}
		}
	}

	paramList := strings.Join(params, ", ")
	callParamList := strings.Join(callParams, ", ")

	// Generate return type
	resultDest := ""
	returnType := mapToCType(fn.ReturnType)
	if returnType != "" {
		returnType = " " + returnType
		resultDest = "__result := "
	}

	// Generate function wrapper
	wrapper := []string{fmt.Sprintf(`
//export %s
func %s(%s)%s {
`, fn.FuncName, fn.FuncName, paramList, returnType)}

	if len(tempVars) > 0 {
		wrapper = append(wrapper, strings.Join(tempVars, ""))
	}

	wrapper = append(wrapper, fmt.Sprintf("\t%s%s(%s)\n", resultDest, fn.ExportName, callParamList))

	if len(assignBack) > 0 {
		wrapper = append(wrapper, strings.Join(assignBack, ""))
	}

	// Handle return conversion
	if resultDest != "" {
		cType := mapToCType(fn.ReturnType)
		varType := mapToFunc(fn.ReturnType)

		switch {
		case fn.ReturnType.IsFunc:
			wrapper = append(wrapper, "\treturn plugify.GetFunctionPointerForDelegate(__result)\n")
		case fn.ReturnType.TypeString == "string" || fn.ReturnType.TypeString == "any":
			wrapper = append(wrapper,
				fmt.Sprintf("\t__return := plugify.Construct%s(__result)\n", varType),
				fmt.Sprintf("\treturn *(*%s)(unsafe.Pointer(&__return))\n", cType))
		case fn.ReturnType.TypeString == "vec2" || fn.ReturnType.TypeString == "vec3" || fn.ReturnType.TypeString == "vec4" || fn.ReturnType.TypeString == "mat4x4":
			wrapper = append(wrapper, fmt.Sprintf("\treturn *(*%s)(unsafe.Pointer(&__result))\n", cType))
		case fn.ReturnType.IsArray:
			typeName := ""
			if fn.ReturnType.ElemType != nil && fn.ReturnType.ElemType.IsEnum {
				typeName = fmt.Sprintf("[%s]", fn.ReturnType.ElemType.EnumTypeName)
			}
			wrapper = append(wrapper,
				fmt.Sprintf("\t__return := plugify.ConstructVector%s%s(__result)\n", varType, typeName),
				fmt.Sprintf("\treturn *(*%s)(unsafe.Pointer(&__return))\n", cType))
		case fn.ReturnType.IsEnum:
			wrapper = append(wrapper, fmt.Sprintf("\treturn %s(__result)\n", fn.ReturnType.EnumTypeName))
		default:
			wrapper = append(wrapper, "\treturn __result\n")
		}
	}

	wrapper = append(wrapper, "}\n")
	return strings.Join(wrapper, "")
}

// Helper functions for type mapping (refactored to use TypeInfo)
func mapToFunc(typeInfo TypeInfo) string {
	// Handle array types - return the element type's function name
	if typeInfo.IsArray && typeInfo.ElemType != nil {
		return mapToFuncBasic(typeInfo.ElemType.TypeString)
	}

	// Handle basic types
	return mapToFuncBasic(typeInfo.TypeString)
}

func mapToFuncBasic(typeStr string) string {
	switch typeStr {
	case "void":
		return ""
	case "bool":
		return "Bool"
	case "char8":
		return "Int8"
	case "char16":
		return "UInt16"
	case "int8":
		return "Int8"
	case "int16":
		return "Int16"
	case "int32":
		return "Int32"
	case "int64":
		return "Int64"
	case "uint8":
		return "UInt8"
	case "uint16":
		return "UInt16"
	case "uint32":
		return "UInt32"
	case "uint64":
		return "UInt64"
	case "ptr64", "ptr32":
		return "Pointer"
	case "float":
		return "Float"
	case "double":
		return "Double"
	case "string":
		return "String"
	case "any":
		return "Variant"
	case "vec2":
		return "Vector2"
	case "vec3":
		return "Vector3"
	case "vec4":
		return "Vector4"
	case "mat4x4":
		return "Matrix4x4"
	default:
		return "Pointer"
	}
}

func mapToCType(typeInfo TypeInfo) string {
	// Handle function types
	if typeInfo.IsFunc {
		return "unsafe.Pointer"
	}

	// Handle array types
	if typeInfo.IsArray {
		return "C.Vector"
	}

	// Handle enum types - use the underlying type
	if typeInfo.IsEnum {
		return mapToCTypeBasic(typeInfo.TypeString)
	}

	// Handle basic types
	return mapToCTypeBasic(typeInfo.TypeString)
}

func mapToCTypeBasic(typeStr string) string {
	switch typeStr {
	case "void":
		return ""
	case "bool":
		return "bool"
	case "char8":
		return "int8"
	case "char16":
		return "uint16"
	case "int8":
		return "int8"
	case "int16":
		return "int16"
	case "int32":
		return "int32"
	case "int64":
		return "int64"
	case "uint8":
		return "uint8"
	case "uint16":
		return "uint16"
	case "uint32":
		return "uint32"
	case "uint64":
		return "uint64"
	case "ptr64", "ptr32":
		return "uintptr"
	case "float":
		return "float32"
	case "double":
		return "float64"
	case "string":
		return "C.String"
	case "any":
		return "C.Variant"
	case "vec2":
		return "C.Vector2"
	case "vec3":
		return "C.Vector3"
	case "vec4":
		return "C.Vector4"
	case "mat4x4":
		return "C.Matrix4x4"
	default:
		return "uintptr"
	}
}

func mapToGoType(typeInfo TypeInfo) string {
	// Handle function types
	if typeInfo.IsFunc && typeInfo.FuncSig != nil {
		return typeInfo.FuncSig.Name
	}

	// Handle array types
	if typeInfo.IsArray && typeInfo.ElemType != nil {
		elemType := mapToGoTypeBasic(typeInfo.ElemType.TypeString)
		// Handle arrays of enums
		if typeInfo.ElemType.IsEnum {
			return "[]" + typeInfo.ElemType.EnumTypeName
		}
		return "[]" + elemType
	}

	// Handle enum types - use the actual enum type name
	if typeInfo.IsEnum {
		return typeInfo.EnumTypeName
	}

	// Handle basic types
	return mapToGoTypeBasic(typeInfo.TypeString)
}

func mapToGoTypeBasic(typeStr string) string {
	switch typeStr {
	case "void":
		return ""
	case "bool":
		return "bool"
	case "char8":
		return "int8"
	case "char16":
		return "uint16"
	case "int8":
		return "int8"
	case "int16":
		return "int16"
	case "int32":
		return "int32"
	case "int64":
		return "int64"
	case "uint8":
		return "uint8"
	case "uint16":
		return "uint16"
	case "uint32":
		return "uint32"
	case "uint64":
		return "uint64"
	case "ptr64", "ptr32":
		return "uintptr"
	case "float":
		return "float32"
	case "double":
		return "float64"
	case "string":
		return "string"
	case "any":
		return "interface{}"
	case "vec2":
		return "plugify.Vector2"
	case "vec3":
		return "plugify.Vector3"
	case "vec4":
		return "plugify.Vector4"
	case "mat4x4":
		return "plugify.Matrix4x4"
	default:
		return "interface{}"
	}
}

// Generate autoexports.h
func generateAutoExportsHeader() error {
	header := `#pragma once
// autoexports.h - Auto-generated by Plugify Go Generator

#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>

typedef unsigned short char16_t;

#ifdef __cplusplus
extern "C" {
#endif

typedef struct String { char* data; size_t size; size_t cap; } String;
typedef struct Vector { void* begin; void* end; void* capacity; } Vector;
typedef struct Vector2 { float x, y; } Vector2;
typedef struct Vector3 { float x, y, z; } Vector3;
typedef struct Vector4 { float x, y, z, w; } Vector4;
typedef struct Matrix4x4 { float m[4][4]; } Matrix4x4;
typedef struct Variant {
	union {
		bool boolean;
		char char8;
		char16_t char16;
		int8_t int8;
		int16_t int16;
		int32_t int32;
		int64_t int64;
		uint8_t uint8;
		uint16_t uint16;
		uint32_t uint32;
		uint64_t uint64;
		void* ptr;
		float flt;
		double dbl;
		String str;
		Vector vec;
		Vector2 vec2;
		Vector3 vec3;
		Vector4 vec4;
	};
#if INTPTR_MAX == INT32_MAX
	volatile char pad[8];
#endif
	uint8_t current;
} Variant;

#ifdef __cplusplus
}
#endif
`
	return os.WriteFile("autoexports.h", []byte(header), 0644)
}
