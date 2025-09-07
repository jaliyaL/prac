package main

import "fmt"

/*  Implement a simple caching mechanism using a map and mutex. */

/*
🔑 Concepts
Map for storage
Go’s built-in map is perfect for storing key-value pairs.
Example: map[string]string or map[string]int.

Mutex for safety
Maps in Go are not safe for concurrent access.
If multiple goroutines read/write a map at the same time, it will panic.
Use sync.Mutex to lock the map while reading or writing.
*/

func main() {
	fmt.Println("hi")
}
