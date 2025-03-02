package plugify

/*
#include <plugify.h>
*/
import "C"
import (
	"fmt"
	"reflect"
	"runtime"
	"unsafe"
)

type ValueType byte

const (
	Invalid ValueType = iota

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
	/*_BaseStart = Void
	_BaseEnd   = Function

	_FloatStart = Float
	_FloatEnd   = Double

	_ObjectStart = String
	_ObjectEnd   = ArrayMatrix4x4

	_ArrayStart = ArrayBool
	_ArrayEnd   = ArrayMatrix4x4

	_StructStart = Vector2
	_StructEnd   = Matrix4x4

	_LastAssigned = Matrix4x4*/
)

type ManagedType struct {
	valueType ValueType
	ref       bool
}

const (
	ExpectedSize = 2                                                 // Expected size in bytes
	_            = uint(unsafe.Sizeof(ManagedType{})) - ExpectedSize // Compile-time check
)

// typeSwitcher maps Go reflect.Types to ValueType
var typeSwitcher = map[reflect.Type]ValueType{
	reflect.TypeOf(nil):                        Void,
	reflect.TypeOf(true):                       Bool,
	reflect.TypeOf(int8(0)):                    Int8,
	reflect.TypeOf(int16(0)):                   Int16,
	reflect.TypeOf(int32(0)):                   Int32,
	reflect.TypeOf(int64(0)):                   Int64,
	reflect.TypeOf(uint8(0)):                   UInt8,
	reflect.TypeOf(uint16(0)):                  UInt16,
	reflect.TypeOf(uint32(0)):                  UInt32,
	reflect.TypeOf(uint64(0)):                  UInt64,
	reflect.TypeOf(uintptr(0)):                 Pointer,
	reflect.TypeOf(float32(0)):                 Float,
	reflect.TypeOf(float64(0)):                 Double,
	reflect.TypeOf(""):                         String,
	reflect.TypeOf([]bool{}):                   ArrayBool,
	reflect.TypeOf([]int8{}):                   ArrayInt8,
	reflect.TypeOf([]int16{}):                  ArrayInt16,
	reflect.TypeOf([]int32{}):                  ArrayInt32,
	reflect.TypeOf([]int64{}):                  ArrayInt64,
	reflect.TypeOf([]uint8{}):                  ArrayUInt8,
	reflect.TypeOf([]uint16{}):                 ArrayUInt16,
	reflect.TypeOf([]uint32{}):                 ArrayUInt32,
	reflect.TypeOf([]uint64{}):                 ArrayUInt64,
	reflect.TypeOf([]uintptr{}):                ArrayPointer,
	reflect.TypeOf([]float32{}):                ArrayFloat,
	reflect.TypeOf([]float64{}):                ArrayDouble,
	reflect.TypeOf([]string{}):                 ArrayString,
	reflect.TypeOf([]interface{}{}):            ArrayAny,
	reflect.TypeOf([]Vector2{}):                ArrayVector2,
	reflect.TypeOf([]Vector3{}):                ArrayVector3,
	reflect.TypeOf([]Vector4{}):                ArrayVector4,
	reflect.TypeOf([]Matrix4x4{}):              ArrayMatrix4x4,
	reflect.TypeOf(Vector2{}):                  Vector2Type,
	reflect.TypeOf(Vector3{}):                  Vector3Type,
	reflect.TypeOf(Vector4{}):                  Vector4Type,
	reflect.TypeOf(Matrix4x4{}):                Matrix4x4Type,
	reflect.TypeOf((*interface{})(nil)).Elem(): Any,
	reflect.TypeOf(reflect.TypeOf(nil)):        Pointer, // For function pointers
}

// maps a reflect.Type to ValueType
func ConvertToValueType(t reflect.Type) ValueType {
	baseType := t
	if t.Kind() == reflect.Ptr {
		baseType = t.Elem()
	}

	if baseType.Kind() == reflect.Array {
		elementType := baseType.Elem()
		if elementType.Kind() == reflect.Int {
			baseType = reflect.SliceOf(elementType)
		}
	}

	if baseType.Kind() == reflect.Func {
		return Function
	}

	if val, ok := typeSwitcher[baseType]; ok {
		return val
	}

	return Invalid
}

