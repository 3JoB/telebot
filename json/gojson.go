package json

import (
	"io"

	"github.com/goccy/go-json"
)

type GoJson struct{}

func NewGoJson() Json {
	return GoJson{}
}

func (GoJson) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (GoJson) MarshalIndent(v any, prefix string, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func (GoJson) Unmarshal(buf []byte, v any) error {
	return json.Unmarshal(buf, v)
}

func (GoJson) NewEncoder(w io.Writer) Encoder {
	return json.NewEncoder(w)
}

func (GoJson) NewDecoder(r io.Reader) Decoder {
	return json.NewDecoder(r)
}
