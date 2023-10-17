package net

import (
	"bytes"
	"io"
	"sync"

	"github.com/valyala/fasthttp"

	"github.com/3JoB/telebot/v2/pkg/json"
)

type FastHTTPRequest struct {
	json         json.Json
	w            *bytes.Buffer
	f            io.ReadWriteCloser
	client       *fasthttp.Client
	request      *fasthttp.Request
	response     *fasthttp.Response
	responsePool *sync.Pool
}

func (f *FastHTTPRequest) acquire() {
	f.request, f.response = fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
}

func (f *FastHTTPRequest) acquireResponse() *FastHTTPResponse {
	v := responsePool.Get()
	if v == nil {
		return &FastHTTPResponse{}
	}
	return v.(*FastHTTPResponse)
}

func (f *FastHTTPRequest) MethodGET() {
	f.request.Header.SetMethod("GET")
}

func (f *FastHTTPRequest) MethodPOST() {
	f.request.Header.SetMethod("POST")
}

func (f *FastHTTPRequest) SetRequestURI(v string) {
	f.request.SetRequestURI(v)
}

func (f *FastHTTPRequest) SetContentType(v string) {
	f.request.Header.Set("Content-Type", v)
}

func (f *FastHTTPRequest) SetWriter(w *bytes.Buffer) {
	f.w = w
}

func (f *FastHTTPRequest) SetWriteCloser(v io.ReadWriteCloser) {
	f.f = v
}

func (f *FastHTTPRequest) Write(b []byte) {
	f.request.BodyWriter().Write(b) //nolint:errcheck
}

func (f *FastHTTPRequest) WriteFile(content string, r io.Reader) error {
	f.SetContentType(content)
	f.MethodPOST()
	_, err := io.Copy(f.request.BodyWriter(), r)
	return err
}

func (f *FastHTTPRequest) WriteJson(v any) error {
	if err := f.json.NewEncoder(f.request.BodyWriter()).Encode(v); err != nil {
		return err
	}
	f.SetContentType("application/json")
	return nil
}

// Body returns writer for populating request body.
func (f *FastHTTPRequest) Body() io.Writer {
	return f.request.BodyWriter()
}

func (f *FastHTTPRequest) Do() (NetResponse, error) {
	var err error
	f.request.Header.Set("User-Agent", UA)

	if err := f.client.Do(f.request, f.response); err != nil {
		return nil, err
	}

	resp := f.acquireResponse()
	resp.code = f.response.StatusCode()

	if resp.IsStatusCode(200) && f.f != nil {
		_, err = f.response.WriteTo(f.f)
		goto END
	}

	if f.w != nil {
		f.w.Write(f.response.Body()) //nolint:errcheck
	} else {
		resp.body = f.response.Body()
	}

END:

	return resp, err
}

func (f *FastHTTPRequest) Reset() {
	fasthttp.ReleaseRequest(f.request)
	fasthttp.ReleaseResponse(f.response)
	f.request = nil
	f.response = nil
	f.client = nil
	f.f = nil
	f.w = nil
}
