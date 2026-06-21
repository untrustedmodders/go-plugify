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

type funcArg struct {
	isRef bool
	t     reflect.Type
}

type function struct {
	fn   any
	addr unsafe.Pointer

	args []funcArg
	ret  funcArg
}

var (
	mu          sync.Mutex
	functionMap = make(map[unsafe.Pointer]function)
)

func raw[T any](val T) unsafe.Pointer {
	return unsafe.Pointer(uintptr(*(*uintptr)(unsafe.Pointer(&val))))
}

func pinT[T any](val T, pool *arena, size int) unsafe.Pointer {
	tmp := (*T)(pool.Alloc(size))
	*tmp = val
	return unsafe.Pointer(tmp)
}

func pin(val unsafe.Pointer, pool *arena, size int) unsafe.Pointer {
	tmp := pool.Alloc(size)
	C.memcpy(tmp, val, C.size_t(size))
	return tmp
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
			ptr := pool.Alloc(size)
			switch retType {
			case Vector2Type, Vector3Type, Vector4Type, Matrix4x4Type:
				break
			case String:
				defer DestroyString((*PlgString)(ptr))
			case Any:
				defer DestroyVariant((*PlgVariant)(ptr))
			case ArrayBool:
				defer DestroyVectorBool((*PlgVector)(ptr))
			case ArrayChar8:
				defer DestroyVectorChar8((*PlgVector)(ptr))
			case ArrayChar16:
				defer DestroyVectorChar16((*PlgVector)(ptr))
			case ArrayInt8:
				defer DestroyVectorInt8((*PlgVector)(ptr))
			case ArrayInt16:
				defer DestroyVectorInt16((*PlgVector)(ptr))
			case ArrayInt32:
				defer DestroyVectorInt32((*PlgVector)(ptr))
			case ArrayInt64:
				defer DestroyVectorInt64((*PlgVector)(ptr))
			case ArrayUInt8:
				defer DestroyVectorUInt8((*PlgVector)(ptr))
			case ArrayUInt16:
				defer DestroyVectorUInt16((*PlgVector)(ptr))
			case ArrayUInt32:
				defer DestroyVectorUInt32((*PlgVector)(ptr))
			case ArrayUInt64:
				defer DestroyVectorUInt64((*PlgVector)(ptr))
			case ArrayPointer:
				defer DestroyVectorPointer((*PlgVector)(ptr))
			case ArrayFloat:
				defer DestroyVectorFloat((*PlgVector)(ptr))
			case ArrayDouble:
				defer DestroyVectorDouble((*PlgVector)(ptr))
			case ArrayString:
				defer DestroyVectorString((*PlgVector)(ptr))
			case ArrayAny:
				defer DestroyVectorVariant((*PlgVector)(ptr))
			case ArrayVector2:
				defer DestroyVectorVector2((*PlgVector)(ptr))
			case ArrayVector3:
				defer DestroyVectorVector3((*PlgVector)(ptr))
			case ArrayVector4:
				defer DestroyVectorVector4((*PlgVector)(ptr))
			case ArrayMatrix4x4:
				defer DestroyVectorMatrix4x4((*PlgVector)(ptr))
			default:
				panicker(fmt.Sprintf("GetDelegateForFunctionPointer defered return type not supported %v", retType))
			}
			params[index] = uint64(uintptr(ptr))
			index++
		}

		for i, arg := range args {
			pt := parameterTypes[i]
			vt := pt.valueType
			size := sizeOfValueType(vt)
			var ptrUnsafe unsafe.Pointer
			if pt.ref {
				switch vt {
				case Bool:
					ptrUnsafe = pin(arg.UnsafePointer(), &pool, size)
					defer func() {
						arg.Elem().SetBool(*(*bool)(ptrUnsafe))
					}()
				case Char8:
					ptrUnsafe = pin(arg.UnsafePointer(), &pool, size)
					defer func() {
						arg.Elem().SetInt(int64(*(*int8)(ptrUnsafe)))
					}()
				case Char16:
					ptrUnsafe = pin(arg.UnsafePointer(), &pool, size)
					defer func() {
						arg.Elem().SetUint(uint64(*(*uint16)(ptrUnsafe)))
					}()
				case Int8:
					ptrUnsafe = pin(arg.UnsafePointer(), &pool, size)
					defer func() {
						arg.Elem().SetInt(int64(*(*int8)(ptrUnsafe)))
					}()
				case Int16:
					ptrUnsafe = pin(arg.UnsafePointer(), &pool, size)
					defer func() {
						arg.Elem().SetInt(int64(*(*int16)(ptrUnsafe)))
					}()
				case Int32:
					ptrUnsafe = pin(arg.UnsafePointer(), &pool, size)
					defer func() {
						arg.Elem().SetInt(int64(*(*int32)(ptrUnsafe)))
					}()
				case Int64:
					ptrUnsafe = pin(arg.UnsafePointer(), &pool, size)
					defer func() {
						arg.Elem().SetInt(*(*int64)(ptrUnsafe))
					}()
				case UInt8:
					ptrUnsafe = pin(arg.UnsafePointer(), &pool, size)
					defer func() {
						arg.Elem().SetUint(uint64(*(*uint8)(ptrUnsafe)))
					}()
				case UInt16:
					ptrUnsafe = pin(arg.UnsafePointer(), &pool, size)
					defer func() {
						arg.Elem().SetUint(uint64(*(*uint16)(ptrUnsafe)))
					}()
				case UInt32:
					ptrUnsafe = pin(arg.UnsafePointer(), &pool, size)
					defer func() {
						arg.Elem().SetUint(uint64(*(*uint32)(ptrUnsafe)))
					}()
				case UInt64:
					ptrUnsafe = pin(arg.UnsafePointer(), &pool, size)
					defer func() {
						arg.Elem().SetUint(*(*uint64)(ptrUnsafe))
					}()
				case Pointer:
					ptrUnsafe = pin(arg.UnsafePointer(), &pool, size)
					defer func() {
						arg.Elem().SetUint(uint64(*(*uintptr)(ptrUnsafe)))
					}()
				case Float:
					ptrUnsafe = pin(arg.UnsafePointer(), &pool, size)
					defer func() {
						arg.Elem().SetFloat(float64(*(*float32)(ptrUnsafe)))
					}()
				case Double:
					ptrUnsafe = pin(arg.UnsafePointer(), &pool, size)
					defer func() {
						arg.Elem().SetFloat(*(*float64)(ptrUnsafe))
					}()
				case Vector2Type:
					ptrUnsafe = pin(arg.UnsafePointer(), &pool, size)
					defer func() {
						C.memcpy(unsafe.Pointer(arg.UnsafePointer()), ptrUnsafe, C.size_t(size))
					}()
				case Vector3Type:
					ptrUnsafe = pin(arg.UnsafePointer(), &pool, size)
					defer func() {
						C.memcpy(unsafe.Pointer(arg.UnsafePointer()), ptrUnsafe, C.size_t(size))
					}()
				case Vector4Type:
					ptrUnsafe = pin(arg.UnsafePointer(), &pool, size)
					defer func() {
						C.memcpy(unsafe.Pointer(arg.UnsafePointer()), ptrUnsafe, C.size_t(size))
					}()
				case Matrix4x4Type:
					ptrUnsafe = pin(arg.UnsafePointer(), &pool, size)
					defer func() {
						C.memcpy(unsafe.Pointer(arg.UnsafePointer()), ptrUnsafe, C.size_t(size))
					}()
				case Function:
					ptrUnsafe = GetFunctionPointerForDelegate(arg)
				case String:
					elem := arg.Elem()
					ptrUnsafe = pinT(ConstructString(elem.String()), &pool, size)
					defer func() {
						elem.SetString(GetStringData[string]((*PlgString)(ptrUnsafe)))
						DestroyString((*PlgString)(ptrUnsafe))
					}()
				case Any:
					elem := arg.Elem()
					ptrUnsafe = pinT(ConstructVariant(elem.Interface()), &pool, size)
					defer func() {
						data := GetVariantData((*PlgVariant)(ptrUnsafe))
						if data != nil {
							elem.Set(reflect.ValueOf(data))
						}
						DestroyVariant((*PlgVariant)(ptrUnsafe))
					}()
				case ArrayBool:
					slice := arg.Elem()
					ptrUnsafe = pinT(constructVectorBoolToSlice(slice), &pool, size)
					defer func() {
						getVectorDataBoolToSlice((*PlgVector)(ptrUnsafe), slice)
						DestroyVectorBool((*PlgVector)(ptrUnsafe))
					}()
				case ArrayChar8:
					slice := arg.Elem()
					ptrUnsafe = pinT(constructVectorChar8ToSlice(slice), &pool, size)
					defer func() {
						getVectorDataChar8ToSlice((*PlgVector)(ptrUnsafe), slice)
						DestroyVectorChar8((*PlgVector)(ptrUnsafe))
					}()
				case ArrayChar16:
					slice := arg.Elem()
					ptrUnsafe = pinT(constructVectorChar16ToSlice(slice), &pool, size)
					defer func() {
						getVectorDataChar16ToSlice((*PlgVector)(ptrUnsafe), slice)
						DestroyVectorChar16((*PlgVector)(ptrUnsafe))
					}()
				case ArrayInt8:
					slice := arg.Elem()
					ptrUnsafe = pinT(constructVectorInt8ToSlice(slice), &pool, size)
					defer func() {
						getVectorDataInt8ToSlice((*PlgVector)(ptrUnsafe), slice)
						DestroyVectorInt8((*PlgVector)(ptrUnsafe))
					}()
				case ArrayInt16:
					slice := arg.Elem()
					ptrUnsafe = pinT(constructVectorInt16ToSlice(slice), &pool, size)
					defer func() {
						getVectorDataInt16ToSlice((*PlgVector)(ptrUnsafe), slice)
						DestroyVectorInt16((*PlgVector)(ptrUnsafe))
					}()
				case ArrayInt32:
					slice := arg.Elem()
					ptrUnsafe = pinT(constructVectorInt32ToSlice(slice), &pool, size)
					defer func() {
						getVectorDataInt32ToSlice((*PlgVector)(ptrUnsafe), slice)
						DestroyVectorInt32((*PlgVector)(ptrUnsafe))
					}()
				case ArrayInt64:
					slice := arg.Elem()
					ptrUnsafe = pinT(constructVectorInt64ToSlice(slice), &pool, size)
					defer func() {
						getVectorDataInt64ToSlice((*PlgVector)(ptrUnsafe), slice)
						DestroyVectorInt32((*PlgVector)(ptrUnsafe))
					}()
				case ArrayUInt8:
					slice := arg.Elem()
					ptrUnsafe = pinT(constructVectorUInt8ToSlice(slice), &pool, size)
					defer func() {
						getVectorDataUInt8ToSlice((*PlgVector)(ptrUnsafe), slice)
						DestroyVectorUInt8((*PlgVector)(ptrUnsafe))
					}()
				case ArrayUInt16:
					slice := arg.Elem()
					ptrUnsafe = pinT(constructVectorUInt16ToSlice(slice), &pool, size)
					defer func() {
						getVectorDataUInt16ToSlice((*PlgVector)(ptrUnsafe), slice)
						DestroyVectorUInt16((*PlgVector)(ptrUnsafe))
					}()
				case ArrayUInt32:
					slice := arg.Elem()
					ptrUnsafe = pinT(constructVectorUInt32ToSlice(slice), &pool, size)
					defer func() {
						getVectorDataUInt32ToSlice((*PlgVector)(ptrUnsafe), slice)
						DestroyVectorUInt32((*PlgVector)(ptrUnsafe))
					}()
				case ArrayUInt64:
					slice := arg.Elem()
					ptrUnsafe = pinT(constructVectorUInt64ToSlice(slice), &pool, size)
					defer func() {
						getVectorDataUInt64ToSlice((*PlgVector)(ptrUnsafe), slice)
						DestroyVectorUInt64((*PlgVector)(ptrUnsafe))
					}()
				case ArrayPointer:
					slice := arg.Elem()
					ptrUnsafe = pinT(constructVectorPointerToSlice(slice), &pool, size)
					defer func() {
						getVectorDataPointerToSlice((*PlgVector)(ptrUnsafe), slice)
						DestroyVectorPointer((*PlgVector)(ptrUnsafe))
					}()
				case ArrayFloat:
					slice := arg.Elem()
					ptrUnsafe = pinT(constructVectorFloatToSlice(slice), &pool, size)
					defer func() {
						getVectorDataFloatToSlice((*PlgVector)(ptrUnsafe), slice)
						DestroyVectorFloat((*PlgVector)(ptrUnsafe))
					}()
				case ArrayDouble:
					slice := arg.Elem()
					ptrUnsafe = pinT(constructVectorDoubleToSlice(slice), &pool, size)
					defer func() {
						getVectorDataDoubleToSlice((*PlgVector)(ptrUnsafe), slice)
						DestroyVectorDouble((*PlgVector)(ptrUnsafe))
					}()
				case ArrayString:
					slice := arg.Elem()
					ptrUnsafe = pinT(constructVectorStringToSlice(slice), &pool, size)
					defer func() {
						getVectorDataStringToSlice((*PlgVector)(ptrUnsafe), slice)
						DestroyVectorString((*PlgVector)(ptrUnsafe))
					}()
				case ArrayAny:
					slice := arg.Elem()
					ptrUnsafe = pinT(constructVectorVariantToSlice(slice), &pool, size)
					defer func() {
						getVectorDataVariantToSlice((*PlgVector)(ptrUnsafe), slice)
						DestroyVectorVariant((*PlgVector)(ptrUnsafe))
					}()
				case ArrayVector2:
					slice := arg.Elem()
					ptrUnsafe = pinT(constructVectorVector2ToSlice(slice), &pool, size)
					defer func() {
						getVectorDataVector2ToSlice((*PlgVector)(ptrUnsafe), slice)
						DestroyVectorVector2((*PlgVector)(ptrUnsafe))
					}()
				case ArrayVector3:
					slice := arg.Elem()
					ptrUnsafe = pinT(constructVectorVector3ToSlice(slice), &pool, size)
					defer func() {
						getVectorDataVector3ToSlice((*PlgVector)(ptrUnsafe), slice)
						DestroyVectorVector3((*PlgVector)(ptrUnsafe))
					}()
				case ArrayVector4:
					slice := arg.Elem()
					ptrUnsafe = pinT(constructVectorVector4ToSlice(slice), &pool, size)
					defer func() {
						getVectorDataVector4ToSlice((*PlgVector)(ptrUnsafe), slice)
						DestroyVectorVector4((*PlgVector)(ptrUnsafe))
					}()
				case ArrayMatrix4x4:
					slice := arg.Elem()
					ptrUnsafe = pinT(constructVectorMatrix4x4ToSlice(slice), &pool, size)
					defer func() {
						getVectorDataMatrix4x4ToSlice((*PlgVector)(ptrUnsafe), slice)
						DestroyVectorMatrix4x4((*PlgVector)(ptrUnsafe))
					}()
				default:
					panicker(fmt.Sprintf("GetDelegateForFunctionPointer parameter type not supported %v", retType))
				}
			} else {
				switch vt {
				case Bool:
					ptrUnsafe = raw(arg.Bool())
				case Char8, Int8:
					ptrUnsafe = raw(int8(arg.Int()))
				case Int16:
					ptrUnsafe = raw(int16(arg.Int()))
				case Int32:
					ptrUnsafe = raw(int32(arg.Int()))
				case Int64:
					ptrUnsafe = raw(int64(arg.Int()))
				case UInt8:
					ptrUnsafe = raw(uint8(arg.Uint()))
				case Char16, UInt16:
					ptrUnsafe = raw(uint16(arg.Uint()))
				case UInt32:
					ptrUnsafe = raw(uint32(arg.Uint()))
				case UInt64:
					ptrUnsafe = raw(uint64(arg.Uint()))
				case Pointer:
					ptrUnsafe = raw(uintptr(arg.Uint()))
				case Float:
					ptrUnsafe = raw(float32(arg.Float()))
				case Double:
					ptrUnsafe = raw(arg.Float())
				case Vector2Type:
					val := reflect.New(arg.Type())
					val.Elem().Set(arg)
					ptrUnsafe = pin(val.UnsafePointer(), &pool, size)
				case Vector3Type:
					val := reflect.New(arg.Type())
					val.Elem().Set(arg)
					ptrUnsafe = pin(val.UnsafePointer(), &pool, size)
				case Vector4Type:
					val := reflect.New(arg.Type())
					val.Elem().Set(arg)
					ptrUnsafe = pin(val.UnsafePointer(), &pool, size)
				case Matrix4x4Type:
					val := reflect.New(arg.Type())
					val.Elem().Set(arg)
					ptrUnsafe = pin(val.UnsafePointer(), &pool, size)
				case Function:
					ptrUnsafe = GetFunctionPointerForDelegate(arg)
				case String:
					ptrUnsafe = pinT(ConstructString(arg.String()), &pool, size)
					defer DestroyString((*PlgString)(ptrUnsafe))
				case Any:
					ptrUnsafe = pinT(ConstructVariant(arg.Interface()), &pool, size)
					defer DestroyVariant((*PlgVariant)(ptrUnsafe))
				case ArrayBool:
					ptrUnsafe = pinT(constructVectorBoolToSlice(arg), &pool, size)
					defer DestroyVectorBool((*PlgVector)(ptrUnsafe))
				case ArrayChar8:
					ptrUnsafe = pinT(constructVectorChar8ToSlice(arg), &pool, size)
					defer DestroyVectorChar8((*PlgVector)(ptrUnsafe))
				case ArrayChar16:
					ptrUnsafe = pinT(constructVectorChar16ToSlice(arg), &pool, size)
					defer DestroyVectorChar16((*PlgVector)(ptrUnsafe))
				case ArrayInt8:
					ptrUnsafe = pinT(constructVectorInt8ToSlice(arg), &pool, size)
					defer DestroyVectorInt8((*PlgVector)(ptrUnsafe))
				case ArrayInt16:
					ptrUnsafe = pinT(constructVectorInt16ToSlice(arg), &pool, size)
					defer DestroyVectorInt16((*PlgVector)(ptrUnsafe))
				case ArrayInt32:
					ptrUnsafe = pinT(constructVectorInt32ToSlice(arg), &pool, size)
					defer DestroyVectorInt32((*PlgVector)(ptrUnsafe))
				case ArrayInt64:
					ptrUnsafe = pinT(constructVectorInt64ToSlice(arg), &pool, size)
					defer DestroyVectorInt64((*PlgVector)(ptrUnsafe))
				case ArrayUInt8:
					ptrUnsafe = pinT(constructVectorUInt8ToSlice(arg), &pool, size)
					defer DestroyVectorUInt8((*PlgVector)(ptrUnsafe))
				case ArrayUInt16:
					ptrUnsafe = pinT(constructVectorUInt16ToSlice(arg), &pool, size)
					defer DestroyVectorUInt16((*PlgVector)(ptrUnsafe))
				case ArrayUInt32:
					ptrUnsafe = pinT(constructVectorUInt32ToSlice(arg), &pool, size)
					defer DestroyVectorUInt32((*PlgVector)(ptrUnsafe))
				case ArrayUInt64:
					ptrUnsafe = pinT(constructVectorUInt64ToSlice(arg), &pool, size)
					defer DestroyVectorUInt64((*PlgVector)(ptrUnsafe))
				case ArrayPointer:
					ptrUnsafe = pinT(constructVectorPointerToSlice(arg), &pool, size)
					defer DestroyVectorPointer((*PlgVector)(ptrUnsafe))
				case ArrayFloat:
					ptrUnsafe = pinT(constructVectorFloatToSlice(arg), &pool, size)
					defer DestroyVectorFloat((*PlgVector)(ptrUnsafe))
				case ArrayDouble:
					ptrUnsafe = pinT(constructVectorDoubleToSlice(arg), &pool, size)
					defer DestroyVectorDouble((*PlgVector)(ptrUnsafe))
				case ArrayString:
					ptrUnsafe = pinT(constructVectorStringToSlice(arg), &pool, size)
					defer DestroyVectorString((*PlgVector)(ptrUnsafe))
				case ArrayAny:
					ptrUnsafe = pinT(constructVectorVariantToSlice(arg), &pool, size)
					defer DestroyVectorVariant((*PlgVector)(ptrUnsafe))
				case ArrayVector2:
					ptrUnsafe = pinT(constructVectorVector2ToSlice(arg), &pool, size)
					defer DestroyVectorVector2((*PlgVector)(ptrUnsafe))
				case ArrayVector3:
					ptrUnsafe = pinT(constructVectorVector3ToSlice(arg), &pool, size)
					defer DestroyVectorVector3((*PlgVector)(ptrUnsafe))
				case ArrayVector4:
					ptrUnsafe = pinT(constructVectorVector4ToSlice(arg), &pool, size)
					defer DestroyVectorVector4((*PlgVector)(ptrUnsafe))
				case ArrayMatrix4x4:
					ptrUnsafe = pinT(constructVectorMatrix4x4ToSlice(arg), &pool, size)
					defer DestroyVectorMatrix4x4((*PlgVector)(ptrUnsafe))
				default:
					panicker(fmt.Sprintf("GetDelegateForFunctionPointer parameter type not supported %v", vt))
				}
			}
			params[index] = (uint64)(uintptr(ptrUnsafe))
			index++
		}

		var retStore C.uint128_t
		C.Plugify_CallFunction(call, (*C.uint64_t)(unsafe.Pointer(unsafe.SliceData(params))), &retStore)
		ret := unsafe.Pointer(&retStore)

		switch retType {
		case Void:
			// skip
		case Bool:
			val := reflect.New(fnType.Out(0)).Elem()
			val.SetBool(*(*bool)(ret))
			results[0] = val
		case Char8:
			val := reflect.New(fnType.Out(0)).Elem()
			val.SetInt(int64(*(*int8)(ret)))
			results[0] = val
		case Char16:
			val := reflect.New(fnType.Out(0)).Elem()
			val.SetUint(uint64(*(*uint16)(ret)))
			results[0] = val
		case Int8:
			val := reflect.New(fnType.Out(0)).Elem()
			val.SetInt(int64(*(*int8)(ret)))
			results[0] = val
		case Int16:
			val := reflect.New(fnType.Out(0)).Elem()
			val.SetInt(int64(*(*int16)(ret)))
			results[0] = val
		case Int32:
			val := reflect.New(fnType.Out(0)).Elem()
			val.SetInt(int64(*(*int32)(ret)))
			results[0] = val
		case Int64:
			val := reflect.New(fnType.Out(0)).Elem()
			val.SetInt(*(*int64)(ret))
			results[0] = val
		case UInt8:
			val := reflect.New(fnType.Out(0)).Elem()
			val.SetUint(uint64(*(*uint8)(ret)))
			results[0] = val
		case UInt16:
			val := reflect.New(fnType.Out(0)).Elem()
			val.SetUint(uint64(*(*uint16)(ret)))
			results[0] = val
		case UInt32:
			val := reflect.New(fnType.Out(0)).Elem()
			val.SetUint(uint64(*(*uint32)(ret)))
			results[0] = val
		case UInt64:
			val := reflect.New(fnType.Out(0)).Elem()
			val.SetUint(*(*uint64)(ret))
			results[0] = val
		case Pointer:
			val := reflect.New(fnType.Out(0)).Elem()
			val.SetUint(uint64(*(*uintptr)(ret)))
			results[0] = val
		case Float:
			val := reflect.New(fnType.Out(0)).Elem()
			val.SetFloat(float64(*(*float32)(ret)))
			results[0] = val
		case Double:
			val := reflect.New(fnType.Out(0)).Elem()
			val.SetFloat(*(*float64)(ret))
			results[0] = val
		case Function:
			val := reflect.New(fnType.Out(0)).Elem()
			val.Set(reflect.ValueOf(GetDelegateForFunctionPointer(*(*unsafe.Pointer)(ret), fnType.Out(0))))
			results[0] = val
		case Vector2Type:
			if hasRet {
				results[0] = reflect.NewAt(fnType.Out(0), *(*unsafe.Pointer)(ret)).Elem()
			} else {
				results[0] = reflect.NewAt(fnType.Out(0), ret).Elem()
			}
		case Vector3Type:
			if hasRet {
				results[0] = reflect.NewAt(fnType.Out(0), *(*unsafe.Pointer)(ret)).Elem()
			} else {
				results[0] = reflect.NewAt(fnType.Out(0), ret).Elem()
			}
		case Vector4Type:
			if hasRet {
				results[0] = reflect.NewAt(fnType.Out(0), *(*unsafe.Pointer)(ret)).Elem()
			} else {
				results[0] = reflect.NewAt(fnType.Out(0), ret).Elem()
			}
		case Matrix4x4Type:
			if hasRet {
				results[0] = reflect.NewAt(fnType.Out(0), *(*unsafe.Pointer)(ret)).Elem()
			} else {
				results[0] = reflect.NewAt(fnType.Out(0), ret).Elem()
			}
		case String:
			val := reflect.New(fnType.Out(0)).Elem()
			val.SetString(GetStringData[string]((*PlgString)(*(*unsafe.Pointer)(ret))))
			results[0] = val
		case Any:
			results[0] = reflect.ValueOf(GetVariantData((*PlgVariant)(*(*unsafe.Pointer)(ret))))
		case ArrayBool:
			results[0] = getVectorDataBoolReturn((*PlgVector)(*(*unsafe.Pointer)(ret)), fnType.Out(0))
		case ArrayChar8:
			results[0] = getVectorDataChar8Return((*PlgVector)(*(*unsafe.Pointer)(ret)), fnType.Out(0))
		case ArrayChar16:
			results[0] = getVectorDataChar16Return((*PlgVector)(*(*unsafe.Pointer)(ret)), fnType.Out(0))
		case ArrayInt8:
			results[0] = getVectorDataInt8Return((*PlgVector)(*(*unsafe.Pointer)(ret)), fnType.Out(0))
		case ArrayInt16:
			results[0] = getVectorDataInt16Return((*PlgVector)(*(*unsafe.Pointer)(ret)), fnType.Out(0))
		case ArrayInt32:
			results[0] = getVectorDataInt32Return((*PlgVector)(*(*unsafe.Pointer)(ret)), fnType.Out(0))
		case ArrayInt64:
			results[0] = getVectorDataInt64Return((*PlgVector)(*(*unsafe.Pointer)(ret)), fnType.Out(0))
		case ArrayUInt8:
			results[0] = getVectorDataUInt8Return((*PlgVector)(*(*unsafe.Pointer)(ret)), fnType.Out(0))
		case ArrayUInt16:
			results[0] = getVectorDataUInt16Return((*PlgVector)(*(*unsafe.Pointer)(ret)), fnType.Out(0))
		case ArrayUInt32:
			results[0] = getVectorDataUInt32Return((*PlgVector)(*(*unsafe.Pointer)(ret)), fnType.Out(0))
		case ArrayUInt64:
			results[0] = getVectorDataUInt64Return((*PlgVector)(*(*unsafe.Pointer)(ret)), fnType.Out(0))
		case ArrayPointer:
			results[0] = getVectorDataPointerReturn((*PlgVector)(*(*unsafe.Pointer)(ret)), fnType.Out(0))
		case ArrayFloat:
			results[0] = getVectorDataFloatReturn((*PlgVector)(*(*unsafe.Pointer)(ret)), fnType.Out(0))
		case ArrayDouble:
			results[0] = getVectorDataDoubleReturn((*PlgVector)(*(*unsafe.Pointer)(ret)), fnType.Out(0))
		case ArrayString:
			results[0] = getVectorDataStringReturn((*PlgVector)(*(*unsafe.Pointer)(ret)), fnType.Out(0))
		case ArrayAny:
			results[0] = getVectorDataAnyReturn((*PlgVector)(*(*unsafe.Pointer)(ret)), fnType.Out(0))
		case ArrayVector2:
			results[0] = getVectorDataVector2Return((*PlgVector)(*(*unsafe.Pointer)(ret)), fnType.Out(0))
		case ArrayVector3:
			results[0] = getVectorDataVector3Return((*PlgVector)(*(*unsafe.Pointer)(ret)), fnType.Out(0))
		case ArrayVector4:
			results[0] = getVectorDataVector4Return((*PlgVector)(*(*unsafe.Pointer)(ret)), fnType.Out(0))
		case ArrayMatrix4x4:
			results[0] = getVectorDataMatrix4x4Return((*PlgVector)(*(*unsafe.Pointer)(ret)), fnType.Out(0))
		default:
			panicker(fmt.Sprintf("GetDelegateForFunctionPointer return type not supported %v", retType))
		}

		return results
	})

	fn := wrapper.Interface()

	fnArgs, fnRet := getFuncData(fnType)
	functionMap[fnPtr] = function{fn, addr, fnArgs, fnRet}

	return fn
}

