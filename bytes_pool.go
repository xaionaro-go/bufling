package bufling

import (
	"sync"
)

type BytesBuffer struct {
	locker sync.Mutex
	Buffer []byte
}

type BytesPool struct {
	cursor cursor
	bufs   []BytesBuffer
}

func NewBytesPool(maxParallel uint) *BytesPool {
	return &BytesPool{
		cursor: *newCursor(maxParallel),
		bufs:   make([]BytesBuffer, maxParallel),
	}
}

func (pool *BytesPool) Next() *BytesBuffer {
	buf := &pool.bufs[pool.cursor.Next()]
	buf.locker.Lock()
	return buf
}

func (buf *BytesBuffer) Unlock() {
	buf.Buffer = buf.Buffer[:0]
	buf.locker.Unlock()
}
