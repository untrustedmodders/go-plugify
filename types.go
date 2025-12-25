package plugify

/*
#include "plugify.h"
*/
import "C"
import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

type valueType byte

const (
	Invalid valueType = iota

	// C types
	Void
	Bool
	Char8
	Char16
	Int8
	Int16
	Int32
	Int64
	UInt8
	UInt16
	UInt32
	UInt64
	Pointer
	Float
	Double
	Function

	// plg::string
	String

	// plg::any
	Any

	// plg::vector
	ArrayBool
	ArrayChar8
	ArrayChar16
	ArrayInt8
	ArrayInt16
	ArrayInt32
	ArrayInt64
	ArrayUInt8
	ArrayUInt16
	ArrayUInt32
	ArrayUInt64
	ArrayPointer
	ArrayFloat
	ArrayDouble
	ArrayString
	ArrayAny
	ArrayVector2
	ArrayVector3
	ArrayVector4
	ArrayMatrix4x4

	// plg:vec
	Vector2Type
	Vector3Type
	Vector4Type

	// plg:mat
	Matrix4x4Type

	// Helpers
	_BaseStart = Void
	_BaseEnd   = Function

	_FloatStart = Float
	_FloatEnd   = Double

	_ObjectStart = String
	_ObjectEnd   = ArrayMatrix4x4

	_ArrayStart = ArrayBool
	_ArrayEnd   = ArrayMatrix4x4

	_StructStart = Vector2Type
	_StructEnd   = Matrix4x4Type

	_LastAssigned = Matrix4x4Type
)

type managedType struct {
	valueType valueType
	ref       bool
}

const (
	expectedSize = 2
	_            = unsafe.Sizeof(managedType{}) - expectedSize
)

func sizeOfValueType(vt valueType) int {
	switch vt {
	case Void:
		return 0
	case Bool, Char8, Int8, UInt8:
		return C.sizeof_bool
	case Char16, Int16, UInt16:
		return C.sizeof_char16_t
	case Int32, UInt32, Float:
		return C.sizeof_float
	case Int64, UInt64, Double:
		return C.sizeof_double
	case Function, Pointer:
		return C.sizeof_ptrdiff_t
	case String:
		return C.sizeof_String
	case Any:
		return C.sizeof_Variant
	case ArrayBool, ArrayChar8, ArrayChar16, ArrayInt8, ArrayInt16, ArrayInt32,
		ArrayInt64, ArrayUInt8, ArrayUInt16, ArrayUInt32, ArrayUInt64, ArrayPointer,
		ArrayFloat, ArrayDouble, ArrayString, ArrayAny, ArrayVector2, ArrayVector3,
		ArrayVector4, ArrayMatrix4x4:
		return C.sizeof_Vector
	case Vector2Type:
		return C.sizeof_Vector2 // 2 floats
	case Vector3Type:
		return C.sizeof_Vector3 // 3 floats
	case Vector4Type:
		return C.sizeof_Vector4 // 4 floats
	case Matrix4x4Type:
		return C.sizeof_Matrix4x4 // 16 floats
	default:
		return 0
	}
}

var valueTypeToReflect = map[valueType]reflect.Type{
	Void:           reflect.TypeOf(nil),
	Bool:           reflect.TypeOf(true),
	Char8:          reflect.TypeOf(int8(0)),
	Char16:         reflect.TypeOf(uint16(0)),
	Int8:           reflect.TypeOf(int8(0)),
	Int16:          reflect.TypeOf(int16(0)),
	Int32:          reflect.TypeOf(int32(0)),
	Int64:          reflect.TypeOf(int64(0)),
	UInt8:          reflect.TypeOf(uint8(0)),
	UInt16:         reflect.TypeOf(uint16(0)),
	UInt32:         reflect.TypeOf(uint32(0)),
	UInt64:         reflect.TypeOf(uint64(0)),
	Pointer:        reflect.TypeOf(uintptr(0)),
	Float:          reflect.TypeOf(float32(0)),
	Double:         reflect.TypeOf(float64(0)),
	String:         reflect.TypeOf(""),
	ArrayBool:      reflect.TypeOf([]bool{}),
	ArrayChar8:     reflect.TypeOf([]int8{}),
	ArrayChar16:    reflect.TypeOf([]uint16{}),
	ArrayInt8:      reflect.TypeOf([]int8{}),
	ArrayInt16:     reflect.TypeOf([]int16{}),
	ArrayInt32:     reflect.TypeOf([]int32{}),
	ArrayInt64:     reflect.TypeOf([]int64{}),
	ArrayUInt8:     reflect.TypeOf([]uint8{}),
	ArrayUInt16:    reflect.TypeOf([]uint16{}),
	ArrayUInt32:    reflect.TypeOf([]uint32{}),
	ArrayUInt64:    reflect.TypeOf([]uint64{}),
	ArrayPointer:   reflect.TypeOf([]uintptr{}),
	ArrayFloat:     reflect.TypeOf([]float32{}),
	ArrayDouble:    reflect.TypeOf([]float64{}),
	ArrayString:    reflect.TypeOf([]string{}),
	ArrayAny:       reflect.TypeOf([]any{}),
	ArrayVector2:   reflect.TypeOf([]Vector2{}),
	ArrayVector3:   reflect.TypeOf([]Vector3{}),
	ArrayVector4:   reflect.TypeOf([]Vector4{}),
	ArrayMatrix4x4: reflect.TypeOf([]Matrix4x4{}),
	Vector2Type:    reflect.TypeOf(Vector2{}),
	Vector3Type:    reflect.TypeOf(Vector3{}),
	Vector4Type:    reflect.TypeOf(Vector4{}),
	Matrix4x4Type:  reflect.TypeOf(Matrix4x4{}),
}

