package main

import "fmt"

func main() {

	fmt.Println("Dividing 2 by 1", division(2, 1))
	fmt.Println("Dividing 2 by 0 ", division(2, 0))
	fmt.Println("Dividing 2 by 2 ", division(2, 2))
}

func division(a, b int) int {
	defer func() {
		r := recover()

		if r != nil {
			fmt.Println(r)
		}
	}()
	if b == 0 {
		panic("division by zero")
	}
	defer func() {
		fmt.Println("Another defer")
	}()

	return a / b

}
