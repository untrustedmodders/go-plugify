package plugify

/*
#include "plugify.h"
*/
import "C"
import (
	"context"
	"runtime"
	"runtime/debug"
	"sync"
	"unsafe"
)

type PluginInfo interface {
	Info() *debug.BuildInfo

	ID() int
	Name() string
	Description() string
	Version() string
	Author() string
	Website() string
	License() string
	Location() string
	Dependencies() []string

	OnInit(handle unsafe.Pointer)
	OnStart() error
	OnUpdate(dt float32) error
	OnEnd() error
	OnShutdown()

	Call(method unsafe.Pointer, data unsafe.Pointer, params unsafe.Pointer, count int, ret unsafe.Pointer)

	Starting() bool
	Updating() bool
	Ending() bool
	Loaded() bool
	Context() context.Context
	Handle() unsafe.Pointer
}

type PluginStart func() error
type PluginUpdate func(dt float32) error
type PluginEnd func() error

type pluginInfo struct {
	info   *debug.BuildInfo
	start  PluginStart
	update PluginUpdate
	end    PluginEnd

	loaded bool
	once   sync.Once
	ctx    context.Context
	cancel context.CancelFunc
	handle unsafe.Pointer

	id           int
	name         string
	description  string
	version      string
	author       string
	website      string
	license      string
	location     string
	dependencies []string
}

var plugins = []pluginInfo{}

func (p *pluginInfo) Info() *debug.BuildInfo {
	return p.info
}

func (p *pluginInfo) ID() int {
	return p.id
}

func (p *pluginInfo) Name() string {
	return p.name
}

func (p *pluginInfo) Description() string {
	return p.description
}

func (p *pluginInfo) Version() string {
	return p.version
}

func (p *pluginInfo) Author() string {
	return p.author
}

func (p *pluginInfo) Website() string {
	return p.website
}

func (p *pluginInfo) License() string {
	return p.license
}

func (p *pluginInfo) Location() string {
	return p.location
}

func (p *pluginInfo) Dependencies() []string {
	return p.dependencies
}

func (p *pluginInfo) OnInit(handle unsafe.Pointer) {
	if handle != nil {
		p.handle = handle
	} else {
		p.handle = unsafe.Pointer(C.Plugify_GetPlugin(p.info.Main.Path))
	}
	h := C.PluginHandle(p.handle)

	p.id = int(C.Plugify_GetPluginId(h))

	nameStr := C.Plugify_GetPluginName(h)
	p.name = GetStringData[string](&nameStr)
	C.Plugify_DestroyString(&nameStr)

	descriptionStr := C.Plugify_GetPluginDescription(h)
	p.description = GetStringData[string](&descriptionStr)
	C.Plugify_DestroyString(&descriptionStr)

	versionsStr := C.Plugify_GetPluginVersion(h)
	p.version = GetStringData[string](&versionsStr)
	C.Plugify_DestroyString(&versionsStr)

	authorStr := C.Plugify_GetPluginAuthor(h)
	p.author = GetStringData[string](&authorStr)
	C.Plugify_DestroyString(&authorStr)

	websiteStr := C.Plugify_GetPluginWebsite(h)
	p.website = GetStringData[string](&websiteStr)
	C.Plugify_DestroyString(&websiteStr)

	licenseStr := C.Plugify_GetPluginLicense(h)
	p.license = GetStringData[string](&licenseStr)
	C.Plugify_DestroyString(&licenseStr)

	locationStr := C.Plugify_GetPluginLocation(h)
	p.location = GetStringData[string](&locationStr)
	C.Plugify_DestroyString(&locationStr)

	dependenciesStr := C.Plugify_GetPluginDependencies(h)
	p.dependencies = GetVectorDataString[string](&dependenciesStr)
	C.Plugify_DestroyVectorString(&dependenciesStr)
}

func (p *pluginInfo) OnStart() error {
	p.loaded = true
	p.ctx, p.cancel = context.WithCancel(context.Background())

	if p.start != nil {
		return p.start()
	}
	return nil
}

func (p *pluginInfo) OnUpdate(dt float32) error {
	if p.update != nil {
		return p.update(dt)
	}
	return nil
}

func (p *pluginInfo) OnEnd() error {
	defer func() {
		p.cancel()
		p.loaded = false
	}()
	if p.end != nil {
		return p.end()
	}
	return nil
}

func (p *pluginInfo) OnShutdown() {
	clear(functionMap)

	for _, v := range calls {
		C.Plugify_DeleteCall(v)
	}
	clear(calls)

	for _, v := range callbacks {
		C.Plugify_DeleteCallback(v)
	}
	clear(callbacks)

	runtime.GC()
	runtime.Gosched()
}

func (p *pluginInfo) Call(method unsafe.Pointer, data unsafe.Pointer, params unsafe.Pointer, count int, ret unsafe.Pointer) {
	internalCall(C.MethodHandle(method), data, (*C.Parameters)(params), C.size_t(count), (*C.Return)(ret))
}

func (p *pluginInfo) Starting() bool {
	return p.start != nil
}

func (p *pluginInfo) Updating() bool {
	return p.update != nil
}

func (p *pluginInfo) Ending() bool {
	return p.end != nil
}

func (p *pluginInfo) Loaded() bool {
	return p.loaded
}

func (p *pluginInfo) Context() context.Context {
	return p.ctx
}

func (p *pluginInfo) Handle() unsafe.Pointer {
	return p.handle
}

func NewPlugin(
	info *debug.BuildInfo,
	start PluginStart,
	update PluginUpdate,
	end PluginEnd,
) PluginInfo {

	plugins = append(plugins, pluginInfo{
		info:   info,
		start:  start,
		update: update,
		end:    end,
	})

	return &plugins[len(plugins)-1]
}
