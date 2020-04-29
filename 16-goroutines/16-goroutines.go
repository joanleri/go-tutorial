package main

import (
	"fmt"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Print("(", n, ":", i, ") ")
	}
}

func main() {
	for i := 0; i < 5; i++ {
		go f(i)
	}

	// time.Sleep(time.Microsecond * 100)
	fmt.Println("Here we Go!")
	var input string
	fmt.Scanln(&input)
}
