package plugify

/*
#include "plugify.h"
*/
import "C"
import (
	"context"
	"runtime/debug"
	"sync"
)

type PluginInfo interface {
	ID() int
	Name() string
	Description() string
	Version() string
	Author() string
	Website() string
	License() string
	Location() string
	Dependencies() []string

	Starting() bool
	Updating() bool
	Ending() bool
	Loaded() bool
	Context() context.Context
}

type PluginStart func() error
type PluginUpdate func(dt float32) error
type PluginEnd func() error

type pluginInfo struct {
	info *debug.BuildInfo

	start  PluginStart
	update PluginUpdate
	end    PluginEnd

	loaded bool
	once   sync.Once
	ctx    context.Context
	cancel context.CancelFunc

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

var pluginMap = make(map[string]*pluginInfo)

func (p *pluginInfo) init() {
	p.once.Do(func() {
		handle := C.Plugify_GetPlugin(p.info.Main.Path)

		p.id = int(C.Plugify_GetPluginId(handle))

		nameStr := C.Plugify_GetPluginName(handle)
		p.name = GetStringData[string](&nameStr)
		C.Plugify_DestroyString(&nameStr)

		descriptionStr := C.Plugify_GetPluginDescription(handle)
		p.description = GetStringData[string](&descriptionStr)
		C.Plugify_DestroyString(&descriptionStr)

		versionsStr := C.Plugify_GetPluginVersion(handle)
		p.version = GetStringData[string](&versionsStr)
		C.Plugify_DestroyString(&versionsStr)

		authorStr := C.Plugify_GetPluginAuthor(handle)
		p.author = GetStringData[string](&authorStr)
		C.Plugify_DestroyString(&authorStr)

		websiteStr := C.Plugify_GetPluginWebsite(handle)
		p.website = GetStringData[string](&websiteStr)
		C.Plugify_DestroyString(&websiteStr)

		licenseStr := C.Plugify_GetPluginLicense(handle)
		p.license = GetStringData[string](&licenseStr)
		C.Plugify_DestroyString(&licenseStr)

		locationStr := C.Plugify_GetPluginLocation(handle)
		p.location = GetStringData[string](&locationStr)
		C.Plugify_DestroyString(&locationStr)

		dependenciesStr := C.Plugify_GetPluginDependencies(handle)
		p.dependencies = GetVectorDataString[string](&dependenciesStr)
		C.Plugify_DestroyVectorString(&dependenciesStr)
	})
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

func NewPlugin(
	info *debug.BuildInfo,
	start PluginStart,
	update PluginUpdate,
	end PluginEnd,
) PluginInfo {

	p := &pluginInfo{
		info:   info,
		start:  start,
		update: update,
		end:    end,
	}

	pluginMap[info.Main.Path] = p

	return p
}
