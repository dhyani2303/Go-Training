package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)

	}
	//creating while loop using for loop

	j := 10
	for j < 0 {
		fmt.Println(j)
		j--
	}

	//to create infinite loop
	//c := 0
	//for {
	//	c++
	//}
	//fmt.Println(c) // this statement is unreachable

	// having multiple variables in a single for loop

	for i, j := 0, 100; i < 10; i, j = i+1, j+1 {
		fmt.Println(i, j)

	}

}
