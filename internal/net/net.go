// net package is a dedicated hybrid network wrapper for TEP.
//
// It may have some impact on performance due to one or more additional allocations.
package net

import (
	"io"
	"sync"

	"github.com/3JoB/telebot/internal/json"
)

const UA = "Mozilla/5.0(compatible; Telebot-Expansion-Pack/v1; +https://github.com/3JoB/telebot)"

var (
	requestPool  sync.Pool
	responsePool sync.Pool
)

type NetFrame interface {
	SetJsonProcessor(v json.Json)
	AcquireRequest() NetRequest // Create a new request object
}

type NetRequest interface {
	MethodPOST()
	MethodGET()
	Body() io.Writer
	SetContentType(v string)
	SetRequestURI(v string)
	SetWriter(w io.Writer)
	Write(b []byte)
	WriteFile(content string, r io.Reader) error
	WriteJson(v any) error
	Do() (NetResponse, error)
	Reset()
	Release()
}

type NetResponse interface {
	StatusCode() int
	IsStatusCode(v int) bool
	Bytes() []byte
	Reset()
	Release()
}