func getFuncData(fnType reflect.Type) ([]funcArg, funcArg) {
	fnArgs := make([]funcArg, fnType.NumIn())
	for i := range fnArgs {
		argType := fnType.In(i)

		isRef := argType.Kind() == reflect.Pointer

		if isRef {
			argType = argType.Elem()
		}

		fnArgs[i] = funcArg{
			isRef: isRef,
			t:     argType,
		}
	}

	var fnRet funcArg
	numOut := fnType.NumOut()
	if numOut > 1 {
		panicker("getFuncData only one returned value supported")
	} else if numOut == 1 {
		out := fnType.Out(0)

		if out.Kind() == reflect.Pointer {
			fnRet.isRef = true
			out = out.Elem()
		}

		fnRet.t = out
	}

	return fnArgs, fnRet
}

func getArgumentPtr(p *C.Parameters, idx C.size_t) unsafe.Pointer {
	return unsafe.Pointer(uintptr(unsafe.Pointer(&p.arguments)) + uintptr(idx)*unsafe.Sizeof(uint64(0)))
}

/*func setArgument[T any](p *C.Parameters, idx C.size_t, val T) {
	ptr := getArgumentPtr(p, idx)
	**(**T)(ptr) = val
}

func setArgumentCopy(p *C.Parameters, idx C.size_t, val reflect.Value) {
	src := reflect.New(val.Type())
	src.Elem().Set(val)

	ptr := getArgumentPtr(p, idx)
	C.memcpy(*(*unsafe.Pointer)(ptr), src.UnsafePointer(), C.size_t(val.Type().Size()))
}*/

