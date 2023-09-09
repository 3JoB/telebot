package net

import (
	"github.com/cornelk/hashmap"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpproxy"
)

type FastHTTP struct {
	client *fasthttp.Client
	header *Header
}

func NewFastHTTPClient() {
	e := &FastHTTP{
		client: &fasthttp.Client{
			NoDefaultUserAgentHeader:      true,
			DisableHeaderNamesNormalizing: false,
			Dial:                          fasthttpproxy.FasthttpProxyHTTPDialer(),
		},
		header: &Header{
			m: hashmap.New[string, string](),
		},
	}
	e.header.m.Set("User-Agent", "Mozilla/5.0(compatible; Telebot-Expansion-Pack/v1; +https://github.com/3JoB/telebot)")
}

func (f *FastHTTP) Header() *hashmap.Map[string, string] {
	return f.header.m
}

func (f *FastHTTP) Acquire() {}
