package main

import (
	"fmt"
	"net/http"
)

func main() {
	// This URL has a typo (.comm vs .com) and will result in an error.
	url := "https://www.google.comm/"

	fmt.Printf("Attempting to fetch: %s\n", url)
	resp, err := http.Get(url)

	// The program enters this block because err is not nil.
	if err != nil {
		fmt.Printf("An error occurred: %v\n", err)
		// The program will continue from here.
		fmt.Println("Program did not panic and will continue.")
		return
	}

	defer resp.Body.Close()
	fmt.Printf("Successfully fetched %s, Status: %s\n", url, resp.Status)
}
