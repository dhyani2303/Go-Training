package main

import "fmt"

type Product struct {
	price float64

	productName string
}

func main() {

	quantity := 1
	price := 20.
	name := "Laptop"
	sold := true
	fmt.Printf("Before calling changeValues function\nquantity:%v\nprice:%v\nname:%v\nsold:%v\n", quantity, price, name, sold)

	changeValues(quantity, price, name, sold)

	fmt.Printf("After calling changeValues function\nquantity:%v\n,price:%v\nname:%v\nsold:%v\n", quantity, price, name, sold)

	fmt.Printf("Before calling changeValuesByPointers function\nquantity:%v\nprice:%v\nname:%v\nsold:%v\n", quantity, price, name, sold)

	changeValuesByPointers(&quantity, &price, &name, &sold)

	fmt.Printf("After calling changeValuesByPointers function\nquantity:%v\n,price:%v\nname:%v\nsold:%v\n", quantity, price, name, sold)

	itemDetails := Product{

		productName: "PC",

		price: 200.0,
	}

	fmt.Println("Before calling changeStruct", itemDetails)

	changeStruct(itemDetails)

	fmt.Println("After calling changeStruct", itemDetails)

	fmt.Println("Before calling changeStructByPointers", itemDetails)

	changeStructByPointer(&itemDetails)

	fmt.Println("After calling changeStructByPointers", itemDetails)

}

func changeValues(quantity int, price float64, name string, sold bool) {

	quantity = 10
	price = 100.20
	name = "Mobile"
	sold = false

	fmt.Printf("Inside the changeValues function\nquantity:%v\nprice:%v\nname:%v\nsold:%v\n", quantity, price, name, sold)
}

func changeValuesByPointers(quantity *int, price *float64, name *string, sold *bool) {

	*quantity = 10
	*price = 100.20
	*name = "Mobile"
	*sold = false

	fmt.Printf("Inside the changeValuesByPointers function\nquantity:%v\nprice:%v\nname:%v\nsold:%v\n", *quantity, *price, *name, *sold)
}

func changeStruct(p Product) {

	p.price = 20.

	p.productName = "Laptop"
}

func changeStructByPointer(p *Product) {

	p.price = 20. // equivalent to (*p).price =20.

	p.productName = "Laptop"
}
