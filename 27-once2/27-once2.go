package main

import (
	"fmt"
	"sync"
)

func main() {

	var count int
	var wg sync.WaitGroup
	var once sync.Once

	increment := func() {
		count++
	}

	reset := func() {
		count = 0
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		once.Do(increment)
	}()

	wg.Wait()

	wg.Add(1)
	go func() {
		defer wg.Done()
		once.Do(reset)
	}()

	wg.Wait()
	fmt.Println("Count is", count)
}
