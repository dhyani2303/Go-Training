package main

import "fmt"

type name []string

//type distance *int
//method cannot be created of named pointer
//func (d distance) print(){
//
//}

func (n name) print() {

	// n is called method's receiver
	// n is the actual copy of the names we're working with and is available in the function.
	// n is like this or self from OOP.
	// any variable of type names can call this function on itself like variable_name.print()

	for i, value := range n {

		fmt.Println(i, value)
	}
}

func main() {

	friends := name{"ABC", "DEF"}

	//calling the method
	friends.print()

	//alternatively method can be called by following way
	name.print(friends) //not very common

}
