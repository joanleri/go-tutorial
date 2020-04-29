package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	counter1 := 0
	counter2 := 0

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("1st goroutine is working very hard...")
		for i := 0; i < 100; i++ {
			time.Sleep(time.Millisecond * 20)
			counter1++
		}
		fmt.Println("1st goroutine has finished...")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("2nd goroutine is working very hard...")
		for i := 0; i < 100; i++ {
			counter2++
		}
		fmt.Println("2nd goroutine has finished...")
	}()

	wg.Wait()
	fmt.Println("All goroutines finished.")
	fmt.Println("Counter 1:", counter1, "Counter 2:", counter2)
}
