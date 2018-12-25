package bufling

import (
	"testing"
)

func BenchmarkParallelBytesNext(b *testing.B) {
	bufs := NewBytesPool(1024)
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bufs.Next().Unlock()
		}
	})
}

func BenchmarkBytesAppend(b *testing.B) {
	bufs := NewBytesPool(2)

	appendingData := []byte(`test string`)
	tryAppend := func() {
		buf := bufs.Next()
		buf.Buffer = append(buf.Buffer, appendingData...)
		buf.Unlock()
	}

	tryAppend()
	tryAppend()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tryAppend()
	}
}
