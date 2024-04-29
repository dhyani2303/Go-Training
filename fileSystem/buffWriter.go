package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile("test.txt",
		os.O_WRONLY|os.O_CREATE|os.O_APPEND,
		0644,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//creating a buffered Writer

	bufferWriter := bufio.NewWriter(file)

	bs := []byte{96, 97, 98}

	bytesWritten, err := bufferWriter.Write(bs)

	if err != nil {
		log.Fatal(err)
	}
	log.Println(bytesWritten)

	bufferWriter.WriteString("\nHello there\n")

	log.Println("Available buffer size ", bufferWriter.Available(), " Buffer to be written ", bufferWriter.Buffered())

	bufferWriter.Flush()

}
