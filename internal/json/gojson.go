package json

import (
	"io"

	"github.com/goccy/go-json"
)

type GoJson struct{}

func NewGoJson() Json {
	return &GoJson{}
}

func (g *GoJson) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (g *GoJson) MarshalIndent(v any, prefix string, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func (g *GoJson) Unmarshal(buf []byte, v any) error {
	return json.Unmarshal(buf, v)
}

func (g *GoJson) NewEncoder(w io.Writer) Encoder {
	return json.NewEncoder(w)
}

func (g *GoJson) NewDecoder(r io.Reader) Decoder {
	return json.NewDecoder(r)
}
