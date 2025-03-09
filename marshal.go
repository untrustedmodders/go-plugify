package plugify

/*
#include "plugify.h"
#cgo noescape aligned_malloc
#cgo noescape aligned_free

#cgo nocallback aligned_malloc
#cgo nocallback aligned_free
*/
import "C"
import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"unsafe"
)

var callbacks []C.JitCallback
var calls []C.JitCall

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

type memoryPool struct {
	pool      unsafe.Pointer // Start of the aligned memory block
	size      int            // Total size of the pool
	nextFree  int            // Next free byte offset
	alignment int            // Alignment requirement (16 bytes)
}

func newMemoryPool(size int, alignment int) *memoryPool {
	// Allocate a block of memory with alignment
	pool := C.aligned_malloc(C.size_t(size), C.size_t(alignment))
	if pool == nil {
		panic("failed to allocate memory pool")
	}

	return &memoryPool{
		pool:      pool,
		size:      size,
		nextFree:  0,
		alignment: alignment,
	}
}

func (p *memoryPool) alloc(size int) unsafe.Pointer {
	// Calculate the next aligned address
	alignedNextFree := (p.nextFree + p.alignment - 1) &^ (p.alignment - 1)

	// Check if there is enough space in the pool
	if alignedNextFree+size > p.size {
		panic("memory pool exhausted")
	}

	// Allocate the memory
	ptr := unsafe.Pointer(uintptr(p.pool) + uintptr(alignedNextFree))
	p.nextFree = alignedNextFree + size

	return ptr
}

func (p *memoryPool) reset() {
	p.nextFree = 0
}

func (p *memoryPool) free() {
	C.aligned_free(p.pool)
	p.pool = nil
}

func sizeOfValueType(vt ValueType) int {
	switch vt {
	case Void:
		return 0
	case Bool, Char8, Int8, UInt8:
		return 1
	case Char16, Int16, UInt16:
		return 2
	case Int32, UInt32, Float:
		return 4
	case Int64, UInt64, Double, Function, Pointer:
		return 8
	case String:
		return 24
	case Any:
		return 32
	case ArrayBool, ArrayChar8, ArrayChar16, ArrayInt8, ArrayInt16, ArrayInt32,
		ArrayInt64, ArrayUInt8, ArrayUInt16, ArrayUInt32, ArrayUInt64, ArrayPointer,
		ArrayFloat, ArrayDouble, ArrayString, ArrayAny, ArrayVector2, ArrayVector3,
		ArrayVector4, ArrayMatrix4x4:
		return 24
	case Vector2Type:
		return 8 // 2 floats
	case Vector3Type:
		return 12 // 3 floats
	case Vector4Type:
		return 16 // 4 floats
	case Matrix4x4Type:
		return 64 // 16 floats
	default:
		return 0
	}
}

func roundUp(size int) int {
	return (size + 15) &^ 15
}

func getValueTypeSize(vt ValueType) int {
	return roundUp(sizeOfValueType(vt))
}

type ManagedType struct {
	valueType ValueType
	ref       bool
}

const (
	ExpectedSize = 2
	_            = uint(unsafe.Sizeof(ManagedType{})) - ExpectedSize
)

