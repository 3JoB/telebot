package fs

import "io"

type FileSystem interface {
	Create(string) (io.ReadWriteCloser, error)
}
