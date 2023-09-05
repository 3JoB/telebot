package telebot

import "github.com/valyala/fasthttp"

// Acquire returns an empty fasthttp instance from request pool.
//
// The returned fasthttp instance may be passed to Release when it is no longer needed.
// This allows Request recycling, reduces GC pressure and usually improves performance.
func acquire() (req *fasthttp.Request, resp *fasthttp.Response) {
	req, resp = fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	return
}

// Release returns req and resp acquired via Acquire to request pool.
//
// It is forbidden accessing req and/or its' members after returning it to request pool.
func release(req *fasthttp.Request, res *fasthttp.Response) {
	fasthttp.ReleaseRequest(req)
	fasthttp.ReleaseResponse(res)
}
