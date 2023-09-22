package json

import (
	"io"

	"github.com/sugawarayuuta/sonnet"
)

type Sonnet struct{}

func NewSonnet() Json {
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

func (Sonnet) NewEncoder(w io.Writer) Encoder {
	return sonnet.NewEncoder(w)
}

func (Sonnet) NewDecoder(r io.Reader) Decoder {
	return sonnet.NewDecoder(r)
}
