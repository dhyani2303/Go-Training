package main

import "fmt"

func main() {

	// anonymous struct
	dhyani := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Dhyani",
		lastName:  "Patel",
		age:       21,
	}
	fmt.Println(dhyani)

	fmt.Printf("%#v\n", dhyani)

	//anonymous fields
	type Book struct {
		string
		float64
		bool
	}
	b1 := Book{"1984 by George Orwell", 10.2, false}

	fmt.Printf("%#v\n", b1)

	fmt.Println(b1.string)

	//there can be mixture of anonymous field and named fields
	type Employee1 struct {
		name   string
		salary int
		bool
	}

	e := Employee1{"John", 40000, false}

	_ = e
}
