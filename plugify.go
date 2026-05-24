package plugify

/*
#cgo LDFLAGS: -L${SRCDIR}/libplugify.a
#include "plugify.h"
*/
import "C"
import (
	"fmt"
	"path/filepath"
	"runtime"
	"runtime/debug"
	"unsafe"
)

const kApiVersion = 3

var context C.PluginContext

//export plugify_PluginInit
func plugify_PluginInit(api []unsafe.Pointer, version int32, handle C.PluginHandle) int32 {
	if version < kApiVersion {
		return kApiVersion
	}
	i := 0
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
	C.Plugify_SetIsLoaded(api[i])
	i++
	C.Plugify_SetLog(api[i])
	i++
	C.Plugify_SetIsLogging(api[i])
	i++
	C.Plugify_SetBeginZone(api[i])
	i++
	C.Plugify_SetEndZone(api[i])
	i++
	C.Plugify_SetIsProfiling(api[i])
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

	baseStr := C.Plugify_GetBaseDir()
	baseDir = GetStringData(&baseStr)
	C.Plugify_DestroyString(&baseStr)

	extensionsStr := C.Plugify_GetExtensionsDir()
	extensionsDir = GetStringData(&extensionsStr)
	C.Plugify_DestroyString(&extensionsStr)

	configsStr := C.Plugify_GetConfigsDir()
	configsDir = GetStringData(&configsStr)
	C.Plugify_DestroyString(&configsStr)

	dataStr := C.Plugify_GetDataDir()
	dataDir = GetStringData(&dataStr)
	C.Plugify_DestroyString(&dataStr)

	logsStr := C.Plugify_GetLogsDir()
	logsDir = GetStringData(&logsStr)
	C.Plugify_DestroyString(&logsStr)

	cacheStr := C.Plugify_GetCacheDir()
	cacheDir = GetStringData(&cacheStr)
	C.Plugify_DestroyString(&cacheStr)

	C.pluginHandle = handle

	plugin.id = int64(C.Plugify_GetPluginId())

	name := C.Plugify_GetPluginName()
	plugin.name = GetStringData(&name)
	C.Plugify_DestroyString(&name)

	description := C.Plugify_GetPluginDescription()
	plugin.description = GetStringData(&description)
	C.Plugify_DestroyString(&description)

	versions := C.Plugify_GetPluginVersion()
	plugin.version = GetStringData(&versions)
	C.Plugify_DestroyString(&versions)

	author := C.Plugify_GetPluginAuthor()
	plugin.author = GetStringData(&author)
	C.Plugify_DestroyString(&author)

	website := C.Plugify_GetPluginWebsite()
	plugin.website = GetStringData(&website)
	C.Plugify_DestroyString(&website)

	license := C.Plugify_GetPluginLicense()
	plugin.license = GetStringData(&license)
	C.Plugify_DestroyString(&license)

	location := C.Plugify_GetPluginLocation()
	plugin.location = GetStringData(&location)
	C.Plugify_DestroyString(&location)

	dependencies := C.Plugify_GetPluginDependencies()
	plugin.dependencies = GetVectorDataString(&dependencies)
	C.Plugify_DestroyVectorString(&dependencies)

	isProfiling = bool(C.Plugify_IsProfiling())
	isLogging = bool(C.Plugify_IsLogging())

	context = C.PluginContext{
		hasUpdate: C.bool(plugin.hasPluginUpdateCallback),
		hasStart:  C.bool(plugin.hasPluginStartCallback),
		hasEnd:    C.bool(plugin.hasPluginEndCallback),
	}

	return 0
}

//export plugify_PluginStart
func plugify_PluginStart() C.PluginResult {
	plugin.loaded = true
	var err error
	Block{
		Try: func() {
			err = plugin.fnPluginStartCallback()
		},
		Catch: func(exc Exception) {
			err = fmt.Errorf("%v", exc)
		},
	}.Do()
	return result(err)
}

//export plugify_PluginUpdate
func plugify_PluginUpdate(dt float32) C.PluginResult {
	var err error
	Block{
		Try: func() {
			err = plugin.fnPluginUpdateCallback(dt)
		},
		Catch: func(exc Exception) {
			err = fmt.Errorf("%v", exc)
		},
	}.Do()
	return result(err)
}

//export plugify_PluginEnd
func plugify_PluginEnd() C.PluginResult {
	var err error
	Block{
		Try: func() {
			err = plugin.fnPluginEndCallback()
		},
		Catch: func(exc Exception) {
			err = fmt.Errorf("%v", exc)
		},
		Finally: func() {
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

			plugin.loaded = false
		},
	}.Do()
	return result(err)
}

//export plugify_PluginContext
func plugify_PluginContext() *C.PluginContext {
	return &context
}

func stacktrace(err any) {
	msg := fmt.Sprintf("%v", err)
	stack := debug.Stack()
	if len(stack) > 0 {
		msg += fmt.Sprintf("\nStack Trace: \n%s", stack)
	}
	Log(msg, Error, 3)
}

func panicker(err any) {
	stacktrace(err)
	panic(err)
}

func caller(skip int) (line int, file string, funk string) {
	pc, file, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}
	return line, filepath.Base(file), runtime.FuncForPC(pc).Name()
}

func result(err error) C.PluginResult {
	if err != nil {
		return C.PluginResult{
			code:    C.Failed,
			message: ConstructString(err.Error()),
		}
	}

	return C.PluginResult{
		code:    C.Ok,
		message: ConstructString(""),
	}
}
