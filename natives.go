package plugify

/*
#include "plugify.h"

*/
import "C"
import (
	"fmt"
	"unsafe"
)

type PlgString = C.String
type PlgVariant = C.Variant
type PlgVector = C.Vector
type PlgVector2 = C.Vector2
type PlgVector3 = C.Vector3
type PlgVector4 = C.Vector4
type PlgMatrix4x4 = C.Matrix4x4

// String functions

func ConstructString(s string) PlgString {
	return C.Plugify_ConstructString(s)
}

func DestroyString(s *PlgString) {
	C.Plugify_DestroyString(s)
}

func GetStringData(s *PlgString) string {
	return C.GoStringN(C.Plugify_GetStringData(s), C.int(C.Plugify_GetStringLength(s)))
}

func GetStringLength(s *PlgString) C.ptrdiff_t {
	return C.Plugify_GetStringLength(s)
}

func AssignString(s *PlgString, str string) {
	C.Plugify_AssignString(s, str)
}

// Variant functions

func GetVariantData(v *PlgVariant) any {
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
		return GetStringData((*PlgString)(unsafe.Pointer(v)))
	case ArrayBool:
		return GetVectorDataBool((*PlgVector)(unsafe.Pointer(v)))
	case ArrayChar8:
		return GetVectorDataChar8((*PlgVector)(unsafe.Pointer(v)))
	case ArrayChar16:
		return GetVectorDataChar16((*PlgVector)(unsafe.Pointer(v)))
	case ArrayInt8:
		return GetVectorDataInt8((*PlgVector)(unsafe.Pointer(v)))
	case ArrayInt16:
		return GetVectorDataInt16((*PlgVector)(unsafe.Pointer(v)))
	case ArrayInt32:
		return GetVectorDataInt32((*PlgVector)(unsafe.Pointer(v)))
	case ArrayInt64:
		return GetVectorDataInt64((*PlgVector)(unsafe.Pointer(v)))
	case ArrayUInt8:
		return GetVectorDataUInt8((*PlgVector)(unsafe.Pointer(v)))
	case ArrayUInt16:
		return GetVectorDataUInt16((*PlgVector)(unsafe.Pointer(v)))
	case ArrayUInt32:
		return GetVectorDataUInt32((*PlgVector)(unsafe.Pointer(v)))
	case ArrayUInt64:
		return GetVectorDataUInt64((*PlgVector)(unsafe.Pointer(v)))
	case ArrayPointer:
		return GetVectorDataPointer((*PlgVector)(unsafe.Pointer(v)))
	case ArrayFloat:
		return GetVectorDataFloat((*PlgVector)(unsafe.Pointer(v)))
	case ArrayDouble:
		return GetVectorDataDouble((*PlgVector)(unsafe.Pointer(v)))
	case ArrayString:
		return GetVectorDataString((*PlgVector)(unsafe.Pointer(v)))
	case ArrayVector2:
		return GetVectorDataVector2((*PlgVector)(unsafe.Pointer(v)))
	case ArrayVector3:
		return GetVectorDataVector3((*PlgVector)(unsafe.Pointer(v)))
	case ArrayVector4:
		return GetVectorDataVector4((*PlgVector)(unsafe.Pointer(v)))
	case ArrayMatrix4x4:
		return GetVectorDataMatrix4x4((*PlgVector)(unsafe.Pointer(v)))
	case Vector2Type:
		return *(*Vector2)(unsafe.Pointer(v))
	case Vector3Type:
		return *(*Vector3)(unsafe.Pointer(v))
	case Vector4Type:
		return *(*Vector4)(unsafe.Pointer(v))
	default:
		panicker(NewTypeNotFoundException("Type not found"))
		return nil
	}
}

