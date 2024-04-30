package main

import "fmt"

// An anonymous function is a function which doesn’t contain any name and is declared inline using a function literal.
// Anonymous functions can be used closures.
func main() {

	a := increment(10)

	fmt.Printf("%T\n", a)

	a() // since a is func() int type we can call it like this. //11

	fmt.Println(a()) //12

}

func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}
