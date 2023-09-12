package bPool

import (
	"bytes"
	"sync"
)

var fPool sync.Pool

type bP struct {
	*bytes.Buffer
}

func New() *bP {
	p := fPool.Get()
	if p == nil {
		return &bP{Buffer: &bytes.Buffer{}}
	}
	return p.(*bP)
}

func Put(b *bP) {
	if b == nil {
		return
	}
	b.Reset()
	fPool.Put(b)
}

func (b *bP) Close() error {
	Put(b)
	return nil
}
