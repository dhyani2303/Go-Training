package main

import (
	"fmt"
	"math"
)

// interface contains the name of the methods with its return type
type shape interface {
	area() float64

	parameter() float64
}

type rectangle struct {
	height, width float64
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {

	return r.height * r.width
}
func (r rectangle) parameter() float64 {

	return 2 * (r.height + r.width)
}
func (c circle) area() float64 {

	return math.Pi * c.radius * c.radius
}

func (c circle) parameter() float64 {

	return 2 * math.Pi * c.radius
}

func printShape(s shape) {

	fmt.Printf("Shape: %#v\n", s)

	fmt.Println("Parameter:", s.parameter())

	fmt.Println("Area:", s.area())
}
func (c circle) volume() float64 {

	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}
func main() {

	var s shape

	fmt.Printf("Type of s %T\n", s) //value is nil

	//calling method from the interface variable will cause runtime error
	//s.parameter()

	rectangle1 := rectangle{

		height: 10.0,
		width:  2.0,
	}

	//this is valid
	//var s1 shape = rectangle1

	//	fmt.Printf("Type of s1: %T\n", s1)
	//	fmt.Println("Value of s1 is", s1)

	fmt.Printf("%T\n", rectangle1) //rectangle

	printShape(rectangle1)

	circle1 := circle{
		radius: 20.0,
	}

	//here we have dynamically changed the type of s from rectangle to circle.
	// this is not possible in variables but possible only in interface
	//i.e interface has dynamic types and can be changed during runtime
	s = circle1

	fmt.Printf("Type of s %T\n", s)

	printShape(circle1)

	//assertion and switch types

	var s1 shape

	//if we try to assert nil assert then it won't be a successful assertion
	s1 = circle1

	//s.volume() // this is not allowed

	//asserting s to circle
	ball, ok := s1.(circle)

	if ok == true {
		fmt.Println("Circle:", ball.volume())
	} else {
		fmt.Println("Something else is wrong")
	}

	//switch types

	s1 = rectangle1

	switch value := s1.(type) {

	case circle:
		fmt.Printf("%#v\n", value)

	case rectangle:
		fmt.Printf("%#v\n", value)
	}

}
