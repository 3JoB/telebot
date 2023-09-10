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
	Acquire() NetRequest // Create a new request object
}

type NetRequest interface {
	MethodPOST()
	MethodGET()
	AddHeader(k, v string)
	AddHeaders(m map[string]string)
	SetHeader(k, v string)
	SetHeaders(m map[string]string)
	SetRequestURI(v string)
	Body() io.Writer
	SendJson() NetResponse
	SendFile() NetResponse
	SendAny() NetResponse
	Reset()
	Release()
}

type NetResponse interface {
	StatusCode() int
	IsStatusCode(v int) bool
	Reader() io.ReadCloser
	Bytes() []byte
	Reset()
	Release()
}
