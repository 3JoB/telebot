// net package is a dedicated hybrid network wrapper for TEP.
package net

import (
	"io"

	"github.com/cornelk/hashmap"

	"github.com/3JoB/telebot/internal/json"
)

type NetFrame interface {
	Header() *Header
	SetJsonProcessor(v json.Json)
	GETFile()
	POSTFile()
	GETJson()
	POSTJson()
	Acquire() NetRequest // Create a new request object
}

type NetRequest interface {
	MethodPOST()
	MethodGET()
	AddHeader(k, v string)
	AddHeaders(m *hashmap.Map[string, string])
	SetHeader(k, v string)
	SetHeaders(m *hashmap.Map[string, string])
	SetRequestURI(v string)
	SendJson(v any) NetResponse
	SendFile() NetResponse
	SendAny() NetResponse
	Release()
}

type NetResponse interface {
	StatusCode() int
	IsStatusCode(v int) bool
	Reader() io.ReadCloser
	Bytes() []byte
	Release()
}

type Header struct {
	m *hashmap.Map[string, string]
}
