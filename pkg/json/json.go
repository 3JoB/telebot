package json

import (
	"bytes"
	"io"
)

// The Json interface is used to customize the json handler.
// Five wrappers are provided by default. For detailed documentation,
// see: https://pkg.go.dev/github.com/3JoB/telebot/json.
//
// Some methods use the default sonnet because they are not under *Bot.
type Json interface {
	Marshal(v any) ([]byte, error)
	MarshalIndent(v any, prefix string, indent string) ([]byte, error)
	Unmarshal(buf []byte, v any) error
	NewEncoder(w io.Writer) Encoder
	NewDecoder(r io.Reader) Decoder
	Indent(dst *bytes.Buffer, src []byte, prefix string, indent string) error
}

type Encoder interface {
	Encode(v any) error
}

type Decoder interface {
	Decode(v any) error  // Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by v.
}
