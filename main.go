package main

import (
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

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
		fmt.Printf("%q is a duplicate of %q\n", path, v)
	} else {
		files[hash] = path
	}

	return nil
}

func main() {
	f, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	if len(os.Args) != 2 {
		fmt.Printf("USAGE : %s <target_directory> \n", os.Args[0])
		os.Exit(0)
	}

	dir := os.Args[1]

	errs := filepath.Walk(dir, checkDuplicate)
	if errs != nil {
		fmt.Println(errs)
		os.Exit(1)
	}

	// files, err := ioutil.ReadDir(os.Args[1])
	// if err != nil {
	// 	log.Fatalf("Could not read directory: %v", err)
	// }

	// for _, file := range files {
	// 	fileInfo, err = os.Open(os.Args[1] + file.Name())
	// 	if err != nil {
	// 		log.Fatalf("Could not readfile: %s - %v", file.Name(), err)
	// 	}
	// 	fileInfo.Close()
	// 	extName, file := createDir(fileInfo.Name())

	// 	moveFile(os.Args[1]+file, "./"+extName+"/"+file)
	// 	continue
	// }
}
