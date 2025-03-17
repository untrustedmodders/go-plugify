package plugify

/*
#cgo LDFLAGS: -L${SRCDIR}/libplugify.a
#include "plugify.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

const kApiVersion = 1

type PluginStartCallback func()
type PluginUpdateCallback func(dt float32)
type PluginEndCallback func()
type PluginPanicCallback func() []byte

type Plugify struct {
	Id           int64
	Name         string
	FullName     string
	Description  string
	Version      string
	Author       string
	Website      string
	BaseDir      string
	ConfigsDir   string
	DataDir      string
	LogsDir      string
	Dependencies []string

	fnPluginStartCallback   PluginStartCallback
	fnPluginUpdateCallback  PluginUpdateCallback
	fnPluginEndCallback     PluginEndCallback
	fnPluginPanicCallback   PluginPanicCallback
	hasPluginStartCallback  bool
	hasPluginUpdateCallback bool
	hasPluginEndCallback    bool
	hasPluginPanicCallback  bool
}

var plugify = Plugify{
	Id:           -1,
	Name:         "",
	FullName:     "",
	Description:  "",
	Version:      "",
	Author:       "",
	Website:      "",
	BaseDir:      "",
	ConfigsDir:   "",
	DataDir:      "",
	LogsDir:      "",
	Dependencies: []string{},

	fnPluginStartCallback:   func() {},
	fnPluginUpdateCallback:  func(dt float32) {},
	fnPluginEndCallback:     func() {},
	fnPluginPanicCallback:   func() []byte { return []byte{} },
	hasPluginStartCallback:  false,
	hasPluginUpdateCallback: false,
	hasPluginEndCallback:    false,
	hasPluginPanicCallback:  false,
}

var context C.PluginContext

func OnPluginStart(fn PluginStartCallback) {
	plugify.fnPluginStartCallback = fn
	plugify.hasPluginStartCallback = true
}

func OnPluginUpdate(fn PluginUpdateCallback) {
	plugify.fnPluginUpdateCallback = fn
	plugify.hasPluginUpdateCallback = true
}

func OnPluginEnd(fn PluginEndCallback) {
	plugify.fnPluginEndCallback = fn
	plugify.hasPluginEndCallback = true
}

func OnPluginPanic(fn PluginPanicCallback) {
	plugify.fnPluginPanicCallback = fn
	plugify.hasPluginPanicCallback = true
}

func (p *Plugify) FindResource(path string) string {
	C_output := C.Plugify_FindPluginResource(path)
	output := C.GoString(C_output)
	C.Plugify_DeleteCStr(C_output)
	return output
}

//export Plugify_Init
func Plugify_Init(api []unsafe.Pointer, version int32, handle C.PluginHandle) int32 {
	if version < kApiVersion {
		return kApiVersion
	}
	i := 0
	C.Plugify_SetGetMethodPtr(api[i])
	i++
	C.Plugify_SetGetMethodPtr2(api[i])
	i++
	C.Plugify_SetGetBaseDir(api[i])
	i++
	C.Plugify_SetIsModuleLoaded(api[i])
	i++
	C.Plugify_SetIsPluginLoaded(api[i])
	i++
	C.Plugify_SetPrintException(api[i])
	i++

	C.Plugify_SetGetPluginId(api[i])
	i++
	C.Plugify_SetGetPluginName(api[i])
	i++
	C.Plugify_SetGetPluginFullName(api[i])
	i++
	C.Plugify_SetGetPluginDescription(api[i])
	i++
	C.Plugify_SetGetPluginVersion(api[i])
	i++
	C.Plugify_SetGetPluginAuthor(api[i])
	i++
	C.Plugify_SetGetPluginWebsite(api[i])
	i++
	C.Plugify_SetGetPluginBaseDir(api[i])
	i++
	C.Plugify_SetGetPluginConfigsDir(api[i])
	i++
	C.Plugify_SetGetPluginDataDir(api[i])
	i++
	C.Plugify_SetGetPluginLogsDir(api[i])
	i++
	C.Plugify_SetGetPluginDependencies(api[i])
	i++
	C.Plugify_SetGetPluginDependenciesSize(api[i])
	i++
	C.Plugify_SetFindPluginResource(api[i])
	i++

	C.Plugify_SetDeleteCStr(api[i])
	i++
	C.Plugify_SetDeleteCStrArr(api[i])
	i++

	C.Plugify_SetConstructString(api[i])
	i++
	C.Plugify_SetDestroyString(api[i])
	i++
	C.Plugify_SetGetStringData(api[i])
	i++
	C.Plugify_SetGetStringLength(api[i])
	i++
	C.Plugify_SetAssignString(api[i])
	i++

	C.Plugify_SetDestroyVariant(api[i])
	i++

	C.Plugify_SetConstructVectorBool(api[i])
	i++
	C.Plugify_SetConstructVectorChar8(api[i])
	i++
	C.Plugify_SetConstructVectorChar16(api[i])
	i++
	C.Plugify_SetConstructVectorInt8(api[i])
	i++
	C.Plugify_SetConstructVectorInt16(api[i])
	i++
	C.Plugify_SetConstructVectorInt32(api[i])
	i++
	C.Plugify_SetConstructVectorInt64(api[i])
	i++
	C.Plugify_SetConstructVectorUInt8(api[i])
	i++
	C.Plugify_SetConstructVectorUInt16(api[i])
	i++
	C.Plugify_SetConstructVectorUInt32(api[i])
	i++
	C.Plugify_SetConstructVectorUInt64(api[i])
	i++
	C.Plugify_SetConstructVectorPointer(api[i])
	i++
	C.Plugify_SetConstructVectorFloat(api[i])
	i++
	C.Plugify_SetConstructVectorDouble(api[i])
	i++
	C.Plugify_SetConstructVectorString(api[i])
	i++
	C.Plugify_SetConstructVectorVariant(api[i])
	i++
	C.Plugify_SetConstructVectorVector2(api[i])
	i++
	C.Plugify_SetConstructVectorVector3(api[i])
	i++
	C.Plugify_SetConstructVectorVector4(api[i])
	i++
	C.Plugify_SetConstructVectorMatrix4x4(api[i])
	i++

	C.Plugify_SetDestroyVectorBool(api[i])
	i++
	C.Plugify_SetDestroyVectorChar8(api[i])
	i++
	C.Plugify_SetDestroyVectorChar16(api[i])
	i++
	C.Plugify_SetDestroyVectorInt8(api[i])
	i++
	C.Plugify_SetDestroyVectorInt16(api[i])
	i++
	C.Plugify_SetDestroyVectorInt32(api[i])
	i++
	C.Plugify_SetDestroyVectorInt64(api[i])
	i++
	C.Plugify_SetDestroyVectorUInt8(api[i])
	i++
	C.Plugify_SetDestroyVectorUInt16(api[i])
	i++
	C.Plugify_SetDestroyVectorUInt32(api[i])
	i++
	C.Plugify_SetDestroyVectorUInt64(api[i])
	i++
	C.Plugify_SetDestroyVectorPointer(api[i])
	i++
	C.Plugify_SetDestroyVectorFloat(api[i])
	i++
	C.Plugify_SetDestroyVectorDouble(api[i])
	i++
	C.Plugify_SetDestroyVectorString(api[i])
	i++
	C.Plugify_SetDestroyVectorVariant(api[i])
	i++
	C.Plugify_SetDestroyVectorVector2(api[i])
	i++
	C.Plugify_SetDestroyVectorVector3(api[i])
	i++
	C.Plugify_SetDestroyVectorVector4(api[i])
	i++
	C.Plugify_SetDestroyVectorMatrix4x4(api[i])
	i++

	C.Plugify_SetGetVectorSizeBool(api[i])
	i++
	C.Plugify_SetGetVectorSizeChar8(api[i])
	i++
	C.Plugify_SetGetVectorSizeChar16(api[i])
	i++
	C.Plugify_SetGetVectorSizeInt8(api[i])
	i++
	C.Plugify_SetGetVectorSizeInt16(api[i])
	i++
	C.Plugify_SetGetVectorSizeInt32(api[i])
	i++
	C.Plugify_SetGetVectorSizeInt64(api[i])
	i++
	C.Plugify_SetGetVectorSizeUInt8(api[i])
	i++
	C.Plugify_SetGetVectorSizeUInt16(api[i])
	i++
	C.Plugify_SetGetVectorSizeUInt32(api[i])
	i++
	C.Plugify_SetGetVectorSizeUInt64(api[i])
	i++
	C.Plugify_SetGetVectorSizePointer(api[i])
	i++
	C.Plugify_SetGetVectorSizeFloat(api[i])
	i++
	C.Plugify_SetGetVectorSizeDouble(api[i])
	i++
	C.Plugify_SetGetVectorSizeString(api[i])
	i++
	C.Plugify_SetGetVectorSizeVariant(api[i])
	i++
	C.Plugify_SetGetVectorSizeVector2(api[i])
	i++
	C.Plugify_SetGetVectorSizeVector3(api[i])
	i++
	C.Plugify_SetGetVectorSizeVector4(api[i])
	i++
	C.Plugify_SetGetVectorSizeMatrix4x4(api[i])
	i++

	C.Plugify_SetGetVectorDataBool(api[i])
	i++
	C.Plugify_SetGetVectorDataChar8(api[i])
	i++
	C.Plugify_SetGetVectorDataChar16(api[i])
	i++
	C.Plugify_SetGetVectorDataInt8(api[i])
	i++
	C.Plugify_SetGetVectorDataInt16(api[i])
	i++
	C.Plugify_SetGetVectorDataInt32(api[i])
	i++
	C.Plugify_SetGetVectorDataInt64(api[i])
	i++
	C.Plugify_SetGetVectorDataUInt8(api[i])
	i++
	C.Plugify_SetGetVectorDataUInt16(api[i])
	i++
	C.Plugify_SetGetVectorDataUInt32(api[i])
	i++
	C.Plugify_SetGetVectorDataUInt64(api[i])
	i++
	C.Plugify_SetGetVectorDataPointer(api[i])
	i++
	C.Plugify_SetGetVectorDataFloat(api[i])
	i++
	C.Plugify_SetGetVectorDataDouble(api[i])
	i++
	C.Plugify_SetGetVectorDataString(api[i])
	i++
	C.Plugify_SetGetVectorDataVariant(api[i])
	i++
	C.Plugify_SetGetVectorDataVector2(api[i])
	i++
	C.Plugify_SetGetVectorDataVector3(api[i])
	i++
	C.Plugify_SetGetVectorDataVector4(api[i])
	i++
	C.Plugify_SetGetVectorDataMatrix4x4(api[i])
	i++

	C.Plugify_SetAssignVectorBool(api[i])
	i++
	C.Plugify_SetAssignVectorChar8(api[i])
	i++
	C.Plugify_SetAssignVectorChar16(api[i])
	i++
	C.Plugify_SetAssignVectorInt8(api[i])
	i++
	C.Plugify_SetAssignVectorInt16(api[i])
	i++
	C.Plugify_SetAssignVectorInt32(api[i])
	i++
	C.Plugify_SetAssignVectorInt64(api[i])
	i++
	C.Plugify_SetAssignVectorUInt8(api[i])
	i++
	C.Plugify_SetAssignVectorUInt16(api[i])
	i++
	C.Plugify_SetAssignVectorUInt32(api[i])
	i++
	C.Plugify_SetAssignVectorUInt64(api[i])
	i++
	C.Plugify_SetAssignVectorPointer(api[i])
	i++
	C.Plugify_SetAssignVectorFloat(api[i])
	i++
	C.Plugify_SetAssignVectorDouble(api[i])
	i++
	C.Plugify_SetAssignVectorString(api[i])
	i++
	C.Plugify_SetAssignVectorVariant(api[i])
	i++
	C.Plugify_SetAssignVectorVector2(api[i])
	i++
	C.Plugify_SetAssignVectorVector3(api[i])
	i++
	C.Plugify_SetAssignVectorVector4(api[i])
	i++
	C.Plugify_SetAssignVectorMatrix4x4(api[i])
	i++

	C.Plugify_SetNewCall(api[i])
	i++
	C.Plugify_SetDeleteCall(api[i])
	i++
	C.Plugify_SetGetCallFunction(api[i])
	i++
	C.Plugify_SetGetCallError(api[i])
	i++

	C.Plugify_SetNewCallback(api[i])
	i++
	C.Plugify_SetDeleteCallback(api[i])
	i++
	C.Plugify_SetGetCallbackFunction(api[i])
	i++
	C.Plugify_SetGetCallbackError(api[i])
	i++

	C.Plugify_SetGetMethodParamCount(api[i])
	i++
	C.Plugify_SetGetMethodParamType(api[i])
	i++
	C.Plugify_SetGetMethodPrototype(api[i])
	i++
	C.Plugify_SetGetMethodEnum(api[i])
	i++

	C.pluginHandle = handle

	plugify.Id = int64(C.Plugify_GetPluginId())
	plugify.Name = C.GoString(C.Plugify_GetPluginName())
	plugify.FullName = C.GoString(C.Plugify_GetPluginFullName())
	plugify.Description = C.GoString(C.Plugify_GetPluginDescription())
	plugify.Version = C.GoString(C.Plugify_GetPluginVersion())
	plugify.Author = C.GoString(C.Plugify_GetPluginAuthor())
	plugify.Website = C.GoString(C.Plugify_GetPluginWebsite())

	baseDir := C.Plugify_GetPluginBaseDir()
	plugify.BaseDir = C.GoString(baseDir)
	C.Plugify_DeleteCStr(baseDir)

	configsDir := C.Plugify_GetPluginConfigsDir()
	plugify.ConfigsDir = C.GoString(configsDir)
	C.Plugify_DeleteCStr(configsDir)

	dataDir := C.Plugify_GetPluginDataDir()
	plugify.DataDir = C.GoString(dataDir)
	C.Plugify_DeleteCStr(dataDir)

	logsDir := C.Plugify_GetPluginLogsDir()
	plugify.LogsDir = C.GoString(logsDir)
	C.Plugify_DeleteCStr(logsDir)

	dependencies := C.Plugify_GetPluginDependencies()
	plugify.Dependencies = make([]string, int(C.Plugify_GetPluginDependenciesSize()))
	for j := range plugify.Dependencies {
		plugify.Dependencies[j] = C.GoString(*(**C.char)(unsafe.Pointer(uintptr(dependencies) + uintptr(j)*C.sizeof_uintptr_t)))
	}
	C.Plugify_DeleteCStrArr(dependencies)

	context = C.PluginContext{
		hasUpdate: C.bool(plugify.hasPluginUpdateCallback),
		hasStart:  C.bool(plugify.hasPluginStartCallback),
		hasEnd:    C.bool(plugify.hasPluginEndCallback),
		hasPanic:  C.bool(plugify.hasPluginPanicCallback),
	}

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

	clear(functionMap)

	for _, v := range calls {
		C.Plugify_DeleteCall(v)
	}
	clear(calls)

	for _, v := range callbacks {
		C.Plugify_DeleteCallback(v)
	}
	clear(callbacks)
}

//export Plugify_PluginContext
func Plugify_PluginContext() *C.PluginContext {
	return &context
}

func panicker(v any) {
	msg := fmt.Sprintf("%v", v)
	stack := plugify.fnPluginPanicCallback()
	if len(stack) > 0 {
		msg += fmt.Sprintf("\nStack Trace: \n%s", stack)
	}
	C.Plugify_PrintException(msg)
	panic(v)
}
