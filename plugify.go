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
	C.Plugify_SetGetMethodPtr2(unsafe.Pointer(api[i])); i++
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
	M[4][4] float32
}