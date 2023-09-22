package pool

import (
	"bytes"
	"sync"
)

var bufferPool = sync.Pool{
	New: func() any {
		return &Buffer{Buffer: &bytes.Buffer{}}
	},
}

type Buffer struct {
	*bytes.Buffer
}

func NewBuffer() *Buffer {
	return bufferPool.Get().(*Buffer)
}

func ReleaseBuffer(b *Buffer) {
	if b == nil {
		return
	}
	b.Reset()
	bufferPool.Put(b)
}

func (b *Buffer) Close() error {
	ReleaseBuffer(b)
	return nil
}
