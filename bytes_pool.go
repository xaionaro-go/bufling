package bufling

import (
	"sync"
)

type BytesBuffer struct {
	sync.Mutex
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
	buf.Lock()
	buf.Buffer = buf.Buffer[:0]
	return buf
}
