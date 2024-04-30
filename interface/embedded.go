package main

import "fmt"

type Reader interface {
	read() string
}

type Writer interface {
	write(str string)
}

type ReadWriter interface {
	Reader
	Writer
}

type Book struct {
	Title string

	Author string

	Pages int
}

func (b Book) read() string {

	return "Reading " + b.Title + "by" + b.Author
}
func (b Book) write(str string) {

	fmt.Println("Writing", str)

}

func main() {

	book := Book{
		Title:  "The Great Gatsby",
		Author: "F.Scott Fitzgerald",
		Pages:  100,
	}

	var rw ReadWriter = book

	fmt.Println("Reading", rw.read())

	rw.write("Hello there")

}
