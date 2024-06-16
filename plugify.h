// plugify.h
#pragma once

#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>

#ifdef __cplusplus
extern "C" {
#endif

enum DataType {
	BOOL,
	CHAR8,
	CHAR16,
	INT8,
	INT16,
	INT32,
	INT64,
	UINT8,
	UINT16,
	UINT32,
	UINT64,
	UINTPTR,
	FLOAT,
	DOUBLE,
	STRING
};

void* Plugify_GetMethodPtr(const char* methodName);
bool Plugify_IsModuleLoaded(const char* moduleName, int requiredVersion, bool minimum); // INT_MAX for latest version
bool Plugify_IsPluginLoaded(const char* pluginName, int requiredVersion, bool minimum); // INT_MAX for latest version

void Plugify_SetPluginHandle(void* handle);
ptrdiff_t Plugify_GetPluginId();
const char* Plugify_GetPluginName();
const char* Plugify_GetPluginFullName();
const char* Plugify_GetPluginDescription();
const char* Plugify_GetPluginVersion();
const char* Plugify_GetPluginAuthor();
const char* Plugify_GetPluginWebsite();
const char* Plugify_GetPluginBaseDir(); // Plugify_DeleteCStr
const char** Plugify_GetPluginDependencies(); // Plugify_DeleteVectorDataCStr
ptrdiff_t Plugify_GetPluginDependenciesSize();
const char* Plugify_FindPluginResource(const char* path); // Plugify_DeleteCStr
void Plugify_DeleteCStr(const char* path);

void* Plugify_AllocateString();
void* Plugify_CreateString(_GoString_ source);
const char* Plugify_GetStringData(void* ptr);
ptrdiff_t Plugify_GetStringLength(void* ptr);
void Plugify_AssignString(void* ptr, _GoString_ source);
void Plugify_FreeString(void* ptr);
void Plugify_DeleteString(void* ptr);

void* Plugify_CreateVector(void* arr, ptrdiff_t len, enum DataType type);
void* Plugify_AllocateVector(enum DataType type);
ptrdiff_t Plugify_GetVectorSize(void* ptr, enum DataType type);
void* Plugify_GetVectorData(void* ptr, enum DataType type); // Plugify_DeleteVectorDataCStr for STRING / Plugify_DeleteVectorDataBool for BOOL
void Plugify_AssignVector(void* ptr, void* arr, ptrdiff_t len, enum DataType type);
void Plugify_DeleteVector(void* ptr, enum DataType type);
void Plugify_FreeVector(void* ptr, enum DataType type);

void Plugify_DeleteVectorDataBool(void* ptr);
void Plugify_DeleteVectorDataCStr(void* ptr);

void Plugify_SetGetMethodPtr(void* func);
void Plugify_SetIsModuleLoaded(void* func);
void Plugify_SetIsPluginLoaded(void* func);
void Plugify_SetGetPluginId(void* func);
void Plugify_SetGetPluginName(void* func);
void Plugify_SetGetPluginFullName(void* func);
void Plugify_SetGetPluginDescription(void* func);
void Plugify_SetGetPluginVersion(void* func);
void Plugify_SetGetPluginAuthor(void* func);
void Plugify_SetGetPluginWebsite(void* func);
void Plugify_SetGetPluginBaseDir(void* func);
void Plugify_SetGetPluginDependencies(void* func);
void Plugify_SetGetPluginDependenciesSize(void* func);
void Plugify_SetFindPluginResource(void* func);
void Plugify_SetDeleteCStr(void* func);
void Plugify_SetAllocateString(void* func);
void Plugify_SetCreateString(void* func);
void Plugify_SetGetStringData(void* func);
void Plugify_SetGetStringLength(void* func);
void Plugify_SetAssignString(void* func);
void Plugify_SetFreeString(void* func);
void Plugify_SetDeleteString(void* func);
void Plugify_SetCreateVector(void* func);
void Plugify_SetAllocateVector(void* func);
void Plugify_SetGetVectorSize(void* func);
void Plugify_SetGetVectorData(void* func);
void Plugify_SetAssignVector(void* func);
void Plugify_SetDeleteVector(void* func);
void Plugify_SetFreeVector(void* func);

void Plugify_SetDeleteVectorDataBool(void* func);
void Plugify_SetDeleteVectorDataCStr(void* func);
#ifdef __cplusplus
}
#endif