package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile(

		"test.txt",

		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,

		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteSlice := []byte("Hello there") // i character is of 1 byte

	bytesWritten, err := file.Write(byteSlice)

	if err != nil {
		log.Fatal(err)

	}
	fmt.Println(bytesWritten)
}
