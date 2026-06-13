package generator

import (
	"go/types"
	"unsafe"
)

func mapBasicType(basic *types.Basic) string {
	switch basic.Kind() {
	case types.Bool:
		return "bool"
	case types.Int8:
		return "int8"
	case types.Int16:
		return "int16"
	case types.Int32:
		return "int32"
	case types.Int64:
		return "int64"
	case types.Uint8:
		return "uint8"
	case types.Uint16:
		return "uint16"
	case types.Uint32:
		return "uint32"
	case types.Uint64:
		return "uint64"
	case types.Float32:
		return "float"
	case types.Float64:
		return "double"
	case types.String:
		return "string"
	case types.UnsafePointer:
		return getPtrType()
	case types.Uintptr:
		return getPtrType()
	case types.UntypedNil:
		return getPtrType()
	case types.Uint:
		return getUIntType()
	case types.Int:
		return getIntType()
	default:
		return basic.String()
	}
}

func getPtrType() string {
	if unsafe.Sizeof(uintptr(0)) == 4 {
		return "ptr32"
	} else {
		return "ptr64"
	}
}

func getIntType() string {
	if unsafe.Sizeof(uintptr(0)) == 4 {
		return "int32"
	} else {
		return "int64"
	}
}

func getUIntType() string {
	if unsafe.Sizeof(uintptr(0)) == 4 {
		return "uint32"
	} else {
		return "uint64"
	}
}
