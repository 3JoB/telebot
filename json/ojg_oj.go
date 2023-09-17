package json

import (
	"io"

	"github.com/goccy/go-json"
	"github.com/ohler55/ojg/oj"
)

type Ojg struct{}

func NewOjg() Json {
	return Ojg{}
}

func (Ojg) Marshal(v any) ([]byte, error) {
	return oj.Marshal(v)
}

func (Ojg) MarshalIndent(v any, prefix string, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func (Ojg) Unmarshal(buf []byte, v any) error {
	return oj.Unmarshal(buf, v)
}

func (Ojg) NewEncoder(w io.Writer) Encoder {
	return json.NewEncoder(w)
}

func (Ojg) NewDecoder(r io.Reader) Decoder {
	return json.NewDecoder(r)
}
