package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := make(chan string)

	ch2 := make(chan string)

	go func() {

		time.Sleep(2 * time.Second)

		ch1 <- "Hello from first channel"
	}()

	go func() {

		time.Sleep(2 * time.Second)

		ch2 <- "Hello from second channel"
	}()

	time.Sleep(2 * time.Second)

	for i := 0; i < 2; i++ {

		select {

		case msg1 := <-ch1:
			fmt.Println(msg1)

		case msg2 := <-ch2:

			fmt.Println(msg2)
		default:
			fmt.Println("No data")
		}
	}

}
