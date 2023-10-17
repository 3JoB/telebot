package net

import "testing"

func TestNilRelease(m *testing.M) {
	cli := NewFastHTTPClient()
	cli.Release(nil, nil)
}
