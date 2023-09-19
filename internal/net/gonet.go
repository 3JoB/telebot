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
	http_proxy, ok := os.LookupEnv("http_proxy")
	if ok {
		g.client = g.client.SetProxy(http_proxy)
	}
	g.json = json.NewGoJson()
	g.client.JSONMarshal = g.json.Marshal
	g.client.JSONUnmarshal = g.json.Unmarshal
	return g
}

func (g *GoNet) SetJsonHandle(v json.Json) {
	g.client.JSONMarshal = v.Marshal
	g.client.JSONUnmarshal = v.Unmarshal
	g.json = v
}

func (g *GoNet) AcquireRequest() NetRequest {
	v := requestPool.Get()
	if v == nil {
		r := &GoNetRequest{}
		r.json = g.json
		r.r = g.client.R()
		return r
	}
	r := v.(*GoNetRequest)
	r.r = g.client.R()
	return r
}
