#include <stddef.h>
#include <assert.h>
typedef struct { const char *p; ptrdiff_t n; } _GoString_;
#include "plugify.h"

_Static_assert(sizeof(Vector2) == 8, "Unsupported Vector2 size");
_Static_assert(sizeof(Vector3) == 12, "Unsupported Vector3 size");
_Static_assert(sizeof(Vector4) == 16, "Unsupported Vector4 size");
_Static_assert(sizeof(Matrix4x4) == 64, "Unsupported Matrix4x4 size");
_Static_assert(sizeof(String) == sizeof(void*) * 3, "Unsupported String size");
_Static_assert(sizeof(Vector) == sizeof(void*) * 3, "Unsupported Vector size");
_Static_assert(sizeof(Variant) == 32, "Unsupported Variant size");
_Static_assert(offsetof(struct Variant, current) == 24, "Unsupported Variant layout");

_Static_assert(sizeof(ManagedType) == 2, "Unsupported ManagedType size");
_Static_assert(sizeof(Parameters) == 8, "Unsupported Parameters size");
_Static_assert(sizeof(Return) == 16, "Unsupported Return size");

PluginHandle pluginHandle = NULL;

// Function pointers
void* (*GetMethodPtr)(const char*) = NULL;
void (*GetMethodPtr2)(const char*, void**) = NULL;
const char* (*GetPluginBaseDir)(PluginHandle) = NULL;
const char* (*GetPluginExtensionsDir)(PluginHandle) = NULL;
const char* (*GetPluginConfigsDir)(PluginHandle) = NULL;
const char* (*GetPluginDataDir)(PluginHandle) = NULL;
const char* (*GetPluginLogsDir)(PluginHandle) = NULL;
const char* (*GetPluginCacheDir)(PluginHandle) = NULL;
bool (*IsExtensionLoaded)(_GoString_, _GoString_) = NULL;
void (*PrintException)(_GoString_) = NULL;

// Function pointers for PluginHandle functions
ptrdiff_t (*GetPluginId)(PluginHandle) = NULL;
const char* (*GetPluginName)(PluginHandle) = NULL;
const char* (*GetPluginDescription)(PluginHandle) = NULL;
const char* (*GetPluginVersion)(PluginHandle) = NULL;
const char* (*GetPluginAuthor)(PluginHandle) = NULL;
const char* (*GetPluginWebsite)(PluginHandle) = NULL;
const char* (*GetPluginLicense)(PluginHandle) = NULL;
const char* (*GetPluginLocation)(PluginHandle) = NULL;
void* (*GetPluginDependencies)(PluginHandle) = NULL;
ptrdiff_t (*GetPluginDependenciesSize)(PluginHandle) = NULL;

// Function pointers for deleting C strings
void (*DeleteCStr)(const char*) = NULL;
void (*DeleteCStrArr)(void*) = NULL;

// Function pointers for String functions
String (*ConstructString)(_GoString_) = NULL;
void (*DestroyString)(String*) = NULL;
const char* (*GetStringData)(String*) = NULL;
ptrdiff_t (*GetStringLength)(String*) = NULL;
void (*AssignString)(String*, _GoString_) = NULL;

// Function pointers for Variant functions
void (*DestroyVariant)(Variant*) = NULL;

// Function pointers for ConstructVector functions
Vector (*ConstructVectorBool)(bool*, ptrdiff_t) = NULL;
Vector (*ConstructVectorChar8)(char*, ptrdiff_t) = NULL;
Vector (*ConstructVectorChar16)(char16_t*, ptrdiff_t) = NULL;
Vector (*ConstructVectorInt8)(int8_t*, ptrdiff_t) = NULL;
Vector (*ConstructVectorInt16)(int16_t*, ptrdiff_t) = NULL;
Vector (*ConstructVectorInt32)(int32_t*, ptrdiff_t) = NULL;
Vector (*ConstructVectorInt64)(int64_t*, ptrdiff_t) = NULL;
Vector (*ConstructVectorUInt8)(uint8_t*, ptrdiff_t) = NULL;
Vector (*ConstructVectorUInt16)(uint16_t*, ptrdiff_t) = NULL;
Vector (*ConstructVectorUInt32)(uint32_t*, ptrdiff_t) = NULL;
Vector (*ConstructVectorUInt64)(uint64_t*, ptrdiff_t) = NULL;
Vector (*ConstructVectorPointer)(uintptr_t*, ptrdiff_t) = NULL;
Vector (*ConstructVectorFloat)(float*, ptrdiff_t) = NULL;
Vector (*ConstructVectorDouble)(double*, ptrdiff_t) = NULL;
Vector (*ConstructVectorString)(_GoString_*, ptrdiff_t) = NULL;
Vector (*ConstructVectorVariant)(ptrdiff_t) = NULL;
Vector (*ConstructVectorVector2)(Vector2*, ptrdiff_t) = NULL;
Vector (*ConstructVectorVector3)(Vector3*, ptrdiff_t) = NULL;
Vector (*ConstructVectorVector4)(Vector4*, ptrdiff_t) = NULL;
Vector (*ConstructVectorMatrix4x4)(Matrix4x4*, ptrdiff_t) = NULL;

// Function pointers for DestroyVector functions
void (*DestroyVectorBool)(Vector*) = NULL;
void (*DestroyVectorChar8)(Vector*) = NULL;
void (*DestroyVectorChar16)(Vector*) = NULL;
void (*DestroyVectorInt8)(Vector*) = NULL;
void (*DestroyVectorInt16)(Vector*) = NULL;
void (*DestroyVectorInt32)(Vector*) = NULL;
void (*DestroyVectorInt64)(Vector*) = NULL;
void (*DestroyVectorUInt8)(Vector*) = NULL;
void (*DestroyVectorUInt16)(Vector*) = NULL;
void (*DestroyVectorUInt32)(Vector*) = NULL;
void (*DestroyVectorUInt64)(Vector*) = NULL;
void (*DestroyVectorPointer)(Vector*) = NULL;
void (*DestroyVectorFloat)(Vector*) = NULL;
void (*DestroyVectorDouble)(Vector*) = NULL;
void (*DestroyVectorString)(Vector*) = NULL;
void (*DestroyVectorVariant)(Vector*) = NULL;
void (*DestroyVectorVector2)(Vector*) = NULL;
void (*DestroyVectorVector3)(Vector*) = NULL;
void (*DestroyVectorVector4)(Vector*) = NULL;
void (*DestroyVectorMatrix4x4)(Vector*) = NULL;

