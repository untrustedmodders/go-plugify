// plugify.h
#pragma once

#include <stdlib.h>
#include <stdint.h>
#include <stdbool.h>
#include <string.h>

typedef unsigned short char16_t;
typedef struct uint128_t {
	uint64_t low;
	uint64_t high;
} uint128_t;

#ifdef __cplusplus
extern "C" {
#endif

typedef struct GoString_ { const char* p; ptrdiff_t n; } GoString_;
typedef struct String { char* data; size_t size; size_t cap; } String;
typedef struct Vector { void* begin; void* end; void* capacity; } Vector;
typedef struct Vector2 { float x, y; } Vector2;
typedef struct Vector3 { float x, y, z; } Vector3;
typedef struct Vector4 { float x, y, z, w; } Vector4;
typedef struct Matrix4x4 { float m[4][4]; } Matrix4x4;
typedef struct Variant {
	union {
		bool boolean;
		char char8;
		char16_t char16;
		int8_t int8;
		int16_t int16;
		int32_t int32;
		int64_t int64;
		uint8_t uint8;
		uint16_t uint16;
		uint32_t uint32;
		uint64_t uint64;
		void* ptr;
		float flt;
		double dbl;
		String str;
		Vector vec;
		Vector2 vec2;
		Vector3 vec3;
		Vector4 vec4;
	};
#if INTPTR_MAX == INT32_MAX
	volatile char pad[8];
#endif
	uint8_t current;
} Variant;

typedef struct ManagedType {
	uint8_t valueType;
	bool ref;
} ManagedType;

typedef struct Parameters {
	uint64_t arguments;
} Parameters;

typedef struct Return {
	uint64_t ret[2];
} Return;

typedef struct PluginContext {
	bool hasUpdate;
	bool hasStart;
	bool hasEnd;
	bool hasPanic;
} PluginContext;

typedef void* PluginHandle;
extern PluginHandle pluginHandle;

// Extern declarations for Plugify_ functions
extern String Plugify_GetBaseDir();
extern String Plugify_GetExtensionsDir();
extern String Plugify_GetConfigsDir();
extern String Plugify_GetDataDir();
extern String Plugify_GetLogsDir();
extern String Plugify_GetCacheDir();
extern bool Plugify_IsExtensionLoaded(_GoString_ name, _GoString_ constraint);
extern void Plugify_PrintException(_GoString_ message);

extern ptrdiff_t Plugify_GetPluginId();
extern String Plugify_GetPluginName();
extern String Plugify_GetPluginLicense();
extern String Plugify_GetPluginDescription();
extern String Plugify_GetPluginVersion();
extern String Plugify_GetPluginAuthor();
extern String Plugify_GetPluginWebsite();
extern String Plugify_GetPluginLocation();
extern Vector Plugify_GetPluginDependencies();

extern String Plugify_ConstructString(_GoString_ source);
extern void Plugify_DestroyString(String* string);
extern const char* Plugify_GetStringData(String* string);
extern ptrdiff_t Plugify_GetStringLength(String* string);
extern void Plugify_AssignString(String* string, _GoString_ source);

extern void Plugify_DestroyVariant(Variant* any);

extern Vector Plugify_ConstructVectorBool(bool* data, ptrdiff_t size);
extern Vector Plugify_ConstructVectorChar8(char* data, ptrdiff_t size);
extern Vector Plugify_ConstructVectorChar16(char16_t* data, ptrdiff_t size);
extern Vector Plugify_ConstructVectorInt8(int8_t* data, ptrdiff_t size);
extern Vector Plugify_ConstructVectorInt16(int16_t* data, ptrdiff_t size);
extern Vector Plugify_ConstructVectorInt32(int32_t* data, ptrdiff_t size);
extern Vector Plugify_ConstructVectorInt64(int64_t* data, ptrdiff_t size);
extern Vector Plugify_ConstructVectorUInt8(uint8_t* data, ptrdiff_t size);
extern Vector Plugify_ConstructVectorUInt16(uint16_t* data, ptrdiff_t size);
extern Vector Plugify_ConstructVectorUInt32(uint32_t* data, ptrdiff_t size);
extern Vector Plugify_ConstructVectorUInt64(uint64_t* data, ptrdiff_t size);
extern Vector Plugify_ConstructVectorPointer(uintptr_t* data, ptrdiff_t size);
extern Vector Plugify_ConstructVectorFloat(float* data, ptrdiff_t size);
extern Vector Plugify_ConstructVectorDouble(double* data, ptrdiff_t size);
extern Vector Plugify_ConstructVectorString(_GoString_* data, ptrdiff_t size);
extern Vector Plugify_ConstructVectorVariant(ptrdiff_t size);
extern Vector Plugify_ConstructVectorVector2(Vector2* data, ptrdiff_t size);
extern Vector Plugify_ConstructVectorVector3(Vector3* data, ptrdiff_t size);
extern Vector Plugify_ConstructVectorVector4(Vector4* data, ptrdiff_t size);
extern Vector Plugify_ConstructVectorMatrix4x4(Matrix4x4* data, ptrdiff_t size);

extern void Plugify_DestroyVectorBool(Vector* vec);
extern void Plugify_DestroyVectorChar8(Vector* vec);
extern void Plugify_DestroyVectorChar16(Vector* vec);
extern void Plugify_DestroyVectorInt8(Vector* vec);
extern void Plugify_DestroyVectorInt16(Vector* vec);
extern void Plugify_DestroyVectorInt32(Vector* vec);
extern void Plugify_DestroyVectorInt64(Vector* vec);
extern void Plugify_DestroyVectorUInt8(Vector* vec);
extern void Plugify_DestroyVectorUInt16(Vector* vec);
extern void Plugify_DestroyVectorUInt32(Vector* vec);
extern void Plugify_DestroyVectorUInt64(Vector* vec);
extern void Plugify_DestroyVectorPointer(Vector* vec);
extern void Plugify_DestroyVectorFloat(Vector* vec);
extern void Plugify_DestroyVectorDouble(Vector* vec);
extern void Plugify_DestroyVectorString(Vector* vec);
extern void Plugify_DestroyVectorVariant(Vector* vec);
extern void Plugify_DestroyVectorVector2(Vector* vec);
extern void Plugify_DestroyVectorVector3(Vector* vec);
extern void Plugify_DestroyVectorVector4(Vector* vec);
extern void Plugify_DestroyVectorMatrix4x4(Vector* vec);

extern ptrdiff_t Plugify_GetVectorSizeBool(Vector* vec);
extern ptrdiff_t Plugify_GetVectorSizeChar8(Vector* vec);
extern ptrdiff_t Plugify_GetVectorSizeChar16(Vector* vec);
extern ptrdiff_t Plugify_GetVectorSizeInt8(Vector* vec);
extern ptrdiff_t Plugify_GetVectorSizeInt16(Vector* vec);
extern ptrdiff_t Plugify_GetVectorSizeInt32(Vector* vec);
extern ptrdiff_t Plugify_GetVectorSizeInt64(Vector* vec);
extern ptrdiff_t Plugify_GetVectorSizeUInt8(Vector* vec);
extern ptrdiff_t Plugify_GetVectorSizeUInt16(Vector* vec);
extern ptrdiff_t Plugify_GetVectorSizeUInt32(Vector* vec);
extern ptrdiff_t Plugify_GetVectorSizeUInt64(Vector* vec);
extern ptrdiff_t Plugify_GetVectorSizePointer(Vector* vec);
extern ptrdiff_t Plugify_GetVectorSizeFloat(Vector* vec);
extern ptrdiff_t Plugify_GetVectorSizeDouble(Vector* vec);
extern ptrdiff_t Plugify_GetVectorSizeString(Vector* vec);
extern ptrdiff_t Plugify_GetVectorSizeVariant(Vector* vec);
extern ptrdiff_t Plugify_GetVectorSizeVector2(Vector* vec);
extern ptrdiff_t Plugify_GetVectorSizeVector3(Vector* vec);
extern ptrdiff_t Plugify_GetVectorSizeVector4(Vector* vec);
extern ptrdiff_t Plugify_GetVectorSizeMatrix4x4(Vector* vec);

extern bool* Plugify_GetVectorDataBool(Vector* vec);
extern char* Plugify_GetVectorDataChar8(Vector* vec);
extern char16_t* Plugify_GetVectorDataChar16(Vector* vec);
extern int8_t* Plugify_GetVectorDataInt8(Vector* vec);
extern int16_t* Plugify_GetVectorDataInt16(Vector* vec);
extern int32_t* Plugify_GetVectorDataInt32(Vector* vec);
extern int64_t* Plugify_GetVectorDataInt64(Vector* vec);
extern uint8_t* Plugify_GetVectorDataUInt8(Vector* vec);
extern uint16_t* Plugify_GetVectorDataUInt16(Vector* vec);
extern uint32_t* Plugify_GetVectorDataUInt32(Vector* vec);
extern uint64_t* Plugify_GetVectorDataUInt64(Vector* vec);
extern uintptr_t* Plugify_GetVectorDataPointer(Vector* vec);
extern float* Plugify_GetVectorDataFloat(Vector* vec);
extern double* Plugify_GetVectorDataDouble(Vector* vec);
extern String* Plugify_GetVectorDataString(Vector* vec);
extern Variant* Plugify_GetVectorDataVariant(Vector* vec, ptrdiff_t index);
extern Vector2* Plugify_GetVectorDataVector2(Vector* vec);
extern Vector3* Plugify_GetVectorDataVector3(Vector* vec);
extern Vector4* Plugify_GetVectorDataVector4(Vector* vec);
extern Matrix4x4* Plugify_GetVectorDataMatrix4x4(Vector* vec);

extern void Plugify_AssignVectorBool(Vector* vec, bool* data, ptrdiff_t size);
extern void Plugify_AssignVectorChar8(Vector* vec, char* data, ptrdiff_t size);
extern void Plugify_AssignVectorChar16(Vector* vec, char16_t* data, ptrdiff_t size);
extern void Plugify_AssignVectorInt8(Vector* vec, int8_t* data, ptrdiff_t size);
extern void Plugify_AssignVectorInt16(Vector* vec, int16_t* data, ptrdiff_t size);
extern void Plugify_AssignVectorInt32(Vector* vec, int32_t* data, ptrdiff_t size);
extern void Plugify_AssignVectorInt64(Vector* vec, int64_t* data, ptrdiff_t size);
extern void Plugify_AssignVectorUInt8(Vector* vec, uint8_t* data, ptrdiff_t size);
extern void Plugify_AssignVectorUInt16(Vector* vec, uint16_t* data, ptrdiff_t size);
extern void Plugify_AssignVectorUInt32(Vector* vec, uint32_t* data, ptrdiff_t size);
extern void Plugify_AssignVectorUInt64(Vector* vec, uint64_t* data, ptrdiff_t size);
extern void Plugify_AssignVectorPointer(Vector* vec, uintptr_t* data, ptrdiff_t size);
extern void Plugify_AssignVectorFloat(Vector* vec, float* data, ptrdiff_t size);
extern void Plugify_AssignVectorDouble(Vector* vec, double* data, ptrdiff_t size);
extern void Plugify_AssignVectorString(Vector* vec, _GoString_* data, ptrdiff_t size);
extern void Plugify_AssignVectorVariant(Vector* vec, ptrdiff_t size);
extern void Plugify_AssignVectorVector2(Vector* vec, Vector2* data, ptrdiff_t size);
extern void Plugify_AssignVectorVector3(Vector* vec, Vector3* data, ptrdiff_t size);
extern void Plugify_AssignVectorVector4(Vector* vec, Vector4* data, ptrdiff_t size);
extern void Plugify_AssignVectorMatrix4x4(Vector* vec, Matrix4x4* data, ptrdiff_t size);

typedef void* JitCall;

extern JitCall Plugify_NewCall(void* target, ManagedType* params, ptrdiff_t count, ManagedType ret);
extern void Plugify_DeleteCall(JitCall call);
extern void* Plugify_GetCallFunction(JitCall call);
extern const char* Plugify_GetCallError(JitCall call);

extern void Plugify_CallFunction(JitCall call, uint64_t* params, uint128_t* ret);

typedef void* JitCallback;

extern JitCallback Plugify_NewCallback(_GoString_ name, void* handle);
extern void Plugify_DeleteCallback(JitCallback callback);
extern void* Plugify_GetCallbackFunction(JitCallback callback);
extern const char* Plugify_GetCallbackError(JitCallback callback);

typedef void* MethodHandle;
typedef void* EnumHandle;

extern ptrdiff_t Plugify_GetMethodParamCount(MethodHandle handle);
extern ManagedType Plugify_GetMethodParamType(MethodHandle handle, ptrdiff_t index);
extern MethodHandle Plugify_GetMethodPrototype(MethodHandle handle, ptrdiff_t index);
extern EnumHandle Plugify_GetMethodEnum(MethodHandle handle, ptrdiff_t index);

extern void Plugify_SetGetBaseDir(void* ptr);
extern void Plugify_SetGetExtensionsDir(void* ptr);
extern void Plugify_SetGetConfigsDir(void* ptr);
extern void Plugify_SetGetDataDir(void* ptr);
extern void Plugify_SetGetLogsDir(void* ptr);
extern void Plugify_SetGetCacheDir(void* ptr);
extern void Plugify_SetIsExtensionLoaded(void* ptr);
extern void Plugify_SetPrintException(void* ptr);
extern void Plugify_SetGetPluginId(void* ptr);
extern void Plugify_SetGetPluginName(void* ptr);
extern void Plugify_SetGetPluginDescription(void* ptr);
extern void Plugify_SetGetPluginVersion(void* ptr);
extern void Plugify_SetGetPluginAuthor(void* ptr);
extern void Plugify_SetGetPluginWebsite(void* ptr);
extern void Plugify_SetGetPluginLicense(void* ptr);
extern void Plugify_SetGetPluginLocation(void* ptr);
extern void Plugify_SetGetPluginDependencies(void* ptr);
extern void Plugify_SetConstructString(void* ptr);
extern void Plugify_SetDestroyString(void* ptr);
extern void Plugify_SetGetStringData(void* ptr);
extern void Plugify_SetGetStringLength(void* ptr);
extern void Plugify_SetAssignString(void* ptr);
extern void Plugify_SetDestroyVariant(void* ptr);
extern void Plugify_SetConstructVectorBool(void* ptr);
extern void Plugify_SetConstructVectorChar8(void* ptr);
extern void Plugify_SetConstructVectorChar16(void* ptr);
extern void Plugify_SetConstructVectorInt8(void* ptr);
extern void Plugify_SetConstructVectorInt16(void* ptr);
extern void Plugify_SetConstructVectorInt32(void* ptr);
extern void Plugify_SetConstructVectorInt64(void* ptr);
extern void Plugify_SetConstructVectorUInt8(void* ptr);
extern void Plugify_SetConstructVectorUInt16(void* ptr);
extern void Plugify_SetConstructVectorUInt32(void* ptr);
extern void Plugify_SetConstructVectorUInt64(void* ptr);
extern void Plugify_SetConstructVectorPointer(void* ptr);
extern void Plugify_SetConstructVectorFloat(void* ptr);
extern void Plugify_SetConstructVectorDouble(void* ptr);
extern void Plugify_SetConstructVectorString(void* ptr);
extern void Plugify_SetConstructVectorVariant(void* ptr);
extern void Plugify_SetConstructVectorVector2(void* ptr);
extern void Plugify_SetConstructVectorVector3(void* ptr);
extern void Plugify_SetConstructVectorVector4(void* ptr);
extern void Plugify_SetConstructVectorMatrix4x4(void* ptr);
extern void Plugify_SetDestroyVectorBool(void* ptr);
extern void Plugify_SetDestroyVectorChar8(void* ptr);
extern void Plugify_SetDestroyVectorChar16(void* ptr);
extern void Plugify_SetDestroyVectorInt8(void* ptr);
extern void Plugify_SetDestroyVectorInt16(void* ptr);
extern void Plugify_SetDestroyVectorInt32(void* ptr);
extern void Plugify_SetDestroyVectorInt64(void* ptr);
extern void Plugify_SetDestroyVectorUInt8(void* ptr);
extern void Plugify_SetDestroyVectorUInt16(void* ptr);
extern void Plugify_SetDestroyVectorUInt32(void* ptr);
extern void Plugify_SetDestroyVectorUInt64(void* ptr);
extern void Plugify_SetDestroyVectorPointer(void* ptr);
extern void Plugify_SetDestroyVectorFloat(void* ptr);
extern void Plugify_SetDestroyVectorDouble(void* ptr);
extern void Plugify_SetDestroyVectorString(void* ptr);
extern void Plugify_SetDestroyVectorVariant(void* ptr);
extern void Plugify_SetDestroyVectorVector2(void* ptr);
extern void Plugify_SetDestroyVectorVector3(void* ptr);
extern void Plugify_SetDestroyVectorVector4(void* ptr);
extern void Plugify_SetDestroyVectorMatrix4x4(void* ptr);
extern void Plugify_SetGetVectorSizeBool(void* ptr);
extern void Plugify_SetGetVectorSizeChar8(void* ptr);
extern void Plugify_SetGetVectorSizeChar16(void* ptr);
extern void Plugify_SetGetVectorSizeInt8(void* ptr);
extern void Plugify_SetGetVectorSizeInt16(void* ptr);
extern void Plugify_SetGetVectorSizeInt32(void* ptr);
extern void Plugify_SetGetVectorSizeInt64(void* ptr);
extern void Plugify_SetGetVectorSizeUInt8(void* ptr);
extern void Plugify_SetGetVectorSizeUInt16(void* ptr);
extern void Plugify_SetGetVectorSizeUInt32(void* ptr);
extern void Plugify_SetGetVectorSizeUInt64(void* ptr);
extern void Plugify_SetGetVectorSizePointer(void* ptr);
extern void Plugify_SetGetVectorSizeFloat(void* ptr);
extern void Plugify_SetGetVectorSizeDouble(void* ptr);
extern void Plugify_SetGetVectorSizeString(void* ptr);
extern void Plugify_SetGetVectorSizeVariant(void* ptr);
extern void Plugify_SetGetVectorSizeVector2(void* ptr);
extern void Plugify_SetGetVectorSizeVector3(void* ptr);
extern void Plugify_SetGetVectorSizeVector4(void* ptr);
extern void Plugify_SetGetVectorSizeMatrix4x4(void* ptr);
extern void Plugify_SetGetVectorDataBool(void* ptr);
extern void Plugify_SetGetVectorDataChar8(void* ptr);
extern void Plugify_SetGetVectorDataChar16(void* ptr);
extern void Plugify_SetGetVectorDataInt8(void* ptr);
extern void Plugify_SetGetVectorDataInt16(void* ptr);
extern void Plugify_SetGetVectorDataInt32(void* ptr);
extern void Plugify_SetGetVectorDataInt64(void* ptr);
extern void Plugify_SetGetVectorDataUInt8(void* ptr);
extern void Plugify_SetGetVectorDataUInt16(void* ptr);
extern void Plugify_SetGetVectorDataUInt32(void* ptr);
extern void Plugify_SetGetVectorDataUInt64(void* ptr);
extern void Plugify_SetGetVectorDataPointer(void* ptr);
extern void Plugify_SetGetVectorDataFloat(void* ptr);
extern void Plugify_SetGetVectorDataDouble(void* ptr);
extern void Plugify_SetGetVectorDataString(void* ptr);
extern void Plugify_SetGetVectorDataVariant(void* ptr);
extern void Plugify_SetGetVectorDataVector2(void* ptr);
extern void Plugify_SetGetVectorDataVector3(void* ptr);
extern void Plugify_SetGetVectorDataVector4(void* ptr);
extern void Plugify_SetGetVectorDataMatrix4x4(void* ptr);
extern void Plugify_SetAssignVectorBool(void* ptr);
extern void Plugify_SetAssignVectorChar8(void* ptr);
extern void Plugify_SetAssignVectorChar16(void* ptr);
extern void Plugify_SetAssignVectorInt8(void* ptr);
extern void Plugify_SetAssignVectorInt16(void* ptr);
extern void Plugify_SetAssignVectorInt32(void* ptr);
extern void Plugify_SetAssignVectorInt64(void* ptr);
extern void Plugify_SetAssignVectorUInt8(void* ptr);
extern void Plugify_SetAssignVectorUInt16(void* ptr);
extern void Plugify_SetAssignVectorUInt32(void* ptr);
extern void Plugify_SetAssignVectorUInt64(void* ptr);
extern void Plugify_SetAssignVectorPointer(void* ptr);
extern void Plugify_SetAssignVectorFloat(void* ptr);
extern void Plugify_SetAssignVectorDouble(void* ptr);
extern void Plugify_SetAssignVectorString(void* ptr);
extern void Plugify_SetAssignVectorVariant(void* ptr);
extern void Plugify_SetAssignVectorVector2(void* ptr);
extern void Plugify_SetAssignVectorVector3(void* ptr);
extern void Plugify_SetAssignVectorVector4(void* ptr);
extern void Plugify_SetAssignVectorMatrix4x4(void* ptr);
extern void Plugify_SetNewCall(void* ptr);
extern void Plugify_SetDeleteCall(void* ptr);
extern void Plugify_SetGetCallFunction(void* ptr);
extern void Plugify_SetGetCallError(void* ptr);
extern void Plugify_SetNewCallback(void* ptr);
extern void Plugify_SetDeleteCallback(void* ptr);
extern void Plugify_SetGetCallbackFunction(void* ptr);
extern void Plugify_SetGetCallbackError(void* ptr);
extern void Plugify_SetGetMethodParamCount(void* ptr);
extern void Plugify_SetGetMethodParamType(void* ptr);
extern void Plugify_SetGetMethodPrototype(void* ptr);
extern void Plugify_SetGetMethodEnum(void* ptr);

extern void* aligned_malloc(size_t size, size_t alignment);
extern void aligned_free(void* ptr);

#ifdef __cplusplus
}
#endif