func getArgument[T any](p *C.Parameters, idx C.size_t) T {
	return *(*T)(getArgumentPtr(p, idx))
}

func getNonRefArgument(p *C.Parameters, idx C.size_t, t reflect.Type) reflect.Value {
	return reflect.NewAt(t, unsafe.Pointer((*uint64)(getArgumentPtr(p, idx)))).Elem()
}

func getNonRefVectorArgument(p *C.Parameters, idx C.size_t, t reflect.Type) reflect.Value {
	return reflect.NewAt(t, unsafe.Pointer(*(**uint64)(getArgumentPtr(p, idx)))).Elem()
}

func getRefArgument(p *C.Parameters, idx C.size_t, t reflect.Type) reflect.Value {
	return reflect.NewAt(t, unsafe.Pointer(*(**uint64)(getArgumentPtr(p, idx))))
}

func getReturnPtr(r *C.Return) unsafe.Pointer {
	return unsafe.Pointer(&r.ret)
}

func setReturn[T any](r *C.Return, val T) {
	ptr := getReturnPtr(r)
	*(*T)(ptr) = val
}

func setReturnCopy(r *C.Return, val reflect.Value) {
	src := reflect.New(val.Type())
	src.Elem().Set(val)

	ptr := getReturnPtr(r)
	C.memcpy(ptr, src.UnsafePointer(), C.size_t(val.Type().Size()))
}

