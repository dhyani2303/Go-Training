package main

import "fmt"

func main() {

	//declaring pointers
	var ptr *float64 // it is uninitialized pointer. value is nil
	fmt.Printf("Type and value of pointer ptr %T %v\n", ptr, ptr)
	x := 100.

	ptr = &x //initialising pointer with the address of x
	fmt.Printf(" Value of pointer ptr and x %v %p\n", ptr, &x)

	p := new(int) // pointer p is initialized and references to an integer whose value is zero

	fmt.Printf("%p", ptr)

	fmt.Printf(" Type and Value of p and value of *p %T  %v %v\n", p, p, *p)

	y := 100

	//another way to initialize the pointer
	ptr1 := &y

	fmt.Printf("Type and value of ptr1: "+
		"%T %v\n", ptr1, ptr1)

	//dereferencing a pointer
	*ptr = 200.

	fmt.Println("*ptr==x:", *ptr == x)

	fmt.Printf("Values of *ptr and x: %v %v\n", *ptr, x)

	//pointer of a pointer

	a := 10

	ptra := &a

	ptr3 := ptra

	fmt.Printf("Type of ptr3: %T , Type of ptra: %T\n", ptr3, ptra)

	fmt.Printf("Value of ptr3 : %v, address of ptra : %p\n", ptr3, &ptra)

	//to access the value of a using ptr3

	fmt.Printf("Value of a using ptr3 %v\n", *ptr3)

	*ptr3++

	fmt.Printf("Value of a is %v\n", a)

	//comparing two pointers

	// two pointers can be same if they are nil or pointing to the same variable

	var1 := 20

	ptr4 := &var1
	ptr5 := &var1

	fmt.Printf("ptr4==ptr5: %v\n", ptr4 == ptr5)

	var2 := 20
	ptr5 = &var2

	//even though the variables value are same but there are two different variables the answer is false
	fmt.Printf("ptr4==ptr5: %v\n", ptr4 == ptr5)

	ptr6 := new(int)

	ptr7 := new(int)
	//since both the pointers point to different memory locations the answer isi false
	fmt.Printf("ptr6==ptr7: %v\n", ptr6 == ptr7)

	var ptr8 *int
	var ptr9 *int

	//since both the pointers are nil the answer is true

	fmt.Printf("ptr8==ptr9: %v\n", ptr8 == ptr9)

	var ptr10 *float64
	_ = ptr10

	//fmt.Printf("ptr9==ptr10",ptr9==ptr10) //compile time error that types doesnt match

}
