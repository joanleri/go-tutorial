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

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			once.Do(increment)
		}()
	}

	wg.Wait()
	fmt.Println("Count is", count)
}
