package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	// READING FROM A FILE
	file, err := os.Open("test.txt")
	if err != nil {
		return
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// read the file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)

	// shorter and efficient way to reading files
	buf, err := ioutil.ReadFile("test.txt")
	if err != nil {
		return
	}
	file_str := string(buf)
	fmt.Println(file_str)

	// CREATING A FILE
	new_file, e := os.Create("new_file.txt")
	if e != nil {
		return
	}
	defer new_file.Close()
	new_file.WriteString("This is a new file")

	// READING AND WALKING A DIRECTORY
	dir, e := os.Open(".")
	if e != nil {
		return
	}
	defer dir.Close()

	fileInfos, e := dir.Readdir(-1)
	if e != nil {
		return
	}
	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

	// Recusively traversinga folder ( inclding subfodlers if any)
	// Use the Walk function in path/filepath package
	// the func that we pass to Walk will be calledfor every file and folder in the root folder
	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})
}
