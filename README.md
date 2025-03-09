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
package main

import (
	"fmt"
	"github.com/untrustedmodders/go-plugify"
	"runtime/debug"
)

func init() {
	plugify.OnPluginStart(func() {
		fmt.Println("OnPluginStart")
	})

	plugify.OnPluginUpdate(func(dt float32) {
		fmt.Println("OnPluginUpdate")
	})

	plugify.OnPluginEnd(func() {
		fmt.Println("OnPluginEnd")
	})

	plugify.OnPluginPanic(func() []byte {
		return debug.Stack() // workaround for could not import runtime/debug inside plugify package
	})
}

func main() {}
```