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

	// amt := time.Duration(rand.Intn(100))
	// time.Sleep(time.Microsecond * amt)
	fmt.Println("Here we Go!")
	var input string
	fmt.Scanln(&input)
}
