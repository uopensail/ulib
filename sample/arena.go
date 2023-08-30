package sample

import (
	"sync"
)

const pageSize uintptr = 32768

type Arena struct {
	sync.RWMutex
	pages [][]byte
	start uintptr
}

func NewArena() *Arena {
	arena := &Arena{
		pages: make([][]byte, 0, 8),
		start: 0,
	}
	arena.pages = append(arena.pages, make([]byte, pageSize))
	return arena
}

func (arena *Arena) allocate(size uintptr) []byte {
	arena.Lock()
	defer arena.Unlock()
	n := len(arena.pages)
	if size > pageSize {
		arena.pages = append(arena.pages, make([]byte, size))
		arena.pages[n], arena.pages[n-1] = arena.pages[n-1], arena.pages[n]
		return arena.pages[n-1]
	}
	remain := pageSize - arena.start
	if remain >= size {
		data := arena.pages[n-1][arena.start : arena.start+size]
		arena.start += size
		return data
	}
	arena.pages = append(arena.pages, make([]byte, pageSize))
	arena.start = size
	return arena.pages[n][:size]
}
