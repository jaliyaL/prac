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

}

// func (c *Counter) Value() {

// }

func main() {

	var Counter int = 0

	start := time.Now()
	for i := 0; i < 1000; i++ {
		Counter++
	}
	fmt.Println(Counter)
	elapsed := time.Since(start)
	fmt.Println("Processed time, ", elapsed)

}