func paramRefToObject(vt valueType, t reflect.Type, p *C.Parameters, i C.size_t) reflect.Value {
	switch vt {
	case Bool:
		return getRefArgument(p, i, t)
	case Char8:
		return getRefArgument(p, i, t)
	case Char16:
		return getRefArgument(p, i, t)
	case Int8:
		return getRefArgument(p, i, t)
	case Int16:
		return getRefArgument(p, i, t)
	case Int32:
		return getRefArgument(p, i, t)
	case Int64:
		return getRefArgument(p, i, t)
	case UInt8:
		return getRefArgument(p, i, t)
	case UInt16:
		return getRefArgument(p, i, t)
	case UInt32:
		return getRefArgument(p, i, t)
	case UInt64:
		return getRefArgument(p, i, t)
	case Pointer:
		return getRefArgument(p, i, t)
	case Float:
		return getRefArgument(p, i, t)
	case Double:
		return getRefArgument(p, i, t)
	case String:
		val := reflect.New(t)
		val.Elem().SetString(GetStringData[string](getArgument[*PlgString](p, i)))
		return val
	case Any:
		val := reflect.New(t)
		data := GetVariantData(getArgument[*PlgVariant](p, i))
		if data != nil {
			val.Elem().Set(reflect.ValueOf(data))
		}
		return val
	case ArrayBool:
		ref := reflect.New(t)
		ref.Elem().Set(getVectorDataChar8Reflect(getArgument[*PlgVector](p, i), t))
		return ref
	case ArrayChar8:
		ref := reflect.New(t)
		ref.Elem().Set(getVectorDataChar8Reflect(getArgument[*PlgVector](p, i), t))
		return ref
	case ArrayChar16:
		ref := reflect.New(t)
		ref.Elem().Set(getVectorDataChar16Reflect(getArgument[*PlgVector](p, i), t))
		return ref
	case ArrayInt8:
		ref := reflect.New(t)
		ref.Elem().Set(getVectorDataInt8Reflect(getArgument[*PlgVector](p, i), t))
		return ref
	case ArrayInt16:
		ref := reflect.New(t)
		ref.Elem().Set(getVectorDataInt16Reflect(getArgument[*PlgVector](p, i), t))
		return ref
	case ArrayInt32:
		ref := reflect.New(t)
		ref.Elem().Set(getVectorDataInt32Reflect(getArgument[*PlgVector](p, i), t))
		return ref
	case ArrayInt64:
		ref := reflect.New(t)
		ref.Elem().Set(getVectorDataInt64Reflect(getArgument[*PlgVector](p, i), t))
		return ref
	case ArrayUInt8:
		ref := reflect.New(t)
		ref.Elem().Set(getVectorDataUInt8Reflect(getArgument[*PlgVector](p, i), t))
		return ref
	case ArrayUInt16:
		ref := reflect.New(t)
		ref.Elem().Set(getVectorDataUInt16Reflect(getArgument[*PlgVector](p, i), t))
		return ref
	case ArrayUInt32:
		ref := reflect.New(t)
		ref.Elem().Set(getVectorDataUInt32Reflect(getArgument[*PlgVector](p, i), t))
		return ref
	case ArrayUInt64:
		ref := reflect.New(t)
		ref.Elem().Set(getVectorDataUInt64Reflect(getArgument[*PlgVector](p, i), t))
		return ref
	case ArrayPointer:
		ref := reflect.New(t)
		ref.Elem().Set(getVectorDataPointerReflect(getArgument[*PlgVector](p, i), t))
		return ref
	case ArrayFloat:
		ref := reflect.New(t)
		ref.Elem().Set(getVectorDataFloatReflect(getArgument[*PlgVector](p, i), t))
		return ref
	case ArrayDouble:
		ref := reflect.New(t)
		ref.Elem().Set(getVectorDataDoubleReflect(getArgument[*PlgVector](p, i), t))
		return ref
	case ArrayString:
		ref := reflect.New(t)
		ref.Elem().Set(getVectorDataStringReflect(getArgument[*PlgVector](p, i), t))
		return ref
	case ArrayAny:
		ref := reflect.New(t)
		ref.Elem().Set(getVectorDataAnyReflect(getArgument[*PlgVector](p, i), t))
		return ref
	case ArrayVector2:
		ref := reflect.New(t)
		ref.Elem().Set(getVectorDataVector2Reflect(getArgument[*PlgVector](p, i), t))
		return ref
	case ArrayVector3:
		ref := reflect.New(t)
		ref.Elem().Set(getVectorDataVector3Reflect(getArgument[*PlgVector](p, i), t))
		return ref
	case ArrayVector4:
		ref := reflect.New(t)
		ref.Elem().Set(getVectorDataVector4Reflect(getArgument[*PlgVector](p, i), t))
		return ref
	case ArrayMatrix4x4:
		ref := reflect.New(t)
		ref.Elem().Set(getVectorDataMatrix4x4Reflect(getArgument[*PlgVector](p, i), t))
		return ref
	case Vector2Type:
		return getRefArgument(p, i, t)
	case Vector3Type:
		return getRefArgument(p, i, t)
	case Vector4Type:
		return getRefArgument(p, i, t)
	case Matrix4x4Type:
		return getRefArgument(p, i, t)
	default:
		panicker(fmt.Sprintf("paramRefToObject unsupported enum type %v", vt))
		return reflect.ValueOf(nil)
	}
}

