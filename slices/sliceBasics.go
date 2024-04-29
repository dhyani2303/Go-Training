package main

import "fmt"

func main() {

	// declaring a string slice, by default is initialized with nil or uninitialized

	var cities []string
	fmt.Println("cities is equal to nil: ", cities == nil)

	// we can not assign elements to nil slice:
	// cities[0] = "Paris" // -> runtime error

	// declaring a slice using a slice literal
	numbers := []int{2, 3, 4, 5} // on the right hand-side of the equal sign is a slice literal
	fmt.Println(numbers)         // => [2 3 4 5]

	nums := make([]int, 2) //2 is the size of slice

	fmt.Println(nums)

	// declaring a slice using a slice literal
	type names []string
	friends := names{"A", "B"}

	fmt.Println(friends)

	// getting an element from the slice
	x := numbers[0]
	fmt.Println("x is", x) // => x is 2

	// modifying an element of the slice
	numbers[1] = 200
	fmt.Printf("%#v\n", numbers) // => []int{2, 200, 4, 5}

	//iterating over a slice
	for index, value := range numbers {
		fmt.Printf("index: %d, value: %v\n", index, value)

	}
	// uninitialized slice, equal to nil
	var nn []int
	fmt.Println(nn == nil) // true

	// empty slice but initialized, not equal to nil
	mm := []int{}
	fmt.Println(mm == nil) //false

	// we can not compare slices using the equal (=) operator
	// fmt.Println(nn == mm) //error -> slice can only be compared to nil

	// to compare 2 slices use a for loop to iterate over the slices and compare element by element
	// it's also necessary to check the length of slices (if a is nil it doesn't enter the for loop)

	a, b := []int{1, 2, 3}, []int{1, 2, 3}

	var eq bool = true
	if len(a) != len(b) {
		eq = false
	}
	for i, valueA := range a {
		if valueA != b[i] {
			eq = false
			break
		}
	}
	if eq {
		fmt.Println("a and b slices are equal")
	} else {
		fmt.Println("a and b slices are not equal")
	}
}
