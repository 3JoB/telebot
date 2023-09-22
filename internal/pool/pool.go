package pool

import (
	"bytes"
	"sync"
)

var bufferPool sync.Pool

type Buffer struct {
	*bytes.Buffer
}

func NewBuffer() *Buffer {
	p := bufferPool.Get()
	if p == nil {
		return &Buffer{Buffer: &bytes.Buffer{}}
	}
	return p.(*Buffer)
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