func paramToObject(vt valueType, t reflect.Type, p *C.Parameters, i C.size_t) reflect.Value {
	switch vt {
	case Bool:
		return getNonRefArgument(p, i, t)
	case Char8:
		return getNonRefArgument(p, i, t)
	case Char16:
		return getNonRefArgument(p, i, t)
	case Int8:
		return getNonRefArgument(p, i, t)
	case Int16:
		return getNonRefArgument(p, i, t)
	case Int32:
		return getNonRefArgument(p, i, t)
	case Int64:
		return getNonRefArgument(p, i, t)
	case UInt8:
		return getNonRefArgument(p, i, t)
	case UInt16:
		return getNonRefArgument(p, i, t)
	case UInt32:
		return getNonRefArgument(p, i, t)
	case UInt64:
		return getNonRefArgument(p, i, t)
	case Pointer:
		return getNonRefArgument(p, i, t)
	case Float:
		return getNonRefArgument(p, i, t)
	case Double:
		return getNonRefArgument(p, i, t)
	case Function:
		val := reflect.New(t).Elem()
		val.Set(reflect.ValueOf(GetDelegateForFunctionPointer(getArgument[unsafe.Pointer](p, i), t)))
		return val
	case String:
		val := reflect.New(t).Elem()
		val.SetString(GetStringData[string](getArgument[*PlgString](p, i)))
		return val
	case Any:
		val := reflect.New(t)
		data := GetVariantData(getArgument[*PlgVariant](p, i))
		if data != nil {
			val.Set(reflect.ValueOf(data))
		}
		return val
	case ArrayBool:
		return getVectorDataBoolReflect(getArgument[*PlgVector](p, i), t)
	case ArrayChar8:
		return getVectorDataChar8Reflect(getArgument[*PlgVector](p, i), t)
	case ArrayChar16:
		return getVectorDataChar16Reflect(getArgument[*PlgVector](p, i), t)
	case ArrayInt8:
		return getVectorDataInt8Reflect(getArgument[*PlgVector](p, i), t)
	case ArrayInt16:
		return getVectorDataInt16Reflect(getArgument[*PlgVector](p, i), t)
	case ArrayInt32:
		return getVectorDataInt32Reflect(getArgument[*PlgVector](p, i), t)
	case ArrayInt64:
		return getVectorDataInt64Reflect(getArgument[*PlgVector](p, i), t)
	case ArrayUInt8:
		return getVectorDataUInt8Reflect(getArgument[*PlgVector](p, i), t)
	case ArrayUInt16:
		return getVectorDataUInt16Reflect(getArgument[*PlgVector](p, i), t)
	case ArrayUInt32:
		return getVectorDataUInt32Reflect(getArgument[*PlgVector](p, i), t)
	case ArrayUInt64:
		return getVectorDataUInt64Reflect(getArgument[*PlgVector](p, i), t)
	case ArrayPointer:
		return getVectorDataPointerReflect(getArgument[*PlgVector](p, i), t)
	case ArrayFloat:
		return getVectorDataFloatReflect(getArgument[*PlgVector](p, i), t)
	case ArrayDouble:
		return getVectorDataDoubleReflect(getArgument[*PlgVector](p, i), t)
	case ArrayString:
		return getVectorDataStringReflect(getArgument[*PlgVector](p, i), t)
	case ArrayAny:
		return getVectorDataAnyReflect(getArgument[*PlgVector](p, i), t)
	case ArrayVector2:
		return getVectorDataVector2Reflect(getArgument[*PlgVector](p, i), t)
	case ArrayVector3:
		return getVectorDataVector3Reflect(getArgument[*PlgVector](p, i), t)
	case ArrayVector4:
		return getVectorDataVector3Reflect(getArgument[*PlgVector](p, i), t)
	case ArrayMatrix4x4:
		return getVectorDataMatrix4x4Reflect(getArgument[*PlgVector](p, i), t)
	case Vector2Type:
		return getNonRefVectorArgument(p, i, t)
	case Vector3Type:
		return getNonRefVectorArgument(p, i, t)
	case Vector4Type:
		return getNonRefVectorArgument(p, i, t)
	case Matrix4x4Type:
		return getNonRefVectorArgument(p, i, t)
	default:
		panicker(fmt.Sprintf("paramToObject unsupported type %v", vt))
	}
	return reflect.ValueOf(nil)
}

