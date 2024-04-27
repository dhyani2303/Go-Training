package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("os.Args:", os.Args) //os.Args is slice of strings ([]string)

	fmt.Println("Argument 1: ", os.Args[0])

	fmt.Println("Argument 2: ", os.Args[1])

	fmt.Println("Argument 3: ", os.Args[2])

}
