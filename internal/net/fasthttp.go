package net

import (
	"github.com/3JoB/telebot/internal/json"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

type FastHTTP struct {
	client *fasthttp.Client
	json json.Json
}

func NewFastHTTPClient() {
	_ = &FastHTTP{
		client: &fasthttp.Client{
			NoDefaultUserAgentHeader:      true,
			DisableHeaderNamesNormalizing: false,
			Dial:                          fasthttpproxy.FasthttpProxyHTTPDialer(),
		},
	}
}

func (f *FastHTTP) SetJsonProcessor(v json.Json) {
	f.json = v
}

func (f *FastHTTP) Acquire() *FastHTTPRequest {
	var r *FastHTTPRequest
	v := requestPool.Get()
	if v == nil {
		r = &FastHTTPRequest{}
	} else {
		r = v.(*FastHTTPRequest)
	}
	if f.json == nil {
		r.json = json.NewGoJson()
	} else {
		r.json = f.json
	}
	return r
}
