package main

import "fmt"

func main() {

	type person struct {
		name   string
		age    int
		gender string
	}

	//can also be declared in this way
	//type person struct {
	//	name,gender string,
	//	age int
	//}

	person1 := person{name: "ABC", age: 21, gender: "Female"}
	//	person2 := person{name: "ABC", age: 21, gender:12345} //compile time error for writing int in place of string

	fmt.Println("Person's name is ", person1.name, "age is ", person1.age, "gender is ", person1.gender)

	//fmt.Println(person1.style) //error is raised since style is not present in struct person
	person2 := person{name: "xyz"}

	person2.gender = "female"
	person2.age = 21

	fmt.Println(person2)

	//copying struct
	person3 := person1

	fmt.Println(person3)
	//comparing struct values
	// two struct values are equal if their corresponding fields are equal.
	fmt.Println(person1 == person3)

}
