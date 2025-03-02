package plugify

/*
#include <plugify.h>
*/
import "C"
import "unsafe"

// String functions

func ConstructString(s string) C.String {
	return C.Plugify_ConstructString(s)
}

func DestroyString(s *C.String) {
	C.Plugify_DestroyString(s)
}

func GetStringData(s *C.String) string {
	return C.GoStringN(C.Plugify_GetStringData(s), C.int(C.Plugify_GetStringLength(s)))
}

func GetStringLength(s *C.String) C.ptrdiff_t {
	return C.Plugify_GetStringLength(s)
}

func AssignString(s *C.String, str string) {
	C.Plugify_AssignString(s, str)
}

// Variant functions

func GetVariantData(v *C.Variant) interface{} {
	switch ValueType(v.current) {
	case Invalid, Void:
		return nil
	case Bool:
		return *(*bool)(unsafe.Pointer(v))
	case Char8:
		return *(*int8)(unsafe.Pointer(v))
	case Char16:
		return *(*uint16)(unsafe.Pointer(v))
	case Int8:
		return *(*int8)(unsafe.Pointer(v))
	case Int16:
		return *(*int16)(unsafe.Pointer(v))
	case Int32:
		return *(*int32)(unsafe.Pointer(v))
	case Int64:
		return *(*int64)(unsafe.Pointer(v))
	case UInt8:
		return *(*uint8)(unsafe.Pointer(v))
	case UInt16:
		return *(*uint16)(unsafe.Pointer(v))
	case UInt32:
		return *(*uint32)(unsafe.Pointer(v))
	case UInt64:
		return *(*uint64)(unsafe.Pointer(v))
	case Pointer:
		return *(*uintptr)(unsafe.Pointer(v))
	case Float:
		return *(*float32)(unsafe.Pointer(v))
	case Double:
		return *(*float64)(unsafe.Pointer(v))
	case String:
		return GetStringData((*C.String)(unsafe.Pointer(v)))
	case ArrayBool:
		return GetVectorDataBool((*C.Vector)(unsafe.Pointer(v)))
	case ArrayChar8:
		return GetVectorDataChar8((*C.Vector)(unsafe.Pointer(v)))
	case ArrayChar16:
		return GetVectorDataChar16((*C.Vector)(unsafe.Pointer(v)))
	case ArrayInt8:
		return GetVectorDataInt8((*C.Vector)(unsafe.Pointer(v)))
	case ArrayInt16:
		return GetVectorDataInt16((*C.Vector)(unsafe.Pointer(v)))
	case ArrayInt32:
		return GetVectorDataInt32((*C.Vector)(unsafe.Pointer(v)))
	case ArrayInt64:
		return GetVectorDataInt64((*C.Vector)(unsafe.Pointer(v)))
	case ArrayUInt8:
		return GetVectorDataUInt8((*C.Vector)(unsafe.Pointer(v)))
	case ArrayUInt16:
		return GetVectorDataUInt16((*C.Vector)(unsafe.Pointer(v)))
	case ArrayUInt32:
		return GetVectorDataUInt32((*C.Vector)(unsafe.Pointer(v)))
	case ArrayUInt64:
		return GetVectorDataUInt64((*C.Vector)(unsafe.Pointer(v)))
	case ArrayPointer:
		return GetVectorDataPointer((*C.Vector)(unsafe.Pointer(v)))
	case ArrayFloat:
		return GetVectorDataFloat((*C.Vector)(unsafe.Pointer(v)))
	case ArrayDouble:
		return GetVectorDataDouble((*C.Vector)(unsafe.Pointer(v)))
	case ArrayString:
		return GetVectorDataString((*C.Vector)(unsafe.Pointer(v)))
	case ArrayVector2:
		return GetVectorDataVector2((*C.Vector)(unsafe.Pointer(v)))
	case ArrayVector3:
		return GetVectorDataVector3((*C.Vector)(unsafe.Pointer(v)))
	case ArrayVector4:
		return GetVectorDataVector4((*C.Vector)(unsafe.Pointer(v)))
	case ArrayMatrix4x4:
		return GetVectorDataMatrix4x4((*C.Vector)(unsafe.Pointer(v)))
	case Vector2Type:
		return *(*Vector2)(unsafe.Pointer(v))
	case Vector3Type:
		return *(*Vector3)(unsafe.Pointer(v))
	case Vector4Type:
		return *(*Vector4)(unsafe.Pointer(v))
	default:
		panic(NewTypeNotFoundException("Type not found"))
	}
}

