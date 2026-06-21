package plugify

/*
#include "plugify.h"
*/
import "C"
import (
	"fmt"
	"runtime"
	"unsafe"
)

const ApiVersion = 3

//export plugify_PluginInit
func plugify_PluginInit(api []unsafe.Pointer, version int, name string) int {
	if version < ApiVersion {
		return ApiVersion
	}

	if len(api) != 0 {
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

		C.Plugify_SetGetPlugin(api[i])
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
		baseDir = GetStringData[string](&baseStr)
		C.Plugify_DestroyString(&baseStr)

		extensionsStr := C.Plugify_GetExtensionsDir()
		extensionsDir = GetStringData[string](&extensionsStr)
		C.Plugify_DestroyString(&extensionsStr)

		configsStr := C.Plugify_GetConfigsDir()
		configsDir = GetStringData[string](&configsStr)
		C.Plugify_DestroyString(&configsStr)

		dataStr := C.Plugify_GetDataDir()
		dataDir = GetStringData[string](&dataStr)
		C.Plugify_DestroyString(&dataStr)

		logsStr := C.Plugify_GetLogsDir()
		logsDir = GetStringData[string](&logsStr)
		C.Plugify_DestroyString(&logsStr)

		cacheStr := C.Plugify_GetCacheDir()
		cacheDir = GetStringData[string](&cacheStr)
		C.Plugify_DestroyString(&cacheStr)

		isProfiling = bool(C.Plugify_IsProfiling())
		isLogging = bool(C.Plugify_IsLogging())
	}

	if len(name) != 0 {
		pluginsMap[name].onInit(name)
	}

	return 0
}

//export plugify_PluginStart
func plugify_PluginStart(name string) C.PluginResult {
	var err error
	Block{
		Try: func() {
			err = pluginsMap[name].onStart()
		},
		Catch: func(exc Exception) {
			err = fmt.Errorf("%v", exc)
		},
	}.Do()
	return result(err)
}

//export plugify_PluginUpdate
func plugify_PluginUpdate(name string, dt float32) C.PluginResult {
	var err error
	Block{
		Try: func() {
			err = pluginsMap[name].onUpdate(dt)
		},
		Catch: func(exc Exception) {
			err = fmt.Errorf("%v", exc)
		},
	}.Do()
	return result(err)
}

//export plugify_PluginEnd
func plugify_PluginEnd(name string) C.PluginResult {
	var err error
	Block{
		Try: func() {
			err = pluginsMap[name].onEnd()
		},
		Catch: func(exc Exception) {
			err = fmt.Errorf("%v", exc)
		},
	}.Do()
	return result(err)
}

//export plugify_PluginContext
func plugify_PluginContext(name string) C.PluginContext {
	return C.PluginContext{
		hasUpdate: C.bool(pluginsMap[name].Updating()),
		hasStart:  C.bool(true),
		hasEnd:    C.bool(true),
	}
}

//export plugify_PluginShutdown
func plugify_PluginShutdown() {
	if len(functionMap) != 0 {
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
	}
}

//export plugify_PluginCall
func plugify_PluginCall(method C.MethodHandle, data unsafe.Pointer, params *C.Parameters, count C.size_t, ret *C.Return) {
	Block{
		Try: func() {
			internalCall(method, data, params, count, ret)
		},
		Catch: func(exc Exception) {
			stacktrace(exc)
		},
	}.Do()
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