func setRefParam(vt valueType, p *C.Parameters, i C.size_t, val reflect.Value) {
	switch vt {
	case Bool:
	case Char8:
	case Char16:
	case Int8:
	case Int16:
	case Int32:
	case Int64:
	case UInt8:
	case UInt16:
	case UInt32:
	case UInt64:
	case Pointer:
	case Float:
	case Double:
	case String:
		AssignString(getArgument[*PlgString](p, i), val.Elem().String())
	case Any:
		AssignVariant(getArgument[*PlgVariant](p, i), val.Elem().Interface())
	case ArrayBool:
		reflectAssignVectorBool(getArgument[*PlgVector](p, i), val.Elem())
	case ArrayChar8:
		reflectAssignVectorChar8(getArgument[*PlgVector](p, i), val.Elem())
	case ArrayChar16:
		reflectAssignVectorChar16(getArgument[*PlgVector](p, i), val.Elem())
	case ArrayInt8:
		reflectAssignVectorInt8(getArgument[*PlgVector](p, i), val.Elem())
	case ArrayInt16:
		reflectAssignVectorInt16(getArgument[*PlgVector](p, i), val.Elem())
	case ArrayInt32:
		reflectAssignVectorInt32(getArgument[*PlgVector](p, i), val.Elem())
	case ArrayInt64:
		reflectAssignVectorInt64(getArgument[*PlgVector](p, i), val.Elem())
	case ArrayUInt8:
		reflectAssignVectorUInt8(getArgument[*PlgVector](p, i), val.Elem())
	case ArrayUInt16:
		reflectAssignVectorUInt16(getArgument[*PlgVector](p, i), val.Elem())
	case ArrayUInt32:
		reflectAssignVectorUInt32(getArgument[*PlgVector](p, i), val.Elem())
	case ArrayUInt64:
		reflectAssignVectorUInt64(getArgument[*PlgVector](p, i), val.Elem())
	case ArrayPointer:
		reflectAssignVectorPointer(getArgument[*PlgVector](p, i), val.Elem())
	case ArrayFloat:
		reflectAssignVectorFloat(getArgument[*PlgVector](p, i), val.Elem())
	case ArrayDouble:
		reflectAssignVectorDouble(getArgument[*PlgVector](p, i), val.Elem())
	case ArrayString:
		reflectAssignVectorString(getArgument[*PlgVector](p, i), val.Elem())
	case ArrayAny:
		reflectAssignVectorVariant(getArgument[*PlgVector](p, i), val.Elem())
	case ArrayVector2:
		reflectAssignVectorVector2(getArgument[*PlgVector](p, i), val.Elem())
	case ArrayVector3:
		reflectAssignVectorVector3(getArgument[*PlgVector](p, i), val.Elem())
	case ArrayVector4:
		reflectAssignVectorVector4(getArgument[*PlgVector](p, i), val.Elem())
	case ArrayMatrix4x4:
		reflectAssignVectorMatrix4x4(getArgument[*PlgVector](p, i), val.Elem())
	case Vector2Type:
	case Vector3Type:
	case Vector4Type:
	case Matrix4x4Type:
	default:
		panicker(fmt.Sprintf("setRefParam unsupported enum type %v", vt))
	}
}

