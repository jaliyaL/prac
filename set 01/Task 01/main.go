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

func (c *Counter) Inc() {

	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter++
}

func (c *Counter) Value() int {

	c.mu.Lock()
	defer c.mu.Unlock()

	return c.counter
}

func main() {

	var wg sync.WaitGroup
	c := Counter{}

	const numIterrations = 1000

	start := time.Now()

	for i := 0; i < numIterrations; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()

	fmt.Println(c.Value())
	elapsed := time.Since(start)
	fmt.Println("Processed time, ", elapsed)
}
