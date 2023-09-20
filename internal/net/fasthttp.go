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

func (f *FastHTTP) SetJsonHandle(v json.Json) {
	f.json = v
}

func (f *FastHTTP) AcquireRequest() NetRequest {
	var r *FastHTTPRequest

	if v := requestPool.Get(); v == nil {
		r = &FastHTTPRequest{}
		r.json = f.json
	} else {
		r = v.(*FastHTTPRequest)
	}

	r.client = f.client
	r.acquire()
	return r
}
