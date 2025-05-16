package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex

	// Shared resource
	counter := 0

	// Start 2 goroutines that try to access the same resource
	for i := 1; i <= 10; i++ {
		go func(id int) {
			mu.Lock() // 👈 Only one goroutine can lock at a time
			fmt.Printf("Goroutine %d is accessing the counter\n", id)
			counter++
			time.Sleep(2 * time.Second)
			fmt.Printf("Goroutine %d is done. Counter = %d\n", id, counter)
			mu.Unlock()
		}(i)
	}

	// Wait for goroutines to finish
	time.Sleep(25 * time.Second)
}
