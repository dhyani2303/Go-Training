package main

import "fmt"

func main() {

	var numbers [4]int

	fmt.Printf("%#v\n", numbers)

	var a1 = [4]float64{}

	var a2 = [3]int{5, -3, 5}

	//initializing only the first 3 elements
	a3 := [4]string{"A", "B", "C"}

	// the ellipsis operator (...) finds out automatically the length of the array
	a4 := [...]int{1, 4, 5}

	// declare an array on multiple lines for better readability
	a5 := [...]int{
		1,
		2,
		3,
	} //the ending comma is mandatory when initializing the array on multiple lines and the closing curly brace is on its own line

	_, _, _, _, _ = a1, a2, a3, a4, a5

	// we can't add or remove elements from the array (it's fixed length)

	numbers[0] = 7 //changing first element (index 0)

	fmt.Printf("%v\n", numbers) // -> [7 0 0 0]

	// compile-time error
	// numbers[5] = 8  // invalid array index 5 (out of bounds for 4-element array)

	// getting an element
	x := numbers[0]

	fmt.Println("x is ", x) // => x is  7

	//two ways to iterate through an array

	for i, v := range numbers {

		fmt.Println("index:", i, "value:", v)
	}
	for i := 0; i < len(numbers); i++ {

		fmt.Println("index:", i, "values:", numbers[i])
	}
	// declaring a multi-dimensional arrays (array of arrays or matrix)
	balances := [2][3]int{
		[3]int{5, 6, 7}, //[3]int is optional
		{8, 9, 10},
	}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println("")
	}

	//  = operator creates a copy of an array.
	// the arrays are not connected and are saved in different memory locations
	m := [3]int{1, 2, 3}
	n := m //n is a copy of m

	fmt.Println("n is equal to m: ", n == m) // => true
	m[0] = -1                                //only m is modified
	fmt.Println("n is equal to m: ", n == m) // => false
}
