package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func(c chan int) {
		for i := 0; i < 5; i++ {
			fmt.Println("Message sent", i)
			c <- i
		}

		fmt.Println("Sent Messages")
		close(c)
	}(ch)

	fmt.Println("Main going to sleep")

	time.Sleep(time.Second * 2)
	fmt.Println("Main coming out of sleep")

	for v := range ch {
		fmt.Println("MessgaeReceived", v)
	}
}
