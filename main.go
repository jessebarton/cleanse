package main

import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// Check for files that are duplicates using md5 and sha256 hashes
// if they are exactly the same put them in the duplicates folder.

// Check for files with extensions that aren't important and then put them in the delete folder.
// List of extensions to not delete (jpg, png, jpeg, cr2)

var filePath = "./files/"

func sha256Hash(f *os.File) {
	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s: %x\n", f.Name(), h.Sum(nil))
}

func md5Hash(f *os.File) {
	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s: %x\n", f.Name(), h.Sum(nil))
}

func moveFile(fileName string, filePath string) {
	err := os.Rename(fileName, filePath)
	if err != nil {
		panic(err)
	}
}

func createDir(path string) {
	extensions := []string{
		".jpg",
		".png",
		".jpeg",
		".cr2",
		".txt",
		".plist",
	}

	ext := filepath.Ext(path)
	for _, extension := range extensions {
		if ext == extension {
			dir := strings.Replace(ext, ".", "", -1)
			if _, err := os.Stat(path + dir); os.IsNotExist(err) {
				os.Mkdir(dir, 0777)
				file := filepath.Base(path)
				fmt.Println(dir, file)
				moveFile(filePath+file, "./"+dir+"/"+file)
			}
		}
	}
}

func main() {
	files, err := ioutil.ReadDir(filePath)
	if err != nil {
		log.Fatalf("Could not read directory: %v", err)
	}
	var f *os.File

	defer f.Close()
	for _, file := range files {
		f, err = os.Open(filePath + file.Name())
		if err != nil {
			log.Fatalf("Could not readfile: %s - %v", file.Name(), err)
		}
		createDir(f.Name())
	}
}
