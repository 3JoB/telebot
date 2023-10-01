package jet

import (
	"bytes"
	"io"

	"github.com/goccy/go-json"
	jet "github.com/wI2L/jettison"

	tjson "github.com/3JoB/telebot/json"
)

type Jet struct{}

// Since Jet only has the Marshal method,
// all other methods are still provided by go-json as is.
func NewJet() tjson.Json {
	return Jet{}
}

func (Jet) Marshal(v any) ([]byte, error) {
	return jet.Marshal(v)
}

func (Jet) MarshalIndent(v any, prefix string, indent string) ([]byte, error) {
	return json.MarshalIndent(v, prefix, indent)
}

func (Jet) Unmarshal(buf []byte, v any) error {
	return json.Unmarshal(buf, v)
}

func (Jet) NewEncoder(w io.Writer) tjson.Encoder {
	return json.NewEncoder(w)
}

func (Jet) NewDecoder(r io.Reader) tjson.Decoder {
	return json.NewDecoder(r)
}

func (Jet) Indent(dst *bytes.Buffer, src []byte, prefix string, indent string) error {
	return json.Indent(dst, src, prefix, indent)
}
