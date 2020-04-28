package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ExpensiveStruct definition
type ExpensiveStruct struct {
	a int
}

func printRandomInt(es *ExpensiveStruct) { fmt.Println(es.a) }

func main() {

	var wg sync.WaitGroup
	var structsCreated int
	const numberOfProcesses = 1000

	myPool := &sync.Pool{
		New: func() interface{} {
			anExpensiveStruct := ExpensiveStruct{a: rand.Intn(1000)}
			structsCreated++
			return anExpensiveStruct
		},
	}

	wg.Add(numberOfProcesses)
	for i := 0; i < numberOfProcesses; i++ {
		go func() {
			defer wg.Done()
			instance := myPool.Get().(ExpensiveStruct)
			printRandomInt(&instance)
			myPool.Put(instance)
		}()
		time.Sleep(time.Microsecond * 1)
	}

	wg.Wait()
	fmt.Println(structsCreated, "expensive instances were created")

}
