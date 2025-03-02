package plugify

/*
#cgo LDFLAGS: -L${SRCDIR}/libplugify.a
#include <plugify.h>
*/
import "C"
import "unsafe"

const kApiVersion = 1

type PluginStartCallback func()
type PluginUpdateCallback func(dt float32)
type PluginEndCallback func()

type Plugify struct {
	Id                     int64
	Name                   string
	FullName               string
	Description            string
	Version                string
	Author                 string
	Website                string
	BaseDir                string
	Dependencies           []string
	fnPluginStartCallback  PluginStartCallback
	fnPluginUpdateCallback PluginUpdateCallback
	fnPluginEndCallback    PluginEndCallback
}

var plugify Plugify = Plugify{
	Id:                    -1,
	Name:                  "",
	FullName:              "",
	Description:           "",
	Version:               "",
	Author:                "",
	Website:               "",
	BaseDir:               "",
	Dependencies:          []string{},
	fnPluginStartCallback: func() {},
	fnPluginEndCallback:   func() {},
}

var BaseDir string = ""

func OnPluginStart(fn PluginStartCallback) {
	plugify.fnPluginStartCallback = fn
}

func OnPluginUpdate(fn PluginUpdateCallback) {
	plugify.fnPluginUpdateCallback = fn
}

func OnPluginEnd(fn PluginEndCallback) {
	plugify.fnPluginEndCallback = fn
}

func (p *Plugify) FindResource(path string) string {
	C_output := C.Plugify_FindPluginResource(path)
	output := C.GoString(C_output)
	C.Plugify_DeleteCStr(C_output)
	return output
}

