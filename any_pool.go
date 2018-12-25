package bufling

import (
	"sync"
)

type Resetter interface {
	Reset()
}

type AnyBuffer struct {
	locker sync.Mutex
	Buffer Resetter
}

type AnyPool struct {
	cursor cursor
	bufs   []AnyBuffer
}

func NewAnyPool(maxParallel uint, initFunc func(*AnyBuffer)) *AnyPool {
	pool := &AnyPool{
		cursor: *newCursor(maxParallel),
		bufs:   make([]AnyBuffer, maxParallel),
	}
	if initFunc != nil {
		for idx, _ := range pool.bufs {
			initFunc(&pool.bufs[idx])
		}
	}
	return pool
}

func (pool *AnyPool) Next() *AnyBuffer {
	buf := &pool.bufs[pool.cursor.Next()]
	buf.locker.Lock()
	return buf
}

func (buf *AnyBuffer) Unlock() {
	buf.Buffer.Reset()
	buf.locker.Unlock()
}
