package json

import (
	"io"

	"github.com/bytedance/sonic"
)

type Sonic struct {
	std sonic.API
}

func NewSonic() Json {
	return &Sonic{
		std: sonic.ConfigFastest,
	}
}

func (s Sonic) Marshal(v any) ([]byte, error) {
	return s.std.Marshal(v)
}

func (s Sonic) Unmarshal(buf []byte, v any) error {
	return s.std.Unmarshal(buf, v)
}

func (s Sonic) MarshalIndent(v any, prefix, indent string) ([]byte, error) {
	return s.std.MarshalIndent(v, prefix, indent)
}

func (s Sonic) NewDecoder(r io.Reader) Decoder {
	return s.std.NewDecoder(r)
}

func (s Sonic) NewEncoder(w io.Writer) Encoder {
	return s.std.NewEncoder(w)
}
