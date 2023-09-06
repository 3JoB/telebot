package json

import (
	"io"

	"github.com/bytedance/sonic"
	"github.com/bytedance/sonic/decoder"
	"github.com/bytedance/sonic/encoder"
)

type Sonic struct {
	std sonic.API
}

func NewSonic() Json {
	return &Sonic{
		std: sonic.ConfigStd,
	}
}

func (s *Sonic) Marshal(v any) ([]byte, error) {
	return s.std.Marshal(v)
}

func (s *Sonic) Unmarshal(buf []byte, v any) error {
	return s.std.Unmarshal(buf, v)
}

func (s *Sonic) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return s.std.MarshalIndent(v, prefix, indent)
}

func (s *Sonic) NewDecoder(r io.Reader) Decoder {
	return decoder.NewStreamDecoder(r)
}

func (s *Sonic) NewEncoder(w io.Writer) Encoder {
	return encoder.NewStreamEncoder(w)
}
