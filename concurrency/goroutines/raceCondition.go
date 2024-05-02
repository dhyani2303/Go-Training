package main

import (
	"fmt"
	"sync"
	"time"
)

// if lock is not released then deadlock occurs fatal error: all goroutines are asleep - deadlock!
func main() {

	n := 200

	var wg sync.WaitGroup

	var lock sync.Mutex

	wg.Add(10 * 2)
	for i := 0; i < 10; i++ {

		go func() {

			time.Sleep(time.Second / 10)

			lock.Lock()
			//if two locks and only one unlock deadlock takes place
			//fatal error: all goroutines are asleep - deadlock!
			n++

			lock.Unlock() // if there are two unlock for one lock we get following error
			//fatal error: sync: unlock of unlocked mutex

			wg.Done()
		}()
		go func() {
			time.Sleep(time.Second / 10)

			lock.TryLock()
			n--
			lock.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(n)
}
