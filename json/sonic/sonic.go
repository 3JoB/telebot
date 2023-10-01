package sonic

import (
	"bytes"
	"io"

	"github.com/bytedance/sonic"
	"github.com/goccy/go-json"

	tjson "github.com/3JoB/telebot/json"
)

type Sonic struct {
	std sonic.API
}

func New() tjson.Json {
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

func (s Sonic) NewDecoder(r io.Reader) tjson.Decoder {
	return s.std.NewDecoder(r)
}

func (s Sonic) NewEncoder(w io.Writer) tjson.Encoder {
	return s.std.NewEncoder(w)
}

func (s Sonic) Indent(dst *bytes.Buffer, src []byte, prefix string, indent string) error {
	return json.Indent(dst, src, prefix, indent)
}
