package plugify

type PluginInfo interface {
	ID() int64
	Name() string
	Description() string
	Version() string
	Author() string
	Website() string
	License() string
	Location() string
	Dependencies() []string

	Loaded() bool
}

type PluginStartCallback func() error
type PluginUpdateCallback func(dt float32) error
type PluginEndCallback func() error
type PluginPanicCallback func() []byte

type pluginInfo struct {
	id           int64
	name         string
	description  string
	version      string
	author       string
	website      string
	license      string
	location     string
	dependencies []string

	fnPluginStartCallback  PluginStartCallback
	fnPluginUpdateCallback PluginUpdateCallback
	fnPluginEndCallback    PluginEndCallback

	loaded bool
}

var plugin = pluginInfo{
	id:           -1,
	name:         "",
	description:  "",
	version:      "",
	author:       "",
	website:      "",
	license:      "",
	dependencies: []string{},

	fnPluginStartCallback:  nil,
	fnPluginUpdateCallback: nil,
	fnPluginEndCallback:    nil,

	loaded: false,
}

func (p *pluginInfo) ID() int64 {
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

func (p *pluginInfo) Loaded() bool {
	return p.loaded
}

func Plugin() PluginInfo {
	return &plugin
}

func OnPluginStart(fn PluginStartCallback) {
	plugin.fnPluginStartCallback = fn
}

func OnPluginUpdate(fn PluginUpdateCallback) {
	plugin.fnPluginUpdateCallback = fn
}

func OnPluginEnd(fn PluginEndCallback) {
	plugin.fnPluginEndCallback = fn
}