func AssignVariant(v *PlgVariant, param any) {
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
	case int:
		valueType = Int64
		*(*C.int64_t)(unsafe.Pointer(v)) = C.int64_t(param.(int))
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
	case uint:
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
		*(*PlgString)(unsafe.Pointer(v)) = C.Plugify_ConstructString(param.(string))
	case []bool:
		valueType = ArrayBool
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorBool(param.([]bool))
	/*case []byte:
		valueType = ArrayChar8
		arr := param.([]byte)
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorChar8((*C.int8_t)(unsafe.Pointer(unsafe.SliceData(arr))), C.ptrdiff_t(len(arr)))
	case []rune:
		valueType = ArrayChar16
		arr := param.([]rune)
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorChar16((*C.uint16_t)(unsafe.Pointer(unsafe.SliceData(arr))), C.ptrdiff_t(len(arr)))*/
	case []int8:
		valueType = ArrayInt8
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorInt8(param.([]int8))
	case []int16:
		valueType = ArrayInt16
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorInt16(param.([]int16))
	case []int32:
		valueType = ArrayInt32
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorInt32(param.([]int32))
	case []int64:
		valueType = ArrayInt64
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorInt64(param.([]int64))
	case []int:
		valueType = ArrayInt64
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorInt(param.([]int))
	case []uint8:
		valueType = ArrayUInt8
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorUInt8(param.([]uint8))
	case []uint16:
		valueType = ArrayUInt16
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorUInt16(param.([]uint16))
	case []uint32:
		valueType = ArrayUInt32
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorUInt32(param.([]uint32))
	case []uint64:
		valueType = ArrayUInt64
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorUInt64(param.([]uint64))
	case []uint:
		valueType = ArrayUInt64
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorUInt(param.([]uint))
	case []uintptr:
		valueType = ArrayPointer
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorPointer(param.([]uintptr))
	case []float32:
		valueType = ArrayFloat
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorFloat(param.([]float32))
	case []float64:
		valueType = ArrayDouble
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorDouble(param.([]float64))
	case []string:
		valueType = ArrayString
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorString(param.([]string))
	case []Vector2:
		valueType = ArrayVector2
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorVector2(param.([]Vector2))
	case []Vector3:
		valueType = ArrayVector3
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorVector3(param.([]Vector3))
	case []Vector4:
		valueType = ArrayVector4
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorVector4(param.([]Vector4))
	case []Matrix4x4:
		valueType = ArrayMatrix4x4
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorMatrix4x4(param.([]Matrix4x4))
	case Vector2:
		valueType = Vector2Type
		vec := param.(Vector2)
		*(*PlgVector2)(unsafe.Pointer(v)) = *(*PlgVector2)(unsafe.Pointer(&vec))
	case Vector3:
		valueType = Vector3Type
		vec := param.(Vector3)
		*(*PlgVector3)(unsafe.Pointer(v)) = *(*PlgVector3)(unsafe.Pointer(&vec))
	case Vector4:
		valueType = Vector4Type
		vec := param.(Vector4)
		*(*PlgVector4)(unsafe.Pointer(v)) = *(*PlgVector4)(unsafe.Pointer(&vec))
	case Matrix4x4:
		valueType = Matrix4x4Type
		vec := param.(Matrix4x4)
		*(*PlgMatrix4x4)(unsafe.Pointer(v)) = *(*PlgMatrix4x4)(unsafe.Pointer(&vec))
	default:
		panicker(NewTypeNotFoundException(fmt.Sprintf("Type not found: %T", param)))
	}

	v.current = C.uint8_t(valueType)
}

func ConstructVariant(v any) PlgVariant {
	var variant PlgVariant
	AssignVariant(&variant, v)
	return variant
}

func DestroyVariant(v *PlgVariant) {
	C.Plugify_DestroyVariant(v)
}

// Vector functions

