package json

import (
	"bytes"
	"io"

	"github.com/goccy/go-json"

	tjson "github.com/3JoB/telebot/json"
)

type GoJson struct{}

func New() tjson.Json {
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

func (GoJson) NewEncoder(w io.Writer) tjson.Encoder {
	return json.NewEncoder(w)
}

func (GoJson) NewDecoder(r io.Reader) tjson.Decoder {
	return json.NewDecoder(r)
}

func (GoJson) Indent(dst *bytes.Buffer, src []byte, prefix string, indent string) error {
	return json.Indent(dst, src, prefix, indent)
}