var reflectToValueType = map[reflect.Type]valueType{
	reflect.TypeOf(nil):                 Void,
	reflect.TypeOf(true):                Bool,
	reflect.TypeOf(int8(0)):             Int8,
	reflect.TypeOf(int16(0)):            Int16,
	reflect.TypeOf(int32(0)):            Int32,
	reflect.TypeOf(int64(0)):            Int64,
	reflect.TypeOf(uint8(0)):            UInt8,
	reflect.TypeOf(uint16(0)):           UInt16,
	reflect.TypeOf(uint32(0)):           UInt32,
	reflect.TypeOf(uint64(0)):           UInt64,
	reflect.TypeOf(uintptr(0)):          Pointer,
	reflect.TypeOf(float32(0)):          Float,
	reflect.TypeOf(float64(0)):          Double,
	reflect.TypeOf(""):                  String,
	reflect.TypeOf([]bool{}):            ArrayBool,
	reflect.TypeOf([]int8{}):            ArrayInt8,
	reflect.TypeOf([]int16{}):           ArrayInt16,
	reflect.TypeOf([]int32{}):           ArrayInt32,
	reflect.TypeOf([]int64{}):           ArrayInt64,
	reflect.TypeOf([]uint8{}):           ArrayUInt8,
	reflect.TypeOf([]uint16{}):          ArrayUInt16,
	reflect.TypeOf([]uint32{}):          ArrayUInt32,
	reflect.TypeOf([]uint64{}):          ArrayUInt64,
	reflect.TypeOf([]uintptr{}):         ArrayPointer,
	reflect.TypeOf([]float32{}):         ArrayFloat,
	reflect.TypeOf([]float64{}):         ArrayDouble,
	reflect.TypeOf([]string{}):          ArrayString,
	reflect.TypeOf([]any{}):             ArrayAny,
	reflect.TypeOf([]Vector2{}):         ArrayVector2,
	reflect.TypeOf([]Vector3{}):         ArrayVector3,
	reflect.TypeOf([]Vector4{}):         ArrayVector4,
	reflect.TypeOf([]Matrix4x4{}):       ArrayMatrix4x4,
	reflect.TypeOf(Vector2{}):           Vector2Type,
	reflect.TypeOf(Vector3{}):           Vector3Type,
	reflect.TypeOf(Vector4{}):           Vector4Type,
	reflect.TypeOf(Matrix4x4{}):         Matrix4x4Type,
	reflect.TypeOf((*any)(nil)).Elem():  Any,
	reflect.TypeOf(reflect.TypeOf(nil)): Pointer, // For function pointers
}

func convertToReflectType(m C.MethodHandle, i int) reflect.Type {
	mt := C.Plugify_GetMethodParamType(m, C.ptrdiff_t(i))

	if mt.ref {
		return reflect.TypeOf((*interface{})(nil)).Elem()
	}

	vt := valueType(mt.valueType)

	if val, ok := valueTypeToReflect[vt]; ok {
		return val
	}

	if vt == Function {
		return createFunctionType(C.Plugify_GetMethodPrototype(m, C.ptrdiff_t(i)))
	}

	return reflect.TypeOf((*interface{})(nil)).Elem()
}

func createFunctionType(method C.MethodHandle) reflect.Type {
	if method == nil {
		panicker("expected a function")
	}

	count := int(C.Plugify_GetMethodParamCount(method))
	in := make([]reflect.Type, count)
	for i := range in {
		in[i] = convertToReflectType(method, i)
	}
	out := []reflect.Type{convertToReflectType(method, -1)}

	return reflect.FuncOf(in, out, false)
}

func createManagedType(t reflect.Type) (managedType, error) {
	baseType := t

	if baseType.Kind() == reflect.Func {
		return managedType{Function, false}, nil
	}

	ref := t.Kind() == reflect.Ptr
	if ref {
		baseType = t.Elem()
	}

	if val, ok := reflectToValueType[baseType]; ok {
		return managedType{val, ref}, nil
	}

	return managedType{}, fmt.Errorf("unsupported type: %v", t)
}

const isWindows bool = runtime.GOOS == "windows" && runtime.GOARCH != "arm64"
const is32bit bool = runtime.GOARCH == "386" || runtime.GOARCH == "arm"

func hasReturnType(returnType managedType) bool {
	hasRet := returnType.valueType >= _ObjectStart && returnType.valueType <= _ObjectEnd // params which pass by refs by default
	if !hasRet {
		var firstHidden valueType
		if isWindows || is32bit {
			firstHidden = Vector3Type
		} else {
			firstHidden = Matrix4x4Type
		}
		hasRet = returnType.valueType >= firstHidden && returnType.valueType <= _StructEnd
	}
	return hasRet
}

func getParameterTypes(fnType reflect.Type) ([]managedType, error) {
	numIn := fnType.NumIn()
	parameterTypes := make([]managedType, numIn)
	for i := 0; i < numIn; i++ {
		mt, err := createManagedType(fnType.In(i))
		if err != nil {
			return nil, fmt.Errorf("parameter %d: %w", i, err)
		}
		parameterTypes[i] = mt
	}
	return parameterTypes, nil
}

func getReturnType(fnType reflect.Type) (managedType, int, error) {
	numOut := fnType.NumOut()
	if numOut > 0 {
		mt, err := createManagedType(fnType.Out(0))
		if err != nil {
			return managedType{}, 0, fmt.Errorf("return type: %w", err)
		}
		return mt, numOut, nil
	}
	return managedType{Void, false}, 0, nil
}
