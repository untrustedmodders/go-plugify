package plugify

import (
	"reflect"
)

// Grow or shrink the slice to the given length.
func sliceSize(slice reflect.Value, size int) {
	sliceCap := slice.Cap()

	if size > sliceCap {
		slice.Grow(size - sliceCap)
	}

	slice.SetLen(size)
}
