package main

import (
	"fmt"
	"math"
)

func main() {

	var a uint8 = 255
	var b int8 = 127
	var d int8 = -128
	d--
	a++
	b++

	var c float32 = math.MaxFloat32
	c = c * 1.2

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
