//go:generate generator.go

package plugify

import "C"
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Plugin represents the structure of the .pplugin file
type Plugin struct {
	Methods []Method `json:"methods"`
}

// Method represents a single exported method
type Method struct {
	Name        string         `json:"name"`
	FuncName    string         `json:"funcName"`
	ParamTypes  []PropertyType `json:"paramTypes"`
	RetType     PropertyType   `json:"retType"`
	Group       string         `json:"group,omitempty"`
	Description string         `json:"description,omitempty"`
}

// EnumValue represents a single enumeration value
type EnumValue struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Value       int64  `json:"value"`
}

// Enum represents an enumeration
type Enum struct {
	Name        string      `json:"name"`
	Description string      `json:"description,omitempty"`
	Values      []EnumValue `json:"values"`
}

// PropertyType represents a parameter type
type PropertyType struct {
	Type        string  `json:"type"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Ref         bool    `json:"ref,omitempty"`
	Prototype   *Method `json:"prototype,omitempty"`
	Enumerator  *Enum   `json:"enum,omitempty"`
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

// generateName returns a valid Go identifier by appending an underscore if the name is reserved
func generateName(name string) string {
	if _, exists := invalidNames[name]; exists {
		return name + "_"
	}
	return name
}

func Generate(packageName string, rootFolder string) {
	code := []string{fmt.Sprintf(`package %s

// #include "autoexports.h"
import "C"
import (
	"github.com/untrustedmodders/go-plugify"
	"reflect"
	"unsafe"
)

// Exported methods
`, packageName)}

	// Walk through the directory recursively
	err := filepath.Walk(rootFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file is a .pplugin file
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".pplugin") {
			// Read the file
			fileContent, err := ioutil.ReadFile(path)
			if err != nil {
				return fmt.Errorf("error reading file %s: %v", path, err)
			}

			// Parse the JSON content
			var plugin Plugin
			if err := json.Unmarshal(fileContent, &plugin); err != nil {
				return fmt.Errorf("error parsing JSON in file %s: %v", path, err)
			}

			// Generate export function wrappers
			for _, method := range plugin.Methods {
				code = append(code, generateWrapper(method))
			}
		}

		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path %v: %v\n", rootFolder, err)
	}

	// Write the generated file
	err = os.WriteFile("autoexports.go", []byte(strings.Join(code, "")), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		os.Exit(1)
	}

	header := `#pragma once
// autoexports.h
#pragma once

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
	// Write the generated file
	err = os.WriteFile("autoexports.h", []byte(header), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		os.Exit(1)
	}

	fmt.Println("Generated autoexports.go!")
}

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
	case "ptr32[]":
		return "Pointer"
	case "ptr64[]":
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
		return "Function" // Function
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
		return "unsafe.Pointer" // Function
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
		return "[]plugify.Vector2"
	case "vec3[]":
		return "[]plugify.Vector3"
	case "vec4[]":
		return "[]plugify.Vector4"
	case "mat4x4[]":
		return "[]plugify.Matrix4x4"
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

