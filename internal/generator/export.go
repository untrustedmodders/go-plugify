package generator

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// Generate autoexports.go
func (g *Package) generateAutoExports(path string) error {

	buffer := bytes.Buffer{}
	buffer.WriteString("package " + g.Pkg.Name + "\n\n")
	if len(g.Exports) > 0 {
		buffer.WriteString("//#include \"autoexports.h\"\nimport \"C\"\n")
	}
	buffer.WriteString(`import (
	"reflect"
	"unsafe"

	"github.com/untrustedmodders/go-plugify"`)

	for _, imp := range g.imports {

		var impName string
		if imp.name != "" {
			impName = imp.name + " "
		}

		buffer.WriteString("\n\t" + impName + "\"" + imp.path + "\"")
	}

	buffer.WriteString("\n)\n")

	buffer.WriteString(`
var _ = reflect.TypeOf(0)
var _ = unsafe.Sizeof(0)
var _ = plugify.Plugin()
`)

	for _, fn := range g.Exports {
		buffer.WriteString(generateWrapper(fn))
	}

	return os.WriteFile(filepath.Join(path, "autoexports.go"), buffer.Bytes(), 0644)
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
		name, _ := generateName(param.Name)

		importName := getImportName(param.Type.packageImport)

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
			case "any":
				var castType string
				if param.Type.IsEnum {
					castType = importName + param.Type.EnumTypeName + "("
				}

				paramCall := fmt.Sprintf("plugify.GetVariantData(%s)", castName)
				if castType != "" {
					paramCall = castType + paramCall + ")"
				}

				tempVars = append(tempVars, fmt.Sprintf("\t%s := %s\n", tempVar, paramCall))
				assignBack = append(assignBack, fmt.Sprintf("\tplugify.AssignVariant(%s, %s)\n", castName, tempVar))
			case "string":
				typeName := "[string]"
				if param.Type.IsEnum {
					typeName = fmt.Sprintf("[%s%s]", importName, param.Type.EnumTypeName)
				}
				tempVars = append(tempVars, fmt.Sprintf("\t%s := plugify.GetStringData%s(%s)\n", tempVar, typeName, castName))
				assignBack = append(assignBack, fmt.Sprintf("\tplugify.AssignString(%s, %s)\n", castName, tempVar))
			default:
				if param.Type.IsArray {

					// Type name of elements
					var elemTypeName string
					if param.Type.ElemType != nil {
						if param.Type.ElemType.IsEnum {
							elemTypeName = getImportName(param.Type.ElemType.packageImport) + param.Type.ElemType.EnumTypeName
						} else {
							elemTypeName = param.Type.ElemType.GoBaseType
						}
					}

					// Type name of array itself
					var typeName string
					if param.Type.IsEnum {
						typeName = importName + param.Type.EnumTypeName
					} else {
						typeName, _ = strings.CutPrefix(param.Type.GoBaseType, "*")
					}

					var genericType string
					if varType == "Variant" {
						genericType = fmt.Sprintf("[%s, %s]", elemTypeName, typeName)
					} else {
						genericType = fmt.Sprintf("[%s]", elemTypeName)
					}

					var cast string
					if param.Type.IsEnum {
						cast = fmt.Sprintf("(%s%s)(", importName, param.Type.EnumTypeName)
					}

					paramCall := fmt.Sprintf("plugify.GetVectorData%s%s(%s)", varType, genericType, castName)
					if cast != "" {
						paramCall = cast + paramCall + ")"
					}

					paramCall = fmt.Sprintf("\t%s := %s\n", tempVar, paramCall)

					tempVars = append(tempVars, paramCall)
					assignBack = append(assignBack, fmt.Sprintf("\tplugify.AssignVector%s(%s, %s)\n", varType, castName, tempVar))
				}
			}

			// Pass address of temporary variable
			callParams = append(callParams, fmt.Sprintf("&%s", tempVar))
		} else {
			// Handle non-ref parameters
			switch {
			case param.Type.IsFunc:
				//funcName := generateName(param.Type.FuncSig.Name)

				funcName := importName + param.Type.FuncSig.Name

				callParams = append(callParams, fmt.Sprintf("plugify.GetDelegateForFunctionPointer(%s, reflect.TypeOf(%s(nil))).(%s)", name, funcName, funcName))
			case param.Type.TypeString == "any":
				var castType string
				if param.Type.IsEnum {
					castType = importName + param.Type.EnumTypeName + "("
				}

				paramCall := fmt.Sprintf("plugify.GetVariantData(%s)", castName)
				if castType != "" {
					paramCall = castType + paramCall + ")"
				}

				callParams = append(callParams, paramCall)
			case param.Type.TypeString == "string":
				typeName := "[string]"
				if param.Type.IsEnum {
					typeName = fmt.Sprintf("[%s%s]", importName, param.Type.EnumTypeName)
				}
				callParams = append(callParams, fmt.Sprintf("plugify.GetStringData%s(%s)", typeName, castName))
			case param.Type.TypeString == "vec2" || param.Type.TypeString == "vec3" || param.Type.TypeString == "vec4" || param.Type.TypeString == "mat4x4":
				deref := "*"
				if param.Type.IsRef {
					deref = ""
				}
				callParams = append(callParams, fmt.Sprintf("%s(*%s%s)(unsafe.Pointer(%s))", deref, importName, goType, name))
			case param.Type.IsArray:
				var elemTypeName string
				if param.Type.ElemType != nil {
					if param.Type.ElemType.IsEnum {
						elemTypeName = importName + param.Type.ElemType.EnumTypeName
					} else {
						elemTypeName = param.Type.ElemType.GoBaseType
					}
				}

				var typeName string
				if param.Type.IsEnum {
					typeName = getImportName(param.Type.ElemType.packageImport) + param.Type.EnumTypeName
				} else {
					typeName, _ = strings.CutPrefix(param.Type.GoBaseType, "*")
				}

				var genericType string
				if varType == "Variant" {
					genericType = fmt.Sprintf("[%s, %s]", elemTypeName, typeName)
				} else {
					genericType = fmt.Sprintf("[%s]", elemTypeName)
				}

				var cast string
				if param.Type.IsEnum {
					cast = fmt.Sprintf("(%s%s)(", importName, param.Type.EnumTypeName)
				}

				paramCall := fmt.Sprintf("plugify.GetVectorData%s%s(%s)", varType, genericType, castName)
				if cast != "" {
					paramCall = cast + paramCall + ")"
				}

				callParams = append(callParams, paramCall)
			case param.Type.IsEnum:
				if param.Type.IsRef {
					callParams = append(callParams, fmt.Sprintf("(*%s%s)(unsafe.Pointer(%s))", importName, param.Type.EnumTypeName, name))
				} else {
					callParams = append(callParams, fmt.Sprintf("%s%s(%s)", importName, param.Type.EnumTypeName, name))
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

	fnImportName := getImportName(fn.packageImport)

	// Generate function wrapper
	wrapper := []string{fmt.Sprintf(`
//export %s
func %s(%s)%s {
`, fn.FuncName, fn.FuncName, paramList, returnType)}

	if len(tempVars) > 0 {
		wrapper = append(wrapper, strings.Join(tempVars, ""))
	}

	wrapper = append(wrapper, fmt.Sprintf("\t%s%s%s(%s)\n", resultDest, fnImportName, fn.ExportName, callParamList))

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
				typeName = fmt.Sprintf("[%s%s]", getImportName(fn.ReturnType.ElemType.packageImport), fn.ReturnType.ElemType.EnumTypeName)
			}
			wrapper = append(wrapper,
				fmt.Sprintf("\t__return := plugify.ConstructVector%s%s(__result)\n", varType, typeName),
				fmt.Sprintf("\treturn *(*%s)(unsafe.Pointer(&__return))\n", cType))
		case fn.ReturnType.IsEnum:
			wrapper = append(wrapper, fmt.Sprintf("\treturn %s(__result)\n", fn.ReturnType.GoBaseType))
		default:
			wrapper = append(wrapper, "\treturn __result\n")
		}
	}

	wrapper = append(wrapper, "}\n")
	return strings.Join(wrapper, "")
}

// Generate autoexports.h
func GenerateAutoExportsHeader(path string) error {
	header := `#pragma once
// autoexports.h - Auto-generated by Plugify Go Generator

#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>

typedef unsigned short char16_t;

#ifdef __cplusplus
extern "C" {
#endif

typedef struct String { uintptr_t data; size_t size; size_t cap; } String;
typedef struct Vector { uintptr_t begin; uintptr_t end; uintptr_t capacity; } Vector;
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
	return os.WriteFile(filepath.Join(path, "autoexports.h"), []byte(header), 0644)
}

func getImportName(packageImport packageImport) string {
	if packageImport.name != "" {
		return packageImport.name + "."
	}
	return ""
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
	case "int":
		/* if is32bit {
			return "int32"
		} */
		return "int64"
	case "uint8":
		return "uint8"
	case "uint16":
		return "uint16"
	case "uint32":
		return "uint32"
	case "uint64":
		return "uint64"
	case "uint":
		/* if is32bit {
			return "uint32"
		} */
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
		return getImportName(typeInfo.packageImport) + typeInfo.FuncSig.Name
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

func generateName(name string) (string, bool) {
	if _, exists := invalidNames[name]; exists {
		return name + "_", true
	}
	return name, false
}
