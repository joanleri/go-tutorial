package main

import "fmt"

func main() {
	// normal for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// while loop
	i := 10
	for i <= 80 {
		fmt.Println(i)
		i *= 2
	}
}
