package main

import "fmt"

func main() {
	i := 1
	switch i {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
		// fallthrough
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("A really big number")
	}
}
