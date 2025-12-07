package plugify

/*
#cgo LDFLAGS: -L${SRCDIR}/libplugify.a
#include "plugify.h"
*/
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

const kApiVersion = 1

type PluginStartCallback func()
type PluginUpdateCallback func(dt float32)
type PluginEndCallback func()
type PluginPanicCallback func() []byte

type PluginInfo struct {
	Id           int64
	Name         string
	Description  string
	Version      string
	Author       string
	Website      string
	License      string
	Location     string
	Dependencies []string

	fnPluginStartCallback   PluginStartCallback
	fnPluginUpdateCallback  PluginUpdateCallback
	fnPluginEndCallback     PluginEndCallback
	fnPluginPanicCallback   PluginPanicCallback
	hasPluginStartCallback  bool
	hasPluginUpdateCallback bool
	hasPluginEndCallback    bool
	hasPluginPanicCallback  bool

	Loaded bool
}

var Plugin = PluginInfo{
	Id:           -1,
	Name:         "",
	Description:  "",
	Version:      "",
	Author:       "",
	Website:      "",
	License:      "",
	Dependencies: []string{},

	fnPluginStartCallback:   func() {},
	fnPluginUpdateCallback:  func(dt float32) {},
	fnPluginEndCallback:     func() {},
	fnPluginPanicCallback:   func() []byte { return []byte{} },
	hasPluginStartCallback:  false,
	hasPluginUpdateCallback: false,
	hasPluginEndCallback:    false,
	hasPluginPanicCallback:  false,

	Loaded: false,
}

var context C.PluginContext

func OnPluginStart(fn PluginStartCallback) {
	Plugin.fnPluginStartCallback = fn
	Plugin.hasPluginStartCallback = true
}

func OnPluginUpdate(fn PluginUpdateCallback) {
	Plugin.fnPluginUpdateCallback = fn
	Plugin.hasPluginUpdateCallback = true
}

func OnPluginEnd(fn PluginEndCallback) {
	Plugin.fnPluginEndCallback = fn
	Plugin.hasPluginEndCallback = true
}

func OnPluginPanic(fn PluginPanicCallback) {
	Plugin.fnPluginPanicCallback = fn
	Plugin.hasPluginPanicCallback = true
}

var BaseDir = ""
var ExtensionsDir = ""
var ConfigsDir = ""
var DataDir = ""
var LogsDir = ""
var CacheDir = ""

func IsExtensionLoaded(name string, constraint string) bool {
	return bool(C.Plugify_IsExtensionLoaded(name, constraint))
}

func PrintException(msg string) {
	C.Plugify_PrintException(msg)
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
	C.Plugify_SetGetExtensionsDir(api[i])
	i++
	C.Plugify_SetGetConfigsDir(api[i])
	i++
	C.Plugify_SetGetDataDir(api[i])
	i++
	C.Plugify_SetGetLogsDir(api[i])
	i++
	C.Plugify_SetGetCacheDir(api[i])
	i++
	C.Plugify_SetIsExtensionLoaded(api[i])
	i++
	C.Plugify_SetPrintException(api[i])
	i++

	C.Plugify_SetGetPluginId(api[i])
	i++
	C.Plugify_SetGetPluginName(api[i])
	i++
	C.Plugify_SetGetPluginDescription(api[i])
	i++
	C.Plugify_SetGetPluginVersion(api[i])
	i++
	C.Plugify_SetGetPluginAuthor(api[i])
	i++
	C.Plugify_SetGetPluginWebsite(api[i])
	i++
	C.Plugify_SetGetPluginLicense(api[i])
	i++
	C.Plugify_SetGetPluginLocation(api[i])
	i++
	C.Plugify_SetGetPluginDependencies(api[i])
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

	baseDir := C.Plugify_GetBaseDir()
	BaseDir = GetStringData(&baseDir)
	C.Plugify_DestroyString(&baseDir)

	extensionsDir := C.Plugify_GetExtensionsDir()
	ExtensionsDir = GetStringData(&extensionsDir)
	C.Plugify_DestroyString(&extensionsDir)

	configsDir := C.Plugify_GetConfigsDir()
	ConfigsDir = GetStringData(&configsDir)
	C.Plugify_DestroyString(&configsDir)

	dataDir := C.Plugify_GetDataDir()
	DataDir = GetStringData(&dataDir)
	C.Plugify_DestroyString(&dataDir)

	logsDir := C.Plugify_GetLogsDir()
	LogsDir = GetStringData(&logsDir)
	C.Plugify_DestroyString(&logsDir)

	cacheDir := C.Plugify_GetCacheDir()
	CacheDir = GetStringData(&cacheDir)
	C.Plugify_DestroyString(&cacheDir)

	C.pluginHandle = handle

	Plugin.Id = int64(C.Plugify_GetPluginId())
	name := C.Plugify_GetPluginName()
	Plugin.Name = GetStringData(&name)
	C.Plugify_DestroyString(&name)

	description := C.Plugify_GetPluginDescription()
	Plugin.Description = GetStringData(&description)
	C.Plugify_DestroyString(&description)

	versions := C.Plugify_GetPluginVersion()
	Plugin.Version = GetStringData(&versions)
	C.Plugify_DestroyString(&versions)

	author := C.Plugify_GetPluginAuthor()
	Plugin.Author = GetStringData(&author)
	C.Plugify_DestroyString(&author)

	website := C.Plugify_GetPluginWebsite()
	Plugin.Website = GetStringData(&website)
	C.Plugify_DestroyString(&website)

	license := C.Plugify_GetPluginLicense()
	Plugin.License = GetStringData(&license)
	C.Plugify_DestroyString(&license)

	location := C.Plugify_GetPluginLocation()
	Plugin.Location = GetStringData(&location)
	C.Plugify_DestroyString(&location)

	dependencies := C.Plugify_GetPluginDependencies()
	Plugin.Dependencies = GetVectorDataString(&dependencies)
	C.Plugify_DestroyVectorString(&dependencies)

	context = C.PluginContext{
		hasUpdate: C.bool(Plugin.hasPluginUpdateCallback),
		hasStart:  C.bool(Plugin.hasPluginStartCallback),
		hasEnd:    C.bool(Plugin.hasPluginEndCallback),
		hasPanic:  C.bool(Plugin.hasPluginPanicCallback),
	}

	return 0
}

//export Plugify_PluginStart
func Plugify_PluginStart() {
	Plugin.Loaded = true

	Plugin.fnPluginStartCallback()
}

//export Plugify_PluginUpdate
func Plugify_PluginUpdate(dt float32) {
	Plugin.fnPluginUpdateCallback(dt)
}

//export Plugify_PluginEnd
func Plugify_PluginEnd() {
	Plugin.fnPluginEndCallback()

	clear(functionMap)

	for _, v := range calls {
		C.Plugify_DeleteCall(v)
	}
	clear(calls)

	for _, v := range callbacks {
		C.Plugify_DeleteCallback(v)
	}
	clear(callbacks)

	runtime.GC()
	runtime.Gosched()

	Plugin.Loaded = false
}

//export Plugify_PluginContext
func Plugify_PluginContext() *C.PluginContext {
	return &context
}

func panicker(v any) {
	msg := fmt.Sprintf("%v", v)
	stack := Plugin.fnPluginPanicCallback()
	if len(stack) > 0 {
		msg += fmt.Sprintf("\nStack Trace: \n%s", stack)
	}
	PrintException(msg)
	panic(v)
}
