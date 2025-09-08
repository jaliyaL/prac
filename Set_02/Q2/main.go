package main

import (
	"fmt"
	"sync"
)

/*
Problem 2: Implement a worker pool where multiple
goroutines process tasks concurrently.
*/
func worker(workers int, jobs <-chan int, wg *sync.WaitGroup) {

	for j := range jobs {
		defer wg.Done()
		fmt.Printf("worker %d got Job %d\n ", workers, j)
	}
}

func main() {

	var wg sync.WaitGroup
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)

	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		wg.Add(1)
		jobs <- j
	}
	close(jobs)

	wg.Wait()
}
