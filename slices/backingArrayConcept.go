package main

import "fmt"

func main() {

	cars := []string{"Ford", "Honda"}

	newCars := []string{}

	newCars = append(newCars, cars...)

	cars[0] = "Nissan"
	fmt.Println(cars, newCars)

	newCars = cars[:]

	fmt.Println(newCars, cars)

	s1 := []int{10, 20, 30, 40, 50}
	newSlice := s1[0:3]

	fmt.Println(len(newSlice), cap(newSlice))

	newSlice1 := s1[0:5]
	//newSlice1[0] = 4
	//fmt.Println(newSlice[2], newSlice1[2])
	fmt.Println(len(newSlice1), cap(newSlice1))

	newSlice2 := s1[3:4]
	fmt.Println(len(newSlice2), cap(newSlice2))
}
