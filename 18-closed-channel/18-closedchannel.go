package main

import "fmt"

func anotherPinger(c chan<- string, repetitions int) {
	defer close(c)
	for i := 0; i < repetitions; i++ {
		c <- "ping"
	}
}

func anotherPrinter(c chan string) {
	defer fmt.Println("(Done printing!)")
	for message, ok := <-c; ok; {
		fmt.Print(message, " ")
		message, ok = <-c
	}
}

func main() {
	var c chan string = make(chan string)
	go anotherPinger(c, 5)
	go anotherPrinter(c)

	var input string
	fmt.Scanln(&input)
}