func NewManagedType(t reflect.Type) ManagedType {
	return ManagedType{
		valueType: ConvertToValueType(t),
		ref:       t.Kind() == reflect.Ptr,
	}
}

// ExternalInvoke creates a Go function that calls a C++ function
/*func ExternalInvoke(funcAddress uintptr, fnType reflect.Type) interface{} {
	if funcAddress == 0 {
		return nil
	}

	if fnType.Kind() != reflect.Func {
		return nil
	}

	parameterTypes := make([]ManagedType, fnType.NumIn())
	for i := 0; i < fnType.NumIn(); i++ {
		parameterTypes[i] = NewManagedType(fnType.In(i))
	}
	returnType := NewManagedType(fnType.Out(0))

	var paramsPtr *C.ManagedType
	if fnType.NumIn() > 0 {
		paramsPtr = (*C.ManagedType)(unsafe.Pointer(&parameterTypes[0]))
	} else {
		paramsPtr = nil
	}

	call := C.Plugify_NewCall(unsafe.Pointer(funcAddress), paramsPtr, C.ptrdiff_t(fnType.NumIn()), *(*C.ManagedType)(unsafe.Pointer(&returnType)))
	if call == nil || C.Plugify_GetCallFunction(call) == nil {
		return nil
	}

	// Create a wrapper function using reflect.MakeFunc
	wrapper := reflect.MakeFunc(fnType, func(args []reflect.Value) (results []reflect.Value) {
		params := make([]uintptr, len(args))

		for i, arg := range args {
			paramType := parameterTypes[i]
			valueType := paramType.valueType
			if paramType.ref {
				switch valueType {
				case Bool:
					params[i] = Pin(arg)
					pin++
				case Char8:
					params[i] = Pin(arg)
					pin++
				case Char16:
					params[i] = Pin(arg)
					pin++
				case Int8:
					params[i] = Pin(arg)
					pin++
				case Int16:
					params[i] = Pin(arg)
					pin++
				case Int32:
					params[i] = Pin(arg)
					pin++
				case Int64:
					params[i] = Pin(arg)
					pin++
				case UInt8:
					params[i] = Pin(arg)
					pin++
				case UInt16:
					params[i] = Pin(arg)
					pin++
				case UInt32:
					params[i] = Pin(arg)
					pin++
				case UInt64:
					params[i] = Pin(arg)
					pin++
				case Pointer:
					params[i] = Pin(arg)
					pin++
				case Float:
					params[i] = Pin(arg)
					pin++
				case Double:
					params[i] = Pin(arg)
					pin++
				case Vector2Type:
					params[i] = Pin(arg)
					pin++
				case Vector3Type:
					params[i] = Pin(arg)
					pin++
				case Vector4Type:
					params[i] = Pin(arg)
					pin++
				case Matrix4x4Type:
					params[i] = Pin(arg)
					pin++
				case Function:
					ptr := GetFunctionPointerForDelegate(arg)
					params[i] = Pin(ptr)
					pin++
				case String:
					tmp := NativeMethods.ConstructString(paramValue.(string))
					ptr := Pin(tmp)
					handlers[handle] = PtrValueType{ptr, valueType}
					handle++
					params[i] = ptr
					pin++
				case Any:
					tmp := NativeMethods.ConstructVariant(arg)
					ptr := Pin(tmp)
					handlers[handle] = PtrValueType{ptr, valueType}
					handle++
					params[i] = ptr
					pin++
				case ArrayBool:
					arr := paramValue.([]bool)
					tmp := NativeMethods.ConstructVectorBool(arr, len(arr))
					ptr := Pin(tmp)
					handlers[handle] = PtrValueType{ptr, valueType}
					handle++
					params[i] = ptr
					pin++
				case ArrayChar8:
					arr := paramValue.([]int8)
					tmp := NativeMethods.ConstructVectorChar8(arr, len(arr))
					ptr := Pin(tmp)
					handlers[handle] = PtrValueType{ptr, valueType}
					handle++
					params[i] = ptr
					pin++
				case ArrayChar16:
					arr := paramValue.([]int16)
					tmp := NativeMethods.ConstructVectorChar16(arr, len(arr))
					ptr := Pin(tmp)
					handlers[handle] = PtrValueType{ptr, valueType}
					handle++
					params[i] = ptr
					pin++
				case ArrayInt8:
					arr := paramValue.([]int8)
					tmp := NativeMethods.ConstructVectorInt8(arr, len(arr))
					ptr := Pin(tmp)
					handlers[handle] = PtrValueType{ptr, valueType}
					handle++
					params[i] = ptr
					pin++
				case ArrayInt16:
					arr := paramValue.([]int16)
					tmp := NativeMethods.ConstructVectorInt16(arr, len(arr))
					ptr := Pin(tmp)
					handlers[handle] = PtrValueType{ptr, valueType}
					handle++
					params[i] = ptr
					pin++
				case ArrayInt32:
					arr := paramValue.([]int32)
					tmp := NativeMethods.ConstructVectorInt32(arr, len(arr))
					ptr := Pin(tmp)
					handlers[handle] = PtrValueType{ptr, valueType}
					handle++
					params[i] = ptr
					pin++
				case ArrayInt64:
					arr := paramValue.([]int64)
					tmp := NativeMethods.ConstructVectorInt64(arr, len(arr))
					ptr := Pin(tmp)
					handlers[handle] = PtrValueType{ptr, valueType}
					handle++
					params[i] = ptr
					pin++
				case ArrayUInt8:
					arr := paramValue.([]uint8)
					tmp := NativeMethods.ConstructVectorUInt8(arr, len(arr))
					ptr := Pin(tmp)
					handlers[handle] = PtrValueType{ptr, valueType}
					handle++
					params[i] = ptr
					pin++
				case ArrayUInt16:
					arr := paramValue.([]uint16)
					tmp := NativeMethods.ConstructVectorUInt16(arr, len(arr))
					ptr := Pin(tmp)
					handlers[handle] = PtrValueType{ptr, valueType}
					handle++
					params[i] = ptr
					pin++
				case ArrayUInt32:
					arr := paramValue.([]uint32)
					tmp := NativeMethods.ConstructVectorUInt32(arr, len(arr))
					ptr := Pin(tmp)
					handlers[handle] = PtrValueType{ptr, valueType}
					handle++
					params[i] = ptr
					pin++
				case ArrayUInt64:
					arr := paramValue.([]uint64)
					tmp := NativeMethods.ConstructVectorUInt64(arr, len(arr))
					ptr := Pin(tmp)
					handlers[handle] = PtrValueType{ptr, valueType}
					handle++
					params[i] = ptr
					pin++
				case ArrayPointer:
					arr := paramValue.([]nint)
					tmp := NativeMethods.ConstructVectorIntPtr(arr, len(arr))
					ptr := Pin(tmp)
					handlers[handle] = PtrValueType{ptr, valueType}
					handle++
					params[i] = ptr
					pin++
				case ArrayFloat:
					arr := paramValue.([]float32)
					tmp := NativeMethods.ConstructVectorFloat(arr, len(arr))
					ptr := Pin(tmp)
					handlers[handle] = PtrValueType{ptr, valueType}
					handle++
					params[i] = ptr
					pin++
				case ArrayDouble:
					arr := paramValue.([]float64)
					tmp := NativeMethods.ConstructVectorDouble(arr, len(arr))
					ptr := Pin(tmp)
					handlers[handle] = PtrValueType{ptr, valueType}
					handle++
					params[i] = ptr
					pin++
				case ArrayString:
					arr := paramValue.([]string)
					tmp := NativeMethods.ConstructVectorString(arr, len(arr))
					ptr := Pin(tmp)
					handlers[handle] = PtrValueType{ptr, valueType}
					handle++
					params[i] = ptr
					pin++
				case ArrayAny:
					arr := paramValue.([]interface{})
					tmp := NativeMethods.ConstructVectorVariant(arr, len(arr))
					ptr := Pin(tmp)
					handlers[handle] = PtrValueType{ptr, valueType}
					handle++
					params[i] = ptr
					pin++
				case ArrayVector2:
					arr := paramValue.([]Vector2)
					tmp := NativeMethods.ConstructVectorVector2(arr, len(arr))
					ptr := Pin(tmp)
					handlers[handle] = PtrValueType{ptr, valueType}
					handle++
					params[i] = ptr
					pin++
				case ArrayVector3:
					arr := paramValue.([]Vector3)
					tmp := NativeMethods.ConstructVectorVector3(arr, len(arr))
					ptr := Pin(tmp)
					handlers[handle] = PtrValueType{ptr, valueType}
					handle++
					params[i] = ptr
					pin++
				case ArrayVector4:
					arr := paramValue.([]Vector4)
					tmp := NativeMethods.ConstructVectorVector4(arr, len(arr))
					ptr := Pin(tmp)
					handlers[handle] = PtrValueType{ptr, valueType}
					handle++
					params[i] = ptr
					pin++
				case ArrayMatrix4x4:
					arr := paramValue.([]Matrix4x4)
					tmp := NativeMethods.ConstructVectorMatrix4x4(arr, len(arr))
					ptr := Pin(tmp)
					handlers[handle] = PtrValueType{ptr, valueType}
					handle++
					params[i] = ptr
					pin++
				default:
					panic(fmt.Sprintf("Parameter '%v' uses not supported type for marshalling!", valueType))
				}
			} else {
				switch valueType {
				case Bool:
					params[i] = Pin(arg)
				case Char8:
					params[i] = Pin(arg)
				case Char16:
					params[i] = Pin(arg)
				case Int8:
					params[i] = Pin(arg)
				case Int16:
					params[i] = Pin(arg)
				case Int32:
					params[i] = Pin(arg)
				case Int64:
					params[i] = Pin(arg)
				case UInt8:
					params[i] = Pin(arg)
				case UInt16:
					params[i] = Pin(arg)
				case UInt32:
					params[i] = Pin(arg)
				case UInt64:
					params[i] = Pin(arg)
				case Pointer:
					params[i] = Pin(arg)
				case Float:
					params[i] = Pin(arg)
				case Double:
					params[i] = Pin(arg)
				case Vector2:
					params[i] = Pin(arg)
				case Vector3:
					params[i] = Pin(arg)
				case Vector4:
					params[i] = Pin(arg)
				case Matrix4x4:
					params[i] = Pin(arg)
				case Function:
					ptr := GetFunctionPointerForDelegate(arg)
					params[i] = Pin(ptr)
				}
			}
		}

		//var retVals []uintptr
		//call.Function(cArgs, &retVals)

		// Prepare return values
		results = make([]reflect.Value, fnType.NumOut())
		if fnType.NumOut() > 0 {
			//results[0] = reflect.ValueOf(int(retVals[0])) // Convert result to Go type
		}

		return results
	})

	return wrapper.Interface()
}
*/