// Function pointers for GetVectorSize functions
ptrdiff_t (*GetVectorSizeBool)(Vector*) = NULL;
ptrdiff_t (*GetVectorSizeChar8)(Vector*) = NULL;
ptrdiff_t (*GetVectorSizeChar16)(Vector*) = NULL;
ptrdiff_t (*GetVectorSizeInt8)(Vector*) = NULL;
ptrdiff_t (*GetVectorSizeInt16)(Vector*) = NULL;
ptrdiff_t (*GetVectorSizeInt32)(Vector*) = NULL;
ptrdiff_t (*GetVectorSizeInt64)(Vector*) = NULL;
ptrdiff_t (*GetVectorSizeUInt8)(Vector*) = NULL;
ptrdiff_t (*GetVectorSizeUInt16)(Vector*) = NULL;
ptrdiff_t (*GetVectorSizeUInt32)(Vector*) = NULL;
ptrdiff_t (*GetVectorSizeUInt64)(Vector*) = NULL;
ptrdiff_t (*GetVectorSizePointer)(Vector*) = NULL;
ptrdiff_t (*GetVectorSizeFloat)(Vector*) = NULL;
ptrdiff_t (*GetVectorSizeDouble)(Vector*) = NULL;
ptrdiff_t (*GetVectorSizeString)(Vector*) = NULL;
ptrdiff_t (*GetVectorSizeVariant)(Vector*) = NULL;
ptrdiff_t (*GetVectorSizeVector2)(Vector*) = NULL;
ptrdiff_t (*GetVectorSizeVector3)(Vector*) = NULL;
ptrdiff_t (*GetVectorSizeVector4)(Vector*) = NULL;
ptrdiff_t (*GetVectorSizeMatrix4x4)(Vector*) = NULL;

// Function pointers for GetVectorData functions
bool* (*GetVectorDataBool)(Vector*) = NULL;
char* (*GetVectorDataChar8)(Vector*) = NULL;
char16_t* (*GetVectorDataChar16)(Vector*) = NULL;
int8_t* (*GetVectorDataInt8)(Vector*) = NULL;
int16_t* (*GetVectorDataInt16)(Vector*) = NULL;
int32_t* (*GetVectorDataInt32)(Vector*) = NULL;
int64_t* (*GetVectorDataInt64)(Vector*) = NULL;
uint8_t* (*GetVectorDataUInt8)(Vector*) = NULL;
uint16_t* (*GetVectorDataUInt16)(Vector*) = NULL;
uint32_t* (*GetVectorDataUInt32)(Vector*) = NULL;
uint64_t* (*GetVectorDataUInt64)(Vector*) = NULL;
uintptr_t* (*GetVectorDataPointer)(Vector*) = NULL;
float* (*GetVectorDataFloat)(Vector*) = NULL;
double* (*GetVectorDataDouble)(Vector*) = NULL;
String* (*GetVectorDataString)(Vector*) = NULL;
Variant* (*GetVectorDataVariant)(Vector*, ptrdiff_t) = NULL;
Vector2* (*GetVectorDataVector2)(Vector*) = NULL;
Vector3* (*GetVectorDataVector3)(Vector*) = NULL;
Vector4* (*GetVectorDataVector4)(Vector*) = NULL;
Matrix4x4* (*GetVectorDataMatrix4x4)(Vector*) = NULL;

// Function pointers for AssignVector functions
void (*AssignVectorBool)(Vector*, bool*, ptrdiff_t) = NULL;
void (*AssignVectorChar8)(Vector*, char*, ptrdiff_t) = NULL;
void (*AssignVectorChar16)(Vector*, char16_t*, ptrdiff_t) = NULL;
void (*AssignVectorInt8)(Vector*, int8_t*, ptrdiff_t) = NULL;
void (*AssignVectorInt16)(Vector*, int16_t*, ptrdiff_t) = NULL;
void (*AssignVectorInt32)(Vector*, int32_t*, ptrdiff_t) = NULL;
void (*AssignVectorInt64)(Vector*, int64_t*, ptrdiff_t) = NULL;
void (*AssignVectorUInt8)(Vector*, uint8_t*, ptrdiff_t) = NULL;
void (*AssignVectorUInt16)(Vector*, uint16_t*, ptrdiff_t) = NULL;
void (*AssignVectorUInt32)(Vector*, uint32_t*, ptrdiff_t) = NULL;
void (*AssignVectorUInt64)(Vector*, uint64_t*, ptrdiff_t) = NULL;
void (*AssignVectorPointer)(Vector*, uintptr_t*, ptrdiff_t) = NULL;
void (*AssignVectorFloat)(Vector*, float*, ptrdiff_t) = NULL;
void (*AssignVectorDouble)(Vector*, double*, ptrdiff_t) = NULL;
void (*AssignVectorString)(Vector*, _GoString_*, ptrdiff_t) = NULL;
void (*AssignVectorVariant)(Vector*, ptrdiff_t) = NULL;
void (*AssignVectorVector2)(Vector*, Vector2*, ptrdiff_t) = NULL;
void (*AssignVectorVector3)(Vector*, Vector3*, ptrdiff_t) = NULL;
void (*AssignVectorVector4)(Vector*, Vector4*, ptrdiff_t) = NULL;
void (*AssignVectorMatrix4x4)(Vector*, Matrix4x4*, ptrdiff_t) = NULL;

// Function pointers for Call functions
JitCall (*NewCall)(void*, ManagedType*, ptrdiff_t, ManagedType) = NULL;
void (*DeleteCall)(JitCall) = NULL;
void* (*GetCallFunction)(JitCall) = NULL;
const char* (*GetCallError)(JitCall) = NULL;

// Function pointers for Callback functions
JitCallback (*NewCallback)(PluginHandle, _GoString_, void*) = NULL;
void (*DeleteCallback)(JitCallback) = NULL;
void* (*GetCallbackFunction)(JitCallback) = NULL;
const char* (*GetCallbackError)(JitCallback) = NULL;

// Function pointers for MethodHandle functions
ptrdiff_t (*GetMethodParamCount)(MethodHandle) = NULL;
ManagedType (*GetMethodParamType)(MethodHandle, ptrdiff_t) = NULL;
MethodHandle (*GetMethodPrototype)(MethodHandle, ptrdiff_t) = NULL;
EnumHandle (*GetMethodEnum)(MethodHandle, ptrdiff_t) = NULL;