func setObjReturn(vt valueType, r *C.Return, ret reflect.Value) {
	switch vt {
	case Void:
		// Do nothing
	case Bool:
		setReturn(r, ret.Bool())
	case Char8:
		setReturn(r, int8(ret.Int()))
	case Char16:
		setReturn(r, uint16(ret.Uint()))
	case Int8:
		setReturn(r, int8(ret.Int()))
	case Int16:
		setReturn(r, int16(ret.Int()))
	case Int32:
		setReturn(r, int32(ret.Int()))
	case Int64:
		setReturn(r, int64(ret.Int()))
	case UInt8:
		setReturn(r, uint8(ret.Uint()))
	case UInt16:
		setReturn(r, uint16(ret.Uint()))
	case UInt32:
		setReturn(r, uint32(ret.Uint()))
	case UInt64:
		setReturn(r, uint64(ret.Uint()))
	case Pointer:
		setReturn(r, uintptr(ret.Uint()))
	case Float:
		setReturn(r, float32(ret.Float()))
	case Double:
		setReturn(r, float64(ret.Float()))
	case Function:
		setReturn(r, GetFunctionPointerForDelegate(ret.Interface()))
	case String:
		setReturn(r, ConstructString(ret.String()))
	case Any:
		setReturn(r, ConstructVariant(ret.Interface()))
	case ArrayBool:
		setReturn(r, constructVectorBoolToSlice(ret))
	case ArrayChar8:
		setReturn(r, constructVectorChar8ToSlice(ret))
	case ArrayChar16:
		setReturn(r, constructVectorChar16ToSlice(ret))
	case ArrayInt8:
		setReturn(r, constructVectorInt8ToSlice(ret))
	case ArrayInt16:
		setReturn(r, constructVectorInt16ToSlice(ret))
	case ArrayInt32:
		setReturn(r, constructVectorInt32ToSlice(ret))
	case ArrayInt64:
		setReturn(r, constructVectorInt64ToSlice(ret))
	case ArrayUInt8:
		setReturn(r, constructVectorUInt8ToSlice(ret))
	case ArrayUInt16:
		setReturn(r, constructVectorUInt16ToSlice(ret))
	case ArrayUInt32:
		setReturn(r, constructVectorUInt32ToSlice(ret))
	case ArrayUInt64:
		setReturn(r, constructVectorUInt64ToSlice(ret))
	case ArrayPointer:
		setReturn(r, constructVectorPointerToSlice(ret))
	case ArrayFloat:
		setReturn(r, constructVectorFloatToSlice(ret))
	case ArrayDouble:
		setReturn(r, constructVectorDoubleToSlice(ret))
	case ArrayString:
		setReturn(r, constructVectorStringToSlice(ret))
	case ArrayAny:
		setReturn(r, constructVectorVariantToSlice(ret))
	case ArrayVector2:
		setReturn(r, constructVectorVector2ToSlice(ret))
	case ArrayVector3:
		setReturn(r, constructVectorVector3ToSlice(ret))
	case ArrayVector4:
		setReturn(r, constructVectorVector4ToSlice(ret))
	case ArrayMatrix4x4:
		setReturn(r, constructVectorMatrix4x4ToSlice(ret))
	case Vector2Type, Vector3Type, Vector4Type, Matrix4x4Type:
		setReturnCopy(r, ret)
	default:
		panicker(fmt.Sprintf("setReturn unsupported enum type %v", vt))
	}
}

