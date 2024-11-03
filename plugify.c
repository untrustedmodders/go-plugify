#include <stddef.h>
typedef struct { const char *p; ptrdiff_t n; } _GoString_;
#include "plugify.h"
// Function typedefs
typedef void* (*GetMethodPtrFunc)(const char*);
typedef void (*GetMethodPtr2Func)(const char*, void**);
typedef const char* (*GetBaseDirFunc)();
typedef bool (*IsModuleLoadedFunc)(const char*, int, bool);
typedef bool (*IsPluginLoadedFunc)(const char*, int, bool);
	
typedef ptrdiff_t (*GetPluginIdFunc)(void*);
typedef const char* (*GetPluginNameFunc)(void*);
typedef const char* (*GetPluginFullNameFunc)(void*);
typedef const char* (*GetPluginDescriptionFunc)(void*);
typedef const char* (*GetPluginVersionFunc)(void*);
typedef const char* (*GetPluginAuthorFunc)(void*);
typedef const char* (*GetPluginWebsiteFunc)(void*);
typedef const char* (*GetPluginBaseDirFunc)(void*);
typedef void* (*GetPluginDependenciesFunc)(void*);
typedef ptrdiff_t (*GetPluginDependenciesSizeFunc)(void*);
typedef const char* (*FindPluginResourceFunc)(void*, const char*);
typedef void (*DeleteCStrFunc)(const char*);

typedef String (*ConstructStringFunc)(String*, _GoString_);
typedef void (*DestroyStringFunc)(String*);
typedef const char* (*GetStringDataFunc)(String*);
typedef ptrdiff_t (*GetStringLengthFunc)(String*);
typedef void (*AssignStringFunc)(String*, _GoString_);

typedef Vector (*ConstructVectorFunc)(void*, ptrdiff_t, enum DataType);
typedef void (*DestroyVectorFunc)(Vector*, enum DataType);
typedef void* (*GetVectorDataFunc)(Vector*, enum DataType);
typedef ptrdiff_t (*GetVectorSizeFunc)(Vector*, enum DataType);
typedef void (*AssignVectorFunc)(Vector*, void*, ptrdiff_t, enum DataType);

typedef void (*DeleteVectorDataCStrFunc)(void*);

void* pluginHandle = NULL;

// Variable declarations
GetMethodPtrFunc getMethodPtr = NULL;
GetMethodPtr2Func getMethodPtr2 = NULL;
GetBaseDirFunc getBaseDir = NULL;
IsModuleLoadedFunc isModuleLoaded = NULL;
IsPluginLoadedFunc isPluginLoaded = NULL;

GetPluginIdFunc getPluginId = NULL;
GetPluginNameFunc getPluginName = NULL;
GetPluginFullNameFunc getPluginFullName = NULL;
GetPluginDescriptionFunc getPluginDescription = NULL;
GetPluginVersionFunc getPluginVersion = NULL;
GetPluginAuthorFunc getPluginAuthor = NULL;
GetPluginWebsiteFunc getPluginWebsite = NULL;
GetPluginBaseDirFunc getPluginBaseDir = NULL;
GetPluginDependenciesFunc getPluginDependencies = NULL;
GetPluginDependenciesSizeFunc getPluginDependenciesSize = NULL;
FindPluginResourceFunc findPluginResource = NULL;
DeleteCStrFunc deleteCStr = NULL;

ConstructStringFunc constructString = NULL;
DestroyStringFunc destroyString = NULL;
GetStringDataFunc getStringData = NULL;
GetStringLengthFunc getStringLength = NULL;
AssignStringFunc assignString = NULL;

ConstructVectorFunc constructVector = NULL;
DestroyVectorFunc destroyVector = NULL;
GetVectorDataFunc getVectorData = NULL;
GetVectorSizeFunc getVectorSize = NULL;
AssignVectorFunc assignVector = NULL;

DeleteVectorDataCStrFunc deleteVectorDataCStr = NULL;

