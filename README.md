# Plugify Language Package for Go

## Usage

### Initialize your module

```sh
go mod init example.com/my-go-plugin
```

### Get the go-plugify module

Note that you need to include the v in the version tag.

```sh
go get github.com/untrustedmodders/go-plugify@v1.0.0
```

```go
// go build -buildmode="c-shared" -ldflags="-X main.PluginName=example_plugin" ./
package main

import (
	"fmt"
	"github.com/untrustedmodders/go-plugify"
)

var plugin plugify.Plugin
var PluginName string // should match name in manifest

func OnPluginStart() error {
	fmt.Println("Go: OnPluginStart")
	return nil
}

func OnPluginUpdate(dt float32) error {
	fmt.Println("Go: OnPluginUpdate")
	return nil
}

func OnPluginEnd() error {
	fmt.Println("Go: OnPluginEnd")
	return nil
}

func init() {
	plugin = plugify.NewPlugin(PluginName, OnPluginStart, OnPluginUpdate, OnPluginEnd)
}

func main() {}
```
