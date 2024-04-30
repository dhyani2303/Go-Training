package main

import "fmt"

// declaring an empty interface
type emptyInterface interface {
}

// declaring a new struct type which has one field of type empty interface
type person struct {
	info interface{}
}

func main() {

	var empty emptyInterface

	empty = 5

	fmt.Println(empty)

	empty = "Go"

	fmt.Println(empty)

	empty = []int{1, 2, 3}

	fmt.Println(empty)

	you := person{}

	you.info = "Name"
	you.info = "Age=20"

	fmt.Println(you)
}