// Function to call GetMethodPtr
void* Plugify_GetMethodPtr(const char* name) { return GetMethodPtr(name); }
// Function to call GetMethodPtr2
void Plugify_GetMethodPtr2(const char* name, void** dest) { GetMethodPtr2(name, dest); }
// Function to call GetBaseDir
const char* Plugify_GetBaseDir() { return GetBaseDir(); }
// Function to call GetExtensionsDir
const char* Plugify_GetExtensionsDir() { return GetExtensionsDir(); }
// Function to call GetConfigsDir
const char* Plugify_GetConfigsDir() { return GetConfigsDir(); }
// Function to call GetDataDir
const char* Plugify_GetDataDir() { return GetDataDir(); }
// Function to call GetLogsDir
const char* Plugify_GetLogsDir() { return GetLogsDir(); }
// Function to call GetCacheDir
const char* Plugify_GetCacheDir() { return GetCacheDir(); }
// Function to call IsExtensionLoaded
bool Plugify_IsExtensionLoaded(_GoString_ name, _GoString_ constraint) { return IsExtensionLoaded(name, constraint); }
// Function to call PrintException
void Plugify_PrintException(_GoString_ message) { PrintException(message); }
// Function to call GetPluginId
ptrdiff_t Plugify_GetPluginId() { return GetPluginId(pluginHandle); }
// Function to call GetPluginName
const char* Plugify_GetPluginName() { return GetPluginName(pluginHandle); }
// Function to call GetPluginDescription
const char* Plugify_GetPluginDescription() { return GetPluginDescription(pluginHandle); }
// Function to call GetPluginVersion
const char* Plugify_GetPluginVersion() { return GetPluginVersion(pluginHandle); }
// Function to call GetPluginAuthor
const char* Plugify_GetPluginAuthor() { return GetPluginAuthor(pluginHandle); }
// Function to call GetPluginWebsite
const char* Plugify_GetPluginWebsite() { return GetPluginWebsite(pluginHandle); }
// Function to call GetPluginLicense
const char* Plugify_GetPluginLicense() { return GetPluginLicense(pluginHandle); }
// Function to call GetPluginLocation
const char* Plugify_GetPluginLocation() { return GetPluginLocation(pluginHandle); }
// Function to call GetPluginDependencies
void* Plugify_GetPluginDependencies() { return GetPluginDependencies(pluginHandle); }
// Function to call GetPluginDependenciesSize
ptrdiff_t Plugify_GetPluginDependenciesSize() { return GetPluginDependenciesSize(pluginHandle); }
// Function to delete C string
void Plugify_DeleteCStr(const char* str) { DeleteCStr(str); }
// Function to delete C string array
void Plugify_DeleteCStrArr(void* arr) { DeleteCStrArr(arr); }
// Function to construct a string
String Plugify_ConstructString(_GoString_ source) { return ConstructString(source); }
// Function to destroy a string
void Plugify_DestroyString(String* string) { DestroyString(string); }
// Function to get string data
const char* Plugify_GetStringData(String* string) { return GetStringData(string); }
// Function to get string length
ptrdiff_t Plugify_GetStringLength(String* string) { return GetStringLength(string); }
// Function to assign a string
void Plugify_AssignString(String* string, _GoString_ source) { AssignString(string, source); }
// Function to destroy a variant
void Plugify_DestroyVariant(Variant* any) { DestroyVariant(any); }
// Functions to construct vectors
Vector Plugify_ConstructVectorBool(bool* data, ptrdiff_t size) { return ConstructVectorBool(data, size); }
Vector Plugify_ConstructVectorChar8(char* data, ptrdiff_t size) { return ConstructVectorChar8(data, size); }
Vector Plugify_ConstructVectorChar16(char16_t* data, ptrdiff_t size) { return ConstructVectorChar16(data, size); }
Vector Plugify_ConstructVectorInt8(int8_t* data, ptrdiff_t size) { return ConstructVectorInt8(data, size); }
Vector Plugify_ConstructVectorInt16(int16_t* data, ptrdiff_t size) { return ConstructVectorInt16(data, size); }
Vector Plugify_ConstructVectorInt32(int32_t* data, ptrdiff_t size) { return ConstructVectorInt32(data, size); }
Vector Plugify_ConstructVectorInt64(int64_t* data, ptrdiff_t size) { return ConstructVectorInt64(data, size); }
Vector Plugify_ConstructVectorUInt8(uint8_t* data, ptrdiff_t size) { return ConstructVectorUInt8(data, size); }
Vector Plugify_ConstructVectorUInt16(uint16_t* data, ptrdiff_t size) { return ConstructVectorUInt16(data, size); }
Vector Plugify_ConstructVectorUInt32(uint32_t* data, ptrdiff_t size) { return ConstructVectorUInt32(data, size); }
Vector Plugify_ConstructVectorUInt64(uint64_t* data, ptrdiff_t size) { return ConstructVectorUInt64(data, size); }
Vector Plugify_ConstructVectorPointer(uintptr_t* data, ptrdiff_t size) { return ConstructVectorPointer(data, size); }
Vector Plugify_ConstructVectorFloat(float* data, ptrdiff_t size) { return ConstructVectorFloat(data, size); }
Vector Plugify_ConstructVectorDouble(double* data, ptrdiff_t size) { return ConstructVectorDouble(data, size); }
Vector Plugify_ConstructVectorString(_GoString_* data, ptrdiff_t size) { return ConstructVectorString(data, size); }
Vector Plugify_ConstructVectorVariant(ptrdiff_t size) { return ConstructVectorVariant(size); }
Vector Plugify_ConstructVectorVector2(Vector2* data, ptrdiff_t size) { return ConstructVectorVector2(data, size); }
Vector Plugify_ConstructVectorVector3(Vector3* data, ptrdiff_t size) { return ConstructVectorVector3(data, size); }
Vector Plugify_ConstructVectorVector4(Vector4* data, ptrdiff_t size) { return ConstructVectorVector4(data, size); }
Vector Plugify_ConstructVectorMatrix4x4(Matrix4x4* data, ptrdiff_t size) { return ConstructVectorMatrix4x4(data, size); }
// Functions to destroy vectors
void Plugify_DestroyVectorBool(Vector* vec) { DestroyVectorBool(vec); }
void Plugify_DestroyVectorChar8(Vector* vec) { DestroyVectorChar8(vec); }
void Plugify_DestroyVectorChar16(Vector* vec) { DestroyVectorChar16(vec); }
void Plugify_DestroyVectorInt8(Vector* vec) { DestroyVectorInt8(vec); }
void Plugify_DestroyVectorInt16(Vector* vec) { DestroyVectorInt16(vec); }
void Plugify_DestroyVectorInt32(Vector* vec) { DestroyVectorInt32(vec); }
void Plugify_DestroyVectorInt64(Vector* vec) { DestroyVectorInt64(vec); }
void Plugify_DestroyVectorUInt8(Vector* vec) { DestroyVectorUInt8(vec); }
void Plugify_DestroyVectorUInt16(Vector* vec) { DestroyVectorUInt16(vec); }
void Plugify_DestroyVectorUInt32(Vector* vec) { DestroyVectorUInt32(vec); }
void Plugify_DestroyVectorUInt64(Vector* vec) { DestroyVectorUInt64(vec); }
void Plugify_DestroyVectorPointer(Vector* vec) { DestroyVectorPointer(vec); }
void Plugify_DestroyVectorFloat(Vector* vec) { DestroyVectorFloat(vec); }
void Plugify_DestroyVectorDouble(Vector* vec) { DestroyVectorDouble(vec); }
void Plugify_DestroyVectorString(Vector* vec) { DestroyVectorString(vec); }
void Plugify_DestroyVectorVariant(Vector* vec) { DestroyVectorVariant(vec); }
void Plugify_DestroyVectorVector2(Vector* vec) { DestroyVectorVector2(vec); }
void Plugify_DestroyVectorVector3(Vector* vec) { DestroyVectorVector3(vec); }
void Plugify_DestroyVectorVector4(Vector* vec) { DestroyVectorVector4(vec); }
void Plugify_DestroyVectorMatrix4x4(Vector* vec) { DestroyVectorMatrix4x4(vec); }
// Functions to get vector size
ptrdiff_t Plugify_GetVectorSizeBool(Vector* vec) { return GetVectorSizeBool(vec); }
ptrdiff_t Plugify_GetVectorSizeChar8(Vector* vec) { return GetVectorSizeChar8(vec); }
ptrdiff_t Plugify_GetVectorSizeChar16(Vector* vec) { return GetVectorSizeChar16(vec); }
ptrdiff_t Plugify_GetVectorSizeInt8(Vector* vec) { return GetVectorSizeInt8(vec); }
ptrdiff_t Plugify_GetVectorSizeInt16(Vector* vec) { return GetVectorSizeInt16(vec); }
ptrdiff_t Plugify_GetVectorSizeInt32(Vector* vec) { return GetVectorSizeInt32(vec); }
ptrdiff_t Plugify_GetVectorSizeInt64(Vector* vec) { return GetVectorSizeInt64(vec); }
ptrdiff_t Plugify_GetVectorSizeUInt8(Vector* vec) { return GetVectorSizeUInt8(vec); }
ptrdiff_t Plugify_GetVectorSizeUInt16(Vector* vec) { return GetVectorSizeUInt16(vec); }
ptrdiff_t Plugify_GetVectorSizeUInt32(Vector* vec) { return GetVectorSizeUInt32(vec); }
ptrdiff_t Plugify_GetVectorSizeUInt64(Vector* vec) { return GetVectorSizeUInt64(vec); }
ptrdiff_t Plugify_GetVectorSizePointer(Vector* vec) { return GetVectorSizePointer(vec); }
ptrdiff_t Plugify_GetVectorSizeFloat(Vector* vec) { return GetVectorSizeFloat(vec); }
ptrdiff_t Plugify_GetVectorSizeDouble(Vector* vec) { return GetVectorSizeDouble(vec); }
ptrdiff_t Plugify_GetVectorSizeString(Vector* vec) { return GetVectorSizeString(vec); }
ptrdiff_t Plugify_GetVectorSizeVariant(Vector* vec) { return GetVectorSizeVariant(vec); }
ptrdiff_t Plugify_GetVectorSizeVector2(Vector* vec) { return GetVectorSizeVector2(vec); }
ptrdiff_t Plugify_GetVectorSizeVector3(Vector* vec) { return GetVectorSizeVector3(vec); }
ptrdiff_t Plugify_GetVectorSizeVector4(Vector* vec) { return GetVectorSizeVector4(vec); }
ptrdiff_t Plugify_GetVectorSizeMatrix4x4(Vector* vec) { return GetVectorSizeMatrix4x4(vec); }
// Functions to get vector data
bool* Plugify_GetVectorDataBool(Vector* vec) { return GetVectorDataBool(vec); }
char* Plugify_GetVectorDataChar8(Vector* vec) { return GetVectorDataChar8(vec); }
char16_t* Plugify_GetVectorDataChar16(Vector* vec) { return GetVectorDataChar16(vec); }
int8_t* Plugify_GetVectorDataInt8(Vector* vec) { return GetVectorDataInt8(vec); }
int16_t* Plugify_GetVectorDataInt16(Vector* vec) { return GetVectorDataInt16(vec); }
int32_t* Plugify_GetVectorDataInt32(Vector* vec) { return GetVectorDataInt32(vec); }
int64_t* Plugify_GetVectorDataInt64(Vector* vec) { return GetVectorDataInt64(vec); }
uint8_t* Plugify_GetVectorDataUInt8(Vector* vec) { return GetVectorDataUInt8(vec); }
uint16_t* Plugify_GetVectorDataUInt16(Vector* vec) { return GetVectorDataUInt16(vec); }
uint32_t* Plugify_GetVectorDataUInt32(Vector* vec) { return GetVectorDataUInt32(vec); }
uint64_t* Plugify_GetVectorDataUInt64(Vector* vec) { return GetVectorDataUInt64(vec); }
uintptr_t* Plugify_GetVectorDataPointer(Vector* vec) { return GetVectorDataPointer(vec); }
float* Plugify_GetVectorDataFloat(Vector* vec) { return GetVectorDataFloat(vec); }
double* Plugify_GetVectorDataDouble(Vector* vec) { return GetVectorDataDouble(vec); }
String* Plugify_GetVectorDataString(Vector* vec) { return GetVectorDataString(vec); }
Variant* Plugify_GetVectorDataVariant(Vector* vec, ptrdiff_t index) { return GetVectorDataVariant(vec, index); }
Vector2* Plugify_GetVectorDataVector2(Vector* vec) { return GetVectorDataVector2(vec); }
Vector3* Plugify_GetVectorDataVector3(Vector* vec) { return GetVectorDataVector3(vec); }
Vector4* Plugify_GetVectorDataVector4(Vector* vec) { return GetVectorDataVector4(vec); }
Matrix4x4* Plugify_GetVectorDataMatrix4x4(Vector* vec) { return GetVectorDataMatrix4x4(vec); }
// Functions to assign vectors
void Plugify_AssignVectorBool(Vector* vec, bool* data, ptrdiff_t size) { AssignVectorBool(vec, data, size); }
void Plugify_AssignVectorChar8(Vector* vec, char* data, ptrdiff_t size) { AssignVectorChar8(vec, data, size); }
void Plugify_AssignVectorChar16(Vector* vec, char16_t* data, ptrdiff_t size) { AssignVectorChar16(vec, data, size); }
void Plugify_AssignVectorInt8(Vector* vec, int8_t* data, ptrdiff_t size) { AssignVectorInt8(vec, data, size); }
void Plugify_AssignVectorInt16(Vector* vec, int16_t* data, ptrdiff_t size) { AssignVectorInt16(vec, data, size); }
void Plugify_AssignVectorInt32(Vector* vec, int32_t* data, ptrdiff_t size) { AssignVectorInt32(vec, data, size); }
void Plugify_AssignVectorInt64(Vector* vec, int64_t* data, ptrdiff_t size) { AssignVectorInt64(vec, data, size); }
void Plugify_AssignVectorUInt8(Vector* vec, uint8_t* data, ptrdiff_t size) { AssignVectorUInt8(vec, data, size); }
void Plugify_AssignVectorUInt16(Vector* vec, uint16_t* data, ptrdiff_t size) { AssignVectorUInt16(vec, data, size); }
void Plugify_AssignVectorUInt32(Vector* vec, uint32_t* data, ptrdiff_t size) { AssignVectorUInt32(vec, data, size); }
void Plugify_AssignVectorUInt64(Vector* vec, uint64_t* data, ptrdiff_t size) { AssignVectorUInt64(vec, data, size); }
void Plugify_AssignVectorPointer(Vector* vec, uintptr_t* data, ptrdiff_t size) { AssignVectorPointer(vec, data, size); }
void Plugify_AssignVectorFloat(Vector* vec, float* data, ptrdiff_t size) { AssignVectorFloat(vec, data, size); }
void Plugify_AssignVectorDouble(Vector* vec, double* data, ptrdiff_t size) { AssignVectorDouble(vec, data, size); }
void Plugify_AssignVectorString(Vector* vec, _GoString_* data, ptrdiff_t size) { AssignVectorString(vec, data, size); }
void Plugify_AssignVectorVariant(Vector* vec, ptrdiff_t size) { AssignVectorVariant(vec, size); }
void Plugify_AssignVectorVector2(Vector* vec, Vector2* data, ptrdiff_t size) { AssignVectorVector2(vec, data, size); }
void Plugify_AssignVectorVector3(Vector* vec, Vector3* data, ptrdiff_t size) { AssignVectorVector3(vec, data, size); }
void Plugify_AssignVectorVector4(Vector* vec, Vector4* data, ptrdiff_t size) { AssignVectorVector4(vec, data, size); }
void Plugify_AssignVectorMatrix4x4(Vector* vec, Matrix4x4* data, ptrdiff_t size) { AssignVectorMatrix4x4(vec, data, size); }

