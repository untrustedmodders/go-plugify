package plugify

/*
#cgo LDFLAGS: -L${SRCDIR}/libplugify.a
#include <plugify.h>
*/
import "C"
import "unsafe"
import "fmt"

type PluginStartCallback func()
type PluginEndCallback func()

type Plugify struct {
	Id                     int64
	Name                   string
	FullName               string
	Description            string
	Version                string
	Author                 string
	Website                string
	BaseDir                string
	Dependencies           []string
	fnPluginStartCallback  PluginStartCallback
	fnPluginEndCallback    PluginEndCallback
}

var plugify Plugify = Plugify {
	Id:                    -1,
	Name:                  "",
	FullName:              "",
	Description:           "",
	Version:               "",
	Author:                "",
	Website:               "",
	BaseDir:               "",
	Dependencies:          []string{},
	fnPluginStartCallback: func() {},
	fnPluginEndCallback:   func() {},
}

var BaseDir string = ""

const kApiVersion = 1

//export Plugify_Init
func Plugify_Init(api []uintptr, version int32, handle uintptr) int32 {
	if version < kApiVersion {
		return kApiVersion
	}
	i := 0
	C.Plugify_SetGetMethodPtr(unsafe.Pointer(api[i])); i++
	C.Plugify_SetGetBaseDir(unsafe.Pointer(api[i])); i++
	C.Plugify_SetIsModuleLoaded(unsafe.Pointer(api[i])); i++
	C.Plugify_SetIsPluginLoaded(unsafe.Pointer(api[i])); i++
	C.Plugify_SetGetPluginId(unsafe.Pointer(api[i])); i++
	C.Plugify_SetGetPluginName(unsafe.Pointer(api[i])); i++
	C.Plugify_SetGetPluginFullName(unsafe.Pointer(api[i])); i++
	C.Plugify_SetGetPluginDescription(unsafe.Pointer(api[i])); i++
	C.Plugify_SetGetPluginVersion(unsafe.Pointer(api[i])); i++
	C.Plugify_SetGetPluginAuthor(unsafe.Pointer(api[i])); i++
	C.Plugify_SetGetPluginWebsite(unsafe.Pointer(api[i])); i++
	C.Plugify_SetGetPluginBaseDir(unsafe.Pointer(api[i])); i++
	C.Plugify_SetGetPluginDependencies(unsafe.Pointer(api[i])); i++
	C.Plugify_SetGetPluginDependenciesSize(unsafe.Pointer(api[i])); i++
	C.Plugify_SetFindPluginResource(unsafe.Pointer(api[i])); i++
	C.Plugify_SetDeleteCStr(unsafe.Pointer(api[i])); i++
	C.Plugify_SetAllocateString(unsafe.Pointer(api[i])); i++
	C.Plugify_SetCreateString(unsafe.Pointer(api[i])); i++
	C.Plugify_SetGetStringData(unsafe.Pointer(api[i])); i++
	C.Plugify_SetGetStringLength(unsafe.Pointer(api[i])); i++
	C.Plugify_SetConstructString(unsafe.Pointer(api[i])); i++
	C.Plugify_SetAssignString(unsafe.Pointer(api[i])); i++
	C.Plugify_SetFreeString(unsafe.Pointer(api[i])); i++
	C.Plugify_SetDeleteString(unsafe.Pointer(api[i])); i++
	C.Plugify_SetCreateVector(unsafe.Pointer(api[i])); i++
	C.Plugify_SetAllocateVector(unsafe.Pointer(api[i])); i++
	C.Plugify_SetGetVectorSize(unsafe.Pointer(api[i])); i++
	C.Plugify_SetGetVectorData(unsafe.Pointer(api[i])); i++
	C.Plugify_SetConstructVector(unsafe.Pointer(api[i])); i++
	C.Plugify_SetAssignVector(unsafe.Pointer(api[i])); i++
	C.Plugify_SetDeleteVector(unsafe.Pointer(api[i])); i++
	C.Plugify_SetFreeVector(unsafe.Pointer(api[i])); i++
	C.Plugify_SetDeleteVectorDataBool(unsafe.Pointer(api[i])); i++
	C.Plugify_SetDeleteVectorDataCStr(unsafe.Pointer(api[i])); i++
	C.Plugify_SetPluginHandle(unsafe.Pointer(handle))
	
	path := C.Plugify_GetBaseDir()
	BaseDir = C.GoString(path)
	C.Plugify_DeleteCStr(path)
	
	plugify.Id = int64(C.Plugify_GetPluginId())
	plugify.Name = C.GoString(C.Plugify_GetPluginName())
	plugify.FullName = C.GoString(C.Plugify_GetPluginFullName())
	plugify.Description = C.GoString(C.Plugify_GetPluginDescription())
	plugify.Version = C.GoString(C.Plugify_GetPluginVersion())
	plugify.Author = C.GoString(C.Plugify_GetPluginAuthor())
	plugify.Website = C.GoString(C.Plugify_GetPluginWebsite())
	
	pluginPath := C.Plugify_GetPluginBaseDir()
	plugify.BaseDir = C.GoString(pluginPath)
	C.Plugify_DeleteCStr(pluginPath)
	
	dependencies := C.Plugify_GetPluginDependencies()
	plugify.Dependencies = make([]string, C.Plugify_GetPluginDependenciesSize())
	for i := range plugify.Dependencies {
		plugify.Dependencies[i] = C.GoString(*(**C.char)(unsafe.Pointer(uintptr(dependencies) + uintptr(i * C.sizeof_uintptr_t))))
	}
	C.Plugify_DeleteVectorDataCStr(dependencies)

	return 0
}

