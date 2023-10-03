package telebot

import (
	"bufio"
	"io"
	"os"

	"github.com/3JoB/telebot/net"
)

// File object represents any sort of file.
type File struct {
	FileID   string `json:"file_id"`
	UniqueID string `json:"file_unique_id"`
	FileSize int64  `json:"file_size"`

	// FilePath is used for files on Telegram server.
	FilePath string `json:"file_path"`

	// FileLocal is used for files on local file system.
	FileLocal string `json:"file_local"`

	// FileURL is used for file on the internet.
	FileURL string `json:"file_url"`

	// FileReader is used for file backed with io.Reader.
	FileReader io.Reader `json:"-"`

	fileName string
}

// FromDisk constructs a new local (on-disk) file object.
//
// Note, it returns File, not *File for a very good reason:
// in telebot, File is pretty much an embeddable struct,
// so upon uploading media you'll need to set embedded File
// with something. NewFile() returning File makes it a one-liner.
//
//	photo := &tele.Photo{File: tele.FromDisk("chicken.jpg")}
func FromDisk(filename string) File {
	return File{FileLocal: filename}
}

// FromURL constructs a new file on provided HTTP URL.
//
// Note, it returns File, not *File for a very good reason:
// in telebot, File is pretty much an embeddable struct,
// so upon uploading media you'll need to set embedded File
// with something. NewFile() returning File makes it a one-liner.
//
//	photo := &tele.Photo{File: tele.FromURL("https://site.com/picture.jpg")}
func FromURL(url string) File {
	return File{FileURL: url}
}

// FromReader constructs a new file from io.Reader.
//
// Note, it returns File, not *File for a very good reason:
// in telebot, File is pretty much an embeddable struct,
// so upon uploading media you'll need to set embedded File
// with something. NewFile() returning File makes it a one-liner.
//
//	photo := &tele.Photo{File: tele.FromReader(bytes.NewReader(...))}
func FromReader(reader io.Reader) File {
	return File{FileReader: reader}
}

func (f *File) stealRef(g *File) {
	if g.OnDisk() {
		f.FileLocal = g.FileLocal
	}

	if g.FileURL != "" {
		f.FileURL = g.FileURL
	}
}

// InCloud tells whether the file is present on Telegram servers.
func (f *File) InCloud() bool {
	return f.FileID != ""
}

// OnDisk will return true if file is present on disk.
func (f *File) OnDisk() bool {
	_, err := os.Stat(f.FileLocal)
	return err == nil
}

type FileStorage struct {
	Reader *os.File
	ID     string
}

func (s *FileStorage) Copy(dst io.Writer) error {
	r, w := bufio.NewReader(s.Reader), bufio.NewWriter(dst)
	var buf [256]byte
    for{
        n, err := r.Read(buf[:])
        if err == io.EOF {
            break
        }
        if err != nil{
            return err
        }
        if _, err := w.Write(buf[:n]); err != nil {
			return err
		}
		if err := w.Flush(); err != nil {
			return err
		}
    }
	return nil
}

func (s *FileStorage) Close() (err error) {
	if s.ID != "" {
		s.Reader.Close() //nolint:errcheck
		s.Reader = nil
		err = net.RemoveTemp(s.ID)
	} else {
		if s.Reader != nil {
			err = s.Reader.Close()
		}
	}
	return
}