JitCall Plugify_NewCall(void* target, ManagedType* params, ptrdiff_t count, ManagedType ret) { return NewCall(target, params, count, ret); }
void Plugify_DeleteCall(JitCall call) { DeleteCall(call); }
void* Plugify_GetCallFunction(JitCall call) { return GetCallFunction(call); }
const char* Plugify_GetCallError(JitCall call) { return GetCallError(call); }

void Plugify_CallFunction(JitCall call, uint64_t* params, uint128_t* ret) {
	((void(*)(uint64_t*, uint128_t*))GetCallFunction(call))(params, ret);
}

JitCallback Plugify_NewCallback(_GoString_ name, void* handle) { return NewCallback(pluginHandle, name, handle); }
void Plugify_DeleteCallback(JitCallback callback) { return DeleteCallback(callback); }
void* Plugify_GetCallbackFunction(JitCallback callback) { return GetCallbackFunction(callback); }
const char* Plugify_GetCallbackError(JitCallback callback) { return GetCallbackError(callback); }

ptrdiff_t Plugify_GetMethodParamCount(MethodHandle handle) { return GetMethodParamCount(handle); }
ManagedType Plugify_GetMethodParamType(MethodHandle handle, ptrdiff_t index) { return GetMethodParamType(handle, index); }
MethodHandle Plugify_GetMethodPrototype(MethodHandle handle, ptrdiff_t index) { return GetMethodPrototype(handle, index); }
EnumHandle Plugify_GetMethodEnum(MethodHandle handle, ptrdiff_t index) { return GetMethodEnum(handle, index); }

