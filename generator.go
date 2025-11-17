package plugify

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/types"
	"strings"
	"os"
	"golang.org/x/tools/go/packages"
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
	EnumValues   []EnumValue
	IsArray      bool
	ElemType     *TypeInfo
}

type FuncSignature struct {
	Name   string
	Params []ParamInfo
	Return TypeInfo
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

	outputFile := output
	if outputFile == "" {
		outputFile = pluginName + ".pplugin"
	}

	manifest := Manifest{
		Schema:      "https://raw.githubusercontent.com/untrustedmodders/plugify/refs/heads/main/schemas/plugin.schema.json",
		Name:        pluginName,
		Version:     version,
		Description: description,
		Author:      author,
		Website:     website,
		License:     license,
		Entry:       pluginEntry,
		Language:    "golang",
		Methods:     convertToManifestMethods(exportedFuncs),
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
				FuncName:   "__" + funcDecl.Name.Name,
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

func findEnumValues(named *types.Named, info *types.Info) []EnumValue {
	// This is a heuristic approach
	// In Go, enums are typically const groups with the same type
	// We'd need to scan the package scope for const declarations
	// For now, return empty - you could extend this
	return nil
}

func convertToManifestMethods(funcs []ExportedFunction) []Method {
	methods := make([]Method, len(funcs))
	for i, f := range funcs {
		methods[i] = Method{
			Name:       f.ExportName,
			FuncName:   f.FuncName,
			ParamTypes: convertParams(f.Params),
			RetType:    convertReturnType(f.ReturnType),
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
				Type: "function",
				Name: p.Name,
				Prototype: &Method{
					Name:       p.Type.FuncSig.Name,
					ParamTypes: convertParams(p.Type.FuncSig.Params),
					RetType:    convertReturnType(p.Type.FuncSig.Return),
				},
			}
		} else if p.Type.IsEnum {
			// Enum parameter
			result[i] = Property{
				Type: p.Type.TypeString,
				Name: p.Name,
				Enumerator: &EnumObject{
					Name:   p.Type.EnumTypeName,
					Values: p.Type.EnumValues,
				},
			}
		} else {
			// Regular parameter
			result[i] = Property{
				Type: p.Type.TypeString,
				Name: p.Name,
			}
		}
	}
	return result
}

func convertReturnType(t TypeInfo) Property {
	if t.IsEnum {
		return Property{
			Type: t.TypeString,
			Enumerator: &EnumObject{
				Name:   t.EnumTypeName,
				Values: t.EnumValues,
			},
		}
	}

	return Property{
		Type: t.TypeString,
	}
}

// Generate autoexports.go
func generateAutoExports(funcs []ExportedFunction, target string) error {
	code := []string{fmt.Sprintf(`package %s

// #include "autoexports.h"
import "C"
import (
	_ "github.com/untrustedmodders/go-plugify"
	_ "reflect"
	_ "unsafe"
)

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
		cType := mapToCType(param.Type.TypeString)
		goType := mapToGoType(param.Type.TypeString)
		varType := mapToFunc(param.Type.TypeString)
		name := generateName(param.Name)

		// Determine if this is a ref parameter (passed by pointer)
		isRef := false
		if strings.HasPrefix(cType, "C.") {
			cType = "*" + cType
			isRef = true
		}

		params = append(params, fmt.Sprintf("%s %s", name, cType))

		castName := name
		if strings.HasPrefix(cType, "*C.") {
			castName = fmt.Sprintf("(*plugify.Plg%s)(unsafe.Pointer(%s))", cType[3:], name)
		}

		// Handle ref parameters with temporary variables
		if isRef && (cType == "*C.String" || cType == "*C.Vector" || cType == "*C.Variant") {
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
				if isRef {
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
				if isRef {
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
	returnType := mapToCType(fn.ReturnType.TypeString)
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
		cType := mapToCType(fn.ReturnType.TypeString)
		varType := mapToFunc(fn.ReturnType.TypeString)

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

// Helper functions for type mapping (from original autoexports.go)
func mapToFunc(jsonType string) string {
	switch jsonType {
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
	case "bool[]":
		return "Bool"
	case "char8[]":
		return "Char8"
	case "char16[]":
		return "Char16"
	case "int8[]":
		return "Int8"
	case "int16[]":
		return "Int16"
	case "int32[]":
		return "Int32"
	case "int64[]":
		return "Int64"
	case "uint8[]":
		return "UInt8"
	case "uint16[]":
		return "UInt16"
	case "uint32[]":
		return "UInt32"
	case "uint64[]":
		return "UInt64"
	case "ptr32[]", "ptr64[]":
		return "Pointer"
	case "float[]":
		return "Float"
	case "double[]":
		return "Double"
	case "string[]":
		return "String"
	case "any[]":
		return "Variant"
	case "vec2[]":
		return "Vector2"
	case "vec3[]":
		return "Vector3"
	case "vec4[]":
		return "Vector4"
	case "mat4x4[]":
		return "Matrix4x4"
	default:
		return "Function"
	}
}

func mapToCType(jsonType string) string {
	switch jsonType {
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
	case "bool[]", "char8[]", "char16[]", "int8[]", "int16[]", "int32[]", "int64[]", "uint8[]", "uint16[]", "uint32[]", "uint64[]", "ptr32[]", "ptr64[]", "float[]", "double[]", "string[]", "any[]", "vec2[]", "vec3[]", "vec4[]", "mat4x4[]":
		return "C.Vector"
	default:
		return "unsafe.Pointer"
	}
}

func mapToGoType(jsonType string) string {
	switch jsonType {
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
	case "bool[]":
		return "[]bool"
	case "int8[]":
		return "[]int8"
	case "int16[]":
		return "[]int16"
	case "int32[]":
		return "[]int32"
	case "int64[]":
		return "[]int64"
	case "uint8[]":
		return "[]uint8"
	case "uint16[]":
		return "[]uint16"
	case "uint32[]":
		return "[]uint32"
	case "uint64[]":
		return "[]uint64"
	case "ptr32[]", "ptr64[]":
		return "[]uintptr"
	case "float[]":
		return "[]float32"
	case "double[]":
		return "[]float64"
	case "string[]":
		return "[]string"
	case "any[]":
		return "[]interface{}"
	case "vec2[]":
		return "[]Vector2"
	case "vec3[]":
		return "[]Vector3"
	case "vec4[]":
		return "[]Vector4"
	case "mat4x4[]":
		return "[]Matrix4x4"
	case "vec2":
		return "Vector2"
	case "vec3":
		return "Vector3"
	case "vec4":
		return "Vector4"
	case "mat4x4":
		return "Matrix4x4"
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
