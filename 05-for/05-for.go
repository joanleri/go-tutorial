package main

import "fmt"

func main() {
	// normal for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// while loop
	for i := 1; i <= 80; {
		fmt.Println(i)
		i *= 2
	}

	// infinite while loop
	for {
		fmt.Println("Infinity")
	}
}
