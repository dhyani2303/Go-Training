package main

import "fmt"

func main() {
	f1(10, 20, 30, 2., 3.)

	s := f2()

	fmt.Println(s)

	sum := f3(2, 3)

	fmt.Println(sum)
}

// a function can return multiple values
func f1(a, b, c int, d, e float64) {
	fmt.Println(a, b, c, d, e)
}
func f2() int {

	return 100
}

func f3(a int, b int) (s int) {

	return a + b

}
