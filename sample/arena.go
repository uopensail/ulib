package sample

import (
	"errors"
	"sync"
)

const (
	pageSize    uintptr = 4096    // 4 KiB – default OS page size
	maxPageSize uintptr = 1 << 30 // 1 GiB hard cap per allocation
	alignment   uintptr = 8       // all allocations are 8-byte aligned
)

var (
	// ErrInvalidSize is returned when the requested allocation size is zero or
	// exceeds maxPageSize.
	ErrInvalidSize = errors.New("invalid allocation size")
	// ErrOutOfMemory is returned when the arena cannot satisfy an allocation.
	ErrOutOfMemory = errors.New("out of memory")
)

// Arena is a thread-safe bump-pointer allocator that manages memory in
// fixed-size pages. It is designed to hold the binary representations used by
// ImmutableFeature, reducing GC pressure by keeping feature data outside the
// normal heap.
//
// Since Go 1.26 the Green Tea garbage collector greatly improves small-object
// GC throughput, but the Arena remains useful because it enables zero-copy
// string and slice access: strings and slices returned by ImmutableFeature
// point directly into arena pages rather than owning separate heap blocks.
//
// Allocation strategy:
//   - Requests ≤ pageSize are served from the current page. When the current
//     page is full a fresh page is appended and becomes the new current page.
//   - Requests > pageSize receive a dedicated page that is inserted just before
//     the current page so the current page always remains at the tail.
//
// All allocations are aligned to 8 bytes for safe use with unsafe.Pointer casts.
type Arena struct {
	mu    sync.RWMutex
	pages [][]byte // pages[len-1] is always the current small-allocation page
	cur   uintptr  // next free byte offset within the current page
}

// NewArena returns a new Arena pre-warmed with one initial page.
func NewArena() *Arena {
	a := &Arena{pages: make([][]byte, 0, 8)}
	a.pages = append(a.pages, make([]byte, pageSize))
	return a
}

// alignSize rounds size up to the next multiple of alignment (8 bytes).
func alignSize(size uintptr) uintptr {
	return (size + alignment - 1) &^ (alignment - 1)
}

// allocate reserves size bytes from the arena and returns the slice.
// The returned slice is valid for the lifetime of the Arena.
func (a *Arena) allocate(size uintptr) ([]byte, error) {
	if size == 0 || size > maxPageSize {
		return nil, ErrInvalidSize
	}
	size = alignSize(size)

	a.mu.Lock()
	defer a.mu.Unlock()

	n := len(a.pages)

	// Large allocations get a dedicated page inserted before the current page.
	if size > pageSize {
		large := make([]byte, size)
		a.pages = append(a.pages, large)
		// Swap new page with current page so current stays at the tail.
		a.pages[n], a.pages[n-1] = a.pages[n-1], a.pages[n]
		return a.pages[n-1], nil
	}

	// Ensure the cursor is aligned before measuring remaining space.
	a.cur = alignSize(a.cur)
	if pageSize-a.cur >= size {
		// Fits in the current page.
		data := a.pages[n-1][a.cur : a.cur+size]
		a.cur += size
		return data, nil
	}

	// Current page is full — append a fresh one.
	a.pages = append(a.pages, make([]byte, pageSize))
	a.cur = size
	return a.pages[n][:size], nil // a.pages[n] is the newly appended page
}

// Size returns the total number of bytes across all allocated pages.
func (a *Arena) Size() uintptr {
	a.mu.RLock()
	defer a.mu.RUnlock()
	var total uintptr
	for _, p := range a.pages {
		total += uintptr(len(p))
	}
	return total
}

// PageCount returns the current number of pages held by the arena.
func (a *Arena) PageCount() int {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return len(a.pages)
}
