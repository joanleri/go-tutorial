package main

import (
	"fmt"
	"sync"
	"time"
)

// HungryPhilosopher definition
type HungryPhilosopher struct {
	name       string
	bitesTaken int
}

// Eat in a polite manner
func (ph *HungryPhilosopher) eatPolitely(fork *Fork, wg *sync.WaitGroup) {
	defer wg.Done()

	// eating for 5 seconds
	for begin := time.Now(); time.Since(begin) <= time.Second*5; {
		fork.L.Lock()
		ph.bitesTaken++
		fork.L.Unlock()
		time.Sleep(time.Millisecond) // chewing time
	}

	fmt.Printf("%s took %d bites.\n", ph.name, ph.bitesTaken)
}

// Eat in a desperate manner
func (ph *HungryPhilosopher) eatHungrily(fork *Fork, wg *sync.WaitGroup) {
	defer wg.Done()

	// eating for 5 seconds
	for begin := time.Now(); time.Since(begin) <= time.Second*5; {
		fork.L.Lock()
		ph.bitesTaken++
		time.Sleep(time.Millisecond) // chewing time
		fork.L.Unlock()
	}

	fmt.Printf("%s took %d bites.\n", ph.name, ph.bitesTaken)
}

// Fork definition
type Fork struct {
	L sync.Mutex
}

func main() {
	var wg sync.WaitGroup

	var fork Fork
	philosopher1 := HungryPhilosopher{name: "Plato"}
	philosopher2 := HungryPhilosopher{name: "Aristotle"}

	// starting dinner
	wg.Add(2)
	go philosopher1.eatPolitely(&fork, &wg)
	go philosopher2.eatHungrily(&fork, &wg)

	// ending dinner
	wg.Wait()
	fmt.Println("Dinner has ended!")
}
