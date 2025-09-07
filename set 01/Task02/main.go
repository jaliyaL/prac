package main

import (
	"fmt"
	"net/http"
	"sync"
)

/*
Problem 2: Concurrent URL Fetcher (15 min)
•	Write a function to fetch status codes for a list of URLs concurrently.
•	Handle errors gracefully.
•	Discuss alternatives like sync.WaitGroup and buffered channels.
*/

func main() {

	urlArray := []string{
		"https://gobyexample.com",
		"https://www.google.com/",
		"https://edition.cnn.comm/",
	}

	var wg sync.WaitGroup
	ch := make(chan string, len(urlArray))

	for _, url := range urlArray {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			resp, err := http.Get(u)
			if err != nil {
				fmt.Println("error: ", err)
				return
			}
			defer resp.Body.Close()
			ch <- fmt.Sprintf("Url : %s and Status : %s", u, resp.Status)
		}(url)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for res := range ch {
		fmt.Println("response is ", res)
	}

}
