package main

import (
	"fmt"
	"sync"
)

// which go routine will be executed first is not fixed
// if the no of wg.Done  wg.Add(n) then we get the following error
// panic: sync: WaitGroup is reused before previous Wait has returned

// if n i.e add(n) is negative then -panic: sync: negative WaitGroup counter

// if there is panic without recover the program will stop abruptly
func f1(wg *sync.WaitGroup) {

	defer func() {
		r := recover()
		if r != nil {
			fmt.Println(r)

			//	wg.Done() //in this case if panic does not take place then wg.Done() will never be called hence there will be deadlock

		}
		wg.Done() // best place to define Done() method
	}()
	for i := 0; i < 10; i++ {
		if i%2 == 1 {
			panic("Ohhh no!")
		}
		fmt.Printf("%d\n", i)
	}
	//wg.Done() // if wg.Done() is written outside the defer function i.e at the end of the function then if panic occurs the program will not process this
	// statement and hence deadlock will take place

}
func f2(wg *sync.WaitGroup) {

	defer func() {
		wg.Done()
	}()
	for i := 10; i < 20; i++ {
		fmt.Println(i)
	}

}
func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go f1(&wg)
	go f2(&wg)

	wg.Wait()
	fmt.Println("Exiting main")

}
