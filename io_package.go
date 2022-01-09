package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	file, err := os.Open("./dir/file.txt")

	if err != nil {
		fmt.Println("Error on open the file")
		return
	}

	defer file.Close()

	stat, err := file.Stat()

	if err != nil {
		fmt.Println("Error on stat")
		return
	}

	bytes := make([]byte, stat.Size())

	_, err = file.Read(bytes)

	if err != nil {
		fmt.Println("Error on read the file")
		return
	}

	str := string(bytes)

	fmt.Println(str)

	// shorter way to read file
	bs, err := ioutil.ReadFile("./dir/file.txt")

	if err != nil {

		fmt.Println("Error on Read File")

		return
	}

	str2 := string(bs)

	fmt.Println(str2)

	// create file

	newFile, err := os.Create("./test.txt")

	if err != nil {
		fmt.Println("Error on Create File")
		return
	}

	defer newFile.Close()

	newFile.WriteString("write from program")

	// list file in a directory

	dir, err := os.Open("./dir")

	if err != nil {
		return
	}

	defer dir.Close()

	fileInfos, err := dir.Readdir(-1)

	if err != nil {
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}

	// read recursively a directory

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path)
		return nil
	})

}
