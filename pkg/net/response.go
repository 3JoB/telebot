package net

type Response struct {
	code int
	body []byte
}

func (r *Response) StatusCode() int {
	return r.code
}

func (r *Response) IsStatusCode(v int) bool {
	return v == r.code
}

func (r *Response) Bytes() []byte {
	return r.body
}

func (r *Response) Reset() {
	r.code = 0
	r.body = []byte{}
}
