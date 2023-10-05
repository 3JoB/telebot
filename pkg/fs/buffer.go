package fs

import (
	"bytes"
	"errors"
	"io"
	"sync"
)

type BufferClose struct {
	bytes.Buffer
}

var (
	bufferClosePool = &sync.Pool{}

	ErrPtr = errors.New("the incoming pointer cannot be nil")
)

func NewBufferClose() *BufferClose {
	r := bufferClosePool.Get()
	if r == nil {
		return &BufferClose{}
	}
	return r.(*BufferClose)
}

func (b *BufferClose) Close() error {
	b.Reset()
	bufferClosePool.Put(b)
	return nil
}

type Buffer struct{}

func (Buffer) Create(string) (io.ReadWriteCloser, error) {
	return NewBufferClose(), nil
}
