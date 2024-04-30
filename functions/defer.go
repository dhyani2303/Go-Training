package main

import "fmt"

func main() {

	defer foo()

	bar()

	fmt.Println("Hello from main")

}

func foo() {

	fmt.Println("Foo!!")
}

func bar() {

	fmt.Println("Bar!!")

	defer foo()
}