// Call methods
void* Plugify_GetMethodPtr(const char* methodName) {
	return getMethodPtr(methodName);
}
void Plugify_GetMethodPtr2(const char* methodName, void** addressDest) {
	getMethodPtr2(methodName, addressDest);
}
const char* Plugify_GetBaseDir() {
	return getBaseDir();
}
bool Plugify_IsModuleLoaded(const char* moduleName, int requiredVersion, bool minimum) {
	return isModuleLoaded(moduleName, requiredVersion, minimum);
}
bool Plugify_IsPluginLoaded(const char* pluginName, int requiredVersion, bool minimum) {
	return isPluginLoaded(pluginName, requiredVersion, minimum);
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
const char* Plugify_GetPluginBaseDir() {
	return getPluginBaseDir(pluginHandle);
}
void* Plugify_GetPluginDependencies() {
	return getPluginDependencies(pluginHandle);
}
ptrdiff_t Plugify_GetPluginDependenciesSize() {
	return getPluginDependenciesSize(pluginHandle);
}
const char* Plugify_FindPluginResource(const char* path) {
	return findPluginResource(pluginHandle, path);
}
void Plugify_DeleteCStr(const char* str) {
	return deleteCStr(str);
}

String Plugify_ConstructString(_GoString_ source) {
	return constructString(source);
}
void Plugify_DestroyString(String* string) {
	destroyString(string);
}
const char* Plugify_GetStringData(String* string) {
	return getStringData(string);
}
ptrdiff_t Plugify_GetStringLength(String* string) {
	return getStringLength(string);
}
void Plugify_AssignString(String* string, _GoString_ source) {
	assignString(string, source);
}


Vector Plugify_ConstructVector(void* arr, ptrdiff_t len, enum DataType type) {
	return constructVector(arr, len, type);
}
void Plugify_DestroyVector(Vector* vector, enum DataType type) {
	destroyVector(vector, type);
}
void* Plugify_GetVectorData(Vector* vector, enum DataType type) {
	return getVectorData(vector, type);
}
ptrdiff_t Plugify_GetVectorSize(Vector* vector, enum DataType type) {
	return getVectorSize(vector, type);
}
void Plugify_AssignVector(Vector* vector, void* arr, ptrdiff_t len, enum DataType type) {
	assignVector(vector, arr, len, type);
}

void Plugify_DeleteVectorDataCStr(void* arr) {
	deleteVectorDataCStr(arr);
}

// Setter methods
void Plugify_SetGetMethodPtr(void* func) {
	getMethodPtr = (GetMethodPtrFunc)func;
}
void Plugify_SetGetMethodPtr2(void* func) {
	getMethodPtr2 = (GetMethodPtr2Func)func;
}
void Plugify_SetGetBaseDir(void* func) {
	getBaseDir = (GetBaseDirFunc)func;
}
void Plugify_SetIsModuleLoaded(void* func) {
	isModuleLoaded = (IsModuleLoadedFunc)func;
}
void Plugify_SetIsPluginLoaded(void* func) {
	isPluginLoaded = (IsPluginLoadedFunc)func;
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
void Plugify_SetGetPluginBaseDir(void* func) {
	getPluginBaseDir = (GetPluginBaseDirFunc)func;
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
void Plugify_SetDeleteCStr(void* func) {
	deleteCStr = (DeleteCStrFunc)func;
}

void Plugify_SetConstructString(void* func) {
	constructString = (ConstructStringFunc)func;
}
void Plugify_SetDestroyString(void* func) {
	destroyString = (DestroyStringFunc)func;
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

void Plugify_SetConstructVector(void* func) {
	constructVector = (ConstructVectorFunc)func;
}
void Plugify_SetDestroyVector(void* func) {
	destroyVector = (DestroyVectorFunc)func;
}
void Plugify_SetGetVectorData(void* func) {
	getVectorData = (GetVectorDataFunc)func;
}
void Plugify_SetGetVectorSize(void* func) {
	getVectorSize = (GetVectorSizeFunc)func;
}
void Plugify_SetAssignVector(void* func) {
	assignVector = (AssignVectorFunc)func;
}

void Plugify_SetDeleteVectorDataCStr(void* func) {
	deleteVectorDataCStr = (DeleteVectorDataCStrFunc)func;
}