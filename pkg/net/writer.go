package net

import (
	"io"
	"sync"
)

// form fasthttp
func Copy(w io.Writer, r io.Reader) (int64, error) {
	vbuf := cbp.Get()
	buf := vbuf.([]byte)
	n, err := io.CopyBuffer(w, r, buf)
	cbp.Put(vbuf)
	return n, err
}

var cbp = sync.Pool{
	New: func() any {
		return make([]byte, 4096)
	},
}
