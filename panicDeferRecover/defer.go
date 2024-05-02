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
	var i int = 0
	defer fmt.Println(i)
	i++

}
func b() {

	for i := 0; i < 4; i++ {

		defer fmt.Print(i, " ")
	}
}

func c() (i int) {

	defer func() { i++ }()
	fmt.Println("Value of  c", i)
	return 1
}
