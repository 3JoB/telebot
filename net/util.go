package net

import (
	"os"
	"path/filepath"

	"github.com/3JoB/ulib/fsutil"
)

const (
	TempDir string = "/telebot/temp"
)

func GetTempDir() (string, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}
	p := filepath.Join(cacheDir, TempDir)
	if !fsutil.IsExist(p) {
		if err := fsutil.Mkdir(p); err != nil {
			return "", err
		}
	}
	return p, nil
}

func GetTemp(id string) (*os.File, error) {
	p, err := GetTempDir()
	if err != nil {
		return nil, err
	}
	file := filepath.Join(p, "/."+id)
	f, err := fsutil.OpenFile(file, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func SetTemp(path string) (*os.File, error) {
	p, err := GetTempDir()
	if err != nil {
		return nil, err
	}
	file := filepath.Join(p, "/."+path)
	f, err := fsutil.OpenFile(file, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func RemoveTemp(id string) error {
	p, err := GetTempDir()
	if err != nil {
		return err
	}
	p = filepath.Join(p, "/"+id)
	return fsutil.Remove(p)
}

func CleanTemp() error {
	p, err := GetTempDir()
	if err != nil {
		return err
	}
	return fsutil.Remove(p)
}
