package main

import (
	"fmt"
	"net/http"
)

func main() {

	urlArray := []string{"https://gobyexample.com", "https://www.google.comm/", "https://edition.cnn.com/"}

	for _, value := range urlArray {
		resp, err := http.Get(value)
		if err != nil {
			// defer func() {
			// 	if r := recover(); r != nil {
			// 		fmt.Println("recoverd from", r)
			// 	}
			// }()
			fmt.Printf("An error occurred: %v\n", err)
			// The program will continue from here.
			fmt.Println("Program did not panic and will continue.")
			continue
		}
		defer resp.Body.Close()
		fmt.Println("Response from -", value, resp.Status)
	}
}