func AssignVariant(v *C.Variant, param interface{}) {
	var valueType ValueType
	switch param.(type) {
	case nil:
		valueType = Invalid
	case bool:
		valueType = Bool
		*(*C.bool)(unsafe.Pointer(v)) = C.bool(param.(bool))
	/*case byte:
		valueType = Char8
		*(*C.int8_t)(unsafe.Pointer(v)) = C.int8_t(param.(byte))
	case rune:
		valueType = Char16
		*(*C.uint16_t)(unsafe.Pointer(v)) = C.uint16_t(param.(rune))*/
	case int8:
		valueType = Int8
		*(*C.int8_t)(unsafe.Pointer(v)) = C.int8_t(param.(int8))
	case int16:
		valueType = Int16
		*(*C.int16_t)(unsafe.Pointer(v)) = C.int16_t(param.(int16))
	case int32:
		valueType = Int32
		*(*C.int32_t)(unsafe.Pointer(v)) = C.int32_t(param.(int32))
	case int64:
		valueType = Int64
		*(*C.int64_t)(unsafe.Pointer(v)) = C.int64_t(param.(int64))
	case uint8:
		valueType = UInt8
		*(*C.uint8_t)(unsafe.Pointer(v)) = C.uint8_t(param.(uint8))
	case uint16:
		valueType = UInt16
		*(*C.uint16_t)(unsafe.Pointer(v)) = C.uint16_t(param.(uint16))
	case uint32:
		valueType = UInt32
		*(*C.uint32_t)(unsafe.Pointer(v)) = C.uint32_t(param.(uint32))
	case uint64:
		valueType = UInt64
		*(*C.uint64_t)(unsafe.Pointer(v)) = C.uint64_t(param.(uint64))
	case uintptr:
		valueType = Pointer
		*(*C.intptr_t)(unsafe.Pointer(v)) = C.intptr_t(param.(uintptr))
	case float32:
		valueType = Float
		*(*C.float)(unsafe.Pointer(v)) = C.float(param.(float32))
	case float64:
		valueType = Double
		*(*C.double)(unsafe.Pointer(v)) = C.double(param.(float64))
	case string:
		valueType = String
		*(*C.String)(unsafe.Pointer(v)) = C.Plugify_ConstructString(param.(string))
	case []bool:
		valueType = ArrayBool
		arr := param.([]bool)
		*(*C.Vector)(unsafe.Pointer(v)) = C.Plugify_ConstructVectorBool((*C.bool)(unsafe.Pointer(&arr[0])), C.ptrdiff_t(len(arr)))
	/*case []byte:
		valueType = ArrayChar8
		arr := param.([]byte)
		*(*C.Vector)(unsafe.Pointer(v)) = C.Plugify_ConstructVectorChar8((*C.int8_t)(unsafe.Pointer(&arr[0])), C.ptrdiff_t(len(arr)))
	case []rune:
		valueType = ArrayChar16
		arr := param.([]rune)
		*(*C.Vector)(unsafe.Pointer(v)) = C.Plugify_ConstructVectorChar16((*C.uint16_t)(unsafe.Pointer(&arr[0])), C.ptrdiff_t(len(arr)))*/
	case []int8:
		valueType = ArrayInt8
		arr := param.([]int8)
		*(*C.Vector)(unsafe.Pointer(v)) = C.Plugify_ConstructVectorInt8((*C.int8_t)(unsafe.Pointer(&arr[0])), C.ptrdiff_t(len(arr)))
	case []int16:
		valueType = ArrayInt16
		arr := param.([]int16)
		*(*C.Vector)(unsafe.Pointer(v)) = C.Plugify_ConstructVectorInt16((*C.int16_t)(unsafe.Pointer(&arr[0])), C.ptrdiff_t(len(arr)))
	case []int32:
		valueType = ArrayInt32
		arr := param.([]int32)
		*(*C.Vector)(unsafe.Pointer(v)) = C.Plugify_ConstructVectorInt32((*C.int32_t)(unsafe.Pointer(&arr[0])), C.ptrdiff_t(len(arr)))
	case []int64:
		valueType = ArrayInt64
		arr := param.([]int64)
		*(*C.Vector)(unsafe.Pointer(v)) = C.Plugify_ConstructVectorInt64((*C.int64_t)(unsafe.Pointer(&arr[0])), C.ptrdiff_t(len(arr)))
	case []uint8:
		valueType = ArrayUInt8
		arr := param.([]uint8)
		*(*C.Vector)(unsafe.Pointer(v)) = C.Plugify_ConstructVectorUInt8((*C.uint8_t)(unsafe.Pointer(&arr[0])), C.ptrdiff_t(len(arr)))
	case []uint16:
		valueType = ArrayUInt16
		arr := param.([]uint16)
		*(*C.Vector)(unsafe.Pointer(v)) = C.Plugify_ConstructVectorUInt16((*C.uint16_t)(unsafe.Pointer(&arr[0])), C.ptrdiff_t(len(arr)))
	case []uint32:
		valueType = ArrayUInt32
		arr := param.([]uint32)
		*(*C.Vector)(unsafe.Pointer(v)) = C.Plugify_ConstructVectorUInt32((*C.uint32_t)(unsafe.Pointer(&arr[0])), C.ptrdiff_t(len(arr)))
	case []uint64:
		valueType = ArrayUInt64
		arr := param.([]uint64)
		*(*C.Vector)(unsafe.Pointer(v)) = C.Plugify_ConstructVectorUInt64((*C.uint64_t)(unsafe.Pointer(&arr[0])), C.ptrdiff_t(len(arr)))
	case []uintptr:
		valueType = ArrayPointer
		arr := param.([]uintptr)
		*(*C.Vector)(unsafe.Pointer(v)) = C.Plugify_ConstructVectorPointer((*C.uintptr_t)(unsafe.Pointer(&arr[0])), C.ptrdiff_t(len(arr)))
	case []float32:
		valueType = ArrayFloat
		arr := param.([]float32)
		*(*C.Vector)(unsafe.Pointer(v)) = C.Plugify_ConstructVectorFloat((*C.float)(unsafe.Pointer(&arr[0])), C.ptrdiff_t(len(arr)))
	case []float64:
		valueType = ArrayDouble
		arr := param.([]float64)
		*(*C.Vector)(unsafe.Pointer(v)) = C.Plugify_ConstructVectorDouble((*C.double)(unsafe.Pointer(&arr[0])), C.ptrdiff_t(len(arr)))
	case []string:
		valueType = ArrayString
		arr := param.([]string)
		*(*C.Vector)(unsafe.Pointer(v)) = C.Plugify_ConstructVectorString(&arr[0], C.ptrdiff_t(len(arr)))
	case []Vector2:
		valueType = ArrayVector2
		arr := param.([]Vector2)
		*(*C.Vector)(unsafe.Pointer(v)) = C.Plugify_ConstructVectorVector2((*C.Vector2)(unsafe.Pointer(&arr[0])), C.ptrdiff_t(len(arr)))
	case []Vector3:
		valueType = ArrayVector3
		arr := param.([]Vector3)
		*(*C.Vector)(unsafe.Pointer(v)) = C.Plugify_ConstructVectorVector3((*C.Vector3)(unsafe.Pointer(&arr[0])), C.ptrdiff_t(len(arr)))
	case []Vector4:
		valueType = ArrayVector4
		arr := param.([]Vector4)
		*(*C.Vector)(unsafe.Pointer(v)) = C.Plugify_ConstructVectorVector4((*C.Vector4)(unsafe.Pointer(&arr[0])), C.ptrdiff_t(len(arr)))
	case []Matrix4x4:
		valueType = ArrayMatrix4x4
		arr := param.([]Matrix4x4)
		*(*C.Vector)(unsafe.Pointer(v)) = C.Plugify_ConstructVectorMatrix4x4((*C.Matrix4x4)(unsafe.Pointer(&arr[0])), C.ptrdiff_t(len(arr)))
	case Vector2:
		valueType = Vector2Type
		vec := param.(Vector2)
		*(*C.Vector2)(unsafe.Pointer(v)) = *(*C.Vector2)(unsafe.Pointer(&vec))
	case Vector3:
		valueType = Vector3Type
		vec := param.(Vector3)
		*(*C.Vector3)(unsafe.Pointer(v)) = *(*C.Vector3)(unsafe.Pointer(&vec))
	case Vector4:
		valueType = Vector4Type
		vec := param.(Vector4)
		*(*C.Vector4)(unsafe.Pointer(v)) = *(*C.Vector4)(unsafe.Pointer(&vec))
	default:
		panic(NewTypeNotFoundException("Type not found"))
	}

	v.current = C.uint8_t(valueType)
}

