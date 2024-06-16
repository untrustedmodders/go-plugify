#include <stddef.h>
typedef struct { const char *p; ptrdiff_t n; } _GoString_;
#include "plugify.h"
// Function typedefs
typedef void* (*GetMethodPtrFunc)(const char*);

typedef ptrdiff_t (*GetPluginIdFunc)(void*);
typedef const char* (*GetPluginNameFunc)(void*);
typedef const char* (*GetPluginFullNameFunc)(void*);
typedef const char* (*GetPluginDescriptionFunc)(void*);
typedef const char* (*GetPluginVersionFunc)(void*);
typedef const char* (*GetPluginAuthorFunc)(void*);
typedef const char* (*GetPluginWebsiteFunc)(void*);
typedef const char** (*GetPluginDependenciesFunc)(void*);
typedef ptrdiff_t (*GetPluginDependenciesSizeFunc)(void*);
typedef const char* (*FindPluginResourceFunc)(void*, const char*);
typedef void (*FreePluginResourceFunc)(const char*);

typedef void* (*AllocateStringFunc)();
typedef void* (*CreateStringFunc)(_GoString_);
typedef const char* (*GetStringDataFunc)(void*);
typedef ptrdiff_t (*GetStringLengthFunc)(void*);
typedef void (*AssignStringFunc)(void*, _GoString_);
typedef void (*FreeStringFunc)(void*);
typedef void (*DeleteStringFunc)(void*);

typedef void* (*CreateVectorFunc)(void*, ptrdiff_t, enum DataType);
typedef void* (*AllocateVectorFunc)(enum DataType);
typedef ptrdiff_t (*GetVectorSizeFunc)(void*, enum DataType);
typedef void* (*GetVectorDataFunc)(void*, enum DataType);
typedef void (*AssignVectorFunc)(void*, void*, ptrdiff_t, enum DataType);
typedef void (*DeleteVectorFunc)(void*, enum DataType);
typedef void (*FreeVectorFunc)(void*, enum DataType);

typedef void (*DeleteVectorDataBoolFunc)(void*);
typedef void (*DeleteVectorDataCStrFunc)(void*);

void* pluginHandle = NULL;

// Variable declarations
GetMethodPtrFunc getMethodPtr = NULL;

GetPluginIdFunc getPluginId = NULL;
GetPluginNameFunc getPluginName = NULL;
GetPluginFullNameFunc getPluginFullName = NULL;
GetPluginDescriptionFunc getPluginDescription = NULL;
GetPluginVersionFunc getPluginVersion = NULL;
GetPluginAuthorFunc getPluginAuthor = NULL;
GetPluginWebsiteFunc getPluginWebsite = NULL;
GetPluginDependenciesFunc getPluginDependencies = NULL;
GetPluginDependenciesSizeFunc getPluginDependenciesSize = NULL;
FindPluginResourceFunc findPluginResource = NULL;
FreePluginResourceFunc freePluginResource = NULL;

AllocateStringFunc allocateString = NULL;
CreateStringFunc createString = NULL;
GetStringDataFunc getStringData = NULL;
GetStringLengthFunc getStringLength = NULL;
AssignStringFunc assignString = NULL;
FreeStringFunc freeString = NULL;
DeleteStringFunc deleteString = NULL;

CreateVectorFunc createVector = NULL;
AllocateVectorFunc allocateVector = NULL;
GetVectorSizeFunc getVectorSize = NULL;
GetVectorDataFunc getVectorData = NULL;
AssignVectorFunc assignVector = NULL;
DeleteVectorFunc deleteVector = NULL;
FreeVectorFunc freeVector = NULL;

DeleteVectorDataBoolFunc deleteVectorDataBool = NULL;
DeleteVectorDataCStrFunc deleteVectorDataCStr = NULL;

// Call methods
void* Plugify_GetMethodPtr(const char* methodName) {
	return getMethodPtr(methodName);
}

void Plugify_SetPluginHandle(void* handle) {
	pluginHandle = handle;
}
ptrdiff_t Plugify_GetPluginId() {
	return getPluginId(pluginHandle);
}
const char* Plugify_GetPluginName() {
	return getPluginName(pluginHandle);
}
const char* Plugify_GetPluginFullName() {
	return getPluginFullName(pluginHandle);
}
const char* Plugify_GetPluginDescription() {
	return getPluginDescription(pluginHandle);
}
const char* Plugify_GetPluginVersion() {
	return getPluginVersion(pluginHandle);
}
const char* Plugify_GetPluginAuthor() {
	return getPluginAuthor(pluginHandle);
}
const char* Plugify_GetPluginWebsite() {
	return getPluginWebsite(pluginHandle);
}
const char** Plugify_GetPluginDependencies() {
	return getPluginDependencies(pluginHandle);
}
ptrdiff_t Plugify_GetPluginDependenciesSize() {
	return getPluginDependenciesSize(pluginHandle);
}
const char* Plugify_FindPluginResource(const char* path) {
	return findPluginResource(pluginHandle, path);
}
void Plugify_FreePluginResource(const char* path) {
	return freePluginResource(path);
}

