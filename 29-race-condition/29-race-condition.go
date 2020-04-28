package main

import (
	"fmt"
)

func main() {
	var data int
	go func() {
		data++
	}()

	// time.Sleep(time.Nanosecond * 1)

	if data == 0 {
		// time.Sleep(time.Nanosecond * 2)
		fmt.Printf("the value is %v.\n", data)
	}
}

// Solving the race condition with WaitGroup
// func main() {
// 	var wg sync.WaitGroup
// 	var data int

// 	wg.Add(1)
// 	go func() {
// 		defer wg.Done()
// 		data++
// 	}()

// 	wg.Wait()
// 	if data == 0 {
// 		fmt.Printf("the value is %v.\n", data)
// 	}
// }
