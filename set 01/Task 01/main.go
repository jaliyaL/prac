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
	counter int
}

func (c *Counter) Inc() {
	c.counter++
}

func (c *Counter) Value() int {

	return c.counter
}

func main() {

	var wg sync.WaitGroup
	c := Counter{}

	const numIterrations = 1000

	ch := make(chan int, numIterrations)
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
