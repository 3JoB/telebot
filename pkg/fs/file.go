package fs

import (
	"io"
	"os"

	"github.com/3JoB/ulib/fsutil"
)

type File struct{}

func (File) Create(id string) (io.ReadWriteCloser, error) {
	return fsutil.OpenFile(id, os.O_RDWR|os.O_CREATE, 0666)
}