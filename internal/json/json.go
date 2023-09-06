//go:build !(sonic && avx && (linux || windows || darwin) && amd64)

package json

import "github.com/goccy/go-json"

var (
	// Marshal is exported by gin/json package.
	Marshal = json.Marshal

	// Unmarshal is exported by gin/json package.
	Unmarshal = json.Unmarshal

	// MarshalIndent is exported by gin/json package.
	MarshalIndent = json.MarshalIndent

	// NewDecoder is exported by gin/json package.
	NewDecoder = json.NewDecoder

	// NewEncoder is exported by gin/json package.
	NewEncoder = json.NewEncoder
)
