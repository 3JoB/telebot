package net

import "testing"

func Test_NilRelease(m *testing.T) {
	cli := NewFastHTTPClient()
	cli.Release(nil, nil)
}