//export Plugify_Init
func Plugify_Init(api []uintptr, version int32, plugin uintptr, assembly uintptr) int32 {
	if version < kApiVersion {
		return kApiVersion
	}
	i := 0
	C.Plugify_SetGetMethodPtr(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetMethodPtr2(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetBaseDir(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetIsModuleLoaded(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetIsPluginLoaded(unsafe.Pointer(api[i]))
	i++

	C.Plugify_SetGetPluginId(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetPluginName(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetPluginFullName(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetPluginDescription(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetPluginVersion(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetPluginAuthor(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetPluginWebsite(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetPluginBaseDir(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetPluginDependencies(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetPluginDependenciesSize(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetFindPluginResource(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetFindFunctionByName(unsafe.Pointer(api[i]))
	i++

	C.Plugify_SetDeleteCStr(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetDeleteCStrArr(unsafe.Pointer(api[i]))
	i++

	C.Plugify_SetConstructString(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetDestroyString(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetStringData(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetStringLength(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetAssignString(unsafe.Pointer(api[i]))
	i++

	C.Plugify_SetDestroyVariant(unsafe.Pointer(api[i]))
	i++

	C.Plugify_SetConstructVectorBool(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetConstructVectorChar8(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetConstructVectorChar16(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetConstructVectorInt8(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetConstructVectorInt16(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetConstructVectorInt32(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetConstructVectorInt64(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetConstructVectorUInt16(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetConstructVectorUInt32(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetConstructVectorUInt64(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetConstructVectorPointer(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetConstructVectorFloat(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetConstructVectorDouble(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetConstructVectorString(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetConstructVectorVariant(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetConstructVectorVector2(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetConstructVectorVector3(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetConstructVectorVector4(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetConstructVectorMatrix4x4(unsafe.Pointer(api[i]))
	i++

	C.Plugify_SetDestroyVectorBool(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetDestroyVectorChar8(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetDestroyVectorChar16(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetDestroyVectorInt8(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetDestroyVectorInt16(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetDestroyVectorInt32(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetDestroyVectorInt64(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetDestroyVectorUInt8(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetDestroyVectorUInt16(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetDestroyVectorUInt32(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetDestroyVectorUInt64(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetDestroyVectorPointer(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetDestroyVectorFloat(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetDestroyVectorDouble(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetDestroyVectorString(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetDestroyVectorVariant(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetDestroyVectorVector2(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetDestroyVectorVector3(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetDestroyVectorVector4(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetDestroyVectorMatrix4x4(unsafe.Pointer(api[i]))
	i++

	C.Plugify_SetGetVectorSizeBool(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorSizeChar8(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorSizeChar16(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorSizeInt8(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorSizeInt16(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorSizeInt32(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorSizeInt64(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorSizeUInt8(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorSizeUInt16(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorSizeUInt32(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorSizeUInt64(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorSizePointer(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorSizeFloat(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorSizeDouble(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorSizeString(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorSizeVariant(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorSizeVector2(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorSizeVector3(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorSizeVector4(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorSizeMatrix4x4(unsafe.Pointer(api[i]))
	i++

	C.Plugify_SetGetVectorDataBool(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorDataChar8(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorDataChar16(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorDataInt8(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorDataInt16(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorDataInt32(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorDataInt64(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorDataUInt8(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorDataUInt16(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorDataUInt32(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorDataUInt64(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorDataPointer(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorDataFloat(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorDataDouble(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorDataString(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorDataVariant(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorDataVector2(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorDataVector3(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorDataVector4(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetVectorDataMatrix4x4(unsafe.Pointer(api[i]))
	i++

	C.Plugify_SetAssignVectorBool(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetAssignVectorChar8(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetAssignVectorChar16(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetAssignVectorInt8(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetAssignVectorInt16(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetAssignVectorInt32(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetAssignVectorInt64(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetAssignVectorUInt8(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetAssignVectorUInt16(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetAssignVectorUInt32(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetAssignVectorUInt64(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetAssignVectorPointer(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetAssignVectorFloat(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetAssignVectorDouble(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetAssignVectorString(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetAssignVectorVariant(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetAssignVectorVector2(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetAssignVectorVector3(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetAssignVectorVector4(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetAssignVectorMatrix4x4(unsafe.Pointer(api[i]))
	i++

	C.Plugify_SetNewCall(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetDeleteCall(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetCallFunction(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetCallError(unsafe.Pointer(api[i]))
	i++

	C.Plugify_SetNewCallback(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetDeleteCallback(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetCallbackFunction(unsafe.Pointer(api[i]))
	i++
	C.Plugify_SetGetCallbackError(unsafe.Pointer(api[i]))
	i++

	C.pluginHandle = unsafe.Pointer(plugin)
	C.assemblyHandle = unsafe.Pointer(assembly)

	path := C.Plugify_GetBaseDir()
	BaseDir = C.GoString(path)
	C.Plugify_DeleteCStr(path)

	plugify.Id = int64(C.Plugify_GetPluginId())
	plugify.Name = C.GoString(C.Plugify_GetPluginName())
	plugify.FullName = C.GoString(C.Plugify_GetPluginFullName())
	plugify.Description = C.GoString(C.Plugify_GetPluginDescription())
	plugify.Version = C.GoString(C.Plugify_GetPluginVersion())
	plugify.Author = C.GoString(C.Plugify_GetPluginAuthor())
	plugify.Website = C.GoString(C.Plugify_GetPluginWebsite())

	pluginPath := C.Plugify_GetPluginBaseDir()
	plugify.BaseDir = C.GoString(pluginPath)
	C.Plugify_DeleteCStr(pluginPath)

	dependencies := C.Plugify_GetPluginDependencies()
	plugify.Dependencies = make([]string, C.Plugify_GetPluginDependenciesSize())
	for i := range plugify.Dependencies {
		plugify.Dependencies[i] = C.GoString(*(**C.char)(unsafe.Pointer(uintptr(dependencies) + uintptr(i)*C.sizeof_uintptr_t)))
	}
	C.Plugify_DeleteCStrArr(dependencies)

	return 0
}

//export Plugify_PluginStart
func Plugify_PluginStart() {
	plugify.fnPluginStartCallback()
}

//export Plugify_PluginUpdate
func Plugify_PluginUpdate(dt float32) {
	plugify.fnPluginUpdateCallback(dt)
}

//export Plugify_PluginEnd
func Plugify_PluginEnd() {
	plugify.fnPluginEndCallback()
}