func generateWrapper(method Method) string {
	// Generate the function signature
	var params []string
	var callParams []string
	var tempVars []string   // Temporary variables for ref parameters
	var assignBack []string // Code to assign back ref parameters

	for _, param := range method.ParamTypes {
		cType := mapToCType(param.Type)
		goType := mapToGoType(param.Type)
		varType := mapToFunc(param.Type)
		name := generateName(param.Name)

		// Handle non-ref parameters
		if param.Ref || strings.HasPrefix(cType, "C.") {
			cType = "*" + cType
		}
		params = append(params, fmt.Sprintf("%s %s", name, cType))

		castName := fmt.Sprintf("(*plugify.Plg%s)(unsafe.Pointer(%s))", cType[3:], name)

		// Handle ref parameters
		if param.Ref && (cType == "*C.String" || cType == "*C.Vector" || cType == "*C.Variant") {
			// Create a temporary variable for the ref parameter
			tempVar := fmt.Sprintf("_%s", name)

			// Pass the address of the temporary variable to the inner function
			callParams = append(callParams, fmt.Sprintf("&%s", tempVar))

			// Handle special cases for strings, arrays, and functions
			switch param.Type {
			case "string", "any":
				tempVars = append(tempVars, fmt.Sprintf("\t%s := plugify.Get%sData(%s)\n", tempVar, varType, castName))
				assignBack = append(assignBack, fmt.Sprintf("\tplugify.Assign%s(%s, %s)\n", varType, castName, tempVar))
			default:
				if strings.HasSuffix(param.Type, "[]") {
					typeName := ""
					if param.Enumerator != nil {
						typeName = fmt.Sprintf("T[%s]", param.Enumerator.Name)
					}
					tempVars = append(tempVars, fmt.Sprintf("\t%s := plugify.GetVectorData%s%s(%s)\n", tempVar, varType, typeName, castName))
					assignBack = append(assignBack, fmt.Sprintf("\tplugify.AssignVector%s(%s, %s)\n", varType, castName, tempVar))
				} else {
					tempVars = append(tempVars, fmt.Sprintf("\t%s := %s\n", tempVar, name))
					assignBack = append(assignBack, fmt.Sprintf("\t%s = %s\n", name, tempVar))
				}
			}

		} else {
			// Handle special cases for strings, arrays, and functions
			switch param.Type {
			case "function":
				funcName := generateName(param.Prototype.Name)
				callParams = append(callParams, fmt.Sprintf("plugify.GetDelegateForFunctionPointer(%s, reflect.TypeOf(%s(nil))).(%s)", name, funcName, funcName))
			case "string", "any":
				callParams = append(callParams, fmt.Sprintf("plugify.Get%sData(%s)", varType, castName))
			case "vec2", "vec3", "vec4", "mat4x4":
				deref := "*" // Dereference the pointer
				if param.Ref {
					deref = ""
				}
				callParams = append(callParams, fmt.Sprintf("%s(*%s)(unsafe.Pointer(%s))", deref, goType, name))
			default:
				if strings.HasSuffix(param.Type, "[]") {
					typeName := ""
					if param.Enumerator != nil {
						typeName = fmt.Sprintf("T[%s]", param.Enumerator.Name)
					}
					callParams = append(callParams, fmt.Sprintf("plugify.GetVectorData%s%s(%s)", varType, typeName, castName))
				} else {
					if param.Enumerator != nil {
						if param.Ref {
							callParams = append(callParams, fmt.Sprintf("(*%s)(%s)", param.Enumerator.Name, name))
						} else {
							callParams = append(callParams, fmt.Sprintf("%s(%s)", param.Enumerator.Name, name))
						}
					} else {
						callParams = append(callParams, name)
					}
				}
			}
		}
	}
	paramList := strings.Join(params, ", ")
	callParamList := strings.Join(callParams, ", ")

	// Generate the return type
	resultDest := ""
	returnType := mapToCType(method.RetType.Type)
	if returnType != "" {
		returnType = " " + returnType
		resultDest = "__result := "
	}

	// Generate the function wrapper
	wrapper := []string{fmt.Sprintf(`
//export %s
func %s(%s)%s {
`, method.FuncName, method.FuncName, paramList, returnType)}

	// Add temporary variable declarations for ref parameters
	if len(tempVars) > 0 {
		wrapper = append(wrapper, strings.Join(tempVars, ""))
	}

	// Call the inner function
	wrapper = append(wrapper, fmt.Sprintf("\t%s%s(%s)\n", resultDest, method.Name, callParamList))

	// Assign back ref parameters
	if len(assignBack) > 0 {
		wrapper = append(wrapper, strings.Join(assignBack, ""))
	}

	// Handle return type conversion
	if resultDest != "" {
		cType := mapToCType(method.RetType.Type)
		varType := mapToFunc(method.RetType.Type)
		switch method.RetType.Type {
		case "function":
			wrapper = append(wrapper, "\treturn plugify.GetFunctionPointerForDelegate(__result)\n")
		case "string", "any":
			wrapper = append(wrapper,
				fmt.Sprintf("\t__return := plugify.Construct%s(__result)\n", varType),
				fmt.Sprintf("\treturn *(*%s)(unsafe.Pointer(&__return))\n", cType))
		case "vec2", "vec3", "vec4", "mat4x4":
			wrapper = append(wrapper, fmt.Sprintf("\treturn *(*%s)(unsafe.Pointer(&__result))\n", cType))
		default:
			if strings.HasSuffix(method.RetType.Type, "[]") {
				typeName := ""
				if method.RetType.Enumerator != nil {
					typeName = fmt.Sprintf("[%s]", method.RetType.Enumerator.Name)
				}
				wrapper = append(wrapper,
					fmt.Sprintf("\t__return := plugify.ConstructVector%s%s(__result)\n", varType, typeName),
					fmt.Sprintf("\treturn *(*%s)(unsafe.Pointer(&__return))\n", cType))
			} else {
				if method.RetType.Enumerator != nil {
					wrapper = append(wrapper, fmt.Sprintf("\treturn %s(__result)\n", method.RetType.Enumerator.Name))
				} else {
					wrapper = append(wrapper, fmt.Sprintf("\treturn __result\n"))
				}
			}
		}
	}
	wrapper = append(wrapper, "}\n")

	return strings.Join(wrapper, "")
}