func ConstructVariant(v interface{}) C.Variant {
	var variant C.Variant
	AssignVariant(&variant, v)
	return variant
}

func DestroyVariant(v *C.Variant) {
	C.Plugify_DestroyVariant(v)
}

// Vector functions

func ConstructVectorBool[T ~bool](data []T) C.Vector {
	var ptr *C.bool
	if len(data) > 0 {
		ptr = (*C.bool)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	return C.Plugify_ConstructVectorBool(ptr, C.ptrdiff_t(len(data)))
}

func ConstructVectorChar8[T ~int8](data []T) C.Vector {
	var ptr *C.char
	if len(data) > 0 {
		ptr = (*C.char)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	return C.Plugify_ConstructVectorChar8(ptr, C.ptrdiff_t(len(data)))
}

func ConstructVectorChar16[T ~uint16](data []T) C.Vector {
	var ptr *C.char16_t
	if len(data) > 0 {
		ptr = (*C.char16_t)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	return C.Plugify_ConstructVectorChar16(ptr, C.ptrdiff_t(len(data)))
}

func ConstructVectorInt8[T ~int8](data []T) C.Vector {
	var ptr *C.int8_t
	if len(data) > 0 {
		ptr = (*C.int8_t)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	return C.Plugify_ConstructVectorInt8(ptr, C.ptrdiff_t(len(data)))
}

func ConstructVectorInt16[T ~int16](data []T) C.Vector {
	var ptr *C.int16_t
	if len(data) > 0 {
		ptr = (*C.int16_t)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	return C.Plugify_ConstructVectorInt16(ptr, C.ptrdiff_t(len(data)))
}

func ConstructVectorInt32[T ~int32](data []T) C.Vector {
	var ptr *C.int32_t
	if len(data) > 0 {
		ptr = (*C.int32_t)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	return C.Plugify_ConstructVectorInt32(ptr, C.ptrdiff_t(len(data)))
}

func ConstructVectorInt64[T ~int64](data []T) C.Vector {
	var ptr *C.int64_t
	if len(data) > 0 {
		ptr = (*C.int64_t)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	return C.Plugify_ConstructVectorInt64(ptr, C.ptrdiff_t(len(data)))
}

func ConstructVectorUInt8[T ~uint8](data []T) C.Vector {
	var ptr *C.uint8_t
	if len(data) > 0 {
		ptr = (*C.uint8_t)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	return C.Plugify_ConstructVectorUInt8(ptr, C.ptrdiff_t(len(data)))
}

func ConstructVectorUInt16[T ~uint16](data []T) C.Vector {
	var ptr *C.uint16_t
	if len(data) > 0 {
		ptr = (*C.uint16_t)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	return C.Plugify_ConstructVectorUInt16(ptr, C.ptrdiff_t(len(data)))
}

func ConstructVectorUInt32[T ~uint32](data []T) C.Vector {
	var ptr *C.uint32_t
	if len(data) > 0 {
		ptr = (*C.uint32_t)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	return C.Plugify_ConstructVectorUInt32(ptr, C.ptrdiff_t(len(data)))
}

func ConstructVectorUInt64[T ~uint64](data []T) C.Vector {
	var ptr *C.uint64_t
	if len(data) > 0 {
		ptr = (*C.uint64_t)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	return C.Plugify_ConstructVectorUInt64(ptr, C.ptrdiff_t(len(data)))
}

func ConstructVectorPointer(data []uintptr) C.Vector {
	var ptr *C.uintptr_t
	if len(data) > 0 {
		ptr = (*C.uintptr_t)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	return C.Plugify_ConstructVectorPointer(ptr, C.ptrdiff_t(len(data)))
}

func ConstructVectorFloat(data []float32) C.Vector {
	var ptr *C.float
	if len(data) > 0 {
		ptr = (*C.float)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	return C.Plugify_ConstructVectorFloat(ptr, C.ptrdiff_t(len(data)))
}

func ConstructVectorDouble(data []float64) C.Vector {
	var ptr *C.double
	if len(data) > 0 {
		ptr = (*C.double)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	return C.Plugify_ConstructVectorDouble(ptr, C.ptrdiff_t(len(data)))
}

func ConstructVectorString(data []string) C.Vector {
	var ptr *string
	if len(data) > 0 {
		ptr = &data[0]
	}
	return C.Plugify_ConstructVectorString(ptr, C.ptrdiff_t(len(data)))
}

func ConstructVectorVariant(arr []interface{}) C.Vector {
	vec := C.Plugify_ConstructVectorVariant(C.ptrdiff_t(len(arr)))
	AssignVectorVariant(&vec, arr)
	return vec
}

func ConstructVectorVector2(data []Vector2) C.Vector {
	var ptr *C.Vector2
	if len(data) > 0 {
		ptr = (*C.Vector2)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	return C.Plugify_ConstructVectorVector2(ptr, C.ptrdiff_t(len(data)))
}

func ConstructVectorVector3(data []Vector3) C.Vector {
	var ptr *C.Vector3
	if len(data) > 0 {
		ptr = (*C.Vector3)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	return C.Plugify_ConstructVectorVector3(ptr, C.ptrdiff_t(len(data)))
}

func ConstructVectorVector4(data []Vector4) C.Vector {
	var ptr *C.Vector4
	if len(data) > 0 {
		ptr = (*C.Vector4)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	return C.Plugify_ConstructVectorVector4(ptr, C.ptrdiff_t(len(data)))
}

func ConstructVectorMatrix4x4(data []Matrix4x4) C.Vector {
	var ptr *C.Matrix4x4
	if len(data) > 0 {
		ptr = (*C.Matrix4x4)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	return C.Plugify_ConstructVectorMatrix4x4(ptr, C.ptrdiff_t(len(data)))
}

func DestroyVectorBool(v *C.Vector) {
	C.Plugify_DestroyVectorBool(v)
}

func DestroyVectorChar8(v *C.Vector) {
	C.Plugify_DestroyVectorChar8(v)
}

func DestroyVectorChar16(v *C.Vector) {
	C.Plugify_DestroyVectorChar16(v)
}

func DestroyVectorInt8(v *C.Vector) {
	C.Plugify_DestroyVectorInt8(v)
}

func DestroyVectorInt16(v *C.Vector) {
	C.Plugify_DestroyVectorInt16(v)
}

func DestroyVectorInt32(v *C.Vector) {
	C.Plugify_DestroyVectorInt32(v)
}

func DestroyVectorInt64(v *C.Vector) {
	C.Plugify_DestroyVectorInt64(v)
}

func DestroyVectorUInt8(v *C.Vector) {
	C.Plugify_DestroyVectorUInt8(v)
}

func DestroyVectorUInt16(v *C.Vector) {
	C.Plugify_DestroyVectorUInt16(v)
}

func DestroyVectorUInt32(v *C.Vector) {
	C.Plugify_DestroyVectorUInt32(v)
}

func DestroyVectorUInt64(v *C.Vector) {
	C.Plugify_DestroyVectorUInt64(v)
}

func DestroyVectorPointer(v *C.Vector) {
	C.Plugify_DestroyVectorPointer(v)
}

func DestroyVectorFloat(v *C.Vector) {
	C.Plugify_DestroyVectorFloat(v)
}

func DestroyVectorDouble(v *C.Vector) {
	C.Plugify_DestroyVectorDouble(v)
}

func DestroyVectorString(v *C.Vector) {
	C.Plugify_DestroyVectorString(v)
}

func DestroyVectorVariant(v *C.Vector) {
	C.Plugify_DestroyVectorVariant(v)
}

func DestroyVectorVector2(v *C.Vector) {
	C.Plugify_DestroyVectorVector2(v)
}

func DestroyVectorVector3(v *C.Vector) {
	C.Plugify_DestroyVectorVector3(v)
}

func DestroyVectorVector4(v *C.Vector) {
	C.Plugify_DestroyVectorVector4(v)
}

func DestroyVectorMatrix4x4(v *C.Vector) {
	C.Plugify_DestroyVectorMatrix4x4(v)
}

func GetVectorSizeBool(v *C.Vector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeBool(v)
}

func GetVectorSizeChar8(v *C.Vector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeChar8(v)
}

func GetVectorSizeChar16(v *C.Vector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeChar16(v)
}

func GetVectorSizeInt8(v *C.Vector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeInt8(v)
}

func GetVectorSizeInt16(v *C.Vector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeInt16(v)
}

func GetVectorSizeInt32(v *C.Vector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeInt32(v)
}

func GetVectorSizeInt64(v *C.Vector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeInt64(v)
}

func GetVectorSizeUInt8(v *C.Vector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeUInt8(v)
}

func GetVectorSizeUInt16(v *C.Vector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeUInt16(v)
}

func GetVectorSizeUInt32(v *C.Vector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeUInt32(v)
}

func GetVectorSizeUInt64(v *C.Vector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeUInt64(v)
}

func GetVectorSizePointer(v *C.Vector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizePointer(v)
}

func GetVectorSizeFloat(v *C.Vector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeFloat(v)
}

func GetVectorSizeDouble(v *C.Vector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeDouble(v)
}

func GetVectorSizeString(v *C.Vector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeString(v)
}

func GetVectorSizeVariant(v *C.Vector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeVariant(v)
}

func GetVectorSizeVector2(v *C.Vector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeVector2(v)
}

func GetVectorSizeVector3(v *C.Vector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeVector3(v)
}

func GetVectorSizeVector4(v *C.Vector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeVector4(v)
}

func GetVectorSizeMatrix4x4(v *C.Vector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeMatrix4x4(v)
}

func GetVectorDataBool(v *C.Vector) []bool {
	size := int(C.Plugify_GetVectorSizeBool(v))
	arr := make([]bool, size)
	dataPtr := C.Plugify_GetVectorDataBool(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_bool)
	return arr
}

func GetVectorDataChar8(v *C.Vector) []int8 {
	size := int(C.Plugify_GetVectorSizeChar8(v))
	arr := make([]int8, size)
	dataPtr := C.Plugify_GetVectorDataChar8(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_char)
	return arr
}

func GetVectorDataChar16(v *C.Vector) []uint16 {
	size := int(C.Plugify_GetVectorSizeChar16(v))
	arr := make([]uint16, size)
	dataPtr := C.Plugify_GetVectorDataChar16(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_wchar_t)
	return arr
}

func GetVectorDataInt8(v *C.Vector) []int8 {
	size := int(C.Plugify_GetVectorSizeInt8(v))
	arr := make([]int8, size)
	dataPtr := C.Plugify_GetVectorDataInt8(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int8_t)
	return arr
}

func GetVectorDataInt16(v *C.Vector) []int16 {
	size := int(C.Plugify_GetVectorSizeInt16(v))
	arr := make([]int16, size)
	dataPtr := C.Plugify_GetVectorDataInt16(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int16_t)
	return arr
}

func GetVectorDataInt32(v *C.Vector) []int32 {
	size := int(C.Plugify_GetVectorSizeInt32(v))
	arr := make([]int32, size)
	dataPtr := C.Plugify_GetVectorDataInt32(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int32_t)
	return arr
}

func GetVectorDataInt64(v *C.Vector) []int64 {
	size := int(C.Plugify_GetVectorSizeInt64(v))
	arr := make([]int64, size)
	dataPtr := C.Plugify_GetVectorDataInt64(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int64_t)
	return arr
}

func GetVectorDataUInt8(v *C.Vector) []uint8 {
	size := int(C.Plugify_GetVectorSizeUInt8(v))
	arr := make([]uint8, size)
	dataPtr := C.Plugify_GetVectorDataUInt8(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint8_t)
	return arr
}

func GetVectorDataUInt16(v *C.Vector) []uint16 {
	size := int(C.Plugify_GetVectorSizeUInt16(v))
	arr := make([]uint16, size)
	dataPtr := C.Plugify_GetVectorDataUInt16(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint16_t)
	return arr
}

func GetVectorDataUInt32(v *C.Vector) []uint32 {
	size := int(C.Plugify_GetVectorSizeUInt32(v))
	arr := make([]uint32, size)
	dataPtr := C.Plugify_GetVectorDataUInt32(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint32_t)
	return arr
}

func GetVectorDataUInt64(v *C.Vector) []uint64 {
	size := int(C.Plugify_GetVectorSizeUInt64(v))
	arr := make([]uint64, size)
	dataPtr := C.Plugify_GetVectorDataUInt64(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint64_t)
	return arr
}

func GetVectorDataPointer(v *C.Vector) []uintptr {
	size := int(C.Plugify_GetVectorSizePointer(v))
	arr := make([]uintptr, size)
	dataPtr := C.Plugify_GetVectorDataPointer(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_intptr_t)
	return arr
}

func GetVectorDataFloat(v *C.Vector) []float32 {
	size := int(C.Plugify_GetVectorSizeFloat(v))
	arr := make([]float32, size)
	dataPtr := C.Plugify_GetVectorDataFloat(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_float)
	return arr
}

func GetVectorDataDouble(v *C.Vector) []float64 {
	size := int(C.Plugify_GetVectorSizeDouble(v))
	arr := make([]float64, size)
	dataPtr := C.Plugify_GetVectorDataDouble(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_double)
	return arr
}

func GetVectorDataString(v *C.Vector) []string {
	size := int(C.Plugify_GetVectorSizeString(v))
	arr := make([]string, size)
	dataPtr := unsafe.Pointer(C.Plugify_GetVectorDataString(v))
	for i := range arr {
		arr[i] = GetStringData((*C.String)(unsafe.Pointer(uintptr(dataPtr) + uintptr(i)*C.sizeof_String)))
	}
	return arr
}

func GetVectorDataVariant(v *C.Vector) []interface{} {
	size := int(C.Plugify_GetVectorSizeVariant(v))
	arr := make([]interface{}, size)
	for i := 0; i < size; i++ {
		variant := C.Plugify_GetVectorDataVariant(v, C.ptrdiff_t(size))
		arr[i] = GetVariantData(variant)
	}
	return arr
}

func GetVectorDataVector2(v *C.Vector) []Vector2 {
	size := int(C.Plugify_GetVectorSizeVector2(v))
	arr := make([]Vector2, size)
	dataPtr := C.Plugify_GetVectorDataVector2(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector2)
	return arr
}

func GetVectorDataVector3(v *C.Vector) []Vector3 {
	size := int(C.Plugify_GetVectorSizeVector3(v))
	arr := make([]Vector3, size)
	dataPtr := C.Plugify_GetVectorDataVector3(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector3)
	return arr
}

func GetVectorDataVector4(v *C.Vector) []Vector4 {
	size := int(C.Plugify_GetVectorSizeVector4(v))
	arr := make([]Vector4, size)
	dataPtr := C.Plugify_GetVectorDataVector4(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector4)
	return arr
}

func GetVectorDataMatrix4x4(v *C.Vector) []Matrix4x4 {
	size := int(C.Plugify_GetVectorSizeMatrix4x4(v))
	arr := make([]Matrix4x4, size)
	dataPtr := C.Plugify_GetVectorDataMatrix4x4(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Matrix4x4)
	return arr
}

func GetVectorDataBoolT[T ~bool](v *C.Vector) []T {
	size := int(C.Plugify_GetVectorSizeBool(v))
	arr := make([]T, size)
	dataPtr := C.Plugify_GetVectorDataBool(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_bool)
	return arr
}

func GetVectorDataChar8T[T ~int8](v *C.Vector) []T {
	size := int(C.Plugify_GetVectorSizeChar8(v))
	arr := make([]T, size)
	dataPtr := C.Plugify_GetVectorDataChar8(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_char)
	return arr
}

func GetVectorDataChar16T[T ~uint16](v *C.Vector) []T {
	size := int(C.Plugify_GetVectorSizeChar16(v))
	arr := make([]T, size)
	dataPtr := C.Plugify_GetVectorDataChar16(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_wchar_t)
	return arr
}

func GetVectorDataInt8T[T ~int8](v *C.Vector) []T {
	size := int(C.Plugify_GetVectorSizeInt8(v))
	arr := make([]T, size)
	dataPtr := C.Plugify_GetVectorDataInt8(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int8_t)
	return arr
}

func GetVectorDataInt16T[T ~int16](v *C.Vector) []T {
	size := int(C.Plugify_GetVectorSizeInt16(v))
	arr := make([]T, size)
	dataPtr := C.Plugify_GetVectorDataInt16(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int16_t)
	return arr
}

func GetVectorDataInt32T[T ~int32](v *C.Vector) []T {
	size := int(C.Plugify_GetVectorSizeInt32(v))
	arr := make([]T, size)
	dataPtr := C.Plugify_GetVectorDataInt32(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int32_t)
	return arr
}

func GetVectorDataInt64T[T ~int64](v *C.Vector) []T {
	size := int(C.Plugify_GetVectorSizeInt64(v))
	arr := make([]T, size)
	dataPtr := C.Plugify_GetVectorDataInt64(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int64_t)
	return arr
}

func GetVectorDataUInt8T[T ~uint8](v *C.Vector) []T {
	size := int(C.Plugify_GetVectorSizeUInt8(v))
	arr := make([]T, size)
	dataPtr := C.Plugify_GetVectorDataUInt8(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint8_t)
	return arr
}

func GetVectorDataUInt16T[T ~uint16](v *C.Vector) []T {
	size := int(C.Plugify_GetVectorSizeUInt16(v))
	arr := make([]T, size)
	dataPtr := C.Plugify_GetVectorDataUInt16(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint16_t)
	return arr
}

func GetVectorDataUInt32T[T ~uint32](v *C.Vector) []T {
	size := int(C.Plugify_GetVectorSizeUInt32(v))
	arr := make([]T, size)
	dataPtr := C.Plugify_GetVectorDataUInt32(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint32_t)
	return arr
}

func GetVectorDataUInt64T[T ~uint64](v *C.Vector) []T {
	size := int(C.Plugify_GetVectorSizeUInt64(v))
	arr := make([]T, size)
	dataPtr := C.Plugify_GetVectorDataUInt64(v)
	C.memcpy(unsafe.Pointer(&arr[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint64_t)
	return arr
}

//

func GetVectorDataBoolTo[T ~bool](v *C.Vector, arr *[]T) {
	size := int(C.Plugify_GetVectorSizeBool(v))
	if len(*arr) < size {
		*arr = make([]T, size)
	}
	dataPtr := C.Plugify_GetVectorDataBool(v)
	C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_bool)
}

func GetVectorDataChar8To[T ~int8](v *C.Vector, arr *[]T) {
	size := int(C.Plugify_GetVectorSizeChar8(v))
	if len(*arr) < size {
		*arr = make([]T, size)
	}
	dataPtr := C.Plugify_GetVectorDataChar8(v)
	C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_char)
}

func GetVectorDataChar16To[T ~uint16](v *C.Vector, arr *[]T) {
	size := int(C.Plugify_GetVectorSizeChar16(v))
	if len(*arr) < size {
		*arr = make([]T, size)
	}
	dataPtr := C.Plugify_GetVectorDataChar16(v)
	C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_wchar_t)
}

func GetVectorDataInt8To[T ~int8](v *C.Vector, arr *[]T) {
	size := int(C.Plugify_GetVectorSizeInt8(v))
	if len(*arr) < size {
		*arr = make([]T, size)
	}
	dataPtr := C.Plugify_GetVectorDataInt8(v)
	C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int8_t)
}

func GetVectorDataInt16To[T ~int16](v *C.Vector, arr *[]T) {
	size := int(C.Plugify_GetVectorSizeInt16(v))
	if len(*arr) < size {
		*arr = make([]T, size)
	}
	dataPtr := C.Plugify_GetVectorDataInt16(v)
	C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int16_t)
}

func GetVectorDataInt32To[T ~int32](v *C.Vector, arr *[]T) {
	size := int(C.Plugify_GetVectorSizeInt32(v))
	if len(*arr) < size {
		*arr = make([]T, size)
	}
	dataPtr := C.Plugify_GetVectorDataInt32(v)
	C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int32_t)
}

func GetVectorDataInt64To[T ~int64](v *C.Vector, arr *[]T) {
	size := int(C.Plugify_GetVectorSizeInt64(v))
	if len(*arr) < size {
		*arr = make([]T, size)
	}
	dataPtr := C.Plugify_GetVectorDataInt64(v)
	C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int64_t)
}

func GetVectorDataUInt8To[T ~uint8](v *C.Vector, arr *[]T) {
	size := int(C.Plugify_GetVectorSizeUInt8(v))
	if len(*arr) < size {
		*arr = make([]T, size)
	}
	dataPtr := C.Plugify_GetVectorDataUInt8(v)
	C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint8_t)
}

func GetVectorDataUInt16To[T ~uint16](v *C.Vector, arr *[]T) {
	size := int(C.Plugify_GetVectorSizeUInt16(v))
	if len(*arr) < size {
		*arr = make([]T, size)
	}
	dataPtr := C.Plugify_GetVectorDataUInt16(v)
	C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint16_t)
}

func GetVectorDataUInt32To[T ~uint32](v *C.Vector, arr *[]T) {
	size := int(C.Plugify_GetVectorSizeUInt32(v))
	if len(*arr) < size {
		*arr = make([]T, size)
	}
	dataPtr := C.Plugify_GetVectorDataUInt32(v)
	C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint32_t)
}

func GetVectorDataUInt64To[T ~uint64](v *C.Vector, arr *[]T) {
	size := int(C.Plugify_GetVectorSizeUInt64(v))
	if len(*arr) < size {
		*arr = make([]T, size)
	}
	dataPtr := C.Plugify_GetVectorDataUInt64(v)
	C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint64_t)
}

func GetVectorDataPointerTo(v *C.Vector, arr *[]uintptr) {
	size := int(C.Plugify_GetVectorSizePointer(v))
	if len(*arr) < size {
		*arr = make([]uintptr, size)
	}
	dataPtr := C.Plugify_GetVectorDataPointer(v)
	C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_intptr_t)
}

func GetVectorDataFloatTo(v *C.Vector, arr *[]float32) {
	size := int(C.Plugify_GetVectorSizeFloat(v))
	if len(*arr) < size {
		*arr = make([]float32, size)
	}
	dataPtr := C.Plugify_GetVectorDataFloat(v)
	C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_float)
}

func GetVectorDataDoubleTo(v *C.Vector, arr *[]float64) {
	size := int(C.Plugify_GetVectorSizeDouble(v))
	if len(*arr) < size {
		*arr = make([]float64, size)
	}
	dataPtr := C.Plugify_GetVectorDataDouble(v)
	C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_double)
}

func GetVectorDataStringTo(v *C.Vector, arr *[]string) {
	size := int(C.Plugify_GetVectorSizeString(v))
	if len(*arr) < size {
		*arr = make([]string, size)
	}
	dataPtr := unsafe.Pointer(C.Plugify_GetVectorDataString(v))
	for i := range *arr {
		(*arr)[i] = GetStringData((*C.String)(unsafe.Pointer(uintptr(dataPtr) + uintptr(i)*C.sizeof_String)))
	}
}

func GetVectorDataVariantTo(v *C.Vector, arr *[]interface{}) {
	size := int(C.Plugify_GetVectorSizeVariant(v))
	if len(*arr) < size {
		*arr = make([]interface{}, size)
	}
	for i := 0; i < size; i++ {
		variant := C.Plugify_GetVectorDataVariant(v, C.ptrdiff_t(size))
		(*arr)[i] = GetVariantData(variant)
	}
}

func GetVectorDataVector2To(v *C.Vector, arr *[]Vector2) {
	size := int(C.Plugify_GetVectorSizeVector2(v))
	if len(*arr) < size {
		*arr = make([]Vector2, size)
	}
	dataPtr := C.Plugify_GetVectorDataVector2(v)
	C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector2)
}

func GetVectorDataVector3To(v *C.Vector, arr *[]Vector3) {
	size := int(C.Plugify_GetVectorSizeVector3(v))
	if len(*arr) < size {
		*arr = make([]Vector3, size)
	}
	dataPtr := C.Plugify_GetVectorDataVector3(v)
	C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector3)
}

func GetVectorDataVector4To(v *C.Vector, arr *[]Vector4) {
	size := int(C.Plugify_GetVectorSizeVector4(v))
	if len(*arr) < size {
		*arr = make([]Vector4, size)
	}
	dataPtr := C.Plugify_GetVectorDataVector4(v)
	C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector4)
}

func GetVectorDataMatrix4x4To(v *C.Vector, arr *[]Matrix4x4) {
	size := int(C.Plugify_GetVectorSizeMatrix4x4(v))
	if len(*arr) < size {
		*arr = make([]Matrix4x4, size)
	}
	dataPtr := C.Plugify_GetVectorDataMatrix4x4(v)
	C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Matrix4x4)
}

func AssignVectorBool[T ~bool](v *C.Vector, data []T) {
	var ptr *C.bool
	if len(data) > 0 {
		ptr = (*C.bool)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	C.Plugify_AssignVectorBool(v, ptr, C.ptrdiff_t(len(data)))
}

func AssignVectorChar8[T ~int8](v *C.Vector, data []T) {
	var ptr *C.char
	if len(data) > 0 {
		ptr = (*C.char)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	C.Plugify_AssignVectorChar8(v, ptr, C.ptrdiff_t(len(data)))
}

func AssignVectorChar16[T ~uint16](v *C.Vector, data []T) {
	var ptr *C.char16_t
	if len(data) > 0 {
		ptr = (*C.char16_t)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	C.Plugify_AssignVectorChar16(v, ptr, C.ptrdiff_t(len(data)))
}

func AssignVectorInt8[T ~int8](v *C.Vector, data []T) {
	var ptr *C.int8_t
	if len(data) > 0 {
		ptr = (*C.int8_t)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	C.Plugify_AssignVectorInt8(v, ptr, C.ptrdiff_t(len(data)))
}

func AssignVectorInt16[T ~int16](v *C.Vector, data []T) {
	var ptr *C.int16_t
	if len(data) > 0 {
		ptr = (*C.int16_t)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	C.Plugify_AssignVectorInt16(v, ptr, C.ptrdiff_t(len(data)))
}

func AssignVectorInt32[T ~int32](v *C.Vector, data []T) {
	var ptr *C.int32_t
	if len(data) > 0 {
		ptr = (*C.int32_t)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	C.Plugify_AssignVectorInt32(v, ptr, C.ptrdiff_t(len(data)))
}

func AssignVectorInt64[T ~int64](v *C.Vector, data []T) {
	var ptr *C.int64_t
	if len(data) > 0 {
		ptr = (*C.int64_t)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	C.Plugify_AssignVectorInt64(v, ptr, C.ptrdiff_t(len(data)))
}

func AssignVectorUInt8[T ~uint8](v *C.Vector, data []T) {
	var ptr *C.uint8_t
	if len(data) > 0 {
		ptr = (*C.uint8_t)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	C.Plugify_AssignVectorUInt8(v, ptr, C.ptrdiff_t(len(data)))
}

func AssignVectorUInt16[T ~uint16](v *C.Vector, data []T) {
	var ptr *C.uint16_t
	if len(data) > 0 {
		ptr = (*C.uint16_t)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	C.Plugify_AssignVectorUInt16(v, ptr, C.ptrdiff_t(len(data)))
}

func AssignVectorUInt32[T ~uint32](v *C.Vector, data []T) {
	var ptr *C.uint32_t
	if len(data) > 0 {
		ptr = (*C.uint32_t)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	C.Plugify_AssignVectorUInt32(v, ptr, C.ptrdiff_t(len(data)))
}

func AssignVectorUInt64[T ~uint64](v *C.Vector, data []T) {
	var ptr *C.uint64_t
	if len(data) > 0 {
		ptr = (*C.uint64_t)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	C.Plugify_AssignVectorUInt64(v, ptr, C.ptrdiff_t(len(data)))
}

func AssignVectorPointer(v *C.Vector, data []uintptr) {
	var ptr *C.uintptr_t
	if len(data) > 0 {
		ptr = (*C.uintptr_t)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	C.Plugify_AssignVectorPointer(v, ptr, C.ptrdiff_t(len(data)))
}

func AssignVectorFloat(v *C.Vector, data []float32) {
	var ptr *C.float
	if len(data) > 0 {
		ptr = (*C.float)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	C.Plugify_AssignVectorFloat(v, ptr, C.ptrdiff_t(len(data)))
}

func AssignVectorDouble(v *C.Vector, data []float64) {
	var ptr *C.double
	if len(data) > 0 {
		ptr = (*C.double)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	C.Plugify_AssignVectorDouble(v, ptr, C.ptrdiff_t(len(data)))
}

func AssignVectorString(v *C.Vector, data []string) {
	var ptr *string
	if len(data) > 0 {
		ptr = &data[0]
	}
	C.Plugify_AssignVectorString(v, ptr, C.ptrdiff_t(len(data)))
}

func AssignVectorVariant(v *C.Vector, data []interface{}) {
	size := len(data)
	C.Plugify_AssignVectorVariant(v, C.ptrdiff_t(size))
	for i := 0; i < size; i++ {
		variant := C.Plugify_GetVectorDataVariant(v, C.ptrdiff_t(i))
		AssignVariant(variant, data[i])
	}
}

func AssignVectorVector2(v *C.Vector, data []Vector2) {
	var ptr *C.Vector2
	if len(data) > 0 {
		ptr = (*C.Vector2)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	C.Plugify_AssignVectorVector2(v, ptr, C.ptrdiff_t(len(data)))
}

func AssignVectorVector3(v *C.Vector, data []Vector3) {
	var ptr *C.Vector3
	if len(data) > 0 {
		ptr = (*C.Vector3)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	C.Plugify_AssignVectorVector3(v, ptr, C.ptrdiff_t(len(data)))
}

func AssignVectorVector4(v *C.Vector, data []Vector4) {
	var ptr *C.Vector4
	if len(data) > 0 {
		ptr = (*C.Vector4)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	C.Plugify_AssignVectorVector4(v, ptr, C.ptrdiff_t(len(data)))
}

func AssignVectorMatrix4x4(v *C.Vector, data []Matrix4x4) {
	var ptr *C.Matrix4x4
	if len(data) > 0 {
		ptr = (*C.Matrix4x4)(unsafe.Pointer(&data[0]))
	} else {
		ptr = nil
	}
	C.Plugify_AssignVectorMatrix4x4(v, ptr, C.ptrdiff_t(len(data)))
}
