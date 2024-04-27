package main

import "fmt"

func main() {

	i := 0

loop:
	if i < 5 {
		fmt.Println(i)
		i++

		goto loop
	}

	// this is not allowed
	//	goto todo
	//	x := 10
	//todo:
	//
	//	fmt.Println(x)
	//	goto todo  // this is allowed
}
