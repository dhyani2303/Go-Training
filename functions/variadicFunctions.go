package main

import "fmt"

//Variadic functions are functions that take a variable number of arguments.
// Ellipsis prefix (three-dots) in front of the parameter type makes a function variadic.
// The function may be called with zero or more arguments for that parameter.
// If the function takes parameters of different types, only the last parameter of a function can be variadic.

func main() {

	maps := make(map[int]int)

	maps[1] = 2

	mapAsParameter(maps)

	fmt.Println("Map values after the function is returned", maps)

	func1(1, 2, 3, 4)

	func1() //this is also valid since variadic function takes 0 or more arguments

	nums := []int{1, 2}

	func2(nums...)

	fmt.Println(nums) //the value of nums[0] is updated

	sum, product := sumAndProduct(1, 2, 3, 4, 5)

	fmt.Println("Sum is ", sum, "and Product is ", product)
}

func func1(a ...int) {

	fmt.Printf("%T\n", a)

	fmt.Printf("%#v\n", a)

}

func func2(a ...int) {
	a[0] = 50
}
func sumAndProduct(a ...float64) (float64, float64) {

	sum := 0.

	product := 1.0

	for _, v := range a {

		sum += v

		product *= v

	}
	return sum, product
}

func mapAsParameter(x map[int]int) {

	x[0] = 100
}
