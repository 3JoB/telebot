package temp

import (
	"os"
	"path/filepath"

	"github.com/3JoB/ulib/fsutil"
)

const TempDir string = "/telebot/temp"

var Dir = getDir()

func Joi(id string) string {
	if Dir == "" {
		return ""
	}
	return filepath.Join(Dir, "/."+id)
}

func getDir() string {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return ""
	}
	p := filepath.Join(cacheDir, TempDir)
	if !fsutil.IsExist(p) {
		if err := fsutil.Mkdir(p); err != nil {
			return ""
		}
	}
	return p
}

func Get(id string) (*os.File, error) {
	file := Joi(id)
	return fsutil.OpenFile(file, os.O_RDWR|os.O_CREATE, 0666)
}

func Set(id string) (*os.File, error) {
	file := Joi(id)
	f, err := fsutil.OpenFile(file, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func Do(id, now string) error {
	return fsutil.Symlink(Joi(id), filepath.Clean(now))
}

func Remove(id string) error {
	return fsutil.Remove(Joi(id))
}

func Clean() error {
	return fsutil.Remove(Dir)
}
