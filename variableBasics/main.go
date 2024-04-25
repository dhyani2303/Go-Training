package main

import "fmt"

// we can declare package scope variables as below. and we will not get compile time error for not using this variables.
// but we cannot declare short type variables as package scope variable
// y:=100 this is invalid
var y int

func main() {

	var name string

	name = "dhyani"

	age := 20

	gender := "female"

	// we either have to provide the type of var or assign the value so it can infer automatically
	//var a
	//fmt.Println(a)

	a := 100.2
	b := 100

	// will get compile time error about type casting
	//a = b

	a = float64(b)
	fmt.Println(a, b)

	_ = gender

	//ways to declare multiple varaible simultaneously
	car, age := "BMW", 20

	var (
		a1 int

		a2 float32

		a3 string
	)

	var a4, a5, a6 int

	_, _, _, _, _, _ = a1, a2, a3, a4, a5, a6

	sum := 10 + 20

	fmt.Println(name, age, car, sum)
}
