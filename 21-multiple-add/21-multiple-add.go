package main

import (
	"fmt"
	"sync"
)

func hello(wg *sync.WaitGroup, id int) {
	defer wg.Done()
	fmt.Println("Hello from", id)
}

func main() {
	const numGreeters = 5
	var wg sync.WaitGroup
	wg.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}
