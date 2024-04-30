package main

import "fmt"

type car struct {
	brand string

	price int
}

// this is pass by value i.e. there is no change in the original car1
func (c car) changeCar(brand string, price int) {

	c.brand = brand

	c.price = price
}

func (c *car) changeCarByPointer(name string, price int) {

	c.brand = name // (*c).brand = name this both are same

	c.price = price
}
func main() {

	car1 := car{

		brand: "Audi",
		price: 1000,
	}

	fmt.Println(car1)

	car1.changeCar("Abc", 2000)

	fmt.Println(car1)

	car1.changeCarByPointer("DEF", 2000)

	//can also be written like this. if we write just car1 then go internally converts it to &car1
	//(&car1).changeCarByPointer("CVF",1233)
	fmt.Println(car1)

	var ptr *car

	//method can also be called by this manner
	ptr = &car1
	ptr.changeCarByPointer("GHJ", 12334576)
	fmt.Println(car1)
	(*ptr).changeCarByPointer("J", 12334576)
	fmt.Println(car1)

}
