package bufling

import (
	"sync/atomic"
)

type cursor struct {
	maxParallel uint64
	value uint64
}

func (c *cursor) Next() uint64 {
	value := atomic.AddUint64(&c.value, 1)
	if value >= c.maxParallel {
		atomic.StoreUint64(&c.value, 0)
		value = 0
	}

	return value
}

func newCursor(maxParallel uint) *cursor {
	if maxParallel == 0 {
		panic(`maxParallel cannot be 0`)
	}
	return &cursor{
		maxParallel: uint64(maxParallel),
	}
}
