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

type Plugin interface {
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

type plugin struct {
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

var plugins []plugin

func (p *plugin) Info() *debug.BuildInfo {
	return p.info
}

func (p *plugin) ID() int {
	return p.id
}

func (p *plugin) Name() string {
	return p.name
}

func (p *plugin) Description() string {
	return p.description
}

func (p *plugin) Version() string {
	return p.version
}

func (p *plugin) Author() string {
	return p.author
}

func (p *plugin) Website() string {
	return p.website
}

func (p *plugin) License() string {
	return p.license
}

func (p *plugin) Location() string {
	return p.location
}

func (p *plugin) Dependencies() []string {
	return p.dependencies
}

func (p *plugin) OnInit(handle unsafe.Pointer) {
	p.handle = handle

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

func (p *plugin) OnStart() error {
	p.loaded = true
	p.ctx, p.cancel = context.WithCancel(context.Background())

	if p.start != nil {
		return p.start()
	}
	return nil
}

func (p *plugin) OnUpdate(dt float32) error {
	if p.update != nil {
		return p.update(dt)
	}
	return nil
}

func (p *plugin) OnEnd() error {
	defer func() {
		p.cancel()
		p.loaded = false
	}()
	if p.end != nil {
		return p.end()
	}
	return nil
}

func (p *plugin) OnShutdown() {
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

func (p *plugin) Call(method unsafe.Pointer, data unsafe.Pointer, params unsafe.Pointer, count int, ret unsafe.Pointer) {
	internalCall(C.MethodHandle(method), data, (*C.Parameters)(params), C.size_t(count), (*C.Return)(ret))
}

func (p *plugin) Starting() bool {
	return p.start != nil
}

func (p *plugin) Updating() bool {
	return p.update != nil
}

func (p *plugin) Ending() bool {
	return p.end != nil
}

func (p *plugin) Loaded() bool {
	return p.loaded
}

func (p *plugin) Context() context.Context {
	return p.ctx
}

func (p *plugin) Handle() unsafe.Pointer {
	return p.handle
}

func NewPlugin(
	info *debug.BuildInfo,
	start PluginStart,
	update PluginUpdate,
	end PluginEnd,
) Plugin {

	plugins = append(plugins, plugin{
		info:   info,
		start:  start,
		update: update,
		end:    end,
	})

	return &plugins[len(plugins)-1]
}
