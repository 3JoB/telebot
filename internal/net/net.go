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

	// Create a new request object
	AcquireRequest() NetRequest
}

type NetRequest interface {
	MethodPOST()
	MethodGET()

	// Only fasthttp
	Body() io.Writer

	SetContentType(v string)
	SetRequestURI(v string)
	SetWriter(w io.Writer)
	Write(b []byte)
	WriteFile(content string, r io.Reader) error

	// Write the structure directly to the Body as json,
	// which will be processed by the interface.
	WriteJson(v any) error

	// Execute request.
	Do() (NetResponse, error)

	// Release() will clear the data in the current pointer.
	// It is recommended to call it within the Release() method instead
	// of calling it externally.
	Reset()

	// This method is generally not recommended because the built-in methods
	// have automatically called Release() at the end of the Do() method,
	// and only Response needs to be called manually.
	//
	// Release() will clear the data in the current pointer and put it back
	// into the pool. After release, the corresponding pointer should not be used anymore.
	Release()
}

type NetResponse interface {
	StatusCode() int
	IsStatusCode(v int) bool

	// If SetWriter() is called in req, this method will
	// not be set (unless the status code is not 200)
	Bytes() []byte

	// Release() will clear the data in the current pointer.
	// It is recommended to call it within the Release() method instead
	// of calling it externally.
	Reset()

	// Release() will clear the data in the current pointer and put it back
	// into the pool. After release, the corresponding pointer should not be used anymore.
	Release()
}
