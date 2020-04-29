package main

import (
	"fmt"
	"sync"
	"time"
)

// Fork definition
type Fork struct {
	L sync.Mutex
}

// Philosopher definition
type Philosopher struct {
	name string
}

func (ph Philosopher) eat(fork1 *Fork, fork2 *Fork, wg *sync.WaitGroup) {
	defer wg.Done()
	fork1.L.Lock()
	defer fork1.L.Unlock()
	fmt.Println(ph.name, "has grabbed one fork!")

	time.Sleep(time.Second * 2)
	fork2.L.Lock()
	defer fork2.L.Unlock()
	fmt.Println(ph.name, "has grabbed another fork!")
	fmt.Println(ph.name, "is done eating!")
}

func main() {
	var wg sync.WaitGroup
	var f1, f2 Fork
	philosopher1 := Philosopher{"Plato"}
	philosopher2 := Philosopher{"Aristotle"}

	wg.Add(2)
	go philosopher1.eat(&f1, &f2, &wg)
	go philosopher2.eat(&f2, &f1, &wg)
	wg.Wait()
}
