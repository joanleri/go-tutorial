package main

import (
	"fmt"
	"sync"
	"time"
)

// PolitePhilosopher definition
type PolitePhilosopher struct {
	name      string
	isHungry  bool
	companion *PolitePhilosopher
}

func (ph *PolitePhilosopher) eat(fork *Fork, wg *sync.WaitGroup) {
	defer wg.Done()

	for ph.isHungry {
		fork.L.Lock()
		defer fork.L.Unlock()
		fmt.Printf("%s is ready to eat...\n", ph.name)

		if ph.companion.isHungry {
			fork.L.Unlock()
			fmt.Printf("%s leaves the fork so %s can eat...\n", ph.name, ph.companion.name)
		} else {
			ph.isHungry = false
		}
		time.Sleep(time.Second * 1)
	}

	fmt.Printf("%s is done eating...\n", ph.name)
}

// Fork definition
type Fork struct {
	L sync.Mutex
}

func main() {
	var wg sync.WaitGroup
	philosopher1 := PolitePhilosopher{name: "Plato", isHungry: true}
	philosopher2 := PolitePhilosopher{name: "Aristotle", isHungry: true}
	var fork Fork

	// setting companions
	philosopher1.companion = &philosopher2
	philosopher2.companion = &philosopher1

	// starting dinner
	wg.Add(2)
	go philosopher1.eat(&fork, &wg)
	go philosopher2.eat(&fork, &wg)

	// dinner ended
	wg.Wait()
	fmt.Println("Dinner has ended!")
}