type Func struct {
	fn interface{}
	cb C.JitCallback
}

var functionPointerMap = make(map[uintptr]interface{})

//export Plugify_InternalCall
func Plugify_InternalCall(cb unsafe.Pointer, data unsafe.Pointer, args C.size_t, numArgs unsafe.Pointer, ret unsafe.Pointer) {

}

func GetFunctionPointerForDelegate(fn interface{}) uintptr {
	val := reflect.ValueOf(fn)
	if val.Kind() != reflect.Func {
		panic("expected a function")
	}

	for k, v := range functionPointerMap {
		if v == fn {
			return k
		}
	}

	handler := C.Plugify_FindFunctionByName("Plugify_InternalCall")

	name := val.String()
	cb := C.Plugify_NewCallback(name, unsafe.Pointer(handler), unsafe.Pointer(val.Pointer()))
	if cb == nil {
		panic(fmt.Sprintf("Method '%s' has JIT generation error: not found", name))
	}

	f := &Func{fn, cb}

	runtime.SetFinalizer(f, func(f *Func) {
		C.Plugify_DeleteCallback(f.cb)
	})

	addr := C.Plugify_GetCallbackFunction(cb)
	if addr == nil {
		panic(fmt.Sprintf("Method '%s' has JIT generation error: %s", name, C.GoString(C.Plugify_GetCallbackError(cb))))
	}

	return 0
}

/*
	funcType := val.Type()
	numIn := funcType.NumIn()
	numOut := funcType.NumOut()

	fmt.Printf("Function has %d input parameters and %d output parameters\n", numIn, numOut)
	for i := 0; i < numIn; i++ {
		fmt.Printf("Input parameter %d: %s\n", i, funcType.In(i))
	}
	for i := 0; i < numOut; i++ {
		fmt.Printf("Output parameter %d: %s\n", i, funcType.Out(i))
	}
*/
