package plugify

import (
	"reflect"
)

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
