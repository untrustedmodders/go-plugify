package plugify

/*
#include "plugify.h"
//#cgo noescape Plugify_GetMethodPtr
//#cgo noescape Plugify_GetMethodPtr2
//#cgo noescape Plugify_IsModuleLoaded
//#cgo noescape Plugify_IsPluginLoaded
#cgo noescape Plugify_PrintException
#cgo noescape Plugify_GetPluginId
#cgo noescape Plugify_GetPluginName
#cgo noescape Plugify_GetPluginFullName
#cgo noescape Plugify_GetPluginDescription
#cgo noescape Plugify_GetPluginVersion
#cgo noescape Plugify_GetPluginAuthor
#cgo noescape Plugify_GetPluginWebsite
#cgo noescape Plugify_GetPluginBaseDir
#cgo noescape Plugify_GetPluginConfigsDir
#cgo noescape Plugify_GetPluginDataDir
#cgo noescape Plugify_GetPluginLogsDir
#cgo noescape Plugify_GetPluginDependencies
#cgo noescape Plugify_GetPluginDependenciesSize
#cgo noescape Plugify_FindPluginResource
#cgo noescape Plugify_DeleteCStr
#cgo noescape Plugify_DeleteCStrArr
#cgo noescape Plugify_ConstructString
#cgo noescape Plugify_DestroyString
#cgo noescape Plugify_GetStringData
#cgo noescape Plugify_GetStringLength
#cgo noescape Plugify_AssignString
#cgo noescape Plugify_DestroyVariant
#cgo noescape Plugify_ConstructVectorBool
#cgo noescape Plugify_ConstructVectorChar8
#cgo noescape Plugify_ConstructVectorChar16
#cgo noescape Plugify_ConstructVectorInt8
#cgo noescape Plugify_ConstructVectorInt16
#cgo noescape Plugify_ConstructVectorInt32
#cgo noescape Plugify_ConstructVectorInt64
#cgo noescape Plugify_ConstructVectorUInt8
#cgo noescape Plugify_ConstructVectorUInt16
#cgo noescape Plugify_ConstructVectorUInt32
#cgo noescape Plugify_ConstructVectorUInt64
#cgo noescape Plugify_ConstructVectorPointer
#cgo noescape Plugify_ConstructVectorFloat
#cgo noescape Plugify_ConstructVectorDouble
#cgo noescape Plugify_ConstructVectorString
#cgo noescape Plugify_ConstructVectorVariant
#cgo noescape Plugify_ConstructVectorVector2
#cgo noescape Plugify_ConstructVectorVector3
#cgo noescape Plugify_ConstructVectorVector4
#cgo noescape Plugify_ConstructVectorMatrix4x4
#cgo noescape Plugify_DestroyVectorBool
#cgo noescape Plugify_DestroyVectorChar8
#cgo noescape Plugify_DestroyVectorChar16
#cgo noescape Plugify_DestroyVectorInt8
#cgo noescape Plugify_DestroyVectorInt16
#cgo noescape Plugify_DestroyVectorInt32
#cgo noescape Plugify_DestroyVectorInt64
#cgo noescape Plugify_DestroyVectorUInt8
#cgo noescape Plugify_DestroyVectorUInt16
#cgo noescape Plugify_DestroyVectorUInt32
#cgo noescape Plugify_DestroyVectorUInt64
#cgo noescape Plugify_DestroyVectorPointer
#cgo noescape Plugify_DestroyVectorFloat
#cgo noescape Plugify_DestroyVectorDouble
#cgo noescape Plugify_DestroyVectorString
#cgo noescape Plugify_DestroyVectorVariant
#cgo noescape Plugify_DestroyVectorVector2
#cgo noescape Plugify_DestroyVectorVector3
#cgo noescape Plugify_DestroyVectorVector4
#cgo noescape Plugify_DestroyVectorMatrix4x4
#cgo noescape Plugify_GetVectorSizeBool
#cgo noescape Plugify_GetVectorSizeChar8
#cgo noescape Plugify_GetVectorSizeChar16
#cgo noescape Plugify_GetVectorSizeInt8
#cgo noescape Plugify_GetVectorSizeInt16
#cgo noescape Plugify_GetVectorSizeInt32
#cgo noescape Plugify_GetVectorSizeInt64
#cgo noescape Plugify_GetVectorSizeUInt8
#cgo noescape Plugify_GetVectorSizeUInt16
#cgo noescape Plugify_GetVectorSizeUInt32
#cgo noescape Plugify_GetVectorSizeUInt64
#cgo noescape Plugify_GetVectorSizePointer
#cgo noescape Plugify_GetVectorSizeFloat
#cgo noescape Plugify_GetVectorSizeDouble
#cgo noescape Plugify_GetVectorSizeString
#cgo noescape Plugify_GetVectorSizeVariant
#cgo noescape Plugify_GetVectorSizeVector2
#cgo noescape Plugify_GetVectorSizeVector3
#cgo noescape Plugify_GetVectorSizeVector4
#cgo noescape Plugify_GetVectorSizeMatrix4x4
#cgo noescape Plugify_GetVectorDataBool
#cgo noescape Plugify_GetVectorDataChar8
#cgo noescape Plugify_GetVectorDataChar16
#cgo noescape Plugify_GetVectorDataInt8
#cgo noescape Plugify_GetVectorDataInt16
#cgo noescape Plugify_GetVectorDataInt32
#cgo noescape Plugify_GetVectorDataInt64
#cgo noescape Plugify_GetVectorDataUInt8
#cgo noescape Plugify_GetVectorDataUInt16
#cgo noescape Plugify_GetVectorDataUInt32
#cgo noescape Plugify_GetVectorDataUInt64
#cgo noescape Plugify_GetVectorDataPointer
#cgo noescape Plugify_GetVectorDataFloat
#cgo noescape Plugify_GetVectorDataDouble
#cgo noescape Plugify_GetVectorDataString
#cgo noescape Plugify_GetVectorDataVariant
#cgo noescape Plugify_GetVectorDataVector2
#cgo noescape Plugify_GetVectorDataVector3
#cgo noescape Plugify_GetVectorDataVector4
#cgo noescape Plugify_GetVectorDataMatrix4x4
#cgo noescape Plugify_AssignVectorBool
#cgo noescape Plugify_AssignVectorChar8
#cgo noescape Plugify_AssignVectorChar16
#cgo noescape Plugify_AssignVectorInt8
#cgo noescape Plugify_AssignVectorInt16
#cgo noescape Plugify_AssignVectorInt32
#cgo noescape Plugify_AssignVectorInt64
#cgo noescape Plugify_AssignVectorUInt8
#cgo noescape Plugify_AssignVectorUInt16
#cgo noescape Plugify_AssignVectorUInt32
#cgo noescape Plugify_AssignVectorUInt64
#cgo noescape Plugify_AssignVectorPointer
#cgo noescape Plugify_AssignVectorFloat
#cgo noescape Plugify_AssignVectorDouble
#cgo noescape Plugify_AssignVectorString
#cgo noescape Plugify_AssignVectorVariant
#cgo noescape Plugify_AssignVectorVector2
#cgo noescape Plugify_AssignVectorVector3
#cgo noescape Plugify_AssignVectorVector4
#cgo noescape Plugify_AssignVectorMatrix4x4
#cgo noescape Plugify_NewCall
//#cgo noescape Plugify_DeleteCall
#cgo noescape Plugify_GetCallFunction
#cgo noescape Plugify_GetCallError
#cgo noescape Plugify_CallFunction
#cgo noescape Plugify_NewCallback
//#cgo noescape Plugify_DeleteCallback
#cgo noescape Plugify_GetCallbackFunction
#cgo noescape Plugify_GetCallbackError
#cgo noescape Plugify_GetMethodParamCount
#cgo noescape Plugify_GetMethodParamType
#cgo noescape Plugify_GetMethodPrototype
#cgo noescape Plugify_SetGetMethodPtr
#cgo noescape Plugify_SetGetMethodPtr2
#cgo noescape Plugify_SetIsModuleLoaded
#cgo noescape Plugify_SetIsPluginLoaded
#cgo noescape Plugify_SetPrintException
#cgo noescape Plugify_SetGetPluginId
#cgo noescape Plugify_SetGetPluginName
#cgo noescape Plugify_SetGetPluginFullName
#cgo noescape Plugify_SetGetPluginDescription
#cgo noescape Plugify_SetGetPluginVersion
#cgo noescape Plugify_SetGetPluginAuthor
#cgo noescape Plugify_SetGetPluginWebsite
#cgo noescape Plugify_SetGetPluginBaseDir
#cgo noescape Plugify_SetGetPluginConfigsDir
#cgo noescape Plugify_SetGetPluginDataDir
#cgo noescape Plugify_SetGetPluginLogsDir
#cgo noescape Plugify_SetGetPluginDependencies
#cgo noescape Plugify_SetGetPluginDependenciesSize
#cgo noescape Plugify_SetFindPluginResource
#cgo noescape Plugify_SetDeleteCStr
#cgo noescape Plugify_SetDeleteCStrArr
#cgo noescape Plugify_SetConstructString
#cgo noescape Plugify_SetDestroyString
#cgo noescape Plugify_SetGetStringData
#cgo noescape Plugify_SetGetStringLength
#cgo noescape Plugify_SetAssignString
#cgo noescape Plugify_SetDestroyVariant
#cgo noescape Plugify_SetConstructVectorBool
#cgo noescape Plugify_SetConstructVectorChar8
#cgo noescape Plugify_SetConstructVectorChar16
#cgo noescape Plugify_SetConstructVectorInt8
#cgo noescape Plugify_SetConstructVectorInt16
#cgo noescape Plugify_SetConstructVectorInt32
#cgo noescape Plugify_SetConstructVectorInt64
#cgo noescape Plugify_SetConstructVectorUInt8
#cgo noescape Plugify_SetConstructVectorUInt16
#cgo noescape Plugify_SetConstructVectorUInt32
#cgo noescape Plugify_SetConstructVectorUInt64
#cgo noescape Plugify_SetConstructVectorPointer
#cgo noescape Plugify_SetConstructVectorFloat
#cgo noescape Plugify_SetConstructVectorDouble
#cgo noescape Plugify_SetConstructVectorString
#cgo noescape Plugify_SetConstructVectorVariant
#cgo noescape Plugify_SetConstructVectorVector2
#cgo noescape Plugify_SetConstructVectorVector3
#cgo noescape Plugify_SetConstructVectorVector4
#cgo noescape Plugify_SetConstructVectorMatrix4x4
#cgo noescape Plugify_SetDestroyVectorBool
#cgo noescape Plugify_SetDestroyVectorChar8
#cgo noescape Plugify_SetDestroyVectorChar16
#cgo noescape Plugify_SetDestroyVectorInt8
#cgo noescape Plugify_SetDestroyVectorInt16
#cgo noescape Plugify_SetDestroyVectorInt32
#cgo noescape Plugify_SetDestroyVectorInt64
#cgo noescape Plugify_SetDestroyVectorUInt8
#cgo noescape Plugify_SetDestroyVectorUInt16
#cgo noescape Plugify_SetDestroyVectorUInt32
#cgo noescape Plugify_SetDestroyVectorUInt64
#cgo noescape Plugify_SetDestroyVectorPointer
#cgo noescape Plugify_SetDestroyVectorFloat
#cgo noescape Plugify_SetDestroyVectorDouble
#cgo noescape Plugify_SetDestroyVectorString
#cgo noescape Plugify_SetDestroyVectorVariant
#cgo noescape Plugify_SetDestroyVectorVector2
#cgo noescape Plugify_SetDestroyVectorVector3
#cgo noescape Plugify_SetDestroyVectorVector4
#cgo noescape Plugify_SetDestroyVectorMatrix4x4
#cgo noescape Plugify_SetGetVectorSizeBool
#cgo noescape Plugify_SetGetVectorSizeChar8
#cgo noescape Plugify_SetGetVectorSizeChar16
#cgo noescape Plugify_SetGetVectorSizeInt8
#cgo noescape Plugify_SetGetVectorSizeInt16
#cgo noescape Plugify_SetGetVectorSizeInt32
#cgo noescape Plugify_SetGetVectorSizeInt64
#cgo noescape Plugify_SetGetVectorSizeUInt8
#cgo noescape Plugify_SetGetVectorSizeUInt16
#cgo noescape Plugify_SetGetVectorSizeUInt32
#cgo noescape Plugify_SetGetVectorSizeUInt64
#cgo noescape Plugify_SetGetVectorSizePointer
#cgo noescape Plugify_SetGetVectorSizeFloat
#cgo noescape Plugify_SetGetVectorSizeDouble
#cgo noescape Plugify_SetGetVectorSizeString
#cgo noescape Plugify_SetGetVectorSizeVariant
#cgo noescape Plugify_SetGetVectorSizeVector2
#cgo noescape Plugify_SetGetVectorSizeVector3
#cgo noescape Plugify_SetGetVectorSizeVector4
#cgo noescape Plugify_SetGetVectorSizeMatrix4x4
#cgo noescape Plugify_SetGetVectorDataBool
#cgo noescape Plugify_SetGetVectorDataChar8
#cgo noescape Plugify_SetGetVectorDataChar16
#cgo noescape Plugify_SetGetVectorDataInt8
#cgo noescape Plugify_SetGetVectorDataInt16
#cgo noescape Plugify_SetGetVectorDataInt32
#cgo noescape Plugify_SetGetVectorDataInt64
#cgo noescape Plugify_SetGetVectorDataUInt8
#cgo noescape Plugify_SetGetVectorDataUInt16
#cgo noescape Plugify_SetGetVectorDataUInt32
#cgo noescape Plugify_SetGetVectorDataUInt64
#cgo noescape Plugify_SetGetVectorDataPointer
#cgo noescape Plugify_SetGetVectorDataFloat
#cgo noescape Plugify_SetGetVectorDataDouble
#cgo noescape Plugify_SetGetVectorDataString
#cgo noescape Plugify_SetGetVectorDataVariant
#cgo noescape Plugify_SetGetVectorDataVector2
#cgo noescape Plugify_SetGetVectorDataVector3
#cgo noescape Plugify_SetGetVectorDataVector4
#cgo noescape Plugify_SetGetVectorDataMatrix4x4
#cgo noescape Plugify_SetAssignVectorBool
#cgo noescape Plugify_SetAssignVectorChar8
#cgo noescape Plugify_SetAssignVectorChar16
#cgo noescape Plugify_SetAssignVectorInt8
#cgo noescape Plugify_SetAssignVectorInt16
#cgo noescape Plugify_SetAssignVectorInt32
#cgo noescape Plugify_SetAssignVectorInt64
#cgo noescape Plugify_SetAssignVectorUInt8
#cgo noescape Plugify_SetAssignVectorUInt16
#cgo noescape Plugify_SetAssignVectorUInt32
#cgo noescape Plugify_SetAssignVectorUInt64
#cgo noescape Plugify_SetAssignVectorPointer
#cgo noescape Plugify_SetAssignVectorFloat
#cgo noescape Plugify_SetAssignVectorDouble
#cgo noescape Plugify_SetAssignVectorString
#cgo noescape Plugify_SetAssignVectorVariant
#cgo noescape Plugify_SetAssignVectorVector2
#cgo noescape Plugify_SetAssignVectorVector3
#cgo noescape Plugify_SetAssignVectorVector4
#cgo noescape Plugify_SetAssignVectorMatrix4x4
#cgo noescape Plugify_SetNewCall
#cgo noescape Plugify_SetDeleteCall
#cgo noescape Plugify_SetGetCallFunction
#cgo noescape Plugify_SetGetCallError
#cgo noescape Plugify_SetNewCallback
#cgo noescape Plugify_SetDeleteCallback
#cgo noescape Plugify_SetGetCallbackFunction
#cgo noescape Plugify_SetGetCallbackError
#cgo noescape Plugify_SetGetMethodParamCount
#cgo noescape Plugify_SetGetMethodParamType
#cgo noescape Plugify_SetGetMethodPrototype


//#cgo nocallback Plugify_GetMethodPtr
//#cgo nocallback Plugify_GetMethodPtr2
//#cgo nocallback Plugify_IsModuleLoaded
//#cgo nocallback Plugify_IsPluginLoaded
//#cgo nocallback Plugify_PrintException
//#cgo nocallback Plugify_GetPluginId
//#cgo nocallback Plugify_GetPluginName
//#cgo nocallback Plugify_GetPluginFullName
//#cgo nocallback Plugify_GetPluginDescription
//#cgo nocallback Plugify_GetPluginVersion
//#cgo nocallback Plugify_GetPluginAuthor
//#cgo nocallback Plugify_GetPluginWebsite
//#cgo nocallback Plugify_GetPluginBaseDir
//#cgo nocallback Plugify_GetPluginConfigsDir
//#cgo nocallback Plugify_GetPluginDataDir
//#cgo nocallback Plugify_GetPluginLogsDir
//#cgo nocallback Plugify_GetPluginDependencies
//#cgo nocallback Plugify_GetPluginDependenciesSize
//#cgo nocallback Plugify_FindPluginResource
//#cgo nocallback Plugify_DeleteCStr
//#cgo nocallback Plugify_DeleteCStrArr
//#cgo nocallback Plugify_ConstructString
//#cgo nocallback Plugify_DestroyString
//#cgo nocallback Plugify_GetStringData
//#cgo nocallback Plugify_GetStringLength
//#cgo nocallback Plugify_AssignString
//#cgo nocallback Plugify_DestroyVariant
//#cgo nocallback Plugify_ConstructVectorBool
//#cgo nocallback Plugify_ConstructVectorChar8
//#cgo nocallback Plugify_ConstructVectorChar16
//#cgo nocallback Plugify_ConstructVectorInt8
//#cgo nocallback Plugify_ConstructVectorInt16
//#cgo nocallback Plugify_ConstructVectorInt32
//#cgo nocallback Plugify_ConstructVectorInt64
//#cgo nocallback Plugify_ConstructVectorUInt8
//#cgo nocallback Plugify_ConstructVectorUInt16
//#cgo nocallback Plugify_ConstructVectorUInt32
//#cgo nocallback Plugify_ConstructVectorUInt64
//#cgo nocallback Plugify_ConstructVectorPointer
//#cgo nocallback Plugify_ConstructVectorFloat
//#cgo nocallback Plugify_ConstructVectorDouble
//#cgo nocallback Plugify_ConstructVectorString
//#cgo nocallback Plugify_ConstructVectorVariant
//#cgo nocallback Plugify_ConstructVectorVector2
//#cgo nocallback Plugify_ConstructVectorVector3
//#cgo nocallback Plugify_ConstructVectorVector4
//#cgo nocallback Plugify_ConstructVectorMatrix4x4
//#cgo nocallback Plugify_DestroyVectorBool
//#cgo nocallback Plugify_DestroyVectorChar8
//#cgo nocallback Plugify_DestroyVectorChar16
//#cgo nocallback Plugify_DestroyVectorInt8
//#cgo nocallback Plugify_DestroyVectorInt16
//#cgo nocallback Plugify_DestroyVectorInt32
//#cgo nocallback Plugify_DestroyVectorInt64
//#cgo nocallback Plugify_DestroyVectorUInt8
//#cgo nocallback Plugify_DestroyVectorUInt16
//#cgo nocallback Plugify_DestroyVectorUInt32
//#cgo nocallback Plugify_DestroyVectorUInt64
//#cgo nocallback Plugify_DestroyVectorPointer
//#cgo nocallback Plugify_DestroyVectorFloat
//#cgo nocallback Plugify_DestroyVectorDouble
//#cgo nocallback Plugify_DestroyVectorString
//#cgo nocallback Plugify_DestroyVectorVariant
//#cgo nocallback Plugify_DestroyVectorVector2
//#cgo nocallback Plugify_DestroyVectorVector3
//#cgo nocallback Plugify_DestroyVectorVector4
//#cgo nocallback Plugify_DestroyVectorMatrix4x4
//#cgo nocallback Plugify_GetVectorSizeBool
//#cgo nocallback Plugify_GetVectorSizeChar8
//#cgo nocallback Plugify_GetVectorSizeChar16
//#cgo nocallback Plugify_GetVectorSizeInt8
//#cgo nocallback Plugify_GetVectorSizeInt16
//#cgo nocallback Plugify_GetVectorSizeInt32
//#cgo nocallback Plugify_GetVectorSizeInt64
//#cgo nocallback Plugify_GetVectorSizeUInt8
//#cgo nocallback Plugify_GetVectorSizeUInt16
//#cgo nocallback Plugify_GetVectorSizeUInt32
//#cgo nocallback Plugify_GetVectorSizeUInt64
//#cgo nocallback Plugify_GetVectorSizePointer
//#cgo nocallback Plugify_GetVectorSizeFloat
//#cgo nocallback Plugify_GetVectorSizeDouble
//#cgo nocallback Plugify_GetVectorSizeString
//#cgo nocallback Plugify_GetVectorSizeVariant
//#cgo nocallback Plugify_GetVectorSizeVector2
//#cgo nocallback Plugify_GetVectorSizeVector3
//#cgo nocallback Plugify_GetVectorSizeVector4
//#cgo nocallback Plugify_GetVectorSizeMatrix4x4
//#cgo nocallback Plugify_GetVectorDataBool
//#cgo nocallback Plugify_GetVectorDataChar8
//#cgo nocallback Plugify_GetVectorDataChar16
//#cgo nocallback Plugify_GetVectorDataInt8
//#cgo nocallback Plugify_GetVectorDataInt16
//#cgo nocallback Plugify_GetVectorDataInt32
//#cgo nocallback Plugify_GetVectorDataInt64
//#cgo nocallback Plugify_GetVectorDataUInt8
//#cgo nocallback Plugify_GetVectorDataUInt16
//#cgo nocallback Plugify_GetVectorDataUInt32
//#cgo nocallback Plugify_GetVectorDataUInt64
//#cgo nocallback Plugify_GetVectorDataPointer
//#cgo nocallback Plugify_GetVectorDataFloat
//#cgo nocallback Plugify_GetVectorDataDouble
//#cgo nocallback Plugify_GetVectorDataString
//#cgo nocallback Plugify_GetVectorDataVariant
//#cgo nocallback Plugify_GetVectorDataVector2
//#cgo nocallback Plugify_GetVectorDataVector3
//#cgo nocallback Plugify_GetVectorDataVector4
//#cgo nocallback Plugify_GetVectorDataMatrix4x4
//#cgo nocallback Plugify_AssignVectorBool
//#cgo nocallback Plugify_AssignVectorChar8
//#cgo nocallback Plugify_AssignVectorChar16
//#cgo nocallback Plugify_AssignVectorInt8
//#cgo nocallback Plugify_AssignVectorInt16
//#cgo nocallback Plugify_AssignVectorInt32
//#cgo nocallback Plugify_AssignVectorInt64
//#cgo nocallback Plugify_AssignVectorUInt8
//#cgo nocallback Plugify_AssignVectorUInt16
//#cgo nocallback Plugify_AssignVectorUInt32
//#cgo nocallback Plugify_AssignVectorUInt64
//#cgo nocallback Plugify_AssignVectorPointer
//#cgo nocallback Plugify_AssignVectorFloat
//#cgo nocallback Plugify_AssignVectorDouble
//#cgo nocallback Plugify_AssignVectorString
//#cgo nocallback Plugify_AssignVectorVariant
//#cgo nocallback Plugify_AssignVectorVector2
//#cgo nocallback Plugify_AssignVectorVector3
//#cgo nocallback Plugify_AssignVectorVector4
//#cgo nocallback Plugify_AssignVectorMatrix4x4
//#cgo nocallback Plugify_NewCall
//#cgo nocallback Plugify_DeleteCall
//#cgo nocallback Plugify_GetCallFunction
//#cgo nocallback Plugify_GetCallError
//#cgo nocallback Plugify_CallFunction
//#cgo nocallback Plugify_NewCallback
//#cgo nocallback Plugify_DeleteCallback
//#cgo nocallback Plugify_GetCallbackFunction
//#cgo nocallback Plugify_GetCallbackError
//#cgo nocallback Plugify_GetMethodParamCount
//#cgo nocallback Plugify_GetMethodParamType
//#cgo nocallback Plugify_GetMethodPrototype
//#cgo nocallback Plugify_SetGetMethodPtr
//#cgo nocallback Plugify_SetGetMethodPtr2
//#cgo nocallback Plugify_SetIsModuleLoaded
//#cgo nocallback Plugify_SetIsPluginLoaded
//#cgo nocallback Plugify_SetPrintException
//#cgo nocallback Plugify_SetGetPluginId
//#cgo nocallback Plugify_SetGetPluginName
//#cgo nocallback Plugify_SetGetPluginFullName
//#cgo nocallback Plugify_SetGetPluginDescription
//#cgo nocallback Plugify_SetGetPluginVersion
//#cgo nocallback Plugify_SetGetPluginAuthor
//#cgo nocallback Plugify_SetGetPluginWebsite
//#cgo nocallback Plugify_SetGetPluginBaseDir
//#cgo nocallback Plugify_SetGetPluginConfigsDir
//#cgo nocallback Plugify_SetGetPluginDataDir
//#cgo nocallback Plugify_SetGetPluginLogsDir
//#cgo nocallback Plugify_SetGetPluginDependencies
//#cgo nocallback Plugify_SetGetPluginDependenciesSize
//#cgo nocallback Plugify_SetFindPluginResource
//#cgo nocallback Plugify_SetDeleteCStr
//#cgo nocallback Plugify_SetDeleteCStrArr
//#cgo nocallback Plugify_SetConstructString
//#cgo nocallback Plugify_SetDestroyString
//#cgo nocallback Plugify_SetGetStringData
//#cgo nocallback Plugify_SetGetStringLength
//#cgo nocallback Plugify_SetAssignString
//#cgo nocallback Plugify_SetDestroyVariant
//#cgo nocallback Plugify_SetConstructVectorBool
//#cgo nocallback Plugify_SetConstructVectorChar8
//#cgo nocallback Plugify_SetConstructVectorChar16
//#cgo nocallback Plugify_SetConstructVectorInt8
//#cgo nocallback Plugify_SetConstructVectorInt16
//#cgo nocallback Plugify_SetConstructVectorInt32
//#cgo nocallback Plugify_SetConstructVectorInt64
//#cgo nocallback Plugify_SetConstructVectorUInt8
//#cgo nocallback Plugify_SetConstructVectorUInt16
//#cgo nocallback Plugify_SetConstructVectorUInt32
//#cgo nocallback Plugify_SetConstructVectorUInt64
//#cgo nocallback Plugify_SetConstructVectorPointer
//#cgo nocallback Plugify_SetConstructVectorFloat
//#cgo nocallback Plugify_SetConstructVectorDouble
//#cgo nocallback Plugify_SetConstructVectorString
//#cgo nocallback Plugify_SetConstructVectorVariant
//#cgo nocallback Plugify_SetConstructVectorVector2
//#cgo nocallback Plugify_SetConstructVectorVector3
//#cgo nocallback Plugify_SetConstructVectorVector4
//#cgo nocallback Plugify_SetConstructVectorMatrix4x4
//#cgo nocallback Plugify_SetDestroyVectorBool
//#cgo nocallback Plugify_SetDestroyVectorChar8
//#cgo nocallback Plugify_SetDestroyVectorChar16
//#cgo nocallback Plugify_SetDestroyVectorInt8
//#cgo nocallback Plugify_SetDestroyVectorInt16
//#cgo nocallback Plugify_SetDestroyVectorInt32
//#cgo nocallback Plugify_SetDestroyVectorInt64
//#cgo nocallback Plugify_SetDestroyVectorUInt8
//#cgo nocallback Plugify_SetDestroyVectorUInt16
//#cgo nocallback Plugify_SetDestroyVectorUInt32
//#cgo nocallback Plugify_SetDestroyVectorUInt64
//#cgo nocallback Plugify_SetDestroyVectorPointer
//#cgo nocallback Plugify_SetDestroyVectorFloat
//#cgo nocallback Plugify_SetDestroyVectorDouble
//#cgo nocallback Plugify_SetDestroyVectorString
//#cgo nocallback Plugify_SetDestroyVectorVariant
//#cgo nocallback Plugify_SetDestroyVectorVector2
//#cgo nocallback Plugify_SetDestroyVectorVector3
//#cgo nocallback Plugify_SetDestroyVectorVector4
//#cgo nocallback Plugify_SetDestroyVectorMatrix4x4
//#cgo nocallback Plugify_SetGetVectorSizeBool
//#cgo nocallback Plugify_SetGetVectorSizeChar8
//#cgo nocallback Plugify_SetGetVectorSizeChar16
//#cgo nocallback Plugify_SetGetVectorSizeInt8
//#cgo nocallback Plugify_SetGetVectorSizeInt16
//#cgo nocallback Plugify_SetGetVectorSizeInt32
//#cgo nocallback Plugify_SetGetVectorSizeInt64
//#cgo nocallback Plugify_SetGetVectorSizeUInt8
//#cgo nocallback Plugify_SetGetVectorSizeUInt16
//#cgo nocallback Plugify_SetGetVectorSizeUInt32
//#cgo nocallback Plugify_SetGetVectorSizeUInt64
//#cgo nocallback Plugify_SetGetVectorSizePointer
//#cgo nocallback Plugify_SetGetVectorSizeFloat
//#cgo nocallback Plugify_SetGetVectorSizeDouble
//#cgo nocallback Plugify_SetGetVectorSizeString
//#cgo nocallback Plugify_SetGetVectorSizeVariant
//#cgo nocallback Plugify_SetGetVectorSizeVector2
//#cgo nocallback Plugify_SetGetVectorSizeVector3
//#cgo nocallback Plugify_SetGetVectorSizeVector4
//#cgo nocallback Plugify_SetGetVectorSizeMatrix4x4
//#cgo nocallback Plugify_SetGetVectorDataBool
//#cgo nocallback Plugify_SetGetVectorDataChar8
//#cgo nocallback Plugify_SetGetVectorDataChar16
//#cgo nocallback Plugify_SetGetVectorDataInt8
//#cgo nocallback Plugify_SetGetVectorDataInt16
//#cgo nocallback Plugify_SetGetVectorDataInt32
//#cgo nocallback Plugify_SetGetVectorDataInt64
//#cgo nocallback Plugify_SetGetVectorDataUInt8
//#cgo nocallback Plugify_SetGetVectorDataUInt16
//#cgo nocallback Plugify_SetGetVectorDataUInt32
//#cgo nocallback Plugify_SetGetVectorDataUInt64
//#cgo nocallback Plugify_SetGetVectorDataPointer
//#cgo nocallback Plugify_SetGetVectorDataFloat
//#cgo nocallback Plugify_SetGetVectorDataDouble
//#cgo nocallback Plugify_SetGetVectorDataString
//#cgo nocallback Plugify_SetGetVectorDataVariant
//#cgo nocallback Plugify_SetGetVectorDataVector2
//#cgo nocallback Plugify_SetGetVectorDataVector3
//#cgo nocallback Plugify_SetGetVectorDataVector4
//#cgo nocallback Plugify_SetGetVectorDataMatrix4x4
//#cgo nocallback Plugify_SetAssignVectorBool
//#cgo nocallback Plugify_SetAssignVectorChar8
//#cgo nocallback Plugify_SetAssignVectorChar16
//#cgo nocallback Plugify_SetAssignVectorInt8
//#cgo nocallback Plugify_SetAssignVectorInt16
//#cgo nocallback Plugify_SetAssignVectorInt32
//#cgo nocallback Plugify_SetAssignVectorInt64
//#cgo nocallback Plugify_SetAssignVectorUInt8
//#cgo nocallback Plugify_SetAssignVectorUInt16
//#cgo nocallback Plugify_SetAssignVectorUInt32
//#cgo nocallback Plugify_SetAssignVectorUInt64
//#cgo nocallback Plugify_SetAssignVectorPointer
//#cgo nocallback Plugify_SetAssignVectorFloat
//#cgo nocallback Plugify_SetAssignVectorDouble
//#cgo nocallback Plugify_SetAssignVectorString
//#cgo nocallback Plugify_SetAssignVectorVariant
//#cgo nocallback Plugify_SetAssignVectorVector2
//#cgo nocallback Plugify_SetAssignVectorVector3
//#cgo nocallback Plugify_SetAssignVectorVector4
//#cgo nocallback Plugify_SetAssignVectorMatrix4x4
//#cgo nocallback Plugify_SetNewCall
//#cgo nocallback Plugify_SetDeleteCall
//#cgo nocallback Plugify_SetGetCallFunction
//#cgo nocallback Plugify_SetGetCallError
//#cgo nocallback Plugify_SetNewCallback
//#cgo nocallback Plugify_SetDeleteCallback
//#cgo nocallback Plugify_SetGetCallbackFunction
//#cgo nocallback Plugify_SetGetCallbackError
//#cgo nocallback Plugify_SetGetMethodParamCount
//#cgo nocallback Plugify_SetGetMethodParamType
//#cgo nocallback Plugify_SetGetMethodPrototype
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
