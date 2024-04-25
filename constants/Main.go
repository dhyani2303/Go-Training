package main

import "fmt"

func main() {

	//two types of constant -typed and untyped constant
	//no error is generated for unused constants

	const days int = 7 //typed constant
	const pi = 3.13    //untyped constant

	const (
		a = 5
		b // if we don't assign value to b by default the upper value will be considered
		c
	)

	fmt.Println(a, b, c)

	//const result = days * pi  //this is not allowed since days is defined as int

	const result = a * pi // this is valid since both are untyped constants

	// iota is number generator for constants which starts from zero
	// and is incremented by 1 automatically.

	const (
		c1 = iota
		c2 = iota
		c3 = iota
	)
	fmt.Println(c1, c2, c3) // => 0 1 2

	const (
		North = iota //by default 0
		East         //omitting type and value means, repeating its type and value so East = iota = 1 (it increments by 1 automatically)
		South        // -> 2
		West         // -> 3
	)

	// Initializing the constants using a step:
	const (
		c11 = iota * 2 // -> 0

		c22 // -> 2

		c33 // -> 4
	)

	const (
		c41 = iota

		_

		c42
	)

	fmt.Println(c41, c42)
}
