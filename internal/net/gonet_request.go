package net

import (
	"io"

	"github.com/3JoB/resty-ilo"
	"github.com/3JoB/telebot/internal/json"
)

type GoNetRequest struct {
	json json.Json
	uri  string
	r    *resty.Request
	w    io.Writer
}

func (g *GoNetRequest) acquireResponse() *GoNetResponse {
	v := responsePool.Get()
	if v == nil {
		return &GoNetResponse{}
	}
	return v.(*GoNetResponse)
}

func (g *GoNetRequest) MethodGET() {
	g.r.Method = "GET"
}

func (g *GoNetRequest) MethodPOST() {
	g.r.Method = "POST"
}

func (g *GoNetRequest) SetRequestURI(v string) {
	g.uri = v
}

func (g *GoNetRequest) SetContentType(v string) {
	g.r = g.r.SetHeader("Content-Type", v)
}

// If this value is set, when reading data, the Body will be written
// directly to the set Writer interface without returning []byte.
func (g *GoNetRequest) SetWriter(w io.Writer) {
	g.w = w
}

func (g *GoNetRequest) Write(b []byte) {
	g.r = g.r.SetBody(b)
}

func (g *GoNetRequest) WriteFile(content string, r io.Reader) error {
	g.SetContentType(content)
	g.MethodPOST()
	g.r = g.r.SetBody(g.Body())
	return nil
}

func (g *GoNetRequest) WriteJson(v any) error {
	g.r = g.r.SetBody(v)
	g.SetContentType("application/json")
	return nil
}

// Body returns writer for populating request body.
func (g *GoNetRequest) Body() io.Writer {
	return nil
}

func (g *GoNetRequest) Do() (NetResponse, error) {
	defer g.Release()
	var err error
	g.r = g.r.SetHeader("User-Agent", "Mozilla/5.0(compatible; Telebot-Expansion-Pack/v1; +https://github.com/3JoB/telebot)")
	response, err := g.r.Send()
	if err != nil {
		return nil, err
	}
	resp := g.acquireResponse()
	resp.code = response.StatusCode()
	if g.w != nil {
		_, err = io.Copy(g.w, response.RawBody())
	} else {
		resp.body = response.Body()
	}
	_ = response.RawBody().Close()
	return resp, err
}

func (g *GoNetRequest) Reset() {
	g.r = nil
	g.uri = ""
	g.w = nil
}

func (g *GoNetRequest) Release() {
	g.Reset()
	requestPool.Put(g)
}
