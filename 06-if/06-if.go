package main

import "fmt"

func main() {

	i := 11

	if i%2 == 0 {
		fmt.Println("number is divisible by 2")
	} else if i%3 == 0 {
		fmt.Println("number is divisible by 3")
	} else {
		fmt.Println("number is not divisible by 2 or 3")
	}
}
