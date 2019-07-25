package main

import (
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var remove, move *bool
var dir *string
var fileInfo *os.File
var files = make(map[[sha512.Size]byte]string)

func main() {
	f, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	dir = flag.String("directory", "", "String - Directory to Walk: default empty")
	remove = flag.Bool("delete", false, "Bool - Delete files: default false")
	move = flag.Bool("move", false, "Bool - Move files to duplicate directory: default false")

	flag.Parse()

	errs := filepath.Walk(*dir, checkDuplicate)
	if errs != nil {
		fmt.Println(errs)
		os.Exit(1)
	}
	// Organize files into there own folder named after there file extension
	// files, err := ioutil.ReadDir(directoryPath)
	// if err != nil {
	// 	log.Fatalf("Could not read directory: %v", err)
	// }

	// for _, file := range files {
	// 	fileInfo, err = os.Open(directoryPath + file.Name())
	// 	if err != nil {
	// 		log.Fatalf("Could not readfile: %s - %v", file.Name(), err)
	// 	}
	// 	fileInfo.Close()
	// 	extName, file := createDir(fileInfo.Name())

	// 	moveFile(directoryPath+file, "./"+extName+"/"+file)
	// 	continue
	// }
}

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

func checkDuplicate(file string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if info.IsDir() {
		return nil
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	hash := sha512.Sum512(data)
	if v, ok := files[hash]; ok {
		log.Printf("%q is a duplicate of %q\n", file, v)
		handleFile(*remove, *move, *dir, file)
	} else {
		files[hash] = file
	}

	return nil
}

func handleFile(remove, move bool, dir, path string) {
	if remove == true {
		os.Remove(path)
		fmt.Printf("Removed: %v\n", path)
	} else if move == true {
		if _, err := os.Stat("duplicate/" + path); os.IsNotExist(err) {
			fmt.Println(path)
			os.Mkdir("duplicate/"+path, 0777)
		}
		re := regexp.MustCompile(dir)
		file := re.ReplaceAllString(path, "")
		moveFile(path, "./duplicate/"+file)
	} else {
		return
	}

}
