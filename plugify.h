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
void* Plugify_GetVectorData(void* ptr, enum DataType type);
void Plugify_AssignVector(void* ptr, void* arr, ptrdiff_t len, enum DataType type);
void Plugify_DeleteVector(void* ptr, enum DataType type);
void Plugify_FreeVector(void* ptr, enum DataType type);

void Plugify_DeleteVectorDataBool(void* ptr);
void Plugify_DeleteVectorDataCStr(void* ptr);

void Plugify_SetGetMethodPtr(void* func);
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