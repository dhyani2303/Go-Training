package main

import "fmt"

type age int
type oldAge int
type string1 string

func main() {

	type speed uint
	var s1 speed = 10
	var s2 speed = 20

	fmt.Println(s2 - s1)

	var x uint

	//x =s1 // this will give error even though underlying types are same this is incorrect

	x = uint(s1)
	fmt.Println(x)

	var s3 speed = speed(x)
	_ = s3

	var a1 age = 20
	var o1 oldAge = 30

	var x1 string1 = "1234"
	_ = x1

	//fmt.Println(o1-a1) // this cannot be done
	fmt.Println(o1 - oldAge(a1))

	//fmt.Println(o1 - oldAge(x1)) //this is not possible since underlying type are different
	fmt.Println(o1 - oldAge(s1))

}
