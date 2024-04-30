package main

import "fmt"

func main() {
	fmt.Println("Before calling a")
	a()
	fmt.Println("After calling a")
	b()

	fmt.Println("Value returned from c", c())
}
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	defer fmt.Println(i)

	return
}
func b() {

	for i := 0; i < 4; i++ {

		defer fmt.Print(i, " ")
	}
}

func c() (i int) {

	defer func() { i++ }()
	return 1
}
