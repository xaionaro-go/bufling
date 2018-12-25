package bufling

import (
	"sync"
)

type BytesBuffer struct {
	sync.Mutex
	Buffer []byte
}

type Bytes struct {
	cursor cursor
	bufs   []BytesBuffer
}

func NewBytes(maxParallel uint) *Bytes {
	return &Bytes{
		cursor: *newCursor(maxParallel),
		bufs:   make([]BytesBuffer, maxParallel),
	}
}

func (b *Bytes) Next() *BytesBuffer {
	curIdx := b.cursor.Next()
	buf := &b.bufs[curIdx]
	buf.Lock()
	buf.Buffer = buf.Buffer[:0]
	return buf
}
