package main

import (
	"fmt"
	"sync"
	"time"
)

func increment(wg *sync.WaitGroup, lock *sync.RWMutex, countPtr *int) {
	lock.Lock()
	defer lock.Unlock()
	defer wg.Done()
	fmt.Println("Getting lock for writing.")
	*countPtr++
	time.Sleep(time.Second * 5)
	fmt.Println("Releasing lock for writing.")
}

func read(wg *sync.WaitGroup, lock *sync.RWMutex, counterPtr *int, id int) {
	lock.RLock()
	defer lock.RUnlock()
	defer wg.Done()
	fmt.Println("Process", id, "-", "reading counter: ", *counterPtr)
	time.Sleep(time.Second * 1)
	fmt.Println("Process", id, "-", "releasing lock for reading.")
}

func main() {
	var wg sync.WaitGroup
	var lock sync.RWMutex
	const iterations = 5
	counter := 0

	// Writer
	for i := 0; i <= iterations; i++ {
		wg.Add(1)
		go increment(&wg, &lock, &counter)
	}

	// Readers
	const readers = 3
	for i := 0; i < iterations; i++ {
		wg.Add(readers)
		for j := 0; j < readers; j++ {
			go read(&wg, &lock, &counter, j+1)
		}
		wg.Wait()
	}
	wg.Wait()
	fmt.Println("Operations done")
}
