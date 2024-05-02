package main

import (
	"fmt"
)

// if the channel is uninitialised and value is sent then deadlock occurs
// Closing a nil or an already closed channel produces a panic in the current goroutine.
// panic: close of nil channel

//ending a value to a closed channel also produces a panic in the current goroutine.
//panic: send on closed channel

func main() {

	var channel chan int
	//bidirectional channel - it is initialized channel
	ch := make(chan int)

	//receiver channel
	ch1 := make(<-chan int)

	_, _ = ch1, channel

	//send only channel
	ch2 := make(chan<- int)
	_ = ch2

	fmt.Printf("%T\n", ch)
	close(ch)
	fmt.Printf("%T\n", ch)

	go func(n int, c chan<- int) {

		var sum int

		for i := 0; i < n; i++ {
			sum += i
		}
		c <- sum
	}(10, ch)

	//close(channel)

	output, boolean := <-ch //the boolean value is whether the message was received before the chanel was closed or not. True for chanel was not closed

	fmt.Println(output, boolean)
}
