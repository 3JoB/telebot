package net

import "sync"

type FastHTTPResponse struct {
	pool *sync.Pool
	code int
	body []byte
}

func (r *FastHTTPResponse) StatusCode() int {
	return r.code
}

func (r *FastHTTPResponse) IsStatusCode(v int) bool {
	return v == r.code
}

func (r *FastHTTPResponse) Bytes() []byte {
	return r.body
}

func (r *FastHTTPResponse) Reset() {
	r.code = 0
	r.body = nil
}

func (r *FastHTTPResponse) Release() {
	r.Reset()
	r.pool.Put(r)
}
