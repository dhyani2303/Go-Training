package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// declaring a byte slice and initializing it with a length of 2
	//byteSlice := make([]byte, 2) //only 2 bytes of data will be read
	//
	//numRead, err := io.ReadFull(file, byteSlice)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println("Number  of Bytes read:", numRead)
	//log.Println(string(byteSlice))

	//** READING WHOLE FILE INTO A BYTESLICE
	//data, err := io.ReadAll(file)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(string(data))

	// another method to read file
	//data, err := os.ReadFile("test.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Printf("Data read: %s\n", data)

	//reading file line by line

	scanner := bufio.NewScanner(file)
	//scanner.Split(bufio.ScanWords) //data is read word by word

	// the default scanner is bufio.ScanLines and that means it will scan a file line by line.
	// there are also bufio.ScanWords and bufio.ScanRunes.

	success := scanner.Scan()

	if success == false {
		err := scanner.Err()

		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("Scan has been completed")
		}

	}
	fmt.Println("First line:", scanner.Text())

	//to read whole file line by line
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)

	}

}
