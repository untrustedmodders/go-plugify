package plugify

/*
#include "plugify.h"
*/
import "C"
import (
	"fmt"
	"reflect"
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

func ConstructString[T ~string](s T) PlgString {
	return C.Plugify_ConstructString(string(s))
}

func DestroyString(s *PlgString) {
	C.Plugify_DestroyString(s)
}

func GetStringData[T ~string](s *PlgString) T {
	return T(C.GoStringN(C.Plugify_GetStringData(s), C.int(C.Plugify_GetStringLength(s))))
}

func GetStringLength(s *PlgString) C.ptrdiff_t {
	return C.Plugify_GetStringLength(s)
}

func AssignString[T ~string](s *PlgString, str T) {
	C.Plugify_AssignString(s, string(str))
}

// Variant functions
func GetVariantData(v *PlgVariant) any {
	switch valueType(v.current) {
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
		return GetStringData[string]((*PlgString)(unsafe.Pointer(v)))
	case ArrayBool:
		return GetVectorDataBool[bool]((*PlgVector)(unsafe.Pointer(v)))
	case ArrayChar8:
		return GetVectorDataChar8[int8]((*PlgVector)(unsafe.Pointer(v)))
	case ArrayChar16:
		return GetVectorDataChar16[uint16]((*PlgVector)(unsafe.Pointer(v)))
	case ArrayInt8:
		return GetVectorDataInt8[int8]((*PlgVector)(unsafe.Pointer(v)))
	case ArrayInt16:
		return GetVectorDataInt16[int16]((*PlgVector)(unsafe.Pointer(v)))
	case ArrayInt32:
		return GetVectorDataInt32[int32]((*PlgVector)(unsafe.Pointer(v)))
	case ArrayInt64:
		return GetVectorDataInt64[int64]((*PlgVector)(unsafe.Pointer(v)))
	case ArrayUInt8:
		return GetVectorDataUInt8[uint8]((*PlgVector)(unsafe.Pointer(v)))
	case ArrayUInt16:
		return GetVectorDataUInt16[uint16]((*PlgVector)(unsafe.Pointer(v)))
	case ArrayUInt32:
		return GetVectorDataUInt32[uint32]((*PlgVector)(unsafe.Pointer(v)))
	case ArrayUInt64:
		return GetVectorDataUInt64[uint64]((*PlgVector)(unsafe.Pointer(v)))
	case ArrayPointer:
		return GetVectorDataPointer[uintptr]((*PlgVector)(unsafe.Pointer(v)))
	case ArrayFloat:
		return GetVectorDataFloat[float32]((*PlgVector)(unsafe.Pointer(v)))
	case ArrayDouble:
		return GetVectorDataDouble[float64]((*PlgVector)(unsafe.Pointer(v)))
	case ArrayString:
		return GetVectorDataString[string]((*PlgVector)(unsafe.Pointer(v)))
	case ArrayVector2:
		return GetVectorDataVector2[Vector2]((*PlgVector)(unsafe.Pointer(v)))
	case ArrayVector3:
		return GetVectorDataVector3[Vector3]((*PlgVector)(unsafe.Pointer(v)))
	case ArrayVector4:
		return GetVectorDataVector4[Vector4]((*PlgVector)(unsafe.Pointer(v)))
	case ArrayMatrix4x4:
		return GetVectorDataMatrix4x4[Matrix4x4]((*PlgVector)(unsafe.Pointer(v)))
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

func SetVariantData(v *PlgVariant, param any) {
	var valueType valueType
	switch val := param.(type) {
	case nil:
		valueType = Invalid
	case bool:
		valueType = Bool
		*(*C.bool)(unsafe.Pointer(v)) = C.bool(val)
	/*case rune:
	valueType = Char16
	*(*C.uint16_t)(unsafe.Pointer(v)) = C.uint16_t(param.(rune))*/
	case int8:
		valueType = Int8
		*(*C.int8_t)(unsafe.Pointer(v)) = C.int8_t(val)
	case int16:
		valueType = Int16
		*(*C.int16_t)(unsafe.Pointer(v)) = C.int16_t(val)
	case int32:
		valueType = Int32
		*(*C.int32_t)(unsafe.Pointer(v)) = C.int32_t(val)
	case int64:
		valueType = Int64
		*(*C.int64_t)(unsafe.Pointer(v)) = C.int64_t(val)
	case int:
		valueType = C.Int
		*(*C.int64_t)(unsafe.Pointer(v)) = C.int_t(val)
	case uint8:
		valueType = UInt8
		*(*C.uint8_t)(unsafe.Pointer(v)) = C.uint8_t(val)
	case uint16:
		valueType = UInt16
		*(*C.uint16_t)(unsafe.Pointer(v)) = C.uint16_t(val)
	case uint32:
		valueType = UInt32
		*(*C.uint32_t)(unsafe.Pointer(v)) = C.uint32_t(val)
	case uint64:
		valueType = UInt64
		*(*C.uint64_t)(unsafe.Pointer(v)) = C.uint64_t(val)
	case uint:
		valueType = C.UInt
		*(*C.uint64_t)(unsafe.Pointer(v)) = C.uint_t(val)
	case uintptr:
		valueType = Pointer
		*(*C.intptr_t)(unsafe.Pointer(v)) = C.intptr_t(val)
	case float32:
		valueType = Float
		*(*C.float)(unsafe.Pointer(v)) = C.float(val)
	case float64:
		valueType = Double
		*(*C.double)(unsafe.Pointer(v)) = C.double(val)
	case string:
		valueType = String
		*(*PlgString)(unsafe.Pointer(v)) = C.Plugify_ConstructString(val)
	case []bool:
		valueType = ArrayBool
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorBool(val)
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
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorInt8(val)
	case []int16:
		valueType = ArrayInt16
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorInt16(val)
	case []int32:
		valueType = ArrayInt32
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorInt32(val)
	case []int64:
		valueType = ArrayInt64
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorInt64(val)
	case []int:
		valueType = C.ArrayInt
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorInt(val)
	case []uint8:
		valueType = ArrayUInt8
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorUInt8(val)
	case []uint16:
		valueType = ArrayUInt16
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorUInt16(val)
	case []uint32:
		valueType = ArrayUInt32
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorUInt32(val)
	case []uint64:
		valueType = ArrayUInt64
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorUInt64(val)
	case []uint:
		valueType = C.ArrayUInt
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorUInt(val)
	case []uintptr:
		valueType = ArrayPointer
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorPointer(val)
	case []float32:
		valueType = ArrayFloat
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorFloat(val)
	case []float64:
		valueType = ArrayDouble
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorDouble(val)
	case []string:
		valueType = ArrayString
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorString(val)
	case []Vector2:
		valueType = ArrayVector2
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorVector2(val)
	case []Vector3:
		valueType = ArrayVector3
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorVector3(val)
	case []Vector4:
		valueType = ArrayVector4
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorVector4(val)
	case []Matrix4x4:
		valueType = ArrayMatrix4x4
		*(*PlgVector)(unsafe.Pointer(v)) = ConstructVectorMatrix4x4(val)
	case Vector2:
		valueType = Vector2Type
		*(*PlgVector2)(unsafe.Pointer(v)) = *(*PlgVector2)(unsafe.Pointer(&val))
	case Vector3:
		valueType = Vector3Type
		*(*PlgVector3)(unsafe.Pointer(v)) = *(*PlgVector3)(unsafe.Pointer(&val))
	case Vector4:
		valueType = Vector4Type
		*(*PlgVector4)(unsafe.Pointer(v)) = *(*PlgVector4)(unsafe.Pointer(&val))
	case Matrix4x4:
		valueType = Matrix4x4Type
		*(*PlgMatrix4x4)(unsafe.Pointer(v)) = *(*PlgMatrix4x4)(unsafe.Pointer(&val))
	default:
		panicker(NewTypeNotFoundException(fmt.Sprintf("Type not found: %T", param)))
	}

	v.current = C.uint8_t(valueType)
}

func AssignVariant(v *PlgVariant, param any) {
	DestroyVariant(v)
	SetVariantData(v, param)
}

func ConstructVariant(v any) PlgVariant {
	var variant PlgVariant
	SetVariantData(&variant, v)
	return variant
}

func DestroyVariant(v *PlgVariant) {
	C.Plugify_DestroyVariant(v)
}

// Vector functions
func ConstructVectorBool[T ~bool](data []T) PlgVector {
	return C.Plugify_ConstructVectorBool((*C.bool)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func constructVectorBoolToSlice(v reflect.Value) PlgVector {
	return C.Plugify_ConstructVectorBool((*C.bool)(v.UnsafePointer()), C.ptrdiff_t(v.Len()))
}

func ConstructVectorChar8[T ~int8](data []T) PlgVector {
	return C.Plugify_ConstructVectorChar8((*C.char)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func constructVectorChar8ToSlice(v reflect.Value) PlgVector {
	return C.Plugify_ConstructVectorChar8((*C.char)(v.UnsafePointer()), C.ptrdiff_t(v.Len()))
}

func ConstructVectorChar16[T ~uint16](data []T) PlgVector {
	return C.Plugify_ConstructVectorChar16((*C.char16_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func constructVectorChar16ToSlice(v reflect.Value) PlgVector {
	return C.Plugify_ConstructVectorChar16((*C.char16_t)(v.UnsafePointer()), C.ptrdiff_t(v.Len()))
}

func ConstructVectorInt8[T ~int8](data []T) PlgVector {
	return C.Plugify_ConstructVectorInt8((*C.int8_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func constructVectorInt8ToSlice(v reflect.Value) PlgVector {
	return C.Plugify_ConstructVectorInt8((*C.int8_t)(v.UnsafePointer()), C.ptrdiff_t(v.Len()))
}

func ConstructVectorInt16[T ~int16](data []T) PlgVector {
	return C.Plugify_ConstructVectorInt16((*C.int16_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func constructVectorInt16ToSlice(v reflect.Value) PlgVector {
	return C.Plugify_ConstructVectorInt16((*C.int16_t)(v.UnsafePointer()), C.ptrdiff_t(v.Len()))
}

func ConstructVectorInt32[T ~int32](data []T) PlgVector {
	return C.Plugify_ConstructVectorInt32((*C.int32_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func constructVectorInt32ToSlice(v reflect.Value) PlgVector {
	return C.Plugify_ConstructVectorInt32((*C.int32_t)(v.UnsafePointer()), C.ptrdiff_t(v.Len()))
}

func ConstructVectorInt64[T ~int | ~int64](data []T) PlgVector {
	return C.Plugify_ConstructVectorInt64((*C.int64_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func constructVectorInt64ToSlice(v reflect.Value) PlgVector {
	return C.Plugify_ConstructVectorInt64((*C.int64_t)(v.UnsafePointer()), C.ptrdiff_t(v.Len()))
}

func ConstructVectorInt[T ~int](data []T) PlgVector {
	if is32bit {
		return C.Plugify_ConstructVectorInt32((*C.int32_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
	}

	return C.Plugify_ConstructVectorInt64((*C.int64_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func ConstructVectorUInt8[T ~uint8](data []T) PlgVector {
	return C.Plugify_ConstructVectorUInt8((*C.uint8_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func constructVectorUInt8ToSlice(v reflect.Value) PlgVector {
	return C.Plugify_ConstructVectorUInt8((*C.uint8_t)(v.UnsafePointer()), C.ptrdiff_t(v.Len()))
}

func ConstructVectorUInt16[T ~uint16](data []T) PlgVector {
	return C.Plugify_ConstructVectorUInt16((*C.uint16_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func constructVectorUInt16ToSlice(v reflect.Value) PlgVector {
	return C.Plugify_ConstructVectorUInt16((*C.uint16_t)(v.UnsafePointer()), C.ptrdiff_t(v.Len()))
}

func ConstructVectorUInt32[T ~uint32](data []T) PlgVector {
	return C.Plugify_ConstructVectorUInt32((*C.uint32_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func constructVectorUInt32ToSlice(v reflect.Value) PlgVector {
	return C.Plugify_ConstructVectorUInt32((*C.uint32_t)(v.UnsafePointer()), C.ptrdiff_t(v.Len()))
}

func ConstructVectorUInt64[T ~uint64](data []T) PlgVector {
	return C.Plugify_ConstructVectorUInt64((*C.uint64_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func constructVectorUInt64ToSlice(v reflect.Value) PlgVector {
	return C.Plugify_ConstructVectorUInt64((*C.uint64_t)(v.UnsafePointer()), C.ptrdiff_t(v.Len()))
}

func ConstructVectorUInt[T ~uint](data []T) PlgVector {
	if is32bit {
		return C.Plugify_ConstructVectorUInt32((*C.uint32_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
	}

	return C.Plugify_ConstructVectorUInt64((*C.uint64_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func ConstructVectorPointer[T ~uintptr](data []T) PlgVector {
	return C.Plugify_ConstructVectorPointer((*C.uintptr_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func constructVectorPointerToSlice(v reflect.Value) PlgVector {
	return C.Plugify_ConstructVectorPointer((*C.uintptr_t)(v.UnsafePointer()), C.ptrdiff_t(v.Len()))
}

func ConstructVectorFloat[T ~float32](data []T) PlgVector {
	return C.Plugify_ConstructVectorFloat((*C.float)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func constructVectorFloatToSlice(v reflect.Value) PlgVector {
	return C.Plugify_ConstructVectorFloat((*C.float)(v.UnsafePointer()), C.ptrdiff_t(v.Len()))
}

func ConstructVectorDouble[T ~float64](data []T) PlgVector {
	return C.Plugify_ConstructVectorDouble((*C.double)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func constructVectorDoubleToSlice(v reflect.Value) PlgVector {
	return C.Plugify_ConstructVectorDouble((*C.double)(v.UnsafePointer()), C.ptrdiff_t(v.Len()))
}

func ConstructVectorString[T ~string](data []T) PlgVector {
	vec := C.Plugify_ConstructVectorString(C.ptrdiff_t(len(data)))
	AssignVectorString(&vec, data)
	return vec
}

func constructVectorStringToSlice(v reflect.Value) PlgVector {
	vec := C.Plugify_ConstructVectorString(C.ptrdiff_t(v.Len()))
	reflectAssignVectorString(&vec, v)
	return vec
}

func ConstructVectorVariant[T any](arr []T) PlgVector {
	vec := C.Plugify_ConstructVectorVariant(C.ptrdiff_t(len(arr)))
	AssignVectorVariant(&vec, arr)
	return vec
}

func constructVectorVariantToSlice(v reflect.Value) PlgVector {
	vec := C.Plugify_ConstructVectorVariant(C.ptrdiff_t(v.Len()))
	reflectAssignVectorVariant(&vec, v)
	return vec
}

func ConstructVectorVector2[S ~struct{ X, Y float32 }, T ~[]S](data T) PlgVector {
	return C.Plugify_ConstructVectorVector2((*PlgVector2)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func constructVectorVector2ToSlice(v reflect.Value) PlgVector {
	return C.Plugify_ConstructVectorVector2((*PlgVector2)(v.UnsafePointer()), C.ptrdiff_t(v.Len()))
}

func ConstructVectorVector3[S ~struct{ X, Y, Z float32 }, T ~[]S](data T) PlgVector {
	return C.Plugify_ConstructVectorVector3((*PlgVector3)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func constructVectorVector3ToSlice(v reflect.Value) PlgVector {
	return C.Plugify_ConstructVectorVector3((*PlgVector3)(v.UnsafePointer()), C.ptrdiff_t(v.Len()))
}

func ConstructVectorVector4[S ~struct{ X, Y, Z, W float32 }, T ~[]S](data T) PlgVector {
	return C.Plugify_ConstructVectorVector4((*PlgVector4)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func constructVectorVector4ToSlice(v reflect.Value) PlgVector {
	return C.Plugify_ConstructVectorVector4((*PlgVector4)(v.UnsafePointer()), C.ptrdiff_t(v.Len()))
}

func ConstructVectorMatrix4x4[S ~struct{ M [4][4]float32 }, T ~[]S](data T) PlgVector {
	return C.Plugify_ConstructVectorMatrix4x4((*PlgMatrix4x4)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func constructVectorMatrix4x4ToSlice(v reflect.Value) PlgVector {
	return C.Plugify_ConstructVectorMatrix4x4((*PlgMatrix4x4)(v.UnsafePointer()), C.ptrdiff_t(v.Len()))
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

func getVectorDataBoolToSlice(v *PlgVector, s reflect.Value) {
	size := int(C.Plugify_GetVectorSizeBool(v))
	sliceSize(s, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataBool(v)
		C.memcpy(unsafe.Pointer(s.UnsafePointer()), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_bool)
	}
}

func getVectorDataChar8ToSlice(v *PlgVector, s reflect.Value) {
	size := int(C.Plugify_GetVectorSizeChar8(v))
	sliceSize(s, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataChar8(v)
		C.memcpy(unsafe.Pointer(s.UnsafePointer()), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_char)
	}
}

func getVectorDataChar16ToSlice(v *PlgVector, s reflect.Value) {
	size := int(C.Plugify_GetVectorSizeChar16(v))
	sliceSize(s, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataChar16(v)
		C.memcpy(unsafe.Pointer(s.UnsafePointer()), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_char16_t)
	}
}

func getVectorDataInt8ToSlice(v *PlgVector, s reflect.Value) {
	size := int(C.Plugify_GetVectorSizeInt8(v))
	sliceSize(s, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt8(v)
		C.memcpy(unsafe.Pointer(s.UnsafePointer()), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int8_t)
	}
}

func getVectorDataInt16ToSlice(v *PlgVector, s reflect.Value) {
	size := int(C.Plugify_GetVectorSizeInt16(v))
	sliceSize(s, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt16(v)
		C.memcpy(unsafe.Pointer(s.UnsafePointer()), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int16_t)
	}
}

func getVectorDataInt32ToSlice(v *PlgVector, s reflect.Value) {
	size := int(C.Plugify_GetVectorSizeInt32(v))
	sliceSize(s, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt32(v)
		C.memcpy(unsafe.Pointer(s.UnsafePointer()), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int32_t)
	}
}

func getVectorDataBoolReturn(v *PlgVector, out reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeBool(v))
	slice := reflect.MakeSlice(out, size, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataBool(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_bool)
	}

	return slice
}

func getVectorDataChar8Return(v *PlgVector, out reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeChar8(v))
	slice := reflect.MakeSlice(out, size, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataChar8(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_char)
	}

	return slice
}

func getVectorDataChar16Return(v *PlgVector, out reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeChar16(v))
	slice := reflect.MakeSlice(out, size, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataChar16(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_char16_t)
	}

	return slice
}

func getVectorDataInt8Return(v *PlgVector, out reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeInt8(v))
	slice := reflect.MakeSlice(out, size, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt8(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int8_t)
	}

	return slice
}

func getVectorDataInt16Return(v *PlgVector, out reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeInt16(v))
	slice := reflect.MakeSlice(out, size, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt16(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int16_t)
	}

	return slice
}

func getVectorDataInt32Return(v *PlgVector, out reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeInt32(v))
	slice := reflect.MakeSlice(out, size, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt32(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int32_t)
	}

	return slice
}

func getVectorDataInt64Return(v *PlgVector, out reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeInt64(v))
	slice := reflect.MakeSlice(out, size, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt64(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int64_t)
	}

	return slice
}

func getVectorDataUInt8Return(v *PlgVector, out reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeUInt8(v))
	slice := reflect.MakeSlice(out, size, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt8(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint8_t)
	}

	return slice
}

func getVectorDataUInt16Return(v *PlgVector, out reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeUInt16(v))
	slice := reflect.MakeSlice(out, size, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt16(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint16_t)
	}

	return slice
}

func getVectorDataUInt32Return(v *PlgVector, out reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeUInt32(v))
	slice := reflect.MakeSlice(out, size, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt32(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint32_t)
	}

	return slice
}

func getVectorDataUInt64Return(v *PlgVector, out reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeUInt64(v))
	slice := reflect.MakeSlice(out, size, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt64(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint64_t)
	}

	return slice
}

func getVectorDataPointerReturn(v *PlgVector, out reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizePointer(v))
	slice := reflect.MakeSlice(out, size, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataPointer(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_intptr_t)
	}

	return slice
}

func getVectorDataFloatReturn(v *PlgVector, out reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeFloat(v))
	slice := reflect.MakeSlice(out, size, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataFloat(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_float)
	}

	return slice
}

func getVectorDataDoubleReturn(v *PlgVector, out reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeDouble(v))
	slice := reflect.MakeSlice(out, size, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataDouble(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_double)
	}

	return slice
}

func getVectorDataStringReturn(v *PlgVector, out reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeString(v))
	slice := reflect.MakeSlice(out, size, size)
	/* if size > 0 {
		dataPtr := unsafe.Pointer(C.Plugify_GetVectorDataString(v))
		for i := range size {
			slice.Index(i).Set(reflect.ValueOf(GetStringData[string]((*PlgString)(unsafe.Pointer(uintptr(dataPtr) + uintptr(i)*C.sizeof_String)))))
		}
	} */

	if size > 0 {
		for i := range size {
			str := C.Plugify_GetVectorDataString(v, C.ptrdiff_t(i))
			slice.Index(i).SetString(GetStringData[string](str))
		}
	}

	return slice
}

func getVectorDataAnyReturn(v *PlgVector, out reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeVariant(v))
	slice := reflect.MakeSlice(out, size, size)
	/* if size > 0 {
		dataPtr := unsafe.Pointer(C.Plugify_GetVectorDataString(v))
		for i := range size {
			slice.Index(i).Set(reflect.ValueOf(GetStringData[string]((*PlgString)(unsafe.Pointer(uintptr(dataPtr) + uintptr(i)*C.sizeof_String)))))
		}
	} */

	if size > 0 {
		slicePtr := slice.UnsafePointer()

		for i := range size {
			variant := C.Plugify_GetVectorDataVariant(v, C.ptrdiff_t(i))
			*(*any)(unsafe.Add(slicePtr, i)) = GetVariantData(variant)
		}
	}

	return slice
}

func getVectorDataVector2Return(v *PlgVector, out reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeVector2(v))
	slice := reflect.MakeSlice(out, size, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataVector2(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector2)
	}

	return slice
}

func getVectorDataVector3Return(v *PlgVector, out reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeVector3(v))
	slice := reflect.MakeSlice(out, size, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataVector3(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector3)
	}

	return slice
}

func getVectorDataVector4Return(v *PlgVector, out reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeVector4(v))
	slice := reflect.MakeSlice(out, size, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataVector4(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector4)
	}

	return slice
}

func getVectorDataMatrix4x4Return(v *PlgVector, out reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeMatrix4x4(v))
	slice := reflect.MakeSlice(out, size, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataMatrix4x4(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Matrix4x4)
	}

	return slice
}

func getVectorDataInt64ToSlice(v *PlgVector, s reflect.Value) {
	size := int(C.Plugify_GetVectorSizeInt64(v))
	sliceSize(s, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt64(v)
		C.memcpy(unsafe.Pointer(s.UnsafePointer()), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int64_t)
	}
}

func getVectorDataBoolReflect(v *PlgVector, t reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeBool(v))
	slice := reflect.MakeSlice(t, size, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataBool(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_bool)
	}

	return slice
}

func getVectorDataChar8Reflect(v *PlgVector, t reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeChar8(v))
	slice := reflect.MakeSlice(t, size, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataChar8(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_char)
	}

	return slice
}

func getVectorDataChar16Reflect(v *PlgVector, t reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeChar16(v))
	slice := reflect.MakeSlice(t, size, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataChar16(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_char16_t)
	}

	return slice
}

func getVectorDataInt8Reflect(v *PlgVector, t reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeInt8(v))
	slice := reflect.MakeSlice(t, size, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt8(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int8_t)
	}

	return slice
}

func getVectorDataInt16Reflect(v *PlgVector, t reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeInt16(v))
	slice := reflect.MakeSlice(t, size, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt16(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int16_t)
	}

	return slice
}

func getVectorDataInt32Reflect(v *PlgVector, t reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeInt32(v))
	slice := reflect.MakeSlice(t, size, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt32(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int32_t)
	}

	return slice
}

func getVectorDataInt64Reflect(v *PlgVector, t reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeInt64(v))
	slice := reflect.MakeSlice(t, size, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt64(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int64_t)
	}

	return slice
}

func getVectorDataUInt8Reflect(v *PlgVector, t reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeUInt8(v))
	slice := reflect.MakeSlice(t, size, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt8(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint8_t)
	}

	return slice
}

func getVectorDataUInt16Reflect(v *PlgVector, t reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeUInt16(v))
	slice := reflect.MakeSlice(t, size, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt16(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint16_t)
	}

	return slice
}

func getVectorDataUInt32Reflect(v *PlgVector, t reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeUInt32(v))
	slice := reflect.MakeSlice(t, size, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt32(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint32_t)
	}

	return slice
}

func getVectorDataUInt64Reflect(v *PlgVector, t reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeUInt64(v))
	slice := reflect.MakeSlice(t, size, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt64(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint64_t)
	}

	return slice
}

func getVectorDataPointerReflect(v *PlgVector, t reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizePointer(v))
	slice := reflect.MakeSlice(t, size, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataPointer(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_intptr_t)
	}

	return slice
}

func getVectorDataFloatReflect(v *PlgVector, t reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeFloat(v))
	slice := reflect.MakeSlice(t, size, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataFloat(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_float)
	}

	return slice
}

func getVectorDataDoubleReflect(v *PlgVector, t reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeDouble(v))
	slice := reflect.MakeSlice(t, size, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataDouble(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_double)
	}

	return slice
}

func getVectorDataStringReflect(v *PlgVector, t reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeString(v))
	slice := reflect.MakeSlice(t, size, size)

	if size > 0 {
		for i := range size {
			str := C.Plugify_GetVectorDataString(v, C.ptrdiff_t(i))
			slice.Index(i).SetString(GetStringData[string](str))
		}
	}

	return slice
}

func getVectorDataAnyReflect(v *PlgVector, t reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeVariant(v))
	slice := reflect.MakeSlice(t, size, size)

	for i := range size {
		variant := C.Plugify_GetVectorDataVariant(v, C.ptrdiff_t(i))
		slice.Index(i).Set(reflect.ValueOf(GetVariantData(variant)))
	}

	return slice
}

func getVectorDataVector2Reflect(v *PlgVector, t reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeVector2(v))
	slice := reflect.MakeSlice(t, size, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataVector2(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector2)
	}

	return slice
}

func getVectorDataVector3Reflect(v *PlgVector, t reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeVector3(v))
	slice := reflect.MakeSlice(t, size, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataVector3(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector3)
	}

	return slice
}

func getVectorDataVector4Reflect(v *PlgVector, t reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeVector4(v))
	slice := reflect.MakeSlice(t, size, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataVector4(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector4)
	}

	return slice
}

func getVectorDataMatrix4x4Reflect(v *PlgVector, t reflect.Type) reflect.Value {
	size := int(C.Plugify_GetVectorSizeMatrix4x4(v))
	slice := reflect.MakeSlice(t, size, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataMatrix4x4(v)
		C.memcpy(slice.UnsafePointer(), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Matrix4x4)
	}

	return slice
}

func getVectorDataUInt8ToSlice(v *PlgVector, s reflect.Value) {
	size := int(C.Plugify_GetVectorSizeUInt8(v))
	sliceSize(s, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt8(v)
		C.memcpy(unsafe.Pointer(s.UnsafePointer()), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint8_t)
	}
}

func getVectorDataUInt16ToSlice(v *PlgVector, s reflect.Value) {
	size := int(C.Plugify_GetVectorSizeUInt16(v))
	sliceSize(s, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt16(v)
		C.memcpy(unsafe.Pointer(s.UnsafePointer()), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint16_t)
	}
}

func getVectorDataUInt32ToSlice(v *PlgVector, s reflect.Value) {
	size := int(C.Plugify_GetVectorSizeUInt32(v))
	sliceSize(s, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt32(v)
		C.memcpy(unsafe.Pointer(s.UnsafePointer()), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint32_t)
	}
}

func getVectorDataUInt64ToSlice(v *PlgVector, s reflect.Value) {
	size := int(C.Plugify_GetVectorSizeUInt64(v))
	sliceSize(s, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt64(v)
		C.memcpy(unsafe.Pointer(s.UnsafePointer()), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint64_t)
	}
}

func getVectorDataPointerToSlice(v *PlgVector, s reflect.Value) {
	size := int(C.Plugify_GetVectorSizePointer(v))
	sliceSize(s, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataPointer(v)
		C.memcpy(unsafe.Pointer(s.UnsafePointer()), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_intptr_t)
	}
}

func getVectorDataFloatToSlice(v *PlgVector, s reflect.Value) {
	size := int(C.Plugify_GetVectorSizeFloat(v))
	sliceSize(s, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataFloat(v)
		C.memcpy(unsafe.Pointer(s.UnsafePointer()), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_float)
	}
}

func getVectorDataDoubleToSlice(v *PlgVector, s reflect.Value) {
	size := int(C.Plugify_GetVectorSizeDouble(v))
	sliceSize(s, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataDouble(v)
		C.memcpy(unsafe.Pointer(s.UnsafePointer()), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_double)
	}
}

func getVectorDataStringToSlice(v *PlgVector, s reflect.Value) {
	size := int(C.Plugify_GetVectorSizeString(v))
	sliceSize(s, size)

	for i := range size {
		str := C.Plugify_GetVectorDataString(v, C.ptrdiff_t(i))
		s.Index(i).SetString(GetStringData[string](str))
	}
}

func GetVectorDataVariant[V any, T ~[]V](v *PlgVector) T {
	size := int(C.Plugify_GetVectorSizeVariant(v))
	arr := make(T, size)
	/* for i := range size {
		variant := C.Plugify_GetVectorDataVariant(v, C.ptrdiff_t(i))
		//reflect.ValueOf(arr[i]).Set(reflect.ValueOf(GetVariantData(variant)))
	} */

	//slice := reflect.ValueOf(arr)
	//slicePtr := slice.UnsafePointer()

	for i := range size {
		variant := C.Plugify_GetVectorDataVariant(v, C.ptrdiff_t(i))
		data := GetVariantData(variant)

		if data != nil {
			arr[i] = data.(V)
		}

		//*(*uint)(unsafe.Pointer(uintptr(slicePtr) + uintptr(i) * unsafe.Sizeof(any(nil)))) = *(*uint)(unsafe.Pointer(&data))

		// TODO: check memory leak
		// runtime.Pinner
		//C.memcpy(unsafe.Pointer(uintptr(C.size_t(uintptr(slicePtr))+i*C.sizeof_Variant)), unsafe.Pointer(&data), C.sizeof_Variant)
		//entry := unsafe.Pointer(&arr[i])
		//*(*uint)(entry) = *(*uint)(unsafe.Pointer(&data))

		//reflect.ValueOf((arr)[i]).Set(reflect.ValueOf(GetVariantData(variant)))
	}

	return arr
}

func getVectorDataVariantToSlice(v *PlgVector, s reflect.Value) {
	size := int(C.Plugify_GetVectorSizeVariant(v))
	sliceSize(s, size)

	for i := range size {
		variant := C.Plugify_GetVectorDataVariant(v, C.ptrdiff_t(i))
		data := GetVariantData(variant)

		if data != nil {
			s.Index(i).Set(reflect.ValueOf(data))
		}
	}
}

func GetVectorDataVector2[T ~struct{ X, Y float32 }](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeVector2(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataVector2(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector2)
	}
	return arr
}

func getVectorDataVector2ToSlice(v *PlgVector, s reflect.Value) {
	size := int(C.Plugify_GetVectorSizeVector2(v))
	sliceSize(s, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataVector2(v)
		C.memcpy(unsafe.Pointer(s.UnsafePointer()), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector2)
	}
}

func GetVectorDataVector3[T ~struct{ X, Y, Z float32 }](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeVector3(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataVector3(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector3)
	}
	return arr
}

func getVectorDataVector3ToSlice(v *PlgVector, s reflect.Value) {
	size := int(C.Plugify_GetVectorSizeVector3(v))
	sliceSize(s, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataVector3(v)
		C.memcpy(unsafe.Pointer(s.UnsafePointer()), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector3)
	}
}

func GetVectorDataVector4[T ~struct{ X, Y, Z, W float32 }](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeVector4(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataVector4(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector4)
	}
	return arr
}

func getVectorDataVector4ToSlice(v *PlgVector, s reflect.Value) {
	size := int(C.Plugify_GetVectorSizeVector4(v))
	sliceSize(s, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataVector4(v)
		C.memcpy(unsafe.Pointer(s.UnsafePointer()), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector4)
	}
}

func GetVectorDataMatrix4x4[T ~struct{ M [4][4]float32 }](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeMatrix4x4(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataMatrix4x4(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Matrix4x4)
	}
	return arr
}

func getVectorDataMatrix4x4ToSlice(v *PlgVector, s reflect.Value) {
	size := int(C.Plugify_GetVectorSizeMatrix4x4(v))
	sliceSize(s, size)

	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataMatrix4x4(v)
		C.memcpy(unsafe.Pointer(s.UnsafePointer()), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Matrix4x4)
	}
}

func GetVectorDataBool[T ~bool](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeBool(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataBool(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_bool)
	}
	return arr
}

func GetVectorDataChar8[T ~int8](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeChar8(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataChar8(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_char)
	}
	return arr
}

func GetVectorDataChar16[T ~uint16](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeChar16(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataChar16(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_char16_t)
	}
	return arr
}

func GetVectorDataInt8[T ~int8](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeInt8(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt8(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int8_t)
	}
	return arr
}

func GetVectorDataInt16[T ~int16](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeInt16(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt16(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int16_t)
	}
	return arr
}

func GetVectorDataIntT[T ~int](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeInt32(v))
	arr := make([]T, size)

	if size > 0 {
		if is32bit {
			dataPtr := C.Plugify_GetVectorDataInt32(v)
			C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int32_t)
		} else {
			dataPtr := C.Plugify_GetVectorDataInt64(v)
			C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int64_t)
		}
	}

	return arr
}

func GetVectorDataInt32[T ~int32](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeInt32(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt32(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int32_t)
	}
	return arr
}

func GetVectorDataInt64[T ~int64 | ~int](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeInt64(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt64(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int64_t)
	}
	return arr
}

func GetVectorDataUInt8[T ~uint8](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeUInt8(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt8(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint8_t)
	}
	return arr
}

func GetVectorDataUInt16[T ~uint16](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeUInt16(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt16(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint16_t)
	}
	return arr
}

func GetVectorDataUInt32[T ~uint32](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeUInt32(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt32(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint32_t)
	}
	return arr
}

func GetVectorDataUInt64[T ~uint64](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeUInt64(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt64(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint64_t)
	}
	return arr
}

func GetVectorDataPointer[T ~uintptr](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizePointer(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataPointer(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_intptr_t)
	}
	return arr
}

func GetVectorDataFloat[T ~float32](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeFloat(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataFloat(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_float)
	}
	return arr
}

func GetVectorDataDouble[T ~float64](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeDouble(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataDouble(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_double)
	}
	return arr
}

func GetVectorDataString[T ~string](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeString(v))
	arr := make([]T, size)
	for i := range size {
		str := C.Plugify_GetVectorDataString(v, C.ptrdiff_t(i))
		arr[i] = GetStringData[T](str)
	}
	return arr
}

func GetVectorDataVector2T[T any](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeVector2(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataVector2(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector2)
	}
	return arr
}

func GetVectorDataVector3T[T any](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeVector3(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataVector3(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector3)
	}
	return arr
}

func GetVectorDataVector4T[T any](v *PlgVector) []T {
	size := int(C.Plugify_GetVectorSizeVector4(v))
	arr := make([]T, size)
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataVector4(v)
		C.memcpy(unsafe.Pointer(unsafe.SliceData(arr)), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector4)
	}
	return arr
}

func GetVectorDataBoolTo[V ~bool, T ~[]V](v *PlgVector, arr *T) {
	size := int(C.Plugify_GetVectorSizeBool(v))
	if len(*arr) < size {
		*arr = make(T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataBool(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_bool)
	}
}

func GetVectorDataChar8To[V ~int8, T ~[]V](v *PlgVector, arr *T) {
	size := int(C.Plugify_GetVectorSizeChar8(v))
	if len(*arr) < size {
		*arr = make(T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataChar8(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_char)
	}
}

func GetVectorDataChar16To[V ~uint16, T ~[]V](v *PlgVector, arr *T) {
	size := int(C.Plugify_GetVectorSizeChar16(v))
	if len(*arr) < size {
		*arr = make(T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataChar16(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_char16_t)
	}
}

func GetVectorDataInt8To[V ~int8, T ~[]V](v *PlgVector, arr *T) {
	size := int(C.Plugify_GetVectorSizeInt8(v))
	if len(*arr) < size {
		*arr = make(T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt8(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int8_t)
	}
}

func GetVectorDataInt16To[V ~int16, T ~[]V](v *PlgVector, arr *T) {
	size := int(C.Plugify_GetVectorSizeInt16(v))
	if len(*arr) < size {
		*arr = make(T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt16(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int16_t)
	}
}

func GetVectorDataInt32To[V ~int32, T ~[]V](v *PlgVector, arr *T) {
	size := int(C.Plugify_GetVectorSizeInt32(v))
	if len(*arr) < size {
		*arr = make(T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt32(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int32_t)
	}
}

func GetVectorDataInt64To[V ~int64, T ~[]V](v *PlgVector, arr *T) {
	size := int(C.Plugify_GetVectorSizeInt64(v))
	if len(*arr) < size {
		*arr = make(T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataInt64(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int64_t)
	}
}

func GetVectorDataIntTo[V ~int, T ~[]V](v *PlgVector, arr *T) {
	if is32bit {
		size := int(C.Plugify_GetVectorSizeInt32(v))
		if len(*arr) < size {
			*arr = make(T, size)
		}

		dataPtr := C.Plugify_GetVectorDataInt32(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int32_t)
	} else {
		size := int(C.Plugify_GetVectorSizeInt64(v))
		if len(*arr) < size {
			*arr = make(T, size)
		}

		dataPtr := C.Plugify_GetVectorDataInt64(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_int64_t)
	}
}

func GetVectorDataUInt8To[V ~uint8, T ~[]V](v *PlgVector, arr *T) {
	size := int(C.Plugify_GetVectorSizeUInt8(v))
	if len(*arr) < size {
		*arr = make(T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt8(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint8_t)
	}
}

func GetVectorDataUInt16To[V ~uint16, T ~[]V](v *PlgVector, arr *T) {
	size := int(C.Plugify_GetVectorSizeUInt16(v))
	if len(*arr) < size {
		*arr = make(T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt16(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint16_t)
	}
}

func GetVectorDataUInt32To[V ~uint32, T ~[]V](v *PlgVector, arr *T) {
	size := int(C.Plugify_GetVectorSizeUInt32(v))
	if len(*arr) < size {
		*arr = make(T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt32(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint32_t)
	}
}

func GetVectorDataUInt64To[V ~uint64, T ~[]V](v *PlgVector, arr *T) {
	size := int(C.Plugify_GetVectorSizeUInt64(v))
	if len(*arr) < size {
		*arr = make(T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataUInt64(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_uint64_t)
	}
}

func GetVectorDataPointerTo[V ~uintptr, T ~[]V](v *PlgVector, arr *T) {
	size := int(C.Plugify_GetVectorSizePointer(v))
	if len(*arr) < size {
		*arr = make(T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataPointer(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_intptr_t)
	}
}

func GetVectorDataFloatTo[V ~float32, T ~[]V](v *PlgVector, arr *T) {
	size := int(C.Plugify_GetVectorSizeFloat(v))
	if len(*arr) < size {
		*arr = make(T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataFloat(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_float)
	}
}

func GetVectorDataDoubleTo[V ~float64, T ~[]V](v *PlgVector, arr *T) {
	size := int(C.Plugify_GetVectorSizeDouble(v))
	if len(*arr) < size {
		*arr = make(T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataDouble(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_double)
	}
}

func GetVectorDataStringTo[V ~string, T ~[]V](v *PlgVector, arr *T) {
	size := int(C.Plugify_GetVectorSizeString(v))
	if len(*arr) < size {
		*arr = make(T, size)
	}
	for i := range size {
		str := C.Plugify_GetVectorDataString(v, C.ptrdiff_t(i))
		(*arr)[i] = GetStringData[V](str)
	}
}

func GetVectorDataVariantTo[V any, T ~[]V](v *PlgVector, arr *T) {
	size := int(C.Plugify_GetVectorSizeVariant(v))
	// TODO: should use cap
	if len(*arr) < size {
		*arr = make(T, size)
	}

	for i := range size {
		variant := C.Plugify_GetVectorDataVariant(v, C.ptrdiff_t(i))
		data := GetVariantData(variant)

		(*arr)[i] = data.(V)
	}
}

func GetVectorDataVector2To[V ~struct{ X, Y float32 }, T ~[]V](v *PlgVector, arr *T) {
	size := int(C.Plugify_GetVectorSizeVector2(v))
	if len(*arr) < size {
		*arr = make(T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataVector2(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector2)
	}
}

func GetVectorDataVector3To[V ~struct{ X, Y, Z float32 }, T ~[]V](v *PlgVector, arr *T) {
	size := int(C.Plugify_GetVectorSizeVector3(v))
	if len(*arr) < size {
		*arr = make(T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataVector3(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector3)
	}
}

func GetVectorDataVector4To[V ~struct{ X, Y, Z, W float32 }, T ~[]V](v *PlgVector, arr *T) {
	size := int(C.Plugify_GetVectorSizeVector4(v))
	if len(*arr) < size {
		*arr = make(T, size)
	}
	if size > 0 {
		dataPtr := C.Plugify_GetVectorDataVector4(v)
		C.memcpy(unsafe.Pointer(&(*arr)[0]), unsafe.Pointer(dataPtr), C.size_t(size)*C.sizeof_Vector4)
	}
}

func GetVectorDataMatrix4x4To[V ~struct{ M [4][4]float32 }, T ~[]V](v *PlgVector, arr *T) {
	size := int(C.Plugify_GetVectorSizeMatrix4x4(v))
	if len(*arr) < size {
		*arr = make(T, size)
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

func AssignVectorInt32[T ~int32 | ~int](v *PlgVector, data []T) {
	C.Plugify_AssignVectorInt32(v, (*C.int32_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

/* func AssignVectorInt[T ~int](v *PlgVector, data []T) {
	if is32bit {
		C.Plugify_AssignVectorInt32(v, (*C.int32_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
	} else {
		C.Plugify_AssignVectorInt64(v, (*C.int64_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
	}
} */

func AssignVectorInt64[T ~int64 | ~int](v *PlgVector, data []T) {
	C.Plugify_AssignVectorInt64(v, (*C.int64_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func VAssignVectorInt64(v *PlgVector, data reflect.Value) {
	C.Plugify_AssignVectorInt64(v, (*C.int64_t)(data.UnsafePointer()), C.ptrdiff_t(data.Len()))
}

func AssignVectorUInt8[T ~uint8](v *PlgVector, data []T) {
	C.Plugify_AssignVectorUInt8(v, (*C.uint8_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorUInt16[T ~uint16](v *PlgVector, data []T) {
	C.Plugify_AssignVectorUInt16(v, (*C.uint16_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorUInt32[T ~uint32 | ~uint](v *PlgVector, data []T) {
	C.Plugify_AssignVectorUInt32(v, (*C.uint32_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorUInt64[T ~uint64 | ~uint](v *PlgVector, data []T) {
	C.Plugify_AssignVectorUInt64(v, (*C.uint64_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorPointer[T ~uintptr](v *PlgVector, data []T) {
	C.Plugify_AssignVectorPointer(v, (*C.uintptr_t)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorFloat[T ~float32](v *PlgVector, data []T) {
	C.Plugify_AssignVectorFloat(v, (*C.float)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorDouble[T ~float64](v *PlgVector, data []T) {
	C.Plugify_AssignVectorDouble(v, (*C.double)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorString[T ~string](v *PlgVector, data []T) {
	size := len(data)
	C.Plugify_AssignVectorString(v, C.ptrdiff_t(size))
	for i := range size {
		str := C.Plugify_GetVectorDataString(v, C.ptrdiff_t(i))
		AssignString(str, data[i])
	}
}

func reflectAssignVectorString(vec *PlgVector, v reflect.Value) {
	size := v.Len()
	C.Plugify_AssignVectorString(vec, C.ptrdiff_t(size))
	for i := range size {
		str := C.Plugify_GetVectorDataString(vec, C.ptrdiff_t(i))
		AssignString(str, v.Index(i).String())
	}
}
func AssignVectorVariant[T any](v *PlgVector, data []T) {
	size := len(data)
	C.Plugify_AssignVectorVariant(v, C.ptrdiff_t(size))
	for i := range size {
		variant := C.Plugify_GetVectorDataVariant(v, C.ptrdiff_t(i))
		AssignVariant(variant, data[i])
	}
}

func reflectAssignVectorVariant(vec *PlgVector, v reflect.Value) {
	size := v.Len()
	C.Plugify_AssignVectorVariant(vec, C.ptrdiff_t(size))
	for i := range size {
		variant := C.Plugify_GetVectorDataVariant(vec, C.ptrdiff_t(i))
		AssignVariant(variant, v.Index(i).Interface())
	}
}

func AssignVectorVector2[T ~struct{ X, Y float32 }](v *PlgVector, data []T) {
	C.Plugify_AssignVectorVector2(v, (*PlgVector2)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorVector3[T ~struct{ X, Y, Z float32 }](v *PlgVector, data []T) {
	C.Plugify_AssignVectorVector3(v, (*PlgVector3)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorVector4[T ~struct{ X, Y, Z, W float32 }](v *PlgVector, data []T) {
	C.Plugify_AssignVectorVector4(v, (*PlgVector4)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}

func AssignVectorMatrix4x4[T ~struct{ M [4][4]float32 }](v *PlgVector, data []T) {
	C.Plugify_AssignVectorMatrix4x4(v, (*PlgMatrix4x4)(unsafe.Pointer(unsafe.SliceData(data))), C.ptrdiff_t(len(data)))
}
