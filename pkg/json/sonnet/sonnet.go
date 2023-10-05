package sonnet

import (
	"bytes"
	"io"

	"github.com/sugawarayuuta/sonnet"

	"github.com/3JoB/telebot/v2/pkg/json"
)

type Sonnet struct{}

func New() json.Json {
	return Sonnet{}
}

func (Sonnet) Marshal(v any) ([]byte, error) {
	return sonnet.Marshal(v)
}

func (Sonnet) MarshalIndent(v any, prefix string, indent string) ([]byte, error) {
	return sonnet.MarshalIndent(v, prefix, indent)
}

func (Sonnet) Unmarshal(buf []byte, v any) error {
	return sonnet.Unmarshal(buf, v)
}

func (Sonnet) NewEncoder(w io.Writer) json.Encoder {
	return sonnet.NewEncoder(w)
}

func (Sonnet) NewDecoder(r io.Reader) json.Decoder {
	return sonnet.NewDecoder(r)
}

func (Sonnet) Indent(dst *bytes.Buffer, src []byte, prefix string, indent string) error {
	return sonnet.Indent(dst, src, prefix, indent)
}
