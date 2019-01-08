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

func (buf *BytesBuffer) Reset() {
	buf.Buffer = buf.Buffer[:0]
}

func (buf *BytesBuffer) Unlock() {
	buf.Reset()
	buf.locker.Unlock()
}

func (buf *BytesBuffer) Write(appendData []byte) (int, error) {
	buf.Buffer = append(buf.Buffer, appendData...) // see https://github.com/xaionaro-go/benchmarks/tree/master/append-vs-copy
	return len(appendData), nil
}
