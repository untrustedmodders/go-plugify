package plugify

/*
#include "plugify.h"
*/
import "C"
import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"unsafe"
)

var (
	callbacks []C.JitCallback
	calls     []C.JitCall
)

type function struct {
	fn   any
	addr unsafe.Pointer
}

var (
	mu          sync.Mutex
	functionMap = make(map[unsafe.Pointer]function)
)

func raw[T any](val T) uint64 {
	return *(*uint64)(unsafe.Pointer(&val))
}

func pin[T any](val T, pool *arena, size int) uint64 {
	tmp := (*T)(pool.Alloc(size))
	*tmp = val
	return uint64(uintptr(unsafe.Pointer(tmp)))
}

func GetDelegateForFunctionPointer(fnPtr unsafe.Pointer, fnType reflect.Type) any {
	if fnPtr == nil {
		return nil
	}

	mu.Lock()
	defer mu.Unlock()

	el, ok := functionMap[fnPtr]
	if ok {
		return el.fn
	}

	if fnType.Kind() != reflect.Func {
		panicker("expected a function")
	}

	parameterTypes, err := getParameterTypes(fnType)
	if err != nil {
		panicker(err)
	}
	returnType, retCount, err := getReturnType(fnType)
	if err != nil {
		panicker(err)
	}

	hasRet := hasReturnType(returnType)

	paramCount := len(parameterTypes)
	if hasRet {
		paramCount += 1
	}
	paramSize := paramCount * sizeOfValueType(UInt64)

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
		var pool arena

		params := unsafe.Slice((*uint64)(pool.Alloc(paramSize)), paramCount)

		results := make([]reflect.Value, retCount)

		index := 0

		retType := returnType.valueType
		if hasRet {
			size := sizeOfValueType(retType)
			ptr := uint64(uintptr(pool.Alloc(size)))
			switch retType {
			case Vector2Type, Vector3Type, Vector4Type, Matrix4x4Type:
				break
			case String:
				defer DestroyString((*PlgString)(unsafe.Pointer(uintptr(ptr))))
			case Any:
				defer DestroyVariant((*PlgVariant)(unsafe.Pointer(uintptr(ptr))))
			case ArrayBool:
				defer DestroyVectorBool((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
			case ArrayChar8:
				defer DestroyVectorChar8((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
			case ArrayChar16:
				defer DestroyVectorChar16((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
			case ArrayInt8:
				defer DestroyVectorInt8((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
			case ArrayInt16:
				defer DestroyVectorInt16((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
			case ArrayInt32:
				defer DestroyVectorInt32((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
			case ArrayInt64:
				defer DestroyVectorInt64((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
			case ArrayUInt8:
				defer DestroyVectorUInt8((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
			case ArrayUInt16:
				defer DestroyVectorUInt16((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
			case ArrayUInt32:
				defer DestroyVectorUInt32((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
			case ArrayUInt64:
				defer DestroyVectorUInt64((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
			case ArrayPointer:
				defer DestroyVectorPointer((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
			case ArrayFloat:
				defer DestroyVectorFloat((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
			case ArrayDouble:
				defer DestroyVectorDouble((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
			case ArrayString:
				defer DestroyVectorString((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
			case ArrayAny:
				defer DestroyVectorVariant((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
			case ArrayVector2:
				defer DestroyVectorVector2((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
			case ArrayVector3:
				defer DestroyVectorVector3((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
			case ArrayVector4:
				defer DestroyVectorVector4((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
			case ArrayMatrix4x4:
				defer DestroyVectorMatrix4x4((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
			default:
				panicker(fmt.Sprintf("GetDelegateForFunctionPointer return type not supported %v", retType))
			}
			params[index] = ptr
			index++
		}

		for i, arg := range args {
			paramType := parameterTypes[i]
			valueType := paramType.valueType
			size := sizeOfValueType(valueType)
			var ptr uint64
			if paramType.ref {
				switch valueType {
				case Bool:
					val := arg.Interface().(*bool)
					ptr = pin(*val, &pool, size)
					defer func() {
						*val = *(*bool)(unsafe.Pointer(uintptr(ptr)))
					}()
				case Char8:
					val := arg.Interface().(*int8)
					ptr = pin(*val, &pool, size)
					defer func() {
						*val = *(*int8)(unsafe.Pointer(uintptr(ptr)))
					}()
				case Char16:
					val := arg.Interface().(*uint16)
					ptr = pin(*val, &pool, size)
					defer func() {
						*val = *(*uint16)(unsafe.Pointer(uintptr(ptr)))
					}()
				case Int8:
					val := arg.Interface().(*int8)
					ptr = pin(*val, &pool, size)
					defer func() {
						*val = *(*int8)(unsafe.Pointer(uintptr(ptr)))
					}()
				case Int16:
					val := arg.Interface().(*int16)
					ptr = pin(*val, &pool, size)
					defer func() {
						*val = *(*int16)(unsafe.Pointer(uintptr(ptr)))
					}()
				case Int32:
					val := arg.Interface().(*int32)
					ptr = pin(*val, &pool, size)
					defer func() {
						*val = *(*int32)(unsafe.Pointer(uintptr(ptr)))
					}()
				case Int64:
					val := arg.Interface().(*int64)
					ptr = pin(*val, &pool, size)
					defer func() {
						*val = *(*int64)(unsafe.Pointer(uintptr(ptr)))
					}()
				case UInt8:
					val := arg.Interface().(*uint8)
					ptr = pin(*val, &pool, size)
					defer func() {
						*val = *(*uint8)(unsafe.Pointer(uintptr(ptr)))
					}()
				case UInt16:
					val := arg.Interface().(*uint16)
					ptr = pin(*val, &pool, size)
					defer func() {
						*val = *(*uint16)(unsafe.Pointer(uintptr(ptr)))
					}()
				case UInt32:
					val := arg.Interface().(*uint32)
					ptr = pin(*val, &pool, size)
					defer func() {
						*val = *(*uint32)(unsafe.Pointer(uintptr(ptr)))
					}()
				case UInt64:
					val := arg.Interface().(*uint64)
					ptr = pin(*val, &pool, size)
					defer func() {
						*val = *(*uint64)(unsafe.Pointer(uintptr(ptr)))
					}()
				case Pointer:
					val := arg.Interface().(*uintptr)
					ptr = pin(*val, &pool, size)
					defer func() {
						*val = *(*uintptr)(unsafe.Pointer(uintptr(ptr)))
					}()
				case Float:
					val := arg.Interface().(*float32)
					ptr = pin(*val, &pool, size)
					defer func() {
						*val = *(*float32)(unsafe.Pointer(uintptr(ptr)))
					}()
				case Double:
					val := arg.Interface().(*float64)
					ptr = pin(*val, &pool, size)
					defer func() {
						*val = *(*float64)(unsafe.Pointer(uintptr(ptr)))
					}()
				case Vector2Type:
					ptr = pin(*(arg.Interface().(*Vector2)), &pool, size)
					defer func() {
						*(arg.Interface().(*Vector2)) = *(*Vector2)(unsafe.Pointer(uintptr(ptr)))
					}()
				case Vector3Type:
					val := arg.Interface().(*Vector3)
					ptr = pin(*val, &pool, size)
					defer func() {
						*val = *(*Vector3)(unsafe.Pointer(uintptr(ptr)))
					}()
				case Vector4Type:
					val := arg.Interface().(*Vector4)
					ptr = pin(*val, &pool, size)
					defer func() {
						*val = *(*Vector4)(unsafe.Pointer(uintptr(ptr)))
					}()
				case Matrix4x4Type:
					val := arg.Interface().(*Matrix4x4)
					ptr = pin(*val, &pool, size)
					defer func() {
						*val = *(*Matrix4x4)(unsafe.Pointer(uintptr(ptr)))
					}()
				case Function:
					ptr = uint64(uintptr(GetFunctionPointerForDelegate(arg)))
				case String:
					val := arg.Interface().(*string)
					ptr = pin(ConstructString(*val), &pool, size)
					defer func() {
						*val = GetStringData((*PlgString)(unsafe.Pointer(uintptr(ptr))))
						DestroyString((*PlgString)(unsafe.Pointer(uintptr(ptr))))
					}()
				case Any:
					val := arg.Interface().(*any)
					ptr = pin(ConstructVariant(*val), &pool, size)
					defer func() {
						*val = GetVariantData((*PlgVariant)(unsafe.Pointer(uintptr(ptr))))
						DestroyVariant((*PlgVariant)(unsafe.Pointer(uintptr(ptr))))
					}()
				case ArrayBool:
					val := arg.Interface().(*[]bool)
					ptr = pin(ConstructVectorBool(*val), &pool, size)
					defer func() {
						*val = GetVectorDataBool((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
						DestroyVectorBool((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
					}()
				case ArrayChar8:
					val := arg.Interface().(*[]int8)
					ptr = pin(ConstructVectorChar8(*val), &pool, size)
					defer func() {
						*val = GetVectorDataChar8((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
						DestroyVectorChar8((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
					}()
				case ArrayChar16:
					val := arg.Interface().(*[]uint16)
					ptr = pin(ConstructVectorChar16(*val), &pool, size)
					defer func() {
						*val = GetVectorDataChar16((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
						DestroyVectorChar16((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
					}()
				case ArrayInt8:
					val := arg.Interface().(*[]int8)
					ptr = pin(ConstructVectorInt8(*val), &pool, size)
					defer func() {
						*val = GetVectorDataInt8((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
						DestroyVectorInt8((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
					}()
				case ArrayInt16:
					val := arg.Interface().(*[]int16)
					ptr = pin(ConstructVectorInt16(*val), &pool, size)
					defer func() {
						*val = GetVectorDataInt16((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
						DestroyVectorInt16((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
					}()
				case ArrayInt32:
					val := arg.Interface().(*[]int32)
					ptr = pin(ConstructVectorInt32(*val), &pool, size)
					defer func() {
						*val = GetVectorDataInt32((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
						DestroyVectorInt32((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
					}()
				case ArrayInt64:
					val := arg.Interface().(*[]int64)
					ptr = pin(ConstructVectorInt64(*val), &pool, size)
					defer func() {
						*val = GetVectorDataInt64((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
						DestroyVectorInt64((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
					}()
				case ArrayUInt8:
					val := arg.Interface().(*[]uint8)
					ptr = pin(ConstructVectorUInt8(*val), &pool, size)
					defer func() {
						*val = GetVectorDataUInt8((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
						DestroyVectorUInt8((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
					}()
				case ArrayUInt16:
					val := arg.Interface().(*[]uint16)
					ptr = pin(ConstructVectorUInt16(*val), &pool, size)
					defer func() {
						*val = GetVectorDataUInt16((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
						DestroyVectorUInt16((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
					}()
				case ArrayUInt32:
					val := arg.Interface().(*[]uint32)
					ptr = pin(ConstructVectorUInt32(*val), &pool, size)
					defer func() {
						*val = GetVectorDataUInt32((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
						DestroyVectorUInt32((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
					}()
				case ArrayUInt64:
					val := arg.Interface().(*[]uint64)
					ptr = pin(ConstructVectorUInt64(*val), &pool, size)
					defer func() {
						*val = GetVectorDataUInt64((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
						DestroyVectorUInt64((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
					}()
				case ArrayPointer:
					val := arg.Interface().(*[]uintptr)
					ptr = pin(ConstructVectorPointer(*val), &pool, size)
					defer func() {
						*val = GetVectorDataPointer((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
						DestroyVectorPointer((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
					}()
				case ArrayFloat:
					val := arg.Interface().(*[]float32)
					ptr = pin(ConstructVectorFloat(*val), &pool, size)
					defer func() {
						*val = GetVectorDataFloat((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
						DestroyVectorFloat((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
					}()
				case ArrayDouble:
					val := arg.Interface().(*[]float64)
					ptr = pin(ConstructVectorDouble(*val), &pool, size)
					defer func() {
						*val = GetVectorDataDouble((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
						DestroyVectorDouble((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
					}()
				case ArrayString:
					val := arg.Interface().(*[]string)
					ptr = pin(ConstructVectorString(*val), &pool, size)
					defer func() {
						*val = GetVectorDataString((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
						DestroyVectorString((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
					}()
				case ArrayAny:
					val := arg.Interface().(*[]any)
					ptr = pin(ConstructVectorVariant(*val), &pool, size)
					defer func() {
						*val = GetVectorDataVariant((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
						DestroyVectorVariant((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
					}()
				case ArrayVector2:
					val := arg.Interface().(*[]Vector2)
					ptr = pin(ConstructVectorVector2(*val), &pool, size)
					defer func() {
						*val = GetVectorDataVector2((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
						DestroyVectorVector2((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
					}()
				case ArrayVector3:
					val := arg.Interface().(*[]Vector3)
					ptr = pin(ConstructVectorVector3(*val), &pool, size)
					defer func() {
						*val = GetVectorDataVector3((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
						DestroyVectorVector3((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
					}()
				case ArrayVector4:
					val := arg.Interface().(*[]Vector4)
					ptr = pin(ConstructVectorVector4(*val), &pool, size)
					defer func() {
						*val = GetVectorDataVector4((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
						DestroyVectorVector4((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
					}()
				case ArrayMatrix4x4:
					val := arg.Interface().(*[]Matrix4x4)
					ptr = pin(ConstructVectorMatrix4x4(*val), &pool, size)
					defer func() {
						*val = GetVectorDataMatrix4x4((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
						DestroyVectorMatrix4x4((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
					}()
				default:
					panicker(fmt.Sprintf("GetDelegateForFunctionPointer parameter type not supported %v", retType))
				}
			} else {
				switch valueType {
				case Bool:
					ptr = raw(arg.Interface().(bool))
				case Char8:
					ptr = raw(arg.Interface().(int8))
				case Char16:
					ptr = raw(arg.Interface().(uint16))
				case Int8:
					ptr = raw(arg.Interface().(int8))
				case Int16:
					ptr = raw(arg.Interface().(int16))
				case Int32:
					ptr = raw(arg.Interface().(int32))
				case Int64:
					ptr = raw(arg.Interface().(int64))
				case UInt8:
					ptr = raw(arg.Interface().(uint8))
				case UInt16:
					ptr = raw(arg.Interface().(uint16))
				case UInt32:
					ptr = raw(arg.Interface().(uint32))
				case UInt64:
					ptr = raw(arg.Interface().(uint64))
				case Pointer:
					ptr = raw(arg.Interface().(uintptr))
				case Float:
					ptr = raw(arg.Interface().(float32))
				case Double:
					ptr = raw(arg.Interface().(float64))
				case Vector2Type:
					ptr = pin(arg.Interface().(Vector2), &pool, size)
				case Vector3Type:
					ptr = pin(arg.Interface().(Vector3), &pool, size)
				case Vector4Type:
					ptr = pin(arg.Interface().(Vector4), &pool, size)
				case Matrix4x4Type:
					ptr = pin(arg.Interface().(Matrix4x4), &pool, size)
				case Function:
					ptr = uint64(uintptr(GetFunctionPointerForDelegate(arg)))
				case String:
					ptr = pin(ConstructString(arg.Interface().(string)), &pool, size)
					defer DestroyString((*PlgString)(unsafe.Pointer(uintptr(ptr))))
				case Any:
					ptr = pin(ConstructVariant(arg.Interface().(any)), &pool, size)
					defer DestroyVariant((*PlgVariant)(unsafe.Pointer(uintptr(ptr))))
				case ArrayBool:
					ptr = pin(ConstructVectorBool(arg.Interface().([]bool)), &pool, size)
					defer DestroyVectorBool((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
				case ArrayChar8:
					ptr = pin(ConstructVectorChar8(arg.Interface().([]int8)), &pool, size)
					defer DestroyVectorChar8((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
				case ArrayChar16:
					ptr = pin(ConstructVectorChar16(arg.Interface().([]uint16)), &pool, size)
					defer DestroyVectorChar16((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
				case ArrayInt8:
					ptr = pin(ConstructVectorInt8(arg.Interface().([]int8)), &pool, size)
					defer DestroyVectorInt8((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
				case ArrayInt16:
					ptr = pin(ConstructVectorInt16(arg.Interface().([]int16)), &pool, size)
					defer DestroyVectorInt16((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
				case ArrayInt32:
					ptr = pin(ConstructVectorInt32(arg.Interface().([]int32)), &pool, size)
					defer DestroyVectorInt32((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
				case ArrayInt64:
					ptr = pin(ConstructVectorInt64(arg.Interface().([]int64)), &pool, size)
					defer DestroyVectorInt64((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
				case ArrayUInt8:
					ptr = pin(ConstructVectorUInt8(arg.Interface().([]uint8)), &pool, size)
					defer DestroyVectorUInt8((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
				case ArrayUInt16:
					ptr = pin(ConstructVectorUInt16(arg.Interface().([]uint16)), &pool, size)
					defer DestroyVectorUInt16((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
				case ArrayUInt32:
					ptr = pin(ConstructVectorUInt32(arg.Interface().([]uint32)), &pool, size)
					defer DestroyVectorUInt32((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
				case ArrayUInt64:
					ptr = pin(ConstructVectorUInt64(arg.Interface().([]uint64)), &pool, size)
					defer DestroyVectorUInt64((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
				case ArrayPointer:
					ptr = pin(ConstructVectorPointer(arg.Interface().([]uintptr)), &pool, size)
					defer DestroyVectorPointer((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
				case ArrayFloat:
					ptr = pin(ConstructVectorFloat(arg.Interface().([]float32)), &pool, size)
					defer DestroyVectorFloat((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
				case ArrayDouble:
					ptr = pin(ConstructVectorDouble(arg.Interface().([]float64)), &pool, size)
					defer DestroyVectorDouble((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
				case ArrayString:
					ptr = pin(ConstructVectorString(arg.Interface().([]string)), &pool, size)
					defer DestroyVectorString((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
				case ArrayAny:
					ptr = pin(ConstructVectorVariant(arg.Interface().([]any)), &pool, size)
					defer DestroyVectorVariant((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
				case ArrayVector2:
					ptr = pin(ConstructVectorVector2(arg.Interface().([]Vector2)), &pool, size)
					defer DestroyVectorVector2((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
				case ArrayVector3:
					ptr = pin(ConstructVectorVector3(arg.Interface().([]Vector3)), &pool, size)
					defer DestroyVectorVector3((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
				case ArrayVector4:
					ptr = pin(ConstructVectorVector4(arg.Interface().([]Vector4)), &pool, size)
					defer DestroyVectorVector4((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
				case ArrayMatrix4x4:
					ptr = pin(ConstructVectorMatrix4x4(arg.Interface().([]Matrix4x4)), &pool, size)
					defer DestroyVectorMatrix4x4((*PlgVector)(unsafe.Pointer(uintptr(ptr))))
				default:
					panicker(fmt.Sprintf("GetDelegateForFunctionPointer parameter type not supported %v", valueType))
				}
			}
			params[index] = ptr
			index++
		}

		var retStore C.uint128_t
		C.Plugify_CallFunction(call, (*C.uint64_t)(unsafe.Pointer(unsafe.SliceData(params))), &retStore)
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

func paramRefToObject(vt valueType, p *C.Parameters, i C.size_t) reflect.Value {
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

func paramToObject(m C.MethodHandle, vt valueType, p *C.Parameters, i C.size_t) reflect.Value {
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

func setRefParam(vt valueType, p *C.Parameters, i C.size_t, val reflect.Value) {
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

func setObjReturn(vt valueType, r *C.Return, rets []reflect.Value) {
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

//export plugify_InternalCall
func plugify_InternalCall(m C.MethodHandle, data unsafe.Pointer, p *C.Parameters, count C.size_t, r *C.Return) {
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
				vt := valueType(mt.valueType)
				if mt.ref {
					args[i] = paramRefToObject(vt, p, i)
				} else {
					args[i] = paramToObject(m, vt, p, i)
				}
			}

			rets := fnValue.Call(args)

			mt := C.Plugify_GetMethodParamType(m, C.ptrdiff_t(-1))
			vt := valueType(mt.valueType)
			setObjReturn(vt, r, rets)

			for i := C.size_t(0); i < count; i++ {
				mt = C.Plugify_GetMethodParamType(m, C.ptrdiff_t(i))
				vt = valueType(mt.valueType)
				if mt.ref {
					setRefParam(vt, p, i, args[i])
				}
			}
		},
		Catch: func(e Exception) {
			PrintStacktrace(e)
		},
	}.Do()
}

func GetFunctionPointerForDelegate(fn any) unsafe.Pointer {
	if fn == nil {
		return nil
	}

	mu.Lock()
	defer mu.Unlock()

	fnType := reflect.ValueOf(fn)
	if fnType.Kind() != reflect.Func {
		panicker("expected a function")
	}

	valueType := fnType.Type()
	fnPtr := unsafe.Pointer(fnType.Pointer())

	for _, v := range functionMap {
		if unsafe.Pointer(reflect.ValueOf(v.fn).Pointer()) == fnPtr {
			return v.addr
		}
	}

	name, err := getFunctionName(valueType)
	if err != nil {
		panicker(err)
	}

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

func getFunctionName(t reflect.Type) (string, error) {
	pkg, err := normalizePkgName(t)
	if err != nil {
		return "", err
	}
	name := t.Name()
	if name == "" {
		return "", fmt.Errorf("no package name for type %v", t)
	}
	return fmt.Sprintf("%s.%s", pkg, name), nil
}

func normalizePkgName(t reflect.Type) (string, error) {
	path := t.PkgPath()
	if path == "" {
		return "", fmt.Errorf("no package path for type %v", t)
	}
	if path == "main" {
		return Plugin.Name, nil
	}

	parts := splitModulePath(path)

	i := 0
	if isExternalModule(parts) {
		i = 2 // repo name only from example.com/<owner>/<repo>/...
	} else {
		i = len(parts) - 1
	}

	// remove suffix before - or . as they consider invalid in plugify anyway
	return trimSuffix(parts[i]), nil
}

func splitModulePath(path string) []string {
	parts := strings.Split(path, "/")
	i := len(parts) - 1
	if i > 0 {
		if isVersionSegment(parts[i]) {
			parts = parts[:i]
		}
	}
	return parts
}

func isVersionSegment(s string) bool {
	if len(s) < 2 || s[0] != 'v' {
		return false
	}
	_, err := strconv.Atoi(s[1:])
	return err == nil
}

func isExternalModule(parts []string) bool {
	return len(parts) >= 3 && strings.Contains(parts[0], ".")
}

func trimSuffix(s string) string {
	if i := strings.LastIndexAny(s, "-."); i != -1 {
		return s[i+1:]
	}
	return s
}
