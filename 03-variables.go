package main

import "fmt"

func main() {

	// explicit way
	var x string
	x = "Programming Languages Class"

	// implicit
	x2 := "Programming Languages Class"

	// multiple assignments
	var (
		a = 5
		b = 10
		c = 15
	)

	fmt.Println(x, x2, a, b, c)
}
