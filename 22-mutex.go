package main

import (
	"fmt"
	"sync"
)

func increment(lock *sync.Mutex, countPtr *int) {
	lock.Lock()         // goroutine has exclusive use of the critical section by mutual exclusion
	defer lock.Unlock() // goroutine releases access to critical section
	*countPtr++
	fmt.Println("Incrementing", *countPtr)
}

func decrement(lock *sync.Mutex, countPtr *int) {
	lock.Lock()
	defer lock.Unlock()
	*countPtr--
	fmt.Println("Decrementing", *countPtr)
}

func main() {

	var wg sync.WaitGroup
	var lock sync.Mutex
	counter := 0

	// Increment
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			increment(&lock, &counter)
		}()
	}

	// Decrement
	for i := 0; i <= 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			decrement(&lock, &counter)
		}()
	}

	wg.Wait()
	fmt.Println("Operations complete")
}
