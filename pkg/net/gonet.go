package net

import (
	"os"
	"sync"

	"github.com/3JoB/resty-ilo"

	"github.com/3JoB/telebot/v2/pkg/json"
)

type GoNet struct {
	client       *resty.Client
	json         json.Json
	requestPool  *sync.Pool
	responsePool *sync.Pool
}

func NewHTTPClient() NetFrame {
	g := &GoNet{
		client:       resty.New(),
		requestPool:  &sync.Pool{},
		responsePool: &sync.Pool{},
	}
	g.lookProxyEnv()
	return g
}

func (g *GoNet) SetJsonHandle(v json.Json) {
	g.client.JSONMarshal = v.Marshal
	g.client.JSONUnmarshal = v.Unmarshal
	g.json = v
}

func (g *GoNet) Acquire() (NetRequest, NetResponse) {
	var r *GoNetRequest
	if v := requestPool.Get(); v == nil {
		r = &GoNetRequest{}
		r.json = g.json
	} else {
		r = v.(*GoNetRequest)
	}
	r.r = g.client.R()
	r.resp = g.acquireResponse()
	return r, r.resp
}

func (g *GoNet) acquireResponse() *Response {
	v := g.responsePool.Get()
	if v == nil {
		return &Response{}
	}
	return v.(*Response)
}

func (g *GoNet) Release(req NetRequest, resp NetResponse) {
	g.ReleaseRequest(req)
	g.ReleaseResponse(resp)
}

func (g *GoNet) ReleaseRequest(r NetRequest) {
	v, ok := r.(*GoNetRequest)
	if !ok {
		return
	}
	v.Reset()
	g.requestPool.Put(v)
}

func (g *GoNet) ReleaseResponse(r NetResponse) {
	v, ok := r.(*Response)
	if !ok {
		return
	}
	v.Reset()
	g.responsePool.Put(v)
}

func (g *GoNet) lookProxyEnv() {
	if http_proxy, ok := os.LookupEnv("http_proxy"); ok {
		g.client = g.client.SetProxy(http_proxy)
	}
}
