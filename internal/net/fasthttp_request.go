package net

import (
	"io"

	"github.com/valyala/fasthttp"

	"github.com/3JoB/telebot/json"
)

type FastHTTPRequest struct {
	json     json.Json
	client   *fasthttp.Client
	request  *fasthttp.Request
	response *fasthttp.Response
	w        io.Writer
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

// If this value is set, when reading data, the Body will be written
// directly to the set Writer interface without returning []byte.
func (f *FastHTTPRequest) SetWriter(w io.Writer) {
	f.w = w
}

func (f *FastHTTPRequest) Write(b []byte) {
	_, _ = f.request.BodyWriter().Write(b)
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
	defer f.Release()
	var err error
	f.request.Header.Set("User-Agent", "Mozilla/5.0(compatible; Telebot-Expansion-Pack/v1; +https://github.com/3JoB/telebot)")
	if err := f.client.Do(f.request, f.response); err != nil {
		return nil, err
	}
	resp := f.acquireResponse()
	resp.code = f.response.StatusCode()
	if f.response.StatusCode() != 200 {
		resp.body = f.response.Body()
	} else {
		if f.w != nil {
			err = f.response.BodyWriteTo(f.w)
		} else {
			resp.body = f.response.Body()
		}
	}

	return resp, err
}

func (f *FastHTTPRequest) Reset() {
	fasthttp.ReleaseRequest(f.request)
	fasthttp.ReleaseResponse(f.response)
	f.request = nil
	f.response = nil
	f.client = nil
	f.w = nil
}

func (f *FastHTTPRequest) Release() {
	f.Reset()
	requestPool.Put(f)
}
