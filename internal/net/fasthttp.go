package net

import (
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"

	"github.com/3JoB/telebot/json"
)

type FastHTTP struct {
	client *fasthttp.Client
	json   json.Json
}

func NewFastHTTPClient() NetFrame {
	f := &FastHTTP{
		client: &fasthttp.Client{
			NoDefaultUserAgentHeader:      true,
			DisableHeaderNamesNormalizing: false,
			Dial:                          fasthttpproxy.FasthttpProxyHTTPDialer(),
		},
	}
	f.json = json.NewGoJson()
	return f
}

func (f *FastHTTP) SetJsonProcessor(v json.Json) {
	f.json = v
}

func (f *FastHTTP) AcquireRequest() NetRequest {
	v := requestPool.Get()
	if v == nil {
		r := &FastHTTPRequest{}
		r.json = f.json
		r.client = f.client
		r.acquire()
		return r
	}
	r := v.(*FastHTTPRequest)
	r.client = f.client
	r.acquire()
	return r
}
