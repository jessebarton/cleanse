package hash

import (
	"crypto/md5"
	"crypto/sha256"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Md5All(root string) (map[string][md5.Size]byte, error) {
	m := make(map[string][md5.Size]byte)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		m[path] = md5.Sum(data)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return m, nil
}

func Sha256All(root string) (map[string][sha256.Size]byte, error) {
	s := make(map[string][sha256.Size]byte)
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return err
		}
		s[path] = sha256.Sum256(data)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return s, nil
}
