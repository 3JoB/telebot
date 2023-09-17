package json

import (
	"encoding/json"
	"io"
)

type Std struct{}

func NewStd() Json {
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

func (Std) NewEncoder(w io.Writer) Encoder {
	return json.NewEncoder(w)
}

func (Std) NewDecoder(r io.Reader) Decoder {
	return json.NewDecoder(r)
}