func internalCall(m C.MethodHandle, data unsafe.Pointer, p *C.Parameters, count C.size_t, r *C.Return) {
	fn, ok := functionMap[data]
	if !ok {
		panicker(fmt.Sprintf("function %p not found", data))
	}

	fnValue := reflect.ValueOf(fn.fn)

	var args = make([]reflect.Value, int(count))

	/* fnType := reflect.TypeOf(fn.fn)

	numIn := fnType.NumIn()
	if numIn != int(count) {
		panicker(fmt.Sprintf("expected %d parameters, got %d", numIn, count))
	} */

	if len(fn.args) != int(count) {
		panicker(fmt.Sprintf("expected %d parameters, got %d", len(fn.args), int(count)))
	}

	for i := C.size_t(0); i < count; i++ {
		mt := C.Plugify_GetMethodParamType(m, C.ptrdiff_t(i))
		vt := valueType(mt.valueType)

		argType := fn.args[i].t

		if mt.ref {
			args[i] = paramRefToObject(vt, argType, p, i)
		} else {
			args[i] = paramToObject(vt, argType, p, i)
		}
	}

	rets := fnValue.Call(args)

	mt := C.Plugify_GetMethodParamType(m, C.ptrdiff_t(-1))
	vt := valueType(mt.valueType)

	if len(rets) > 0 {
		setObjReturn(vt, r, rets[0])
	}

	for i := C.size_t(0); i < count; i++ {
		mt = C.Plugify_GetMethodParamType(m, C.ptrdiff_t(i))
		vt = valueType(mt.valueType)
		if mt.ref {
			setRefParam(vt, p, i, args[i])
		}
	}
}

func GetFunctionPointerForDelegate(fn any) unsafe.Pointer {
	if fn == nil {
		return nil
	}

	mu.Lock()
	defer mu.Unlock()

	fnVal := reflect.ValueOf(fn)
	if fnVal.Kind() != reflect.Func {
		panicker("expected a function")
	}

	fnType := fnVal.Type()
	fnPtr := unsafe.Pointer(fnVal.Pointer())

	for _, v := range functionMap {
		if unsafe.Pointer(reflect.ValueOf(v.fn).Pointer()) == fnPtr {
			return v.addr
		}
	}

	name, err := getFunctionName(fnType)
	if err != nil {
		panicker(err)
	}

	callback := C.Plugify_NewCallback(C.PluginHandle(plg().handle), name, fnPtr)
	if callback == nil {
		panicker(fmt.Sprintf("%s (jit error: not found)", name))
	}

	callbacks = append(callbacks, callback)

	addr := C.Plugify_GetCallbackFunction(callback)
	if addr == nil {
		panicker(fmt.Sprintf("%s (jit error: %s)", name, string(C.GoString(C.Plugify_GetCallbackError(callback)))))
	}

	fnArgs, fnRet := getFuncData(fnType)
	functionMap[fnPtr] = function{fn, addr, fnArgs, fnRet}

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
		return plg().name, nil
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