func ConstructVectorBool[T ~bool](data []T) PlgVector {
	return C.Plugify_ConstructVectorBool((*C.bool)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func ConstructVectorChar8[T ~int8](data []T) PlgVector {
	return C.Plugify_ConstructVectorChar8((*C.char)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func ConstructVectorChar16[T ~uint16](data []T) PlgVector {
	return C.Plugify_ConstructVectorChar16((*C.char16_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func ConstructVectorInt8[T ~int8](data []T) PlgVector {
	return C.Plugify_ConstructVectorInt8((*C.int8_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func ConstructVectorInt16[T ~int16](data []T) PlgVector {
	return C.Plugify_ConstructVectorInt16((*C.int16_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func ConstructVectorInt32[T ~int32](data []T) PlgVector {
	return C.Plugify_ConstructVectorInt32((*C.int32_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func ConstructVectorInt64[T ~int64](data []T) PlgVector {
	return C.Plugify_ConstructVectorInt64((*C.int64_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func ConstructVectorInt[T ~int](data []T) PlgVector {
	return C.Plugify_ConstructVectorInt64((*C.int64_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func ConstructVectorUInt8[T ~uint8](data []T) PlgVector {
	return C.Plugify_ConstructVectorUInt8((*C.uint8_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func ConstructVectorUInt16[T ~uint16](data []T) PlgVector {
	return C.Plugify_ConstructVectorUInt16((*C.uint16_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func ConstructVectorUInt32[T ~uint32](data []T) PlgVector {
	return C.Plugify_ConstructVectorUInt32((*C.uint32_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func ConstructVectorUInt64[T ~uint64](data []T) PlgVector {
	return C.Plugify_ConstructVectorUInt64((*C.uint64_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func ConstructVectorUInt[T ~uint](data []T) PlgVector {
	return C.Plugify_ConstructVectorUInt64((*C.uint64_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func ConstructVectorPointer(data []uintptr) PlgVector {
	return C.Plugify_ConstructVectorPointer((*C.uintptr_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func ConstructVectorFloat(data []float32) PlgVector {
	return C.Plugify_ConstructVectorFloat((*C.float)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func ConstructVectorDouble(data []float64) PlgVector {
	return C.Plugify_ConstructVectorDouble((*C.double)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func ConstructVectorString(data []string) PlgVector {
	//return C.Plugify_ConstructVectorString((*string)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
	cArray := C.malloc(C.size_t(len(data)) * C.size_t(unsafe.Sizeof(C.GoString_{})))
	defer C.free(cArray)
	arr := ([]C.GoString_)(unsafe.Slice((*C.GoString_)(cArray), len(data)))

	for i, s := range data {
		arr[i].p = (*C.char)(unsafe.Pointer(&[]byte(s)[0]))
		arr[i].n = C.ptrdiff_t(len(s))
	}

	return C.Plugify_ConstructVectorString((*string)(unsafe.Pointer(unsafe.SliceData(arr))), C.ptrdiff_t(len(data)))
}

func ConstructVectorVariant(arr []any) PlgVector {
	vec := C.Plugify_ConstructVectorVariant(C.ptrdiff_t(len(arr)))
	AssignVectorVariant(&vec, arr)
	return vec
}

func ConstructVectorVector2(data []Vector2) PlgVector {
	return C.Plugify_ConstructVectorVector2((*PlgVector2)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func ConstructVectorVector3(data []Vector3) PlgVector {
	return C.Plugify_ConstructVectorVector3((*PlgVector3)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func ConstructVectorVector4(data []Vector4) PlgVector {
	return C.Plugify_ConstructVectorVector4((*PlgVector4)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func ConstructVectorMatrix4x4(data []Matrix4x4) PlgVector {
	return C.Plugify_ConstructVectorMatrix4x4((*PlgMatrix4x4)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func DestroyVectorBool(v *PlgVector) {
	C.Plugify_DestroyVectorBool(v)
}

func DestroyVectorChar8(v *PlgVector) {
	C.Plugify_DestroyVectorChar8(v)
}

func DestroyVectorChar16(v *PlgVector) {
	C.Plugify_DestroyVectorChar16(v)
}

func DestroyVectorInt8(v *PlgVector) {
	C.Plugify_DestroyVectorInt8(v)
}

func DestroyVectorInt16(v *PlgVector) {
	C.Plugify_DestroyVectorInt16(v)
}

func DestroyVectorInt32(v *PlgVector) {
	C.Plugify_DestroyVectorInt32(v)
}

func DestroyVectorInt64(v *PlgVector) {
	C.Plugify_DestroyVectorInt64(v)
}

func DestroyVectorUInt8(v *PlgVector) {
	C.Plugify_DestroyVectorUInt8(v)
}

func DestroyVectorUInt16(v *PlgVector) {
	C.Plugify_DestroyVectorUInt16(v)
}

func DestroyVectorUInt32(v *PlgVector) {
	C.Plugify_DestroyVectorUInt32(v)
}

func DestroyVectorUInt64(v *PlgVector) {
	C.Plugify_DestroyVectorUInt64(v)
}

func DestroyVectorPointer(v *PlgVector) {
	C.Plugify_DestroyVectorPointer(v)
}

func DestroyVectorFloat(v *PlgVector) {
	C.Plugify_DestroyVectorFloat(v)
}

func DestroyVectorDouble(v *PlgVector) {
	C.Plugify_DestroyVectorDouble(v)
}

func DestroyVectorString(v *PlgVector) {
	C.Plugify_DestroyVectorString(v)
}

func DestroyVectorVariant(v *PlgVector) {
	C.Plugify_DestroyVectorVariant(v)
}

func DestroyVectorVector2(v *PlgVector) {
	C.Plugify_DestroyVectorVector2(v)
}

func DestroyVectorVector3(v *PlgVector) {
	C.Plugify_DestroyVectorVector3(v)
}

func DestroyVectorVector4(v *PlgVector) {
	C.Plugify_DestroyVectorVector4(v)
}

func DestroyVectorMatrix4x4(v *PlgVector) {
	C.Plugify_DestroyVectorMatrix4x4(v)
}

func GetVectorSizeBool(v *PlgVector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeBool(v)
}

func GetVectorSizeChar8(v *PlgVector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeChar8(v)
}

func GetVectorSizeChar16(v *PlgVector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeChar16(v)
}

func GetVectorSizeInt8(v *PlgVector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeInt8(v)
}

func GetVectorSizeInt16(v *PlgVector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeInt16(v)
}

func GetVectorSizeInt32(v *PlgVector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeInt32(v)
}

func GetVectorSizeInt64(v *PlgVector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeInt64(v)
}

func GetVectorSizeUInt8(v *PlgVector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeUInt8(v)
}

func GetVectorSizeUInt16(v *PlgVector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeUInt16(v)
}

func GetVectorSizeUInt32(v *PlgVector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeUInt32(v)
}

func GetVectorSizeUInt64(v *PlgVector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeUInt64(v)
}

func GetVectorSizePointer(v *PlgVector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizePointer(v)
}

func GetVectorSizeFloat(v *PlgVector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeFloat(v)
}

func GetVectorSizeDouble(v *PlgVector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeDouble(v)
}

func GetVectorSizeString(v *PlgVector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeString(v)
}

func GetVectorSizeVariant(v *PlgVector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeVariant(v)
}

func GetVectorSizeVector2(v *PlgVector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeVector2(v)
}

func GetVectorSizeVector3(v *PlgVector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeVector3(v)
}

func GetVectorSizeVector4(v *PlgVector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeVector4(v)
}

func GetVectorSizeMatrix4x4(v *PlgVector) C.ptrdiff_t {
	return C.Plugify_GetVectorSizeMatrix4x4(v)
}

func GetVectorDataBool(v *PlgVector) []bool {
	size := int(C.Plugify_GetVectorSizeBool(v))
	arr := make([]bool, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataBool(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_bool)
	}
	return arr
}

func GetVectorDataChar8(v *PlgVector) []int8 {
	size := int(C.Plugify_GetVectorSizeChar8(v))
	arr := make([]int8, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataChar8(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_char)
	}
	return arr
}

func GetVectorDataChar16(v *PlgVector) []uint16 {
	size := int(C.Plugify_GetVectorSizeChar16(v))
	arr := make([]uint16, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataChar16(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_wchar_t)
	}
	return arr
}

func GetVectorDataInt8(v *PlgVector) []int8 {
	size := int(C.Plugify_GetVectorSizeInt8(v))
	arr := make([]int8, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt8(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int8_t)
	}
	return arr
}

func GetVectorDataInt16(v *PlgVector) []int16 {
	size := int(C.Plugify_GetVectorSizeInt16(v))
	arr := make([]int16, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt16(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int16_t)
	}
	return arr
}

func GetVectorDataInt32(v *PlgVector) []int32 {
	size := int(C.Plugify_GetVectorSizeInt32(v))
	arr := make([]int32, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt32(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int32_t)
	}
	return arr
}

func GetVectorDataInt64(v *PlgVector) []int64 {
	size := int(C.Plugify_GetVectorSizeInt64(v))
	arr := make([]int64, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt64(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int64_t)
	}
	return arr
}

func GetVectorDataUInt8(v *PlgVector) []uint8 {
	size := int(C.Plugify_GetVectorSizeUInt8(v))
	arr := make([]uint8, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt8(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint8_t)
	}
	return arr
}

func GetVectorDataUInt16(v *PlgVector) []uint16 {
	size := int(C.Plugify_GetVectorSizeUInt16(v))
	arr := make([]uint16, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt16(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint16_t)
	}
	return arr
}

func GetVectorDataUInt32(v *PlgVector) []uint32 {
	size := int(C.Plugify_GetVectorSizeUInt32(v))
	arr := make([]uint32, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt32(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint32_t)
	}
	return arr
}

func GetVectorDataUInt64(v *PlgVector) []uint64 {
	size := int(C.Plugify_GetVectorSizeUInt64(v))
	arr := make([]uint64, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt64(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint64_t)
	}
	return arr
}

func GetVectorDataPointer(v *PlgVector) []uintptr {
	size := int(C.Plugify_GetVectorSizePointer(v))
	fmt.Println(size)
	arr := make([]uintptr, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataPointer(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_intptr_t)
	}
	return arr
}

func GetVectorDataFloat(v *PlgVector) []float32 {
	size := int(C.Plugify_GetVectorSizeFloat(v))
	arr := make([]float32, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataFloat(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_float)
	}
	return arr
}

func GetVectorDataDouble(v *PlgVector) []float64 {
	size := int(C.Plugify_GetVectorSizeDouble(v))
	arr := make([]float64, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataDouble(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_double)
	}
	return arr
}

func GetVectorDataString(v *PlgVector) []string {
	size := int(C.Plugify_GetVectorSizeString(v))
	arr := make([]string, size)
	dataPtr := unsafe.Pointer(C.Plugify_GetVectorDataString(v))
	for i := range arr {
		arr[i] = GetStringData((*PlgString)(unsafe.Pointer(uintptr(dataPtr) + uintptr(i)*C.sizeof_String)))
	}
	return arr
}

func GetVectorDataVariant(v *PlgVector) []any {
	size := int(C.Plugify_GetVectorSizeVariant(v))
	arr := make([]any, size)
	for i := 0; i < size; i++ {
		variant := C.Plugify_GetVectorDataVariant(v, C.ptrdiff_t(i))
		arr[i] = GetVariantData(variant)
	}
	return arr
}
func GetVectorDataVector2(v *PlgVector) []Vector2 {
	size := int(C.Plugify_GetVectorSizeVector2(v))
	arr := make([]Vector2, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataVector2(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector2)
	}
	return arr
}

func GetVectorDataVector3(v *PlgVector) []Vector3 {
	size := int(C.Plugify_GetVectorSizeVector3(v))
	arr := make([]Vector3, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataVector3(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector3)
	}
	return arr
}

func GetVectorDataVector4(v *PlgVector) []Vector4 {
	size := int(C.Plugify_GetVectorSizeVector4(v))
	arr := make([]Vector4, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataVector4(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector4)
	}
	return arr
}

func GetVectorDataMatrix4x4(v *PlgVector) []Matrix4x4 {
	size := int(C.Plugify_GetVectorSizeMatrix4x4(v))
	arr := make([]Matrix4x4, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataMatrix4x4(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Matrix4x4)
	}
	return arr
}

func GetVectorDataBoolT[T ~bool](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeBool(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataBool(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_bool)
	}
	return arr
}

func GetVectorDataChar8T[T ~int8](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeChar8(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataChar8(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_char)
	}
	return arr
}

func GetVectorDataChar16T[T ~uint16](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeChar16(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataChar16(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_wchar_t)
	}
	return arr
}

func GetVectorDataInt8T[T ~int8](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeInt8(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt8(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int8_t)
	}
	return arr
}

func GetVectorDataInt16T[T ~int16](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeInt16(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt16(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int16_t)
	}
	return arr
}

func GetVectorDataInt32T[T ~int32](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeInt32(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt32(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int32_t)
	}
	return arr
}

func GetVectorDataInt64T[T ~int64](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeInt64(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt64(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int64_t)
	}
	return arr
}

func GetVectorDataUInt8T[T ~uint8](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeUInt8(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt8(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint8_t)
	}
	return arr
}

func GetVectorDataUInt16T[T ~uint16](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeUInt16(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt16(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint16_t)
	}
	return arr
}

func GetVectorDataUInt32T[T ~uint32](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeUInt32(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt32(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint32_t)
	}
	return arr
}

func GetVectorDataUInt64T[T ~uint64](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeUInt64(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt64(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint64_t)
	}
	return arr
}

func GetVectorDataBoolTo[T ~bool](v *PlgVector, arr *[]T) {
	size := int(C.Plugify_GetVectorSizeBool(v))
	if len(*arr) < size {
		*arr = make([]T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataBool(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_bool)
	}
}

func GetVectorDataChar8To[T ~int8](v *PlgVector, arr *[]T) {
	size := int(C.Plugify_GetVectorSizeChar8(v))
	if len(*arr) < size {
		*arr = make([]T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataChar8(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_char)
	}
}

func GetVectorDataChar16To[T ~uint16](v *PlgVector, arr *[]T) {
	size := int(C.Plugify_GetVectorSizeChar16(v))
	if len(*arr) < size {
		*arr = make([]T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataChar16(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_wchar_t)
	}
}

func GetVectorDataInt8To[T ~int8](v *PlgVector, arr *[]T) {
	size := int(C.Plugify_GetVectorSizeInt8(v))
	if len(*arr) < size {
		*arr = make([]T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt8(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int8_t)
	}
}

func GetVectorDataInt16To[T ~int16](v *PlgVector, arr *[]T) {
	size := int(C.Plugify_GetVectorSizeInt16(v))
	if len(*arr) < size {
		*arr = make([]T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt16(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int16_t)
	}
}

func GetVectorDataInt32To[T ~int32](v *PlgVector, arr *[]T) {
	size := int(C.Plugify_GetVectorSizeInt32(v))
	if len(*arr) < size {
		*arr = make([]T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt32(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int32_t)
	}
}

func GetVectorDataInt64To[T ~int64](v *PlgVector, arr *[]T) {
	size := int(C.Plugify_GetVectorSizeInt64(v))
	if len(*arr) < size {
		*arr = make([]T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt64(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int64_t)
	}
}

func GetVectorDataUInt8To[T ~uint8](v *PlgVector, arr *[]T) {
	size := int(C.Plugify_GetVectorSizeUInt8(v))
	if len(*arr) < size {
		*arr = make([]T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt8(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint8_t)
	}
}

func GetVectorDataUInt16To[T ~uint16](v *PlgVector, arr *[]T) {
	size := int(C.Plugify_GetVectorSizeUInt16(v))
	if len(*arr) < size {
		*arr = make([]T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt16(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint16_t)
	}
}

func GetVectorDataUInt32To[T ~uint32](v *PlgVector, arr *[]T) {
	size := int(C.Plugify_GetVectorSizeUInt32(v))
	if len(*arr) < size {
		*arr = make([]T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt32(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint32_t)
	}
}

func GetVectorDataUInt64To[T ~uint64](v *PlgVector, arr *[]T) {
	size := int(C.Plugify_GetVectorSizeUInt64(v))
	if len(*arr) < size {
		*arr = make([]T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt64(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint64_t)
	}
}

func GetVectorDataPointerTo(v *PlgVector, arr *[]uintptr) {
	size := int(C.Plugify_GetVectorSizePointer(v))
	if len(*arr) < size {
		*arr = make([]uintptr, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataPointer(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_intptr_t)
	}
}

func GetVectorDataFloatTo(v *PlgVector, arr *[]float32) {
	size := int(C.Plugify_GetVectorSizeFloat(v))
	if len(*arr) < size {
		*arr = make([]float32, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataFloat(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_float)
	}
}

func GetVectorDataDoubleTo(v *PlgVector, arr *[]float64) {
	size := int(C.Plugify_GetVectorSizeDouble(v))
	if len(*arr) < size {
		*arr = make([]float64, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataDouble(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_double)
	}
}

func GetVectorDataStringTo(v *PlgVector, arr *[]string) {
	size := int(C.Plugify_GetVectorSizeString(v))
	if len(*arr) < size {
		*arr = make([]string, size)
	}
	dataPtr := unsafe.Pointer(C.Plugify_GetVectorDataString(v))
	for i := range *arr {
		(*arr)[i] = GetStringData((*PlgString)(unsafe.Pointer(uintptr(dataPtr) + uintptr(i)*C.sizeof_String)))
	}
}

func GetVectorDataVariantTo(v *PlgVector, arr *[]any) {
	size := int(C.Plugify_GetVectorSizeVariant(v))
	if len(*arr) < size {
		*arr = make([]any, size)
	}
	for i := 0; i < size; i++ {
		variant := C.Plugify_GetVectorDataVariant(v, C.ptrdiff_t(i))
		(*arr)[i] = GetVariantData(variant)
	}
}

func GetVectorDataVector2To(v *PlgVector, arr *[]Vector2) {
	size := int(C.Plugify_GetVectorSizeVector2(v))
	if len(*arr) < size {
		*arr = make([]Vector2, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataVector2(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector2)
	}
}

func GetVectorDataVector3To(v *PlgVector, arr *[]Vector3) {
	size := int(C.Plugify_GetVectorSizeVector3(v))
	if len(*arr) < size {
		*arr = make([]Vector3, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataVector3(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector3)
	}
}

func GetVectorDataVector4To(v *PlgVector, arr *[]Vector4) {
	size := int(C.Plugify_GetVectorSizeVector4(v))
	if len(*arr) < size {
		*arr = make([]Vector4, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataVector4(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector4)
	}
}

func GetVectorDataMatrix4x4To(v *PlgVector, arr *[]Matrix4x4) {
	size := int(C.Plugify_GetVectorSizeMatrix4x4(v))
	if len(*arr) < size {
		*arr = make([]Matrix4x4, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataMatrix4x4(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Matrix4x4)
	}
}

func AssignVectorBool[T ~bool](v *PlgVector, data []T) {
	C.Plugify_AssignVectorBool(v, (*C.bool)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorChar8[T ~int8](v *PlgVector, data []T) {
	C.Plugify_AssignVectorChar8(v, (*C.char)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorChar16[T ~uint16](v *PlgVector, data []T) {
	C.Plugify_AssignVectorChar16(v, (*C.char16_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorInt8[T ~int8](v *PlgVector, data []T) {
	C.Plugify_AssignVectorInt8(v, (*C.int8_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorInt16[T ~int16](v *PlgVector, data []T) {
	C.Plugify_AssignVectorInt16(v, (*C.int16_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorInt32[T ~int32](v *PlgVector, data []T) {
	C.Plugify_AssignVectorInt32(v, (*C.int32_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorInt64[T ~int64](v *PlgVector, data []T) {
	C.Plugify_AssignVectorInt64(v, (*C.int64_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorUInt8[T ~uint8](v *PlgVector, data []T) {
	C.Plugify_AssignVectorUInt8(v, (*C.uint8_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorUInt16[T ~uint16](v *PlgVector, data []T) {
	C.Plugify_AssignVectorUInt16(v, (*C.uint16_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorUInt32[T ~uint32](v *PlgVector, data []T) {
	C.Plugify_AssignVectorUInt32(v, (*C.uint32_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorUInt64[T ~uint64](v *PlgVector, data []T) {
	C.Plugify_AssignVectorUInt64(v, (*C.uint64_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorPointer(v *PlgVector, data []uintptr) {
	C.Plugify_AssignVectorPointer(v, (*C.uintptr_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorFloat(v *PlgVector, data []float32) {
	C.Plugify_AssignVectorFloat(v, (*C.float)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorDouble(v *PlgVector, data []float64) {
	C.Plugify_AssignVectorDouble(v, (*C.double)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorString(v *PlgVector, data []string) {
	//C.Plugify_AssignVectorString(v, (*string)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
	cArray := C.malloc(C.size_t(len(data)) * C.size_t(unsafe.Sizeof(C.GoString_{})))
	defer C.free(cArray)
	arr := ([]C.GoString_)(unsafe.Slice((*C.GoString_)(cArray), len(data)))

	for i, s := range data {
		arr[i].p = (*C.char)(unsafe.Pointer(&[]byte(s)[0]))
		arr[i].n = C.ptrdiff_t(len(s))
	}

	C.Plugify_AssignVectorString(v, (*string)(unsafe.Pointer(unsafe.SliceData(arr))), C.ptrdiff_t(len(data)))
}

func AssignVectorVariant(v *PlgVector, data []any) {
	size := len(data)
	C.Plugify_AssignVectorVariant(v, C.ptrdiff_t(size))
	for i := 0; i < size; i++ {
		variant := C.Plugify_GetVectorDataVariant(v, C.ptrdiff_t(i))
		AssignVariant(variant, data[i])
	}
}

func AssignVectorVector2(v *PlgVector, data []Vector2) {
	C.Plugify_AssignVectorVector2(v, (*PlgVector2)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorVector3(v *PlgVector, data []Vector3) {
	C.Plugify_AssignVectorVector3(v, (*PlgVector3)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorVector4(v *PlgVector, data []Vector4) {
	C.Plugify_AssignVectorVector4(v, (*PlgVector4)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorMatrix4x4(v *PlgVector, data []Matrix4x4) {
	C.Plugify_AssignVectorMatrix4x4(v, (*PlgMatrix4x4)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}
