package net

import (
	"bytes"
	"io"

	"github.com/3JoB/resty-ilo"

	"github.com/3JoB/telebot/pkg/json"
)

type GoNetRequest struct {
	uri    string
	method string
	json   json.Json
	f   io.ReadWriteCloser
	w      *bytes.Buffer
	r      *resty.Request
}

func (g *GoNetRequest) acquireResponse() *GoNetResponse {
	v := responsePool.Get()
	if v == nil {
		return &GoNetResponse{}
	}
	return v.(*GoNetResponse)
}

func (g *GoNetRequest) MethodGET() {
	g.method = "GET"
}

func (g *GoNetRequest) MethodPOST() {
	g.method = "POST"
}

func (g *GoNetRequest) SetRequestURI(v string) {
	g.uri = v
}

func (g *GoNetRequest) SetContentType(v string) {
	g.r = g.r.SetHeader("Content-Type", v)
}

func (g *GoNetRequest) SetWriter(w *bytes.Buffer) {
	g.w = w
}

func (g *GoNetRequest) SetTemp(v io.ReadWriteCloser) {
	g.f = v
}

func (g *GoNetRequest) Write(b []byte) {
	g.r = g.r.SetBody(b)
}

func (g *GoNetRequest) WriteFile(content string, r io.Reader) error {
	g.SetContentType(content)
	g.MethodPOST()
	g.r = g.r.SetBody(r)
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
	var (
		err      error
		response *resty.Response
	)
	g.r = g.r.SetHeader("User-Agent", UA)

	if g.method == "POST" {
		response, err = g.r.Post(g.uri)
	} else {
		response, err = g.r.Get(g.uri)
	}

	if err != nil {
		return nil, err
	}

	resp := g.acquireResponse()
	resp.code = response.StatusCode()

	if resp.IsStatusCode(200) && g.f != nil {
		_, err = Copy(g.f, response.RawBody())
		goto END
	}
	if g.w != nil {
		g.w.Write(response.Body()) //nolint:errcheck
	} else {
		resp.body = response.Body()
	}

END:
	response.RawBody().Close() //nolint:errcheck
	return resp, err
}

func (g *GoNetRequest) Reset() {
	g.uri = ""
	g.method = ""
	g.r = nil
	g.w = nil
	g.f = nil
}

func (g *GoNetRequest) Release() {
	g.Reset()
	requestPool.Put(g)
}
