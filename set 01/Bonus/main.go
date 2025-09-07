package main

import (
	"fmt"
	"sync"
)

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

type Cache struct {
	mu   sync.RWMutex
	data map[string]int
}

func (c *Cache) setCache(name string, age int) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	c.data[name] = age
}

func (c *Cache) getCache(name string) (int, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	age, ok := c.data[name]
	return age, ok
}

func main() {
	c := Cache{data: map[string]int{}}
	c.setCache("Namal", 50)
	c.setCache("Sunil", 20)
	c.setCache("Raj", 10)
	c.setCache("Raju", 180)

	if age, ok := c.getCache("Raj"); ok {
		fmt.Println(c.data)
		fmt.Println(age)
	} else {
		fmt.Println("Not found")
	}

}
