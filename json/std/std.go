package json

import (
	"bytes"
	"encoding/json"
	"io"

	tjson "github.com/3JoB/telebot/json"
)

type Std struct{}

func New() tjson.Json {
	return Std{}
}

func (Std) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (Std) MarshalIndent(v any, prefix string, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func (Std) Unmarshal(buf []byte, v any) error {
	return json.Unmarshal(buf, v)
}

func (Std) NewEncoder(w io.Writer) tjson.Encoder {
	return json.NewEncoder(w)
}

func (Std) NewDecoder(r io.Reader) tjson.Decoder {
	return json.NewDecoder(r)
}

func (Std) Indent(dst *bytes.Buffer, src []byte, prefix string, indent string) error {
	return json.Indent(dst, src, prefix, indent)
}
