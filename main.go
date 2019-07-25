package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var remove *bool
var fileInfo *os.File
var files = make(map[[sha512.Size]byte]string)

func moveFile(fileName string, filePath string) {
	err := os.Rename(fileName, filePath)
	if err != nil {
		log.Println("Could not find file: ", err)
	}
}

func createDir(path string) (string, string) {
	ext := filepath.Ext(path)
	if ext == "" {
		log.Println("No extension deleting file")
		fmt.Println("Removing ", path)
		err := os.Remove(path)
		if err != nil {
			fmt.Printf("Could not delete file: %v", err)
		}
	}
	extName := strings.Replace(ext, ".", "", -1)
	if _, err := os.Stat(path + extName); os.IsNotExist(err) {
		os.Mkdir(extName, 0777)
	}
	file := filepath.Base(path)
	fmt.Println(extName, file)
	return extName, file
}

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

func sha256All(root string) (map[string][sha256.Size]byte, error) {
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

func checkDuplicate(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if info.IsDir() {
		return nil
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	hash := sha512.Sum512(data)
	if v, ok := files[hash]; ok {
		log.Printf("%q is a duplicate of %q\n", path, v)
		deleteDup(*remove, path)
	} else {
		files[hash] = path
	}

	return nil
}

func organizeByExtension() {
	files, err := ioutil.ReadDir(os.Args[1])
	if err != nil {
		log.Fatalf("Could not read directory: %v", err)
	}

	for _, file := range files {
		fileInfo, err := os.Open(os.Args[1] + file.Name())
		if err != nil {
			log.Fatalf("Could not readfile: %s - %v", file.Name(), err)
		}
		fileInfo.Close()
		extName, file := createDir(fileInfo.Name())

		moveFile(os.Args[1]+file, "./"+extName+"/"+file)
		continue
	}
}

func deleteDup(remove bool, v string) {
	if remove == true {
		os.Remove(v)
		fmt.Printf("Removed: %v\n", v)
	} else {
		return
	}
}

func main() {
	f, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	// if len(os.Args) != 2 {
	// 	fmt.Printf("USAGE : %s <target_directory> \n", os.Args[0])
	// 	os.Exit(0)
	// }

	dir := flag.String("directory", "", "Directory to Walk")

	remove = flag.Bool("delete", false, "Delete files.")

	flag.Parse()

	errs := filepath.Walk(*dir, checkDuplicate)
	if errs != nil {
		fmt.Println(errs)
		os.Exit(1)
	}
}
