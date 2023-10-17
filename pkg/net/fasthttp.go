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
		requestPool: &sync.Pool{},
		responsePool: &sync.Pool{},
	}
	return f
}

func (f *FastHTTP) SetJsonHandle(v json.Json) {
	f.json = v
}

func (f *FastHTTP) AcquireRequest() NetRequest {
	var r *FastHTTPRequest

	if v := f.requestPool.Get(); v == nil {
		r = &FastHTTPRequest{}
		r.json = f.json
	} else {
		r = v.(*FastHTTPRequest)
	}

	r.client = f.client
	r.responsePool = f.responsePool
	r.acquire()
	return r
}

func (f *FastHTTP) ReleaseRequest(r NetRequest) {
	r.Reset()
	f.requestPool.Put(r)
}

func (f *FastHTTP) ReleaseResponse(r NetResponse) {
	r.Reset()
	f.responsePool.Put(r)
}