//export Plugify_PluginStart
func Plugify_PluginStart() {
	plugify.fnPluginStartCallback()
}

func OnPluginStart(fn PluginStartCallback) {
	plugify.fnPluginStartCallback = fn
}

//export Plugify_PluginEnd
func Plugify_PluginEnd() {
	plugify.fnPluginEndCallback()
}

func OnPluginEnd(fn PluginEndCallback) {
	plugify.fnPluginEndCallback = fn
}

func (p *Plugify) FindResource(path string) string {
	C_path := C.CString(path)
    C_output := C.Plugify_FindPluginResource(C_path)
	output := C.GoString(C_output)
	C.Plugify_DeleteCStr(C_output)
	C.free(unsafe.Pointer(C_path))
    return output
}

type Vector2 struct {
	X float32
	Y float32
}
type Vector3 struct {
	X float32
	Y float32
	Z float32
}
type Vector4 struct {
	X float32
	Y float32
	Z float32
	W float32
}
type Matrix4x4 struct {
	M00 float32
	M10 float32
	M20 float32
	M30 float32

	M01 float32
	M11 float32
	M21 float32
	M31 float32

	M02 float32
	M12 float32
	M22 float32
	M32 float32

	M03 float32
	M13 float32
	M23 float32
	M33 float32
}

func (v Vector2) String() string {
	return fmt.Sprintf("Vector2(X: %.2f, Y: %.2f)", v.X, v.Y)
}
func (v Vector3) String() string {
	return fmt.Sprintf("Vector3(X: %.2f, Y: %.2f, Z: %.2f)", v.X, v.Y, v.Z)
}
func (v Vector4) String() string {
	return fmt.Sprintf("Vector4(X: %.2f, Y: %.2f, Z: %.2f, W: %.2f)", v.X, v.Y, v.Z, v.W)
}
func (m Matrix4x4) String() string {
	return fmt.Sprintf(
		"Matrix4x4:\n"+
			"%.2f %.2f %.2f %.2f\n"+
			"%.2f %.2f %.2f %.2f\n"+
			"%.2f %.2f %.2f %.2f\n"+
			"%.2f %.2f %.2f %.2f",
		m.M00, m.M01, m.M02, m.M03,
		m.M10, m.M11, m.M12, m.M13,
		m.M20, m.M21, m.M22, m.M23,
		m.M30, m.M31, m.M32, m.M33,
	)
}

func NewMatrix4x4(
	m00, m01, m02, m03,
	m10, m11, m12, m13,
	m20, m21, m22, m23,
	m30, m31, m32, m33 float32,
) Matrix4x4 {
	return Matrix4x4{
		M00: m00, M10: m10, M20: m20, M30: m30,
		M01: m01, M11: m11, M21: m21, M31: m31,
		M02: m02, M12: m12, M22: m22, M32: m32,
		M03: m03, M13: m13, M23: m23, M33: m33,
	}
}