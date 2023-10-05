package net

import (
	"io"
	"sync"
)

/*import (
	"bytes"
	"errors"
	"io"
	"os"
	"sync"

	"github.com/3JoB/resty-ilo"
	"github.com/3JoB/ulib/pool"
	"github.com/valyala/fasthttp"
)

var ErrNoType error = errors.New("unknown input type")

func Writer(r any, w io.ReadWriter) error {
	switch k := r.(type) {
	case *fasthttp.Response:
		return writer_fasthttp(k, w)
	case *resty.Response:
		return writer_resty(k, w)
	default:
		return ErrNoType
	}
}

func writer_fasthttp(r *fasthttp.Response, w io.ReadWriter) error {
	switch p := w.(type) {
	case *os.File:

	case *bytes.Buffer:

	case *pool.BufferClose:
	default:
		return ErrNoType
	}
}

func writer_resty(r *resty.Response, w io.ReadWriter) error {

}

*/

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