void* Plugify_AllocateString() {
	return allocateString();
}
void* Plugify_CreateString(_GoString_ source) {
	return createString(source);
}
const char* Plugify_GetStringData(void* ptr) {
	return getStringData(ptr);
}
ptrdiff_t Plugify_GetStringLength(void* ptr) {
	return getStringLength(ptr);
}
void Plugify_AssignString(void* ptr, _GoString_ source) {
	assignString(ptr, source);
}
void Plugify_FreeString(void* ptr) {
	freeString(ptr);
}

void Plugify_DeleteString(void* ptr) {
	deleteString(ptr);
}


void* Plugify_CreateVector(void* arr, ptrdiff_t len, enum DataType type) {
	return createVector(arr, len, type);
}

void* Plugify_AllocateVector(enum DataType type) {
	return allocateVector(type);
}

ptrdiff_t Plugify_GetVectorSize(void* ptr, enum DataType type) {
	return getVectorSize(ptr, type);
}

void* Plugify_GetVectorData(void* ptr, enum DataType type) {
	return getVectorData(ptr, type);
}

void Plugify_AssignVector(void* ptr, void* arr, ptrdiff_t len, enum DataType type) {
	assignVector(ptr, arr, len, type);
}

void Plugify_DeleteVector(void* ptr, enum DataType type) {
	deleteVector(ptr, type);
}

void Plugify_FreeVector(void* ptr, enum DataType type) {
	freeVector(ptr, type);
}

void Plugify_DeleteVectorDataBool(void* ptr) {
	deleteVectorDataBool(ptr);
}
void Plugify_DeleteVectorDataCStr(void* ptr) {
	deleteVectorDataCStr(ptr);
}

// Setter methods
void Plugify_SetGetMethodPtr(void* func) {
	getMethodPtr = (GetMethodPtrFunc)func;
}

void Plugify_SetGetPluginId(void* func) {
	getPluginId = (GetPluginIdFunc)func;
}

void Plugify_SetGetPluginName(void* func) {
	getPluginName = (GetPluginNameFunc)func;
}

void Plugify_SetGetPluginFullName(void* func) {
	getPluginFullName = (GetPluginFullNameFunc)func;
}

void Plugify_SetGetPluginDescription(void* func) {
	getPluginDescription = (GetPluginDescriptionFunc)func;
}

void Plugify_SetGetPluginVersion(void* func) {
	getPluginVersion = (GetPluginVersionFunc)func;
}

void Plugify_SetGetPluginAuthor(void* func) {
	getPluginAuthor = (GetPluginAuthorFunc)func;
}

void Plugify_SetGetPluginWebsite(void* func) {
	getPluginWebsite = (GetPluginWebsiteFunc)func;
}

void Plugify_SetGetPluginDependencies(void* func) {
	getPluginDependencies = (GetPluginDependenciesFunc)func;
}

void Plugify_SetGetPluginDependenciesSize(void* func) {
	getPluginDependenciesSize = (GetPluginDependenciesSizeFunc)func;
}

void Plugify_SetFindPluginResource(void* func) {
	findPluginResource = (FindPluginResourceFunc)func;
}

void Plugify_SetFreePluginResource(void* func) {
	freePluginResource = (FreePluginResourceFunc)func;
}

void Plugify_SetAllocateString(void* func) {
	allocateString = (AllocateStringFunc)func;
}

void Plugify_SetCreateString(void* func) {
	createString = (CreateStringFunc)func;
}

void Plugify_SetGetStringData(void* func) {
	getStringData = (GetStringDataFunc)func;
}

void Plugify_SetGetStringLength(void* func) {
	getStringLength = (GetStringLengthFunc)func;
}

void Plugify_SetAssignString(void* func) {
	assignString = (AssignStringFunc)func;
}

void Plugify_SetFreeString(void* func) {
	freeString = (FreeStringFunc)func;
}

void Plugify_SetDeleteString(void* func) {
	deleteString = (DeleteStringFunc)func;
}

void Plugify_SetCreateVector(void* func) {
	createVector = (CreateVectorFunc)func;
}

void Plugify_SetAllocateVector(void* func) {
	allocateVector = (AllocateVectorFunc)func;
}

void Plugify_SetGetVectorSize(void* func) {
	getVectorSize = (GetVectorSizeFunc)func;
}

void Plugify_SetGetVectorData(void* func) {
	getVectorData = (GetVectorDataFunc)func;
}

void Plugify_SetAssignVector(void* func) {
	assignVector = (AssignVectorFunc)func;
}

void Plugify_SetDeleteVector(void* func) {
	deleteVector = (DeleteVectorFunc)func;
}

void Plugify_SetFreeVector(void* func) {
	freeVector = (FreeVectorFunc)func;
}

void Plugify_SetDeleteVectorDataBool(void* func) {
	deleteVectorDataBool = (DeleteVectorDataBoolFunc)func;
}

void Plugify_SetDeleteVectorDataCStr(void* func) {
	deleteVectorDataCStr = (DeleteVectorDataCStrFunc)func;
}