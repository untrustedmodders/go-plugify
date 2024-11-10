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
	POINTER,
	FLOAT,
	DOUBLE,
	STRING
};

typedef struct { char* data; size_t size; size_t cap; } String;
typedef struct { void* begin; void* end; void* capacity; } Vector;

void* Plugify_GetMethodPtr(const char* methodName);
void Plugify_GetMethodPtr2(const char* methodName, void** addressDest);
const char* Plugify_GetBaseDir(); // Plugify_DeleteCStr
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
void* Plugify_GetPluginDependencies(); // Plugify_DeleteVectorDataCStr
ptrdiff_t Plugify_GetPluginDependenciesSize();
const char* Plugify_FindPluginResource(const char* path); // Plugify_DeleteCStr
void Plugify_DeleteCStr(const char* str);

String Plugify_ConstructString(_GoString_ source);
void Plugify_DestroyString(String* string);
const char* Plugify_GetStringData(String* string);
ptrdiff_t Plugify_GetStringLength(String* string);
void Plugify_AssignString(String* string, _GoString_ source);

Vector Plugify_ConstructVector(void* arr, ptrdiff_t len, enum DataType type);
void Plugify_DeleteVector(Vector* vector, enum DataType type);
void* Plugify_GetVectorData(Vector* vector, enum DataType type); // Plugify_DeleteVectorDataCStr for STRING
ptrdiff_t Plugify_GetVectorSize(Vector* vector, enum DataType type);
void Plugify_AssignVector(Vector* vector, void* arr, ptrdiff_t len, enum DataType type);

void Plugify_DeleteVectorDataCStr(void* arr);

void Plugify_SetGetMethodPtr(void* func);
void Plugify_SetGetMethodPtr2(void* func);
void Plugify_SetGetBaseDir(void* func);
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
void Plugify_SetConstructString(void* func);
void Plugify_SetDestroyString(void* func);
void Plugify_SetGetStringData(void* func);
void Plugify_SetGetStringLength(void* func);
void Plugify_SetAssignString(void* func);
void Plugify_SetConstructVector(void* func);
void Plugify_SetDestroyVector(void* func);
void Plugify_SetGetVectorData(void* func);
void Plugify_SetGetVectorSize(void* func);
void Plugify_SetAssignVector(void* func);
void Plugify_SetDeleteVectorDataCStr(void* func);
#ifdef __cplusplus
}
#endif