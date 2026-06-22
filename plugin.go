package plugify

/*
#include "plugify.h"
*/
import "C"
import (
	"context"
	"unsafe"
)

type Plugin interface {
	ID() int
	Name() string
	Description() string
	Version() string
	Author() string
	Website() string
	License() string
	Location() string
	Dependencies() []string

	onInit(name string)
	onStart() error
	onUpdate(dt float32) error
	onEnd() error

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
	start  PluginStart
	update PluginUpdate
	end    PluginEnd

	loaded bool
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

var pluginsMap = make(map[string]*plugin)

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

func (p *plugin) onInit(name string) {
	h := C.Plugify_GetPlugin(name)

	p.handle = unsafe.Pointer(h)
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

func (p *plugin) onStart() error {
	p.loaded = true
	p.ctx, p.cancel = context.WithCancel(context.Background())

	if p.start != nil {
		return p.start()
	}
	return nil
}

func (p *plugin) onUpdate(dt float32) error {
	if p.update != nil {
		return p.update(dt)
	}
	return nil
}

func (p *plugin) onEnd() error {
	defer func() {
		p.cancel()
		p.loaded = false
	}()
	if p.end != nil {
		return p.end()
	}
	return nil
}

func (p *plugin) onShutdown() {
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
	name string,
	start PluginStart,
	update PluginUpdate,
	end PluginEnd,
) Plugin {
	p := &plugin{
		start:  start,
		update: update,
		end:    end,
	}
	pluginsMap[name] = p
	return p
}

func plg() *plugin {
	for _, p := range pluginsMap {
		return p
	}
	panic("expected to have a registered plugin")
}