void Plugify_SetGetMethodPtr(void* ptr) { GetMethodPtr = (void* (*)(const char*)) ptr; }
void Plugify_SetGetMethodPtr2(void* ptr) { GetMethodPtr2 = (void (*)(const char*, void**)) ptr; }
void Plugify_SetGetBaseDir(void* ptr) { GetBaseDir = (const char* (*)(PluginHandle)) ptr; }
void Plugify_SetGetExtensionsDir(void* ptr) { GetExtensionsDir = (const char* (*)(PluginHandle)) ptr; }
void Plugify_SetGetConfigsDir(void* ptr) { GetConfigsDir = (const char* (*)(PluginHandle)) ptr; }
void Plugify_SetGetDataDir(void* ptr) { GetDataDir = (const char* (*)(PluginHandle)) ptr; }
void Plugify_SetGetLogsDir(void* ptr) { GetLogsDir = (const char* (*)(PluginHandle)) ptr; }
void Plugify_SetGetCacheDir(void* ptr) { GetCacheDir = (const char* (*)(PluginHandle)) ptr; }
void Plugify_SetIsExtensionLoaded(void* ptr) { IsExtensionLoaded = (bool (*)(_GoString_, _GoString_)) ptr; }
void Plugify_SetPrintException(void* ptr) { PrintException = (void (*)(_GoString_)) ptr; }
void Plugify_SetGetPluginId(void* ptr) { GetPluginId = (ptrdiff_t (*)(PluginHandle)) ptr; }
void Plugify_SetGetPluginName(void* ptr) { GetPluginName = (const char* (*)(PluginHandle)) ptr; }
void Plugify_SetGetPluginDescription(void* ptr) { GetPluginDescription = (const char* (*)(PluginHandle)) ptr; }
void Plugify_SetGetPluginVersion(void* ptr) { GetPluginVersion = (const char* (*)(PluginHandle)) ptr; }
void Plugify_SetGetPluginAuthor(void* ptr) { GetPluginAuthor = (const char* (*)(PluginHandle)) ptr; }
void Plugify_SetGetPluginWebsite(void* ptr) { GetPluginWebsite = (const char* (*)(PluginHandle)) ptr; }
void Plugify_SetGetPluginLicense(void* ptr) { GetPluginLicense = (const char* (*)(PluginHandle)) ptr; }
void Plugify_SetGetPluginLocation(void* ptr) { GetPluginLocation = (const char* (*)(PluginHandle)) ptr; }
void Plugify_SetGetPluginDependencies(void* ptr) { GetPluginDependencies = (void* (*)(PluginHandle)) ptr; }
void Plugify_SetGetPluginDependenciesSize(void* ptr) { GetPluginDependenciesSize = (ptrdiff_t (*)(PluginHandle)) ptr; }
void Plugify_SetDeleteCStr(void* ptr) { DeleteCStr = (void (*)(const char*)) ptr; }
void Plugify_SetDeleteCStrArr(void* ptr) { DeleteCStrArr = (void (*)(void*)) ptr; }
void Plugify_SetConstructString(void* ptr) { ConstructString = (String (*)(_GoString_)) ptr; }
void Plugify_SetDestroyString(void* ptr) { DestroyString = (void (*)(String*)) ptr; }
void Plugify_SetGetStringData(void* ptr) { GetStringData = (const char* (*)(String*)) ptr; }
void Plugify_SetGetStringLength(void* ptr) { GetStringLength = (ptrdiff_t (*)(String*)) ptr; }
void Plugify_SetAssignString(void* ptr) { AssignString = (void (*)(String*, _GoString_)) ptr; }
void Plugify_SetDestroyVariant(void* ptr) { DestroyVariant = (void (*)(Variant*)) ptr; }
void Plugify_SetConstructVectorBool(void* ptr) { ConstructVectorBool = (Vector (*)(bool*, ptrdiff_t)) ptr; }
void Plugify_SetConstructVectorChar8(void* ptr) { ConstructVectorChar8 = (Vector (*)(char*, ptrdiff_t)) ptr; }
void Plugify_SetConstructVectorChar16(void* ptr) { ConstructVectorChar16 = (Vector (*)(char16_t*, ptrdiff_t)) ptr; }
void Plugify_SetConstructVectorInt8(void* ptr) { ConstructVectorInt8 = (Vector (*)(int8_t*, ptrdiff_t)) ptr; }
void Plugify_SetConstructVectorInt16(void* ptr) { ConstructVectorInt16 = (Vector (*)(int16_t*, ptrdiff_t)) ptr; }
void Plugify_SetConstructVectorInt32(void* ptr) { ConstructVectorInt32 = (Vector (*)(int32_t*, ptrdiff_t)) ptr; }
void Plugify_SetConstructVectorInt64(void* ptr) { ConstructVectorInt64 = (Vector (*)(int64_t*, ptrdiff_t)) ptr; }
void Plugify_SetConstructVectorUInt8(void* ptr) { ConstructVectorUInt8 = (Vector (*)(uint8_t*, ptrdiff_t)) ptr; }
void Plugify_SetConstructVectorUInt16(void* ptr) { ConstructVectorUInt16 = (Vector (*)(uint16_t*, ptrdiff_t)) ptr; }
void Plugify_SetConstructVectorUInt32(void* ptr) { ConstructVectorUInt32 = (Vector (*)(uint32_t*, ptrdiff_t)) ptr; }
void Plugify_SetConstructVectorUInt64(void* ptr) { ConstructVectorUInt64 = (Vector (*)(uint64_t*, ptrdiff_t)) ptr; }
void Plugify_SetConstructVectorPointer(void* ptr) { ConstructVectorPointer = (Vector (*)(uintptr_t*, ptrdiff_t)) ptr; }
void Plugify_SetConstructVectorFloat(void* ptr) { ConstructVectorFloat = (Vector (*)(float*, ptrdiff_t)) ptr; }
void Plugify_SetConstructVectorDouble(void* ptr) { ConstructVectorDouble = (Vector (*)(double*, ptrdiff_t)) ptr; }
void Plugify_SetConstructVectorString(void* ptr) { ConstructVectorString = (Vector (*)(_GoString_*, ptrdiff_t)) ptr; }
void Plugify_SetConstructVectorVariant(void* ptr) { ConstructVectorVariant = (Vector (*)(ptrdiff_t)) ptr; }
void Plugify_SetConstructVectorVector2(void* ptr) { ConstructVectorVector2 = (Vector (*)(Vector2*, ptrdiff_t)) ptr; }
void Plugify_SetConstructVectorVector3(void* ptr) { ConstructVectorVector3 = (Vector (*)(Vector3*, ptrdiff_t)) ptr; }
void Plugify_SetConstructVectorVector4(void* ptr) { ConstructVectorVector4 = (Vector (*)(Vector4*, ptrdiff_t)) ptr; }
void Plugify_SetConstructVectorMatrix4x4(void* ptr) { ConstructVectorMatrix4x4 = (Vector (*)(Matrix4x4*, ptrdiff_t)) ptr; }
void Plugify_SetDestroyVectorBool(void* ptr) { DestroyVectorBool = (void (*)(Vector*)) ptr; }
void Plugify_SetDestroyVectorChar8(void* ptr) { DestroyVectorChar8 = (void (*)(Vector*)) ptr; }
void Plugify_SetDestroyVectorChar16(void* ptr) { DestroyVectorChar16 = (void (*)(Vector*)) ptr; }
void Plugify_SetDestroyVectorInt8(void* ptr) { DestroyVectorInt8 = (void (*)(Vector*)) ptr; }
void Plugify_SetDestroyVectorInt16(void* ptr) { DestroyVectorInt16 = (void (*)(Vector*)) ptr; }
void Plugify_SetDestroyVectorInt32(void* ptr) { DestroyVectorInt32 = (void (*)(Vector*)) ptr; }
void Plugify_SetDestroyVectorInt64(void* ptr) { DestroyVectorInt64 = (void (*)(Vector*)) ptr; }
void Plugify_SetDestroyVectorUInt8(void* ptr) { DestroyVectorUInt8 = (void (*)(Vector*)) ptr; }
void Plugify_SetDestroyVectorUInt16(void* ptr) { DestroyVectorUInt16 = (void (*)(Vector*)) ptr; }
void Plugify_SetDestroyVectorUInt32(void* ptr) { DestroyVectorUInt32 = (void (*)(Vector*)) ptr; }
void Plugify_SetDestroyVectorUInt64(void* ptr) { DestroyVectorUInt64 = (void (*)(Vector*)) ptr; }
void Plugify_SetDestroyVectorPointer(void* ptr) { DestroyVectorPointer = (void (*)(Vector*)) ptr; }
void Plugify_SetDestroyVectorFloat(void* ptr) { DestroyVectorFloat = (void (*)(Vector*)) ptr; }
void Plugify_SetDestroyVectorDouble(void* ptr) { DestroyVectorDouble = (void (*)(Vector*)) ptr; }
void Plugify_SetDestroyVectorString(void* ptr) { DestroyVectorString = (void (*)(Vector*)) ptr; }
void Plugify_SetDestroyVectorVariant(void* ptr) { DestroyVectorVariant = (void (*)(Vector*)) ptr; }
void Plugify_SetDestroyVectorVector2(void* ptr) { DestroyVectorVector2 = (void (*)(Vector*)) ptr; }
void Plugify_SetDestroyVectorVector3(void* ptr) { DestroyVectorVector3 = (void (*)(Vector*)) ptr; }
void Plugify_SetDestroyVectorVector4(void* ptr) { DestroyVectorVector4 = (void (*)(Vector*)) ptr; }
void Plugify_SetDestroyVectorMatrix4x4(void* ptr) { DestroyVectorMatrix4x4 = (void (*)(Vector*)) ptr; }
void Plugify_SetGetVectorSizeBool(void* ptr) { GetVectorSizeBool = (ptrdiff_t (*)(Vector*)) ptr; }
void Plugify_SetGetVectorSizeChar8(void* ptr) { GetVectorSizeChar8 = (ptrdiff_t (*)(Vector*)) ptr; }
void Plugify_SetGetVectorSizeChar16(void* ptr) { GetVectorSizeChar16 = (ptrdiff_t (*)(Vector*)) ptr; }
void Plugify_SetGetVectorSizeInt8(void* ptr) { GetVectorSizeInt8 = (ptrdiff_t (*)(Vector*)) ptr; }
void Plugify_SetGetVectorSizeInt16(void* ptr) { GetVectorSizeInt16 = (ptrdiff_t (*)(Vector*)) ptr; }
void Plugify_SetGetVectorSizeInt32(void* ptr) { GetVectorSizeInt32 = (ptrdiff_t (*)(Vector*)) ptr; }
void Plugify_SetGetVectorSizeInt64(void* ptr) { GetVectorSizeInt64 = (ptrdiff_t (*)(Vector*)) ptr; }
void Plugify_SetGetVectorSizeUInt8(void* ptr) { GetVectorSizeUInt8 = (ptrdiff_t (*)(Vector*)) ptr; }
void Plugify_SetGetVectorSizeUInt16(void* ptr) { GetVectorSizeUInt16 = (ptrdiff_t (*)(Vector*)) ptr; }
void Plugify_SetGetVectorSizeUInt32(void* ptr) { GetVectorSizeUInt32 = (ptrdiff_t (*)(Vector*)) ptr; }
void Plugify_SetGetVectorSizeUInt64(void* ptr) { GetVectorSizeUInt64 = (ptrdiff_t (*)(Vector*)) ptr; }
void Plugify_SetGetVectorSizePointer(void* ptr) { GetVectorSizePointer = (ptrdiff_t (*)(Vector*)) ptr; }
void Plugify_SetGetVectorSizeFloat(void* ptr) { GetVectorSizeFloat = (ptrdiff_t (*)(Vector*)) ptr; }
void Plugify_SetGetVectorSizeDouble(void* ptr) { GetVectorSizeDouble = (ptrdiff_t (*)(Vector*)) ptr; }
void Plugify_SetGetVectorSizeString(void* ptr) { GetVectorSizeString = (ptrdiff_t (*)(Vector*)) ptr; }
void Plugify_SetGetVectorSizeVariant(void* ptr) { GetVectorSizeVariant = (ptrdiff_t (*)(Vector*)) ptr; }
void Plugify_SetGetVectorSizeVector2(void* ptr) { GetVectorSizeVector2 = (ptrdiff_t (*)(Vector*)) ptr; }
void Plugify_SetGetVectorSizeVector3(void* ptr) { GetVectorSizeVector3 = (ptrdiff_t (*)(Vector*)) ptr; }
void Plugify_SetGetVectorSizeVector4(void* ptr) { GetVectorSizeVector4 = (ptrdiff_t (*)(Vector*)) ptr; }
void Plugify_SetGetVectorSizeMatrix4x4(void* ptr) { GetVectorSizeMatrix4x4 = (ptrdiff_t (*)(Vector*)) ptr; }
void Plugify_SetGetVectorDataBool(void* ptr) { GetVectorDataBool = (bool* (*)(Vector*)) ptr; }
void Plugify_SetGetVectorDataChar8(void* ptr) { GetVectorDataChar8 = (char* (*)(Vector*)) ptr; }
void Plugify_SetGetVectorDataChar16(void* ptr) { GetVectorDataChar16 = (char16_t* (*)(Vector*)) ptr; }
void Plugify_SetGetVectorDataInt8(void* ptr) { GetVectorDataInt8 = (int8_t* (*)(Vector*)) ptr; }
void Plugify_SetGetVectorDataInt16(void* ptr) { GetVectorDataInt16 = (int16_t* (*)(Vector*)) ptr; }
void Plugify_SetGetVectorDataInt32(void* ptr) { GetVectorDataInt32 = (int32_t* (*)(Vector*)) ptr; }
void Plugify_SetGetVectorDataInt64(void* ptr) { GetVectorDataInt64 = (int64_t* (*)(Vector*)) ptr; }
void Plugify_SetGetVectorDataUInt8(void* ptr) { GetVectorDataUInt8 = (uint8_t* (*)(Vector*)) ptr; }
void Plugify_SetGetVectorDataUInt16(void* ptr) { GetVectorDataUInt16 = (uint16_t* (*)(Vector*)) ptr; }
void Plugify_SetGetVectorDataUInt32(void* ptr) { GetVectorDataUInt32 = (uint32_t* (*)(Vector*)) ptr; }
void Plugify_SetGetVectorDataUInt64(void* ptr) { GetVectorDataUInt64 = (uint64_t* (*)(Vector*)) ptr; }
void Plugify_SetGetVectorDataPointer(void* ptr) { GetVectorDataPointer = (uintptr_t* (*)(Vector*)) ptr; }
void Plugify_SetGetVectorDataFloat(void* ptr) { GetVectorDataFloat = (float* (*)(Vector*)) ptr; }
void Plugify_SetGetVectorDataDouble(void* ptr) { GetVectorDataDouble = (double* (*)(Vector*)) ptr; }
void Plugify_SetGetVectorDataString(void* ptr) { GetVectorDataString = (String* (*)(Vector*)) ptr; }
void Plugify_SetGetVectorDataVariant(void* ptr) { GetVectorDataVariant = (Variant* (*)(Vector*, ptrdiff_t)) ptr; }
void Plugify_SetGetVectorDataVector2(void* ptr) { GetVectorDataVector2 = (Vector2* (*)(Vector*)) ptr; }
void Plugify_SetGetVectorDataVector3(void* ptr) { GetVectorDataVector3 = (Vector3* (*)(Vector*)) ptr; }
void Plugify_SetGetVectorDataVector4(void* ptr) { GetVectorDataVector4 = (Vector4* (*)(Vector*)) ptr; }
void Plugify_SetGetVectorDataMatrix4x4(void* ptr) { GetVectorDataMatrix4x4 = (Matrix4x4* (*)(Vector*)) ptr; }
void Plugify_SetAssignVectorBool(void* ptr) { AssignVectorBool = (void (*)(Vector*, bool*, ptrdiff_t)) ptr; }
void Plugify_SetAssignVectorChar8(void* ptr) { AssignVectorChar8 = (void (*)(Vector*, char*, ptrdiff_t)) ptr; }
void Plugify_SetAssignVectorChar16(void* ptr) { AssignVectorChar16 = (void (*)(Vector*, char16_t*, ptrdiff_t)) ptr; }
void Plugify_SetAssignVectorInt8(void* ptr) { AssignVectorInt8 = (void (*)(Vector*, int8_t*, ptrdiff_t)) ptr; }
void Plugify_SetAssignVectorInt16(void* ptr) { AssignVectorInt16 = (void (*)(Vector*, int16_t*, ptrdiff_t)) ptr; }
void Plugify_SetAssignVectorInt32(void* ptr) { AssignVectorInt32 = (void (*)(Vector*, int32_t*, ptrdiff_t)) ptr; }
void Plugify_SetAssignVectorInt64(void* ptr) { AssignVectorInt64 = (void (*)(Vector*, int64_t*, ptrdiff_t)) ptr; }
void Plugify_SetAssignVectorUInt8(void* ptr) { AssignVectorUInt8 = (void (*)(Vector*, uint8_t*, ptrdiff_t)) ptr; }
void Plugify_SetAssignVectorUInt16(void* ptr) { AssignVectorUInt16 = (void (*)(Vector*, uint16_t*, ptrdiff_t)) ptr; }
void Plugify_SetAssignVectorUInt32(void* ptr) { AssignVectorUInt32 = (void (*)(Vector*, uint32_t*, ptrdiff_t)) ptr; }
void Plugify_SetAssignVectorUInt64(void* ptr) { AssignVectorUInt64 = (void (*)(Vector*, uint64_t*, ptrdiff_t)) ptr; }
void Plugify_SetAssignVectorPointer(void* ptr) { AssignVectorPointer = (void (*)(Vector*, uintptr_t*, ptrdiff_t)) ptr; }
void Plugify_SetAssignVectorFloat(void* ptr) { AssignVectorFloat = (void (*)(Vector*, float*, ptrdiff_t)) ptr; }
void Plugify_SetAssignVectorDouble(void* ptr) { AssignVectorDouble = (void (*)(Vector*, double*, ptrdiff_t)) ptr; }
void Plugify_SetAssignVectorString(void* ptr) { AssignVectorString = (void (*)(Vector*, _GoString_*, ptrdiff_t)) ptr; }
void Plugify_SetAssignVectorVariant(void* ptr) { AssignVectorVariant = (void (*)(Vector*, ptrdiff_t)) ptr; }
void Plugify_SetAssignVectorVector2(void* ptr) { AssignVectorVector2 = (void (*)(Vector*, Vector2*, ptrdiff_t)) ptr; }
void Plugify_SetAssignVectorVector3(void* ptr) { AssignVectorVector3 = (void (*)(Vector*, Vector3*, ptrdiff_t)) ptr; }
void Plugify_SetAssignVectorVector4(void* ptr) { AssignVectorVector4 = (void (*)(Vector*, Vector4*, ptrdiff_t)) ptr; }
void Plugify_SetAssignVectorMatrix4x4(void* ptr) { AssignVectorMatrix4x4 = (void (*)(Vector*, Matrix4x4*, ptrdiff_t)) ptr; }
void Plugify_SetNewCall(void* ptr) { NewCall = (JitCall (*)(void*, ManagedType*, ptrdiff_t, ManagedType)) ptr; }
void Plugify_SetDeleteCall(void* ptr) { DeleteCall = (void (*)(JitCall)) ptr; }
void Plugify_SetGetCallFunction(void* ptr) { GetCallFunction = (void* (*)(JitCall)) ptr; }
void Plugify_SetGetCallError(void* ptr) { GetCallError = (const char* (*)(JitCall)) ptr; }
void Plugify_SetNewCallback(void* ptr) { NewCallback = (JitCallback (*)(PluginHandle, _GoString_, void*)) ptr; }
void Plugify_SetDeleteCallback(void* ptr) { DeleteCallback = (void (*)(JitCallback)) ptr; }
void Plugify_SetGetCallbackFunction(void* ptr) { GetCallbackFunction = (void* (*)(JitCallback)) ptr; }
void Plugify_SetGetCallbackError(void* ptr) { GetCallbackError = (const char* (*)(JitCallback)) ptr; }
void Plugify_SetGetMethodParamCount(void* ptr) { GetMethodParamCount = (ptrdiff_t (*)(MethodHandle)) ptr; }
void Plugify_SetGetMethodParamType(void* ptr) { GetMethodParamType = (ManagedType (*)(MethodHandle, ptrdiff_t)) ptr; }
void Plugify_SetGetMethodPrototype(void* ptr) { GetMethodPrototype = (MethodHandle (*)(MethodHandle, ptrdiff_t)) ptr; }
void Plugify_SetGetMethodEnum(void* ptr) { GetMethodEnum = (EnumHandle (*)(MethodHandle, ptrdiff_t)) ptr; }

#ifdef _WIN32
#include <malloc.h>
void* aligned_malloc(size_t size, size_t alignment) {
	void* ptr = _aligned_malloc(size, alignment);
	memset(ptr, 0, size);
	return ptr;
}

void aligned_free(void* ptr) {
	_aligned_free(ptr);
}
#else
void* aligned_malloc(size_t size, size_t alignment) {
	void* ptr = aligned_alloc(alignment, size);
	memset(ptr, 0, size);
	return ptr;
}

void aligned_free(void* ptr) {
	free(ptr);
}
#endif
