package sample

import (
	"errors"
	"sync"
)

const (
	pageSize    uintptr = 4096    // 4KB page size for memory allocation
	maxPageSize uintptr = 1 << 30 // 1GB maximum allocation size
	alignment   uintptr = 8       // 8-byte memory alignment
)

var (
	ErrInvalidSize = errors.New("invalid allocation size")
	ErrOutOfMemory = errors.New("out of memory")
)

/**
 * @brief Arena is a memory pool allocator that manages memory in fixed-size pages
 *
 * The Arena allocator provides efficient memory allocation by pre-allocating
 * large chunks of memory (pages) and serving allocation requests from these pages.
 * It supports both small allocations (served from current page) and large
 * allocations (dedicated pages).
 *
 * Thread Safety: The Arena is thread-safe using read-write mutex protection.
 *
 * Memory Layout:
 * - Small allocations: Allocated sequentially from current page
 * - Large allocations: Get dedicated pages
 * - All allocations are 8-byte aligned for optimal performance
 */
type Arena struct {
	mu    sync.RWMutex // Read-write mutex for thread safety
	pages [][]byte     // Collection of allocated memory pages
	cur   uintptr      // Current offset in the last page
}

/**
 * @brief Creates a new Arena allocator with an initial page
 *
 * @return Pointer to the newly created Arena
 *
 * The Arena is initialized with:
 * - Empty pages slice with capacity for 8 pages
 * - One initial page of pageSize (4KB)
 * - Current offset set to 0
 */
func NewArena() *Arena {
	arena := &Arena{
		pages: make([][]byte, 0, 8),
		cur:   0,
	}
	arena.pages = append(arena.pages, make([]byte, pageSize))
	return arena
}

/**
 * @brief Aligns size to 8-byte boundary for optimal memory access
 *
 * @param size The size to align
 * @return The aligned size (rounded up to next 8-byte boundary)
 *
 * Uses bit manipulation for efficient alignment:
 * (size + alignment - 1) & ^(alignment - 1)
 */
func alignSize(size uintptr) uintptr {
	return (size + alignment - 1) &^ (alignment - 1)
}

/**
 * @brief Allocates aligned memory from the arena
 *
 * @param size Number of bytes to allocate
 * @return Slice of allocated bytes and error if allocation fails
 *
 * Allocation Strategy:
 * 1. Validate size (0 < size <= maxPageSize)
 * 2. Align size to 8-byte boundary
 * 3. For large allocations (> pageSize): create dedicated page
 * 4. For small allocations: try current page, create new page if needed
 *
 * Thread Safety: Protected by write lock during allocation
 */
func (arena *Arena) allocate(size uintptr) ([]byte, error) {
	if size == 0 || size > maxPageSize {
		return nil, ErrInvalidSize
	}

	size = alignSize(size)

	arena.mu.Lock()
	defer arena.mu.Unlock()

	n := len(arena.pages)

	// For large allocations, create a dedicated page
	if size > pageSize {
		newPage := make([]byte, size)
		arena.pages = append(arena.pages, newPage)
		// Move the new page to second-to-last position to keep current page last
		arena.pages[n], arena.pages[n-1] = arena.pages[n-1], arena.pages[n]
		return arena.pages[n-1], nil
	}

	// Ensure current position is aligned
	arena.cur = alignSize(arena.cur)
	remain := pageSize - arena.cur

	if remain >= size {
		// Allocate from current page
		data := arena.pages[n-1][arena.cur : arena.cur+size]
		arena.cur += size
		return data, nil
	}

	// Need a new page
	arena.pages = append(arena.pages, make([]byte, pageSize))
	arena.cur = size
	return arena.pages[n][:size], nil
}

/**
 * @brief Returns total allocated memory size across all pages
 *
 * @return Total memory size in bytes
 *
 * Thread Safety: Protected by read lock
 */
func (arena *Arena) Size() uintptr {
	arena.mu.RLock()
	defer arena.mu.RUnlock()

	var total uintptr
	for _, page := range arena.pages {
		total += uintptr(len(page))
	}
	return total
}

/**
 * @brief Returns the number of pages currently allocated
 *
 * @return Number of pages
 *
 * Thread Safety: Protected by read lock
 */
func (arena *Arena) PageCount() int {
	arena.mu.RLock()
	defer arena.mu.RUnlock()
	return len(arena.pages)
}
