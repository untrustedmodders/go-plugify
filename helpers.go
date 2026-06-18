package plugify

import (
	"fmt"
	"reflect"
	"runtime/debug"
)

var buildInfo, _ = debug.ReadBuildInfo()

// Grow or shrink the slice to the given length.
func sliceSize(slice reflect.Value, size int) {
	sliceLen := slice.Len() // TODO: should use cap

	if sliceLen < size {
		reqLen := size - sliceLen

		slice.Grow(reqLen)
		slice.SetLen(size)
	} else if sliceLen > size {
		slice.SetLen(size)
	}
}

func stacktrace(err any) {
	msg := fmt.Sprintf("%v", err)
	stack := debug.Stack()
	if len(stack) > 0 {
		msg += fmt.Sprintf("\nStack Trace: \n%s", stack)
	}
	Log(msg, Error, buildInfo, 3)
}

func panicker(err any) {
	stacktrace(err)
	panic(err)
}
