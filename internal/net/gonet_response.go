package net

type GoNetResponse struct{
	code int
	body []byte
}

func (g *GoNetResponse) StatusCode() int {
	return g.code
}

func (g *GoNetResponse) IsStatusCode(v int) bool {
	return v == g.code
}

func (g *GoNetResponse) Bytes() []byte {
	return g.body
}

func (g *GoNetResponse) Reset() {
	g.code = 0
	g.body = nil
}

func (g *GoNetResponse) Release() {
	g.Reset()
	responsePool.Put(g)
}
