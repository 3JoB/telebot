package ojg

import (
	"bytes"
	"io"

	"github.com/goccy/go-json"
	"github.com/ohler55/ojg/oj"

	tjson "github.com/3JoB/telebot/json"
)

type Ojg struct{}

func New() tjson.Json {
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

func (Ojg) NewEncoder(w io.Writer) tjson.Encoder {
	return json.NewEncoder(w)
}

func (Ojg) NewDecoder(r io.Reader) tjson.Decoder {
	return json.NewDecoder(r)
}

func (Ojg) Indent(dst *bytes.Buffer, src []byte, prefix string, indent string) error {
	return json.Indent(dst, src, prefix, indent)
}
