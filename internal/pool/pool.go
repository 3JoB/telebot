package pool

import (
	"bytes"
	"sync"

	"github.com/colega/zeropool"
)

var (
	bufferPool sync.Pool
	mapPool    zeropool.Pool[map[string]any]
)

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

func NewMapper() map[string]any {
	p := mapPool.Get()
	if p == nil {
		return make(map[string]any)
	}
	return p
}

func ReleaseMapper(m map[string]any) {
	if m == nil {
		return
	}
	clear(m)
	mapPool.Put(m)
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
