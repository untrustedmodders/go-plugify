#include <stddef.h>
typedef struct { const char *p; ptrdiff_t n; } _GoString_;
#include "plugify.h"
// Function typedefs
typedef void* (*GetMethodPtrFunc)(const char*);

typedef void* (*AllocateStringFunc)();
typedef void* (*CreateStringFunc)(_GoString_ source);
typedef const char* (*GetStringDataFunc)(void* ptr);
typedef ptrdiff_t (*GetStringLengthFunc)(void* ptr);
typedef void (*AssignStringFunc)(void* ptr, _GoString_ source);
typedef void (*FreeStringFunc)(void* ptr);
typedef void (*DeleteStringFunc)(void* ptr);

typedef void* (*CreateVectorFunc)(void* arr, ptrdiff_t len, enum DataType type);
typedef void* (*AllocateVectorFunc)(enum DataType type);
typedef ptrdiff_t (*GetVectorSizeFunc)(void* ptr, enum DataType type);
typedef void* (*GetVectorDataFunc)(void* ptr, enum DataType type);
typedef void (*AssignVectorFunc)(void* ptr, void* arr, ptrdiff_t len, enum DataType type);
typedef void (*DeleteVectorFunc)(void* ptr, enum DataType type);
typedef void (*FreeVectorFunc)(void* ptr, enum DataType type);

typedef void (*DeleteVectorDataBoolFunc)(void* ptr);
typedef void (*DeleteVectorDataCStrFunc)(void* ptr);

// Variable declarations
GetMethodPtrFunc getMethodPtr = NULL;

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