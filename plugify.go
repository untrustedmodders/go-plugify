package plugify

/*
#cgo LDFLAGS: -L${SRCDIR}/libplugify.a
#include <plugify.h>
*/
import "C"
import "unsafe"

type PluginStartCallback func()
type PluginEndCallback func()

type Plugify struct {
	fnPluginStartCallback PluginStartCallback
	fnPluginEndCallback   PluginEndCallback
}

var plugify Plugify = Plugify {
	fnPluginStartCallback: func() {},
	fnPluginEndCallback:   func() {},
}

const kApiVersion = 1

//export Plugify_Init
func Plugify_Init(api []uintptr, version int32) int32 {
	if version < kApiVersion {
		return kApiVersion
	}
	C.Plugify_SetGetMethodPtr(unsafe.Pointer(api[0]));
	C.Plugify_SetAllocateString(unsafe.Pointer(api[1]));
	C.Plugify_SetCreateString(unsafe.Pointer(api[2]));
	C.Plugify_SetGetStringData(unsafe.Pointer(api[3]));
	C.Plugify_SetGetStringLength(unsafe.Pointer(api[4]));
	C.Plugify_SetAssignString(unsafe.Pointer(api[5]));
	C.Plugify_SetFreeString(unsafe.Pointer(api[6]));
	C.Plugify_SetDeleteString(unsafe.Pointer(api[7]));
	C.Plugify_SetCreateVector(unsafe.Pointer(api[8]));
	C.Plugify_SetAllocateVector(unsafe.Pointer(api[9]));
	C.Plugify_SetGetVectorSize(unsafe.Pointer(api[10]));
	C.Plugify_SetGetVectorData(unsafe.Pointer(api[11]));
	C.Plugify_SetAssignVector(unsafe.Pointer(api[12]));
	C.Plugify_SetDeleteVector(unsafe.Pointer(api[13]));
	C.Plugify_SetFreeVector(unsafe.Pointer(api[14]));
	C.Plugify_SetDeleteVectorDataBool(unsafe.Pointer(api[15]));
	C.Plugify_SetDeleteVectorDataCStr(unsafe.Pointer(api[16]));
	return 0
}

//export Plugify_PluginStart
func Plugify_PluginStart() {
	plugify.fnPluginStartCallback()
}

func OnPluginStart(fn PluginStartCallback) {
	plugify.fnPluginStartCallback = fn
}

//export Plugify_PluginEnd
func Plugify_PluginEnd() {
	plugify.fnPluginEndCallback()
}

func OnPluginEnd(fn PluginEndCallback) {
	plugify.fnPluginEndCallback = fn
}

type Vector2 struct {
	X float32
	Y float32
}
type Vector3 struct {
	X float32
	Y float32
	Z float32
}
type Vector4 struct {
	X float32
	Y float32
	Z float32
	W float32
}
