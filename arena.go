package plugify

import "unsafe"

const (
	// defaultArenaSize is the default size of the stack buffer (4KB)
	defaultArenaSize = 4096
	// alignment for allocated memory (16 bytes for 64-bit systems)
	alignment = 16
)

func alignUp(val int) int {
	return (val + alignment - 1) &^ (alignment - 1)
}

type arena struct {
	buffer [defaultArenaSize]byte
	offset int
}

func (a *arena) Alloc(size int) unsafe.Pointer {
	aligned := alignUp(a.offset)
	newOffset := aligned + size

	if newOffset > len(a.buffer) {
		panicker("memory pool exhausted")
	}

	ptr := unsafe.Pointer(&a.buffer[aligned])
	a.offset = newOffset
	return ptr
}

func (a *arena) Reset() {
	a.offset = 0
}

func (a *arena) Size() int {
	return len(a.buffer)
}

func (a *arena) Remaining() int {
	return len(a.buffer) - a.offset
}

func (a *arena) Used() int {
	return a.offset
}