var valueTypeToReflect = map[ValueType]reflect.Type{
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

func convertToReflectType(m C.MethodHandle, i int) reflect.Type {
	mt := C.Plugify_GetMethodParamType(m, C.ptrdiff_t(i))

	if mt.ref {
		return reflect.TypeOf((*interface{})(nil)).Elem()
	}

	vt := ValueType(mt.valueType)

	if val, ok := valueTypeToReflect[vt]; ok {
		return val
	}

	if vt == Function {
		return createFunctionType(C.Plugify_GetMethodPrototype(m, C.ptrdiff_t(i)))
	}

	return reflect.TypeOf((*interface{})(nil)).Elem()
}

var reflectToValueType = map[reflect.Type]ValueType{
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

func createManagedType(t reflect.Type) ManagedType {
	baseType := t

	if baseType.Kind() == reflect.Func {
		return ManagedType{Function, false}
	}

	ref := t.Kind() == reflect.Ptr
	if ref {
		baseType = t.Elem()
	}

	if val, ok := reflectToValueType[baseType]; ok {
		return ManagedType{val, ref}
	}

	return ManagedType{Invalid, false}
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

type function struct {
	fn   any
	addr unsafe.Pointer
}

var functionMap = make(map[unsafe.Pointer]function)

func cast[T any](val T) unsafe.Pointer {
	return *(*unsafe.Pointer)(unsafe.Pointer(&val))
}

func pin[T any](val T, pool *memoryPool, size int) unsafe.Pointer {
	tmp := (*T)(pool.alloc(size))
	*tmp = val
	return unsafe.Pointer(tmp)
}

func getParameterTypes(fnType reflect.Type) []ManagedType {
	parameterTypes := make([]ManagedType, fnType.NumIn())
	for i := 0; i < fnType.NumIn(); i++ {
		parameterTypes[i] = createManagedType(fnType.In(i))
	}
	return parameterTypes
}

func getReturnType(fnType reflect.Type) (ManagedType, int) {
	if fnType.NumOut() > 0 {
		return createManagedType(fnType.Out(0)), 1
	}
	return ManagedType{Void, false}, 0
}

func hasReturnType(returnType ManagedType) bool {
	hasRet := returnType.valueType >= _ObjectStart && returnType.valueType <= _ObjectEnd // params which pass by refs by default
	if !hasRet {
		var firstHidden ValueType
		if runtime.GOOS == "windows" {
			firstHidden = Vector3Type
		} else {
			firstHidden = Matrix4x4Type
		}
		hasRet = returnType.valueType >= firstHidden && returnType.valueType <= _StructEnd
	}
	return hasRet
}

func calculatePoolSize(parameterTypes []ManagedType, hasRet bool, returnType ManagedType) (int, int, int) {
	paramCount := len(parameterTypes)
	if hasRet {
		paramCount += 1
	}

	paramSize := roundUp(paramCount * sizeOfValueType(Pointer))
	poolSize := paramSize

	for _, t := range parameterTypes {
		if t.ref || t.valueType >= _ObjectStart && t.valueType <= _ObjectEnd {
			poolSize += getValueTypeSize(t.valueType)
		}
	}

	if hasRet {
		poolSize += getValueTypeSize(returnType.valueType)
	}

	return roundUp(poolSize), paramSize, paramCount
}

func GetDelegateForFunctionPointer(fnPtr unsafe.Pointer, fnType reflect.Type) any {
	if fnPtr == nil {
		return nil
	}

	el, ok := functionMap[fnPtr]
	if ok {
		return el.fn
	}

	if fnType.Kind() != reflect.Func {
		panicker("expected a function")
	}

	parameterTypes := getParameterTypes(fnType)
	returnType, retCount := getReturnType(fnType)

	hasRet := hasReturnType(returnType)

	poolSize, paramSize, paramCount := calculatePoolSize(parameterTypes, hasRet, returnType)

	call := C.Plugify_NewCall(fnPtr, (*C.ManagedType)(unsafe.Pointer(unsafe.SliceData(parameterTypes))), C.ptrdiff_t(fnType.NumIn()), *(*C.ManagedType)(unsafe.Pointer(&returnType)))
	if call == nil {
		panicker(fmt.Sprintf("%s (jit error: not found)", fnType.Name()))
	}

	calls = append(calls, call)

	addr := C.Plugify_GetCallFunction(call)
	if addr == nil {
		panicker(fmt.Sprintf("%s (jit error: %s)", fnType.Name(), string(C.GoString(C.Plugify_GetCallError(call)))))
	}

	wrapper := reflect.MakeFunc(fnType, func(args []reflect.Value) []reflect.Value {
		pool := newMemoryPool(poolSize, 16)
		defer pool.free()

		params := unsafe.Slice((*unsafe.Pointer)(pool.alloc(paramSize)), paramCount)

		results := make([]reflect.Value, retCount)

		index := 0

		retType := returnType.valueType
		if hasRet {
			var ptr unsafe.Pointer
			switch retType {
			case Vector2Type, Vector3Type, Vector4Type:
				ptr = pool.alloc(16)
			case Matrix4x4Type:
				ptr = pool.alloc(64)
			case String:
				ptr = pool.alloc(32)
				defer DestroyString((*PlgString)(ptr))
			case Any:
				ptr = pool.alloc(32)
				defer DestroyVariant((*PlgVariant)(ptr))
			case ArrayBool:
				ptr = pool.alloc(32)
				defer DestroyVectorBool((*PlgVector)(ptr))
			case ArrayChar8:
				ptr = pool.alloc(32)
				defer DestroyVectorChar8((*PlgVector)(ptr))
			case ArrayChar16:
				ptr = pool.alloc(32)
				defer DestroyVectorChar16((*PlgVector)(ptr))
			case ArrayInt8:
				ptr = pool.alloc(32)
				defer DestroyVectorInt8((*PlgVector)(ptr))
			case ArrayInt16:
				ptr = pool.alloc(32)
				defer DestroyVectorInt16((*PlgVector)(ptr))
			case ArrayInt32:
				ptr = pool.alloc(32)
				defer DestroyVectorInt32((*PlgVector)(ptr))
			case ArrayInt64:
				ptr = pool.alloc(32)
				defer DestroyVectorInt64((*PlgVector)(ptr))
			case ArrayUInt8:
				ptr = pool.alloc(32)
				defer DestroyVectorUInt8((*PlgVector)(ptr))
			case ArrayUInt16:
				ptr = pool.alloc(32)
				defer DestroyVectorUInt16((*PlgVector)(ptr))
			case ArrayUInt32:
				ptr = pool.alloc(32)
				defer DestroyVectorUInt32((*PlgVector)(ptr))
			case ArrayUInt64:
				ptr = pool.alloc(32)
				defer DestroyVectorUInt64((*PlgVector)(ptr))
			case ArrayPointer:
				ptr = pool.alloc(32)
				defer DestroyVectorPointer((*PlgVector)(ptr))
			case ArrayFloat:
				ptr = pool.alloc(32)
				defer DestroyVectorFloat((*PlgVector)(ptr))
			case ArrayDouble:
				ptr = pool.alloc(32)
				defer DestroyVectorDouble((*PlgVector)(ptr))
			case ArrayString:
				ptr = pool.alloc(32)
				defer DestroyVectorString((*PlgVector)(ptr))
			case ArrayAny:
				ptr = pool.alloc(32)
				defer DestroyVectorVariant((*PlgVector)(ptr))
			case ArrayVector2:
				ptr = pool.alloc(32)
				defer DestroyVectorVector2((*PlgVector)(ptr))
			case ArrayVector3:
				ptr = pool.alloc(32)
				defer DestroyVectorVector3((*PlgVector)(ptr))
			case ArrayVector4:
				ptr = pool.alloc(32)
				defer DestroyVectorVector4((*PlgVector)(ptr))
			case ArrayMatrix4x4:
				ptr = pool.alloc(32)
				defer DestroyVectorMatrix4x4((*PlgVector)(ptr))
			default:
				panicker(fmt.Sprintf("GetDelegateForFunctionPointer return type not supported %v", retType))
			}
			params[index] = ptr
			index++
		}

		for i, arg := range args {
			paramType := parameterTypes[i]
			valueType := paramType.valueType
			var ptr unsafe.Pointer
			if paramType.ref {
				switch valueType {
				case Bool:
					val := arg.Interface().(*bool)
					ptr = pin(*val, pool, 16)
					defer func() {
						*val = *(*bool)(ptr)
					}()
				case Char8:
					val := arg.Interface().(*int8)
					ptr = pin(*val, pool, 16)
					defer func() {
						*val = *(*int8)(ptr)
					}()
				case Char16:
					val := arg.Interface().(*uint16)
					ptr = pin(*val, pool, 16)
					defer func() {
						*val = *(*uint16)(ptr)
					}()
				case Int8:
					val := arg.Interface().(*int8)
					ptr = pin(*val, pool, 16)
					defer func() {
						*val = *(*int8)(ptr)
					}()
				case Int16:
					val := arg.Interface().(*int16)
					ptr = pin(*val, pool, 16)
					defer func() {
						*val = *(*int16)(ptr)
					}()
				case Int32:
					val := arg.Interface().(*int32)
					ptr = pin(*val, pool, 16)
					defer func() {
						*val = *(*int32)(ptr)
					}()
				case Int64:
					val := arg.Interface().(*int64)
					ptr = pin(*val, pool, 16)
					defer func() {
						*val = *(*int64)(ptr)
					}()
				case UInt8:
					val := arg.Interface().(*uint8)
					ptr = pin(*val, pool, 16)
					defer func() {
						*val = *(*uint8)(ptr)
					}()
				case UInt16:
					val := arg.Interface().(*uint16)
					ptr = pin(*val, pool, 16)
					defer func() {
						*val = *(*uint16)(ptr)
					}()
				case UInt32:
					val := arg.Interface().(*uint32)
					ptr = pin(*val, pool, 16)
					defer func() {
						*val = *(*uint32)(ptr)
					}()
				case UInt64:
					val := arg.Interface().(*uint64)
					ptr = pin(*val, pool, 16)
					defer func() {
						*val = *(*uint64)(ptr)
					}()
				case Pointer:
					val := arg.Interface().(*uintptr)
					ptr = pin(*val, pool, 16)
					defer func() {
						*val = *(*uintptr)(ptr)
					}()
				case Float:
					val := arg.Interface().(*float32)
					ptr = pin(*val, pool, 16)
					defer func() {
						*val = *(*float32)(ptr)
					}()
				case Double:
					val := arg.Interface().(*float64)
					ptr = pin(*val, pool, 16)
					defer func() {
						*val = *(*float64)(ptr)
					}()
				case Vector2Type:
					ptr = pin(*(arg.Interface().(*Vector2)), pool, 16)
					defer func() {
						*(arg.Interface().(*Vector2)) = *(*Vector2)(ptr)
					}()
				case Vector3Type:
					val := arg.Interface().(*Vector3)
					ptr = pin(*val, pool, 16)
					defer func() {
						*val = *(*Vector3)(ptr)
					}()
				case Vector4Type:
					val := arg.Interface().(*Vector4)
					ptr = pin(*val, pool, 16)
					defer func() {
						*val = *(*Vector4)(ptr)
					}()
				case Matrix4x4Type:
					val := arg.Interface().(*Matrix4x4)
					ptr = pin(*val, pool, 64)
					defer func() {
						*val = *(*Matrix4x4)(ptr)
					}()
				case Function:
					ptr = GetFunctionPointerForDelegate(arg)
				case String:
					val := arg.Interface().(*string)
					ptr = pin(ConstructString(*val), pool, 32)
					defer func() {
						*val = GetStringData((*PlgString)(ptr))
						DestroyString((*PlgString)(ptr))
					}()
				case Any:
					val := arg.Interface().(*any)
					ptr = pin(ConstructVariant(*val), pool, 32)
					defer func() {
						*val = GetVariantData((*PlgVariant)(ptr))
						DestroyVariant((*PlgVariant)(ptr))
					}()
				case ArrayBool:
					val := arg.Interface().(*[]bool)
					ptr = pin(ConstructVectorBool(*val), pool, 32)
					defer func() {
						*val = GetVectorDataBool((*PlgVector)(ptr))
						DestroyVectorBool((*PlgVector)(ptr))
					}()
				case ArrayChar8:
					val := arg.Interface().(*[]int8)
					ptr = pin(ConstructVectorChar8(*val), pool, 32)
					defer func() {
						*val = GetVectorDataChar8((*PlgVector)(ptr))
						DestroyVectorChar8((*PlgVector)(ptr))
					}()
				case ArrayChar16:
					val := arg.Interface().(*[]uint16)
					ptr = pin(ConstructVectorChar16(*val), pool, 32)
					defer func() {
						*val = GetVectorDataChar16((*PlgVector)(ptr))
						DestroyVectorChar16((*PlgVector)(ptr))
					}()
				case ArrayInt8:
					val := arg.Interface().(*[]int8)
					ptr = pin(ConstructVectorInt8(*val), pool, 32)
					defer func() {
						*val = GetVectorDataInt8((*PlgVector)(ptr))
						DestroyVectorInt8((*PlgVector)(ptr))
					}()
				case ArrayInt16:
					val := arg.Interface().(*[]int16)
					ptr = pin(ConstructVectorInt16(*val), pool, 32)
					defer func() {
						*val = GetVectorDataInt16((*PlgVector)(ptr))
						DestroyVectorInt16((*PlgVector)(ptr))
					}()
				case ArrayInt32:
					val := arg.Interface().(*[]int32)
					ptr = pin(ConstructVectorInt32(*val), pool, 32)
					defer func() {
						*val = GetVectorDataInt32((*PlgVector)(ptr))
						DestroyVectorInt32((*PlgVector)(ptr))
					}()
				case ArrayInt64:
					val := arg.Interface().(*[]int64)
					ptr = pin(ConstructVectorInt64(*val), pool, 32)
					defer func() {
						*val = GetVectorDataInt64((*PlgVector)(ptr))
						DestroyVectorInt64((*PlgVector)(ptr))
					}()
				case ArrayUInt8:
					val := arg.Interface().(*[]uint8)
					ptr = pin(ConstructVectorUInt8(*val), pool, 32)
					defer func() {
						*val = GetVectorDataUInt8((*PlgVector)(ptr))
						DestroyVectorUInt8((*PlgVector)(ptr))
					}()
				case ArrayUInt16:
					val := arg.Interface().(*[]uint16)
					ptr = pin(ConstructVectorUInt16(*val), pool, 32)
					defer func() {
						*val = GetVectorDataUInt16((*PlgVector)(ptr))
						DestroyVectorUInt16((*PlgVector)(ptr))
					}()
				case ArrayUInt32:
					val := arg.Interface().(*[]uint32)
					ptr = pin(ConstructVectorUInt32(*val), pool, 32)
					defer func() {
						*val = GetVectorDataUInt32((*PlgVector)(ptr))
						DestroyVectorUInt32((*PlgVector)(ptr))
					}()
				case ArrayUInt64:
					val := arg.Interface().(*[]uint64)
					ptr = pin(ConstructVectorUInt64(*val), pool, 32)
					defer func() {
						*val = GetVectorDataUInt64((*PlgVector)(ptr))
						DestroyVectorUInt64((*PlgVector)(ptr))
					}()
				case ArrayPointer:
					val := arg.Interface().(*[]uintptr)
					ptr = pin(ConstructVectorPointer(*val), pool, 32)
					defer func() {
						*val = GetVectorDataPointer((*PlgVector)(ptr))
						DestroyVectorPointer((*PlgVector)(ptr))
					}()
				case ArrayFloat:
					val := arg.Interface().(*[]float32)
					ptr = pin(ConstructVectorFloat(*val), pool, 32)
					defer func() {
						*val = GetVectorDataFloat((*PlgVector)(ptr))
						DestroyVectorFloat((*PlgVector)(ptr))
					}()
				case ArrayDouble:
					val := arg.Interface().(*[]float64)
					ptr = pin(ConstructVectorDouble(*val), pool, 32)
					defer func() {
						*val = GetVectorDataDouble((*PlgVector)(ptr))
						DestroyVectorDouble((*PlgVector)(ptr))
					}()
				case ArrayString:
					val := arg.Interface().(*[]string)
					ptr = pin(ConstructVectorString(*val), pool, 32)
					defer func() {
						*val = GetVectorDataString((*PlgVector)(ptr))
						DestroyVectorString((*PlgVector)(ptr))
					}()
				case ArrayAny:
					val := arg.Interface().(*[]any)
					ptr = pin(ConstructVectorVariant(*val), pool, 32)
					defer func() {
						*val = GetVectorDataVariant((*PlgVector)(ptr))
						DestroyVectorVariant((*PlgVector)(ptr))
					}()
				case ArrayVector2:
					val := arg.Interface().(*[]Vector2)
					ptr = pin(ConstructVectorVector2(*val), pool, 32)
					defer func() {
						*val = GetVectorDataVector2((*PlgVector)(ptr))
						DestroyVectorVector2((*PlgVector)(ptr))
					}()
				case ArrayVector3:
					val := arg.Interface().(*[]Vector3)
					ptr = pin(ConstructVectorVector3(*val), pool, 32)
					defer func() {
						*val = GetVectorDataVector3((*PlgVector)(ptr))
						DestroyVectorVector3((*PlgVector)(ptr))
					}()
				case ArrayVector4:
					val := arg.Interface().(*[]Vector4)
					ptr = pin(ConstructVectorVector4(*val), pool, 32)
					defer func() {
						*val = GetVectorDataVector4((*PlgVector)(ptr))
						DestroyVectorVector4((*PlgVector)(ptr))
					}()
				case ArrayMatrix4x4:
					val := arg.Interface().(*[]Matrix4x4)
					ptr = pin(ConstructVectorMatrix4x4(*val), pool, 32)
					defer func() {
						*val = GetVectorDataMatrix4x4((*PlgVector)(ptr))
						DestroyVectorMatrix4x4((*PlgVector)(ptr))
					}()
				default:
					panicker(fmt.Sprintf("GetDelegateForFunctionPointer parameter type not supported %v", retType))
				}
			} else {
				switch valueType {
				case Bool:
					ptr = cast(arg.Interface().(bool))
				case Char8:
					ptr = cast(arg.Interface().(int8))
				case Char16:
					ptr = cast(arg.Interface().(uint16))
				case Int8:
					ptr = cast(arg.Interface().(int8))
				case Int16:
					ptr = cast(arg.Interface().(int16))
				case Int32:
					ptr = cast(arg.Interface().(int32))
				case Int64:
					ptr = cast(arg.Interface().(int64))
				case UInt8:
					ptr = cast(arg.Interface().(uint8))
				case UInt16:
					ptr = cast(arg.Interface().(uint16))
				case UInt32:
					ptr = cast(arg.Interface().(uint32))
				case UInt64:
					ptr = cast(arg.Interface().(uint64))
				case Pointer:
					ptr = cast(arg.Interface().(uintptr))
				case Float:
					ptr = cast(arg.Interface().(float32))
				case Double:
					ptr = cast(arg.Interface().(float64))
				case Vector2Type:
					ptr = pin(arg.Interface().(Vector2), pool, 16)
				case Vector3Type:
					ptr = pin(arg.Interface().(Vector3), pool, 16)
				case Vector4Type:
					ptr = pin(arg.Interface().(Vector4), pool, 16)
				case Matrix4x4Type:
					ptr = pin(arg.Interface().(Matrix4x4), pool, 64)
				case Function:
					ptr = GetFunctionPointerForDelegate(arg)
				case String:
					ptr = pin(ConstructString(arg.Interface().(string)), pool, 32)
					defer DestroyString((*PlgString)(ptr))
				case Any:
					ptr = pin(ConstructVariant(arg.Interface().(any)), pool, 32)
					defer DestroyVariant((*PlgVariant)(ptr))
				case ArrayBool:
					ptr = pin(ConstructVectorBool(arg.Interface().([]bool)), pool, 32)
					defer DestroyVectorBool((*PlgVector)(ptr))
				case ArrayChar8:
					ptr = pin(ConstructVectorChar8(arg.Interface().([]int8)), pool, 32)
					defer DestroyVectorChar8((*PlgVector)(ptr))
				case ArrayChar16:
					ptr = pin(ConstructVectorChar16(arg.Interface().([]uint16)), pool, 32)
					defer DestroyVectorChar16((*PlgVector)(ptr))
				case ArrayInt8:
					ptr = pin(ConstructVectorInt8(arg.Interface().([]int8)), pool, 32)
					defer DestroyVectorInt8((*PlgVector)(ptr))
				case ArrayInt16:
					ptr = pin(ConstructVectorInt16(arg.Interface().([]int16)), pool, 32)
					defer DestroyVectorInt16((*PlgVector)(ptr))
				case ArrayInt32:
					ptr = pin(ConstructVectorInt32(arg.Interface().([]int32)), pool, 32)
					defer DestroyVectorInt32((*PlgVector)(ptr))
				case ArrayInt64:
					ptr = pin(ConstructVectorInt64(arg.Interface().([]int64)), pool, 32)
					defer DestroyVectorInt64((*PlgVector)(ptr))
				case ArrayUInt8:
					ptr = pin(ConstructVectorUInt8(arg.Interface().([]uint8)), pool, 32)
					defer DestroyVectorUInt8((*PlgVector)(ptr))
				case ArrayUInt16:
					ptr = pin(ConstructVectorUInt16(arg.Interface().([]uint16)), pool, 32)
					defer DestroyVectorUInt16((*PlgVector)(ptr))
				case ArrayUInt32:
					ptr = pin(ConstructVectorUInt32(arg.Interface().([]uint32)), pool, 32)
					defer DestroyVectorUInt32((*PlgVector)(ptr))
				case ArrayUInt64:
					ptr = pin(ConstructVectorUInt64(arg.Interface().([]uint64)), pool, 32)
					defer DestroyVectorUInt64((*PlgVector)(ptr))
				case ArrayPointer:
					ptr = pin(ConstructVectorPointer(arg.Interface().([]uintptr)), pool, 32)
					defer DestroyVectorPointer((*PlgVector)(ptr))
				case ArrayFloat:
					ptr = pin(ConstructVectorFloat(arg.Interface().([]float32)), pool, 32)
					defer DestroyVectorFloat((*PlgVector)(ptr))
				case ArrayDouble:
					ptr = pin(ConstructVectorDouble(arg.Interface().([]float64)), pool, 32)
					defer DestroyVectorDouble((*PlgVector)(ptr))
				case ArrayString:
					ptr = pin(ConstructVectorString(arg.Interface().([]string)), pool, 32)
					defer DestroyVectorString((*PlgVector)(ptr))
				case ArrayAny:
					ptr = pin(ConstructVectorVariant(arg.Interface().([]any)), pool, 32)
					defer DestroyVectorVariant((*PlgVector)(ptr))
				case ArrayVector2:
					ptr = pin(ConstructVectorVector2(arg.Interface().([]Vector2)), pool, 32)
					defer DestroyVectorVector2((*PlgVector)(ptr))
				case ArrayVector3:
					ptr = pin(ConstructVectorVector3(arg.Interface().([]Vector3)), pool, 32)
					defer DestroyVectorVector3((*PlgVector)(ptr))
				case ArrayVector4:
					ptr = pin(ConstructVectorVector4(arg.Interface().([]Vector4)), pool, 32)
					defer DestroyVectorVector4((*PlgVector)(ptr))
				case ArrayMatrix4x4:
					ptr = pin(ConstructVectorMatrix4x4(arg.Interface().([]Matrix4x4)), pool, 32)
					defer DestroyVectorMatrix4x4((*PlgVector)(ptr))
				default:
					panicker(fmt.Sprintf("GetDelegateForFunctionPointer parameter type not supported %v", valueType))
				}
			}
			params[index] = ptr
			index++
		}

		var retStore C.__int128
		C.Plugify_CallFunction(call, unsafe.SliceData(params), &retStore)
		ret := unsafe.Pointer(&retStore)

		switch retType {
		case Void:
			// skip
		case Bool:
			results[0] = reflect.ValueOf(*(*bool)(ret))
		case Char8:
			results[0] = reflect.ValueOf(*(*int8)(ret))
		case Char16:
			results[0] = reflect.ValueOf(*(*uint16)(ret))
		case Int8:
			results[0] = reflect.ValueOf(*(*int8)(ret))
		case Int16:
			results[0] = reflect.ValueOf(*(*int16)(ret))
		case Int32:
			results[0] = reflect.ValueOf(*(*int32)(ret))
		case Int64:
			results[0] = reflect.ValueOf(*(*int64)(ret))
		case UInt8:
			results[0] = reflect.ValueOf(*(*uint8)(ret))
		case UInt16:
			results[0] = reflect.ValueOf(*(*uint16)(ret))
		case UInt32:
			results[0] = reflect.ValueOf(*(*uint32)(ret))
		case UInt64:
			results[0] = reflect.ValueOf(*(*uint64)(ret))
		case Pointer:
			results[0] = reflect.ValueOf(*(*uintptr)(ret))
		case Float:
			results[0] = reflect.ValueOf(*(*float32)(ret))
		case Double:
			results[0] = reflect.ValueOf(*(*float64)(ret))
		case Function:
			results[0] = reflect.ValueOf(GetDelegateForFunctionPointer(*(*unsafe.Pointer)(ret), fnType.Out(0)))
		case Vector2Type:
			results[0] = reflect.ValueOf(*(*Vector2)(ret))
		case Vector3Type:
			if hasRet {
				results[0] = reflect.ValueOf(*(*Vector3)(*(*unsafe.Pointer)(ret)))
			} else {
				results[0] = reflect.ValueOf(*(*Vector3)(ret))
			}
		case Vector4Type:
			if hasRet {
				results[0] = reflect.ValueOf(*(*Vector4)(*(*unsafe.Pointer)(ret)))
			} else {
				results[0] = reflect.ValueOf(*(*Vector4)(ret))
			}
		case Matrix4x4Type:
			results[0] = reflect.ValueOf(*(*Matrix4x4)(*(*unsafe.Pointer)(ret)))
		case String:
			results[0] = reflect.ValueOf(GetStringData((*PlgString)(*(*unsafe.Pointer)(ret))))
		case Any:
			results[0] = reflect.ValueOf(GetVariantData((*PlgVariant)(*(*unsafe.Pointer)(ret))))
		case ArrayBool:
			results[0] = reflect.ValueOf(GetVectorDataBool((*PlgVector)(*(*unsafe.Pointer)(ret))))
		case ArrayChar8:
			results[0] = reflect.ValueOf(GetVectorDataChar8((*PlgVector)(*(*unsafe.Pointer)(ret))))
		case ArrayChar16:
			results[0] = reflect.ValueOf(GetVectorDataChar16((*PlgVector)(*(*unsafe.Pointer)(ret))))
		case ArrayInt8:
			results[0] = reflect.ValueOf(GetVectorDataInt8((*PlgVector)(*(*unsafe.Pointer)(ret))))
		case ArrayInt16:
			results[0] = reflect.ValueOf(GetVectorDataInt16((*PlgVector)(*(*unsafe.Pointer)(ret))))
		case ArrayInt32:
			results[0] = reflect.ValueOf(GetVectorDataInt32((*PlgVector)(*(*unsafe.Pointer)(ret))))
		case ArrayInt64:
			results[0] = reflect.ValueOf(GetVectorDataInt64((*PlgVector)(*(*unsafe.Pointer)(ret))))
		case ArrayUInt8:
			results[0] = reflect.ValueOf(GetVectorDataUInt8((*PlgVector)(*(*unsafe.Pointer)(ret))))
		case ArrayUInt16:
			results[0] = reflect.ValueOf(GetVectorDataUInt16((*PlgVector)(*(*unsafe.Pointer)(ret))))
		case ArrayUInt32:
			results[0] = reflect.ValueOf(GetVectorDataUInt32((*PlgVector)(*(*unsafe.Pointer)(ret))))
		case ArrayUInt64:
			results[0] = reflect.ValueOf(GetVectorDataUInt64((*PlgVector)(*(*unsafe.Pointer)(ret))))
		case ArrayPointer:
			results[0] = reflect.ValueOf(GetVectorDataPointer((*PlgVector)(*(*unsafe.Pointer)(ret))))
		case ArrayFloat:
			results[0] = reflect.ValueOf(GetVectorDataFloat((*PlgVector)(*(*unsafe.Pointer)(ret))))
		case ArrayDouble:
			results[0] = reflect.ValueOf(GetVectorDataDouble((*PlgVector)(*(*unsafe.Pointer)(ret))))
		case ArrayString:
			results[0] = reflect.ValueOf(GetVectorDataString((*PlgVector)(*(*unsafe.Pointer)(ret))))
		case ArrayAny:
			results[0] = reflect.ValueOf(GetVectorDataVariant((*PlgVector)(*(*unsafe.Pointer)(ret))))
		case ArrayVector2:
			results[0] = reflect.ValueOf(GetVectorDataVector2((*PlgVector)(*(*unsafe.Pointer)(ret))))
		case ArrayVector3:
			results[0] = reflect.ValueOf(GetVectorDataVector3((*PlgVector)(*(*unsafe.Pointer)(ret))))
		case ArrayVector4:
			results[0] = reflect.ValueOf(GetVectorDataVector4((*PlgVector)(*(*unsafe.Pointer)(ret))))
		case ArrayMatrix4x4:
			results[0] = reflect.ValueOf(GetVectorDataMatrix4x4((*PlgVector)(*(*unsafe.Pointer)(ret))))
		default:
			panicker(fmt.Sprintf("GetDelegateForFunctionPointer return type not supported %v", retType))
		}

		return results
	})

	fn := wrapper.Interface()

	functionMap[fnPtr] = function{fn, addr}

	return fn
}

func setArgument[T any](p *C.Parameters, idx C.size_t, val T) {
	ptr := getArgumentPtr(p, idx)
	*(*T)(ptr) = val
}

func getArgument[T any](p *C.Parameters, idx C.size_t) T {
	return *(*T)(getArgumentPtr(p, idx))
}

func getArgumentPtr(p *C.Parameters, idx C.size_t) unsafe.Pointer {
	return unsafe.Pointer(uintptr(unsafe.Pointer(&p.arguments)) + uintptr(idx)*unsafe.Sizeof(uint64(0)))
}

func setReturn[T any](r *C.Return, val T) {
	ptr := getReturnPtr(r)
	*(*T)(ptr) = val
}

func getReturn[T any](r *C.Return) T {
	return *(*T)(getReturnPtr(r))
}

func getReturnPtr(r *C.Return) unsafe.Pointer {
	return unsafe.Pointer(&r.ret)
}

func heap[T any](v T) *T {
	return &v
}

func paramRefToObject(vt ValueType, p *C.Parameters, i C.size_t) reflect.Value {
	switch vt {
	case Bool:
		return reflect.ValueOf(getArgument[*bool](p, i))
	case Char8:
		return reflect.ValueOf(getArgument[*int8](p, i))
	case Char16:
		return reflect.ValueOf(getArgument[*uint16](p, i))
	case Int8:
		return reflect.ValueOf(getArgument[*int8](p, i))
	case Int16:
		return reflect.ValueOf(getArgument[*int16](p, i))
	case Int32:
		return reflect.ValueOf(getArgument[*int32](p, i))
	case Int64:
		return reflect.ValueOf(getArgument[*int64](p, i))
	case UInt8:
		return reflect.ValueOf(getArgument[*uint8](p, i))
	case UInt16:
		return reflect.ValueOf(getArgument[*uint16](p, i))
	case UInt32:
		return reflect.ValueOf(getArgument[*uint32](p, i))
	case UInt64:
		return reflect.ValueOf(getArgument[*uint64](p, i))
	case Pointer:
		return reflect.ValueOf(getArgument[*uintptr](p, i))
	case Float:
		return reflect.ValueOf(getArgument[*float32](p, i))
	case Double:
		return reflect.ValueOf(getArgument[*float64](p, i))
	case String:
		return reflect.ValueOf(heap[string](GetStringData(getArgument[*PlgString](p, i))))
	case Any:
		return reflect.ValueOf(heap[any](GetVariantData(getArgument[*PlgVariant](p, i))))
	case ArrayBool:
		return reflect.ValueOf(heap[[]bool](GetVectorDataBool(getArgument[*PlgVector](p, i))))
	case ArrayChar8:
		return reflect.ValueOf(heap[[]int8](GetVectorDataChar8(getArgument[*PlgVector](p, i))))
	case ArrayChar16:
		return reflect.ValueOf(heap[[]uint16](GetVectorDataChar16(getArgument[*PlgVector](p, i))))
	case ArrayInt8:
		return reflect.ValueOf(heap[[]int8](GetVectorDataInt8(getArgument[*PlgVector](p, i))))
	case ArrayInt16:
		return reflect.ValueOf(heap[[]int16](GetVectorDataInt16(getArgument[*PlgVector](p, i))))
	case ArrayInt32:
		return reflect.ValueOf(heap[[]int32](GetVectorDataInt32(getArgument[*PlgVector](p, i))))
	case ArrayInt64:
		return reflect.ValueOf(heap[[]int64](GetVectorDataInt64(getArgument[*PlgVector](p, i))))
	case ArrayUInt8:
		return reflect.ValueOf(heap[[]uint8](GetVectorDataUInt8(getArgument[*PlgVector](p, i))))
	case ArrayUInt16:
		return reflect.ValueOf(heap[[]uint16](GetVectorDataUInt16(getArgument[*PlgVector](p, i))))
	case ArrayUInt32:
		return reflect.ValueOf(heap[[]uint32](GetVectorDataUInt32(getArgument[*PlgVector](p, i))))
	case ArrayUInt64:
		return reflect.ValueOf(heap[[]uint64](GetVectorDataUInt64(getArgument[*PlgVector](p, i))))
	case ArrayPointer:
		return reflect.ValueOf(heap[[]uintptr](GetVectorDataPointer(getArgument[*PlgVector](p, i))))
	case ArrayFloat:
		return reflect.ValueOf(heap[[]float32](GetVectorDataFloat(getArgument[*PlgVector](p, i))))
	case ArrayDouble:
		return reflect.ValueOf(heap[[]float64](GetVectorDataDouble(getArgument[*PlgVector](p, i))))
	case ArrayString:
		return reflect.ValueOf(heap[[]string](GetVectorDataString(getArgument[*PlgVector](p, i))))
	case ArrayAny:
		return reflect.ValueOf(heap[[]any](GetVectorDataVariant(getArgument[*PlgVector](p, i))))
	case ArrayVector2:
		return reflect.ValueOf(heap[[]Vector2](GetVectorDataVector2(getArgument[*PlgVector](p, i))))
	case ArrayVector3:
		return reflect.ValueOf(heap[[]Vector3](GetVectorDataVector3(getArgument[*PlgVector](p, i))))
	case ArrayVector4:
		return reflect.ValueOf(heap[[]Vector4](GetVectorDataVector4(getArgument[*PlgVector](p, i))))
	case ArrayMatrix4x4:
		return reflect.ValueOf(heap[[]Matrix4x4](GetVectorDataMatrix4x4(getArgument[*PlgVector](p, i))))
	case Vector2Type:
		return reflect.ValueOf(getArgument[*Vector2](p, i))
	case Vector3Type:
		return reflect.ValueOf(getArgument[*Vector3](p, i))
	case Vector4Type:
		return reflect.ValueOf(getArgument[*Vector4](p, i))
	case Matrix4x4Type:
		return reflect.ValueOf(getArgument[*Matrix4x4](p, i))
	default:
		panicker(fmt.Sprintf("paramRefToObject unsupported enum type %v", vt))
		return reflect.ValueOf(nil)
	}
}

func paramToObject(m C.MethodHandle, vt ValueType, p *C.Parameters, i C.size_t) reflect.Value {
	switch vt {
	case Bool:
		return reflect.ValueOf(getArgument[bool](p, i))
	case Char8:
		return reflect.ValueOf(getArgument[int8](p, i))
	case Char16:
		return reflect.ValueOf(getArgument[uint16](p, i))
	case Int8:
		return reflect.ValueOf(getArgument[int8](p, i))
	case Int16:
		return reflect.ValueOf(getArgument[int16](p, i))
	case Int32:
		return reflect.ValueOf(getArgument[int32](p, i))
	case Int64:
		return reflect.ValueOf(getArgument[int64](p, i))
	case UInt8:
		return reflect.ValueOf(getArgument[uint8](p, i))
	case UInt16:
		return reflect.ValueOf(getArgument[uint16](p, i))
	case UInt32:
		return reflect.ValueOf(getArgument[uint32](p, i))
	case UInt64:
		return reflect.ValueOf(getArgument[uint64](p, i))
	case Pointer:
		return reflect.ValueOf(getArgument[uintptr](p, i))
	case Float:
		return reflect.ValueOf(getArgument[float32](p, i))
	case Double:
		return reflect.ValueOf(getArgument[float64](p, i))
	case Function:
		return reflect.ValueOf(GetDelegateForFunctionPointer(getArgument[unsafe.Pointer](p, i), createFunctionType(m)))
	case String:
		return reflect.ValueOf(GetStringData(getArgument[*PlgString](p, i)))
	case Any:
		return reflect.ValueOf(GetVariantData(getArgument[*PlgVariant](p, i)))
	case ArrayBool:
		return reflect.ValueOf(GetVectorDataBool(getArgument[*PlgVector](p, i)))
	case ArrayChar8:
		return reflect.ValueOf(GetVectorDataChar8(getArgument[*PlgVector](p, i)))
	case ArrayChar16:
		return reflect.ValueOf(GetVectorDataChar16(getArgument[*PlgVector](p, i)))
	case ArrayInt8:
		return reflect.ValueOf(GetVectorDataInt8(getArgument[*PlgVector](p, i)))
	case ArrayInt16:
		return reflect.ValueOf(GetVectorDataInt16(getArgument[*PlgVector](p, i)))
	case ArrayInt32:
		return reflect.ValueOf(GetVectorDataInt32(getArgument[*PlgVector](p, i)))
	case ArrayInt64:
		return reflect.ValueOf(GetVectorDataInt64(getArgument[*PlgVector](p, i)))
	case ArrayUInt8:
		return reflect.ValueOf(GetVectorDataUInt8(getArgument[*PlgVector](p, i)))
	case ArrayUInt16:
		return reflect.ValueOf(GetVectorDataUInt16(getArgument[*PlgVector](p, i)))
	case ArrayUInt32:
		return reflect.ValueOf(GetVectorDataUInt32(getArgument[*PlgVector](p, i)))
	case ArrayUInt64:
		return reflect.ValueOf(GetVectorDataUInt64(getArgument[*PlgVector](p, i)))
	case ArrayPointer:
		return reflect.ValueOf(GetVectorDataPointer(getArgument[*PlgVector](p, i)))
	case ArrayFloat:
		return reflect.ValueOf(GetVectorDataFloat(getArgument[*PlgVector](p, i)))
	case ArrayDouble:
		return reflect.ValueOf(GetVectorDataDouble(getArgument[*PlgVector](p, i)))
	case ArrayString:
		return reflect.ValueOf(GetVectorDataString(getArgument[*PlgVector](p, i)))
	case ArrayAny:
		return reflect.ValueOf(GetVectorDataVariant(getArgument[*PlgVector](p, i)))
	case ArrayVector2:
		return reflect.ValueOf(GetVectorDataVector2(getArgument[*PlgVector](p, i)))
	case ArrayVector3:
		return reflect.ValueOf(GetVectorDataVector3(getArgument[*PlgVector](p, i)))
	case ArrayVector4:
		return reflect.ValueOf(GetVectorDataVector4(getArgument[*PlgVector](p, i)))
	case ArrayMatrix4x4:
		return reflect.ValueOf(GetVectorDataMatrix4x4(getArgument[*PlgVector](p, i)))
	case Vector2Type:
		return reflect.ValueOf(*getArgument[*Vector2](p, i))
	case Vector3Type:
		return reflect.ValueOf(*getArgument[*Vector3](p, i))
	case Vector4Type:
		return reflect.ValueOf(*getArgument[*Vector4](p, i))
	case Matrix4x4Type:
		return reflect.ValueOf(*getArgument[*Matrix4x4](p, i))
	default:
		panicker(fmt.Sprintf("paramToObject unsupported enum type %v", vt))
	}
	return reflect.ValueOf(nil)
}

func setRefParam(vt ValueType, p *C.Parameters, i C.size_t, val reflect.Value) {
	switch vt {
	case Bool:
		setArgument(p, i, val.Interface().(*bool))
	case Char8:
		setArgument(p, i, val.Interface().(*int8))
	case Char16:
		setArgument(p, i, val.Interface().(*uint16))
	case Int8:
		setArgument(p, i, val.Interface().(*int8))
	case Int16:
		setArgument(p, i, val.Interface().(*int16))
	case Int32:
		setArgument(p, i, val.Interface().(*int32))
	case Int64:
		setArgument(p, i, val.Interface().(*int64))
	case UInt8:
		setArgument(p, i, val.Interface().(*uint8))
	case UInt16:
		setArgument(p, i, val.Interface().(*uint16))
	case UInt32:
		setArgument(p, i, val.Interface().(*uint32))
	case UInt64:
		setArgument(p, i, val.Interface().(*uint64))
	case Pointer:
		setArgument(p, i, val.Interface().(*uintptr))
	case Float:
		setArgument(p, i, val.Interface().(*float32))
	case Double:
		setArgument(p, i, val.Interface().(*float64))
	case String:
		AssignString(getArgument[*PlgString](p, i), *val.Interface().(*string))
	case Any:
		AssignVariant(getArgument[*PlgVariant](p, i), *val.Interface().(*any))
	case ArrayBool:
		AssignVectorBool(getArgument[*PlgVector](p, i), *val.Interface().(*[]bool))
	case ArrayChar8:
		AssignVectorChar8(getArgument[*PlgVector](p, i), *val.Interface().(*[]int8))
	case ArrayChar16:
		AssignVectorChar16(getArgument[*PlgVector](p, i), *val.Interface().(*[]uint16))
	case ArrayInt8:
		AssignVectorInt8(getArgument[*PlgVector](p, i), *val.Interface().(*[]int8))
	case ArrayInt16:
		AssignVectorInt16(getArgument[*PlgVector](p, i), *val.Interface().(*[]int16))
	case ArrayInt32:
		AssignVectorInt32(getArgument[*PlgVector](p, i), *val.Interface().(*[]int32))
	case ArrayInt64:
		AssignVectorInt64(getArgument[*PlgVector](p, i), *val.Interface().(*[]int64))
	case ArrayUInt8:
		AssignVectorUInt8(getArgument[*PlgVector](p, i), *val.Interface().(*[]uint8))
	case ArrayUInt16:
		AssignVectorUInt16(getArgument[*PlgVector](p, i), *val.Interface().(*[]uint16))
	case ArrayUInt32:
		AssignVectorUInt32(getArgument[*PlgVector](p, i), *val.Interface().(*[]uint32))
	case ArrayUInt64:
		AssignVectorUInt64(getArgument[*PlgVector](p, i), *val.Interface().(*[]uint64))
	case ArrayPointer:
		AssignVectorPointer(getArgument[*PlgVector](p, i), *val.Interface().(*[]uintptr))
	case ArrayFloat:
		AssignVectorFloat(getArgument[*PlgVector](p, i), *val.Interface().(*[]float32))
	case ArrayDouble:
		AssignVectorDouble(getArgument[*PlgVector](p, i), *val.Interface().(*[]float64))
	case ArrayString:
		AssignVectorString(getArgument[*PlgVector](p, i), *val.Interface().(*[]string))
	case ArrayAny:
		AssignVectorVariant(getArgument[*PlgVector](p, i), *val.Interface().(*[]any))
	case ArrayVector2:
		AssignVectorVector2(getArgument[*PlgVector](p, i), *val.Interface().(*[]Vector2))
	case ArrayVector3:
		AssignVectorVector3(getArgument[*PlgVector](p, i), *val.Interface().(*[]Vector3))
	case ArrayVector4:
		AssignVectorVector4(getArgument[*PlgVector](p, i), *val.Interface().(*[]Vector4))
	case ArrayMatrix4x4:
		AssignVectorMatrix4x4(getArgument[*PlgVector](p, i), *val.Interface().(*[]Matrix4x4))
	case Vector2Type:
		setArgument(p, i, val.Interface().(*Vector2))
	case Vector3Type:
		setArgument(p, i, val.Interface().(*Vector3))
	case Vector4Type:
		setArgument(p, i, val.Interface().(*Vector4))
	case Matrix4x4Type:
		setArgument(p, i, val.Interface().(*Matrix4x4))
	default:
		panicker(fmt.Sprintf("setRefParam unsupported enum type %v", vt))
	}
}

func setObjReturn(vt ValueType, r *C.Return, rets []reflect.Value) {
	switch vt {
	case Void:
		// Do nothing
	case Bool:
		setReturn(r, rets[0].Interface().(bool))
	case Char8:
		setReturn(r, rets[0].Interface().(int8))
	case Char16:
		setReturn(r, rets[0].Interface().(uint16))
	case Int8:
		setReturn(r, rets[0].Interface().(int8))
	case Int16:
		setReturn(r, rets[0].Interface().(int16))
	case Int32:
		setReturn(r, rets[0].Interface().(int32))
	case Int64:
		setReturn(r, rets[0].Interface().(int64))
	case UInt8:
		setReturn(r, rets[0].Interface().(uint8))
	case UInt16:
		setReturn(r, rets[0].Interface().(uint16))
	case UInt32:
		setReturn(r, rets[0].Interface().(uint32))
	case UInt64:
		setReturn(r, rets[0].Interface().(uint64))
	case Pointer:
		setReturn(r, rets[0].Interface().(uintptr))
	case Float:
		setReturn(r, rets[0].Interface().(float32))
	case Double:
		setReturn(r, rets[0].Interface().(float64))
	case Function:
		setReturn(r, GetFunctionPointerForDelegate(rets[0].Interface()))
	case String:
		setReturn(r, ConstructString(rets[0].Interface().(string)))
	case Any:
		setReturn(r, ConstructVariant(rets[0].Interface().(any)))
	case ArrayBool:
		setReturn(r, ConstructVectorBool(rets[0].Interface().([]bool)))
	case ArrayChar8:
		setReturn(r, ConstructVectorChar8(rets[0].Interface().([]int8)))
	case ArrayChar16:
		setReturn(r, ConstructVectorChar16(rets[0].Interface().([]uint16)))
	case ArrayInt8:
		setReturn(r, ConstructVectorInt8(rets[0].Interface().([]int8)))
	case ArrayInt16:
		setReturn(r, ConstructVectorInt16(rets[0].Interface().([]int16)))
	case ArrayInt32:
		setReturn(r, ConstructVectorInt32(rets[0].Interface().([]int32)))
	case ArrayInt64:
		setReturn(r, ConstructVectorInt64(rets[0].Interface().([]int64)))
	case ArrayUInt8:
		setReturn(r, ConstructVectorUInt8(rets[0].Interface().([]uint8)))
	case ArrayUInt16:
		setReturn(r, ConstructVectorUInt16(rets[0].Interface().([]uint16)))
	case ArrayUInt32:
		setReturn(r, ConstructVectorUInt32(rets[0].Interface().([]uint32)))
	case ArrayUInt64:
		setReturn(r, ConstructVectorUInt64(rets[0].Interface().([]uint64)))
	case ArrayPointer:
		setReturn(r, ConstructVectorPointer(rets[0].Interface().([]uintptr)))
	case ArrayFloat:
		setReturn(r, ConstructVectorFloat(rets[0].Interface().([]float32)))
	case ArrayDouble:
		setReturn(r, ConstructVectorDouble(rets[0].Interface().([]float64)))
	case ArrayString:
		setReturn(r, ConstructVectorString(rets[0].Interface().([]string)))
	case ArrayAny:
		setReturn(r, ConstructVectorVariant(rets[0].Interface().([]any)))
	case ArrayVector2:
		setReturn(r, ConstructVectorVector2(rets[0].Interface().([]Vector2)))
	case ArrayVector3:
		setReturn(r, ConstructVectorVector3(rets[0].Interface().([]Vector3)))
	case ArrayVector4:
		setReturn(r, ConstructVectorVector4(rets[0].Interface().([]Vector4)))
	case ArrayMatrix4x4:
		setReturn(r, ConstructVectorMatrix4x4(rets[0].Interface().([]Matrix4x4)))
	case Vector2Type:
		setReturn(r, rets[0].Interface().(Vector2))
	case Vector3Type:
		setReturn(r, rets[0].Interface().(Vector3))
	case Vector4Type:
		setReturn(r, rets[0].Interface().(Vector4))
	case Matrix4x4Type:
		setReturn(r, rets[0].Interface().(Matrix4x4))
	default:
		panicker(fmt.Sprintf("setReturn unsupported enum type %v", vt))
	}
}

//export Plugify_InternalCall
func Plugify_InternalCall(m C.MethodHandle, data unsafe.Pointer, p *C.Parameters, count C.size_t, r *C.Return) {
	Block{
		Try: func() {
			fn, ok := functionMap[data]
			if !ok {
				panicker(fmt.Sprintf("function %p not found", data))
			}

			fnValue := reflect.ValueOf(fn.fn)

			var args = make([]reflect.Value, int(count))

			for i := C.size_t(0); i < count; i++ {
				mt := C.Plugify_GetMethodParamType(m, C.ptrdiff_t(i))
				vt := ValueType(mt.valueType)
				if mt.ref {
					args[i] = paramRefToObject(vt, p, i)
				} else {
					args[i] = paramToObject(m, vt, p, i)
				}
			}

			rets := fnValue.Call(args)

			mt := C.Plugify_GetMethodParamType(m, C.ptrdiff_t(-1))
			vt := ValueType(mt.valueType)
			setObjReturn(vt, r, rets)

			for i := C.size_t(0); i < count; i++ {
				mt = C.Plugify_GetMethodParamType(m, C.ptrdiff_t(i))
				vt = ValueType(mt.valueType)
				if mt.ref {
					setRefParam(vt, p, i, args[i])
				}
			}
		},
		Catch: func(e Exception) {
			msg := fmt.Sprintf("%v", e)
			C.Plugify_PrintException(msg)
		},
	}.Do()
}

func GetFunctionPointerForDelegate(fn any) unsafe.Pointer {
	fnType := reflect.ValueOf(fn)
	if fnType.Kind() != reflect.Func {
		panicker("expected a function")
	}

	for _, v := range functionMap {
		if v.fn == fn {
			return v.addr
		}
	}

	valueType := fnType.Type()
	fnPtr := unsafe.Pointer(fnType.Pointer())

	pkgPath := valueType.PkgPath()
	if pkgPath == "main" {
		pkgPath = plugify.Name
	} else {
		lastSlashIndex := strings.LastIndex(pkgPath, "/")
		if lastSlashIndex != -1 {
			pkgPath = pkgPath[lastSlashIndex+1:]
		}
	}

	name := fmt.Sprintf("%s.%s", pkgPath, valueType.Name())
	callback := C.Plugify_NewCallback(name, fnPtr)
	if callback == nil {
		panicker(fmt.Sprintf("%s (jit error: not found)", name))
	}

	callbacks = append(callbacks, callback)

	addr := C.Plugify_GetCallbackFunction(callback)
	if addr == nil {
		panicker(fmt.Sprintf("%s (jit error: %s)", name, string(C.GoString(C.Plugify_GetCallbackError(callback)))))
	}

	functionMap[fnPtr] = function{fn, addr}

	return addr
}
