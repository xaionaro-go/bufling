package bufling

import (
	"sync"
)

type Resetter interface {
	Reset()
}

type AnyBuffer struct {
	sync.Mutex
	Buffer Resetter
}

type Any struct {
	cursor cursor
	bufs   []AnyBuffer
}

func NewAny(maxParallel uint) *Any {
	return &Any{
		cursor: *newCursor(maxParallel),
		bufs:   make([]AnyBuffer, maxParallel),
	}
}

func (b *Any) Next() *AnyBuffer {
	curIdx := b.cursor.Next()
	buf := &b.bufs[curIdx]
	buf.Lock()
	buf.Buffer.Reset()
	return buf
}
