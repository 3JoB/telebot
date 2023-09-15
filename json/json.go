package json

import "io"

type Json interface {
	Marshal(v any) ([]byte, error)
	MarshalIndent(v any, prefix string, indent string) ([]byte, error)
	Unmarshal(buf []byte, v any) error
	NewEncoder(w io.Writer) Encoder
	NewDecoder(r io.Reader) Decoder
}

type Encoder interface {
	Encode(v any) error
}

type Decoder interface {
	Decode(v any) error  // Decode reads the next JSON-encoded value from its input and stores it in the value pointed to by v.
	Buffered() io.Reader // Buffered returns a reader of the data remaining in the Decoder's buffer. The reader is valid until the next call to Decode.
}
