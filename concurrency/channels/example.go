package main

import (
	"fmt"
	"sync"
)

func main() {

	//ch := make(chan string)

	var wg sync.WaitGroup
	//
	wg.Add(2)
	//
	//go func() {
	//	defer wg.Done()
	//	ch <- "Hello"
	//}()
	//
	//go func() {
	//	defer wg.Done()
	//	fmt.Println(<-ch)
	//}()
	//

	ch := make(chan []int)
	// concurrent map writes error
	map1 := make(map[int]int)
	_ = map1
	slice := []int{}

	go func() {
		for i := 0; i < 1000; i++ {

			slice = append(slice, i)

		}
		wg.Done()
		ch <- slice
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			slice = append(slice, i)

		}
		wg.Done()
		ch <- slice
	}()

	wg.Wait()

	fmt.Println(slice)

}
