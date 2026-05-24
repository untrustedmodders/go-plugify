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
)

func init() {
	plugify.OnPluginStart(func() error {
		fmt.Println("OnPluginStart")
		return nil
	})

	plugify.OnPluginUpdate(func(dt float32) error {
		fmt.Println("OnPluginUpdate")
		return nil
	})

	plugify.OnPluginEnd(func() error {
		fmt.Println("OnPluginEnd")
		return nil
	})
}

func main() {}
```
