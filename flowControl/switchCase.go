package main

import "fmt"

func main() {

	language := "golang"

	switch language {
	case "Java":
		fmt.Println("Java")

	case "Go", "golang":

		fmt.Println("Go")

	default:
		fmt.Println("Every language has its own perks")

	}

	x := 10

	// writing switch alone means true so we will have cases that will return boolean value

	//switch true // writing switch with true and without it is same
	switch {

	case x%2 == 0:

		fmt.Println("Even Number!")

	case x%2 == 0:
		fmt.Println("Odd number")

	}

}
