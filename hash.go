package main

import (
	"crypto/md5"
	"crypto/sha1"
	"io/ioutil"
	"os"
	"path/filepath"
)

func md5All(root string) (map[string][md5.Size]byte, error) {
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

func sha1All(root string) (map[string][sha1.Size]byte, error) {
	m := make(map[string][sha1.Size]byte)
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
		m[path] = sha1.Sum(data)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return m, nil
}
