package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// os.Create() function creates a file if it doesn't already exist. If it exists, the file is truncated.
	// it returns a file descriptor which is a pointer to os.File and an error value.

	var newFile *os.File

	var err error

	newFile, err = os.Create("test.txt")
	if err != nil {
		log.Fatal(err) //program will be terminated
	}

	err = os.Truncate("test.txt", 0) //0 means completely empty the file.

	// error handling
	if err != nil {
		log.Fatal(err)
	}

	newFile.Close()

	file, err := os.Open("test.txt") //file is opened in read only mode

	file, err = os.OpenFile("test.txt", os.O_APPEND, 0644)

	// We can Use opening attributes individually or combined
	// using an OR between them
	// e.g. os.O_CREATE|os.O_APPEND
	// or os.O_CREATE|os.O_TRUNC|os.O_WRONLY
	// os.O_RDONLY // Read only
	// os.O_WRONLY // Write only
	// os.O_RDWR // Read and write
	// os.O_APPEND // Append to end of file
	// os.O_CREATE // Create is none exist
	// os.O_TRUNC // Truncate file when opening
	if err != nil {
		fmt.Println(err)

	}
	file.Close()

	var fileInfo os.FileInfo

	fileInfo, err = os.Stat("test.txt")

	p := fmt.Println

	p("File Name:", fileInfo.Name())
	p("Size in bytes:", fileInfo.Size())

	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("The file does not exist")
		}
	}

	//err = os.Rename("test.txt", "a.txt")
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//

	err = os.Remove("a.txt")
	if err != nil {
		log.Fatal(err)
	}
	p("File is removed")

}
