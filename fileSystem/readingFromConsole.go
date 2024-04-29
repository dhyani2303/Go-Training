package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	//fmt.Println("Enter the first number")
	//
	//scanner.Scan()
	//
	//firstNumber, err := strconv.Atoi(scanner.Text())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Enter the second number")
	//
	//scanner.Scan()
	//
	//secondNumber, err := strconv.Atoi(scanner.Text())
	//if err != nil {
	//	log.Fatal(err)
	//}
	//sum := firstNumber + secondNumber
	//
	//fmt.Printf("The sum of the first number is: %d", sum)

	//to read multiple lines

	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println("You entered ", text)
		if text == "exit" {
			fmt.Println("Bye bye")
			break

		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
