package main

import (
	"fmt"
	"sync"
	"time"
)

// AnInteger definition
type AnInteger struct {
	L     sync.Mutex
	value int
}

func main() {
	var wg sync.WaitGroup

	printSum := func(i1, i2 *AnInteger) {
		defer wg.Done()
		i1.L.Lock()
		defer i1.L.Unlock()

		time.Sleep(time.Second * 2)
		i2.L.Lock()
		defer i2.L.Unlock()

		fmt.Printf("i1 + i2 = %v\n", i1.value+i2.value)
	}

	var a, b AnInteger
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}
