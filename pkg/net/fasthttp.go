package net

import (
	"sync"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"

	"github.com/3JoB/telebot/v2/pkg/json"
)

type FastHTTP struct {
	client       *fasthttp.Client
	json         json.Json
	requestPool  *sync.Pool
	responsePool *sync.Pool
}

func NewFastHTTPClient() NetFrame {
	f := &FastHTTP{
		client: &fasthttp.Client{
			NoDefaultUserAgentHeader:      true,
			DisableHeaderNamesNormalizing: false,
			Dial:                          fasthttpproxy.FasthttpProxyHTTPDialer(),
		},
		requestPool:  &sync.Pool{},
		responsePool: &sync.Pool{},
	}
	return f
}

func (f *FastHTTP) SetJsonHandle(v json.Json) {
	f.json = v
}

func (f *FastHTTP) Acquire() (NetRequest, NetResponse) {
	var r *FastHTTPRequest

	if v := f.requestPool.Get(); v == nil {
		r = &FastHTTPRequest{}
		r.json = f.json
	} else {
		r = v.(*FastHTTPRequest)
	}

	r.resp = f.acquireResponse()
	r.client = f.client
	r.acquire()
	return r, r.resp
}

func (f *FastHTTP) acquireResponse() *Response {
	r := f.responsePool.Get()
	if r == nil {
		return &Response{}
	}
	return r.(*Response)
}

func (f *FastHTTP) Release(req NetRequest, resp NetResponse) {
	f.ReleaseRequest(req)
	f.ReleaseResponse(resp)
}

func (f *FastHTTP) ReleaseRequest(r NetRequest) {
	v, ok := r.(*FastHTTPRequest)
	if !ok {
		return
	}
	v.Reset()
	f.requestPool.Put(v)
}

func (f *FastHTTP) ReleaseResponse(r NetResponse) {
	v, ok := r.(*Response)
	if !ok {
		return
	}
	v.Reset()
	f.responsePool.Put(v)
}
