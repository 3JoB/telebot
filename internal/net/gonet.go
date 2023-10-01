package net

import (
	"os"

	"github.com/3JoB/resty-ilo"

	"github.com/3JoB/telebot/json"
)

type GoNet struct {
	client *resty.Client
	json   json.Json
}

func NewHTTPClient() NetFrame {
	g := &GoNet{
		client: resty.New(),
	}
	g.lookProxyEnv()
	return g
}

func (g *GoNet) SetJsonHandle(v json.Json) {
	g.client.JSONMarshal = v.Marshal
	g.client.JSONUnmarshal = v.Unmarshal
	g.json = v
}

func (g *GoNet) AcquireRequest() NetRequest {
	var r *GoNetRequest
	if v := requestPool.Get(); v == nil {
		r = &GoNetRequest{}
		r.json = g.json
	} else {
		r = v.(*GoNetRequest)
	}
	r.r = g.client.R()
	return r
}

func (g *GoNet) lookProxyEnv() {
	if http_proxy, ok := os.LookupEnv("http_proxy"); ok {
		g.client = g.client.SetProxy(http_proxy)
	}
}
