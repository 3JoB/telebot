package json

import (
	"io"

	"github.com/goccy/go-json"
	jet "github.com/wI2L/jettison"
)

type Jet struct{}

// Since Jet only has the Marshal method,
// all other methods are still provided by go-json as is.
func NewJet() Json {
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

func (Jet) NewEncoder(w io.Writer) Encoder {
	return json.NewEncoder(w)
}

func (Jet) NewDecoder(r io.Reader) Decoder {
	return json.NewDecoder(r)
}
