package main

import "fmt"

func main() {

	//each key corresponds to an index of array. the index starts with 0

	grades := [3]int{ //the keyed elements can be in any order
		1: 10,
		0: 5,
		2: 7,
	}
	fmt.Println(grades) //[5,10,7]

	accounts := [3]int{
		2: 50,
	}
	fmt.Println(accounts)

	names := [...]string{ // if key is not specified then element gets its index from the last keyed element
		5: "London",
		"Paris",
	}
	fmt.Printf("%#v\n", names)
}
