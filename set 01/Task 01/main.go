package main

import (
	"fmt"
	"sync"
	"time"
)

/*
   	Problem 1: Thread-Safe Counter (15 min)
   •	Implement a thread-safe counter with methods Inc() and Value().
   •	Optional: Implement it using channels instead of mutex.
*/

type Counter struct {
	mu      sync.Mutex
	counter int
}

func (c *Counter) Inc(n int, wg *sync.WaitGroup) {

	defer wg.Done()

	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter++

}

func main() {

	var wg sync.WaitGroup
	c := Counter{}

	start := time.Now()

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go c.Inc(i, &wg)
	}
	wg.Wait()

	fmt.Println(c.counter)
	elapsed := time.Since(start)
	fmt.Println("Processed time, ", elapsed)
}
