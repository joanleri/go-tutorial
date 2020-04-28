package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	c := sync.NewCond(&sync.Mutex{})
	var name string
	const numberOfProcesses = 3

	wg.Add(numberOfProcesses)
	for i := 0; i < numberOfProcesses; i++ {
		go func() {
			defer wg.Done()
			c.L.Lock()
			fmt.Println("Waiting to be signaled...")
			c.Wait()
			c.L.Unlock()
			fmt.Println("Greetings", name)
		}()
	}

	time.Sleep(time.Second * 1)
	fmt.Println("Making goroutines wait for 5 seconds...")
	time.Sleep(time.Second * 5)
	c.L.Lock()
	name = "John Bachus"
	c.L.Unlock()

	c.Broadcast()
	wg.Wait()

}
