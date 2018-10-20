package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Check for files that are duplicates using md5 and sha256 hashes
// if they are exactly the same put them in the duplicates folder.

var fileInfo *os.File
var beginning = "P:\\WDBlackDrive\\full_bring\\recup_dir."
var end = "\\"

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

func main() {
	start := time.Now()

	f, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	m, err := md5All(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
	for _, path := range paths {
		fmt.Printf("%x  %s\n", m[path], path)
	}

	for i := 1; i < 520; i++ {
		fmt.Println(strconv.Itoa(i))
		files, err := ioutil.ReadDir(beginning + strconv.Itoa(i) + end)
		if err != nil {
			log.Fatalf("Could not read directory: %v", err)
		}

		for _, file := range files {
			fileInfo, err = os.Open(beginning + strconv.Itoa(i) + end + file.Name())
			if err != nil {
				log.Fatalf("Could not readfile: %s - %v", file.Name(), err)
			}
			fileInfo.Close()
			extName, file := createDir(fileInfo.Name())

			moveFile(beginning+strconv.Itoa(i)+end+file, "./"+extName+"/"+file)
			continue
		}
	}
	elapsed := time.Since(start)
	log.Printf("Finished in: %s", elapsed)
}
