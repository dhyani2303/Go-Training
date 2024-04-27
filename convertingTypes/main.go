package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x = 9
	var y = 10.0

	//on converting y to int the actual type of y is unchanged,instead a new value is returned when conversion is done
	fmt.Println(x * int(y))
	fmt.Printf("%T\n", y)

	//In Go types with different names are different types.
	var c int = 5
	var d int64 = 7

	//int and int64 are not the same type
	fmt.Printf("%T, %T\n", c, d)

	//string conversion to int and float

	s := string(99)

	fmt.Println(s) //Ascii code of 99 is returned

	//s1:=string(32.99) // this is not allowed

	//to convert float to string we have Sprintf()
	var s2 = fmt.Sprintf("%f", 5.12)
	fmt.Println(s2)        // value will be 5.12
	fmt.Printf("%T\n", s2) // string is the type

	//same can be applied for int

	var s3 = fmt.Sprintf("%d\n", 33)

	fmt.Printf("%T %s\n", s3, s3)

	//we have  strconv that provides 5 methods - ParseFloat,ParseInt,ParseBool,ParseComplex,ParseUint

	var result, err = strconv.ParseFloat("3.1642", 64)
	_ = err
	fmt.Println(result)
	var result2, err2 = strconv.ParseInt("2", 8, 8)

	_ = err2
	fmt.Println(result2)

	//Atoi(String(ascii to int) and Itoa (int to string (Ascii))

	i, err3 := strconv.Atoi("-50")
	_ = err3
	s = strconv.Itoa(20)
	fmt.Printf("i Types is %T, i value is %v\n", i, i)

	fmt.Printf("s Types is %T, s value is %v\n", s, s)
}
