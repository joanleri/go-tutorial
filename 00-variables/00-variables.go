package main

import "fmt"

func main() {

	// explicit way
	var x string
	x = "Programming Languages Class"

	// this is equivalent to
	var x2 string = "Programming Languages Class"

	// implicit
	x3 := "Programming Languages Class"

	// multiple declarations
	var d, e, f int
	fmt.Println(d, e, f)

	// multiple assignments
	var (
		a = 5
		b = 10
		c = 15.0
	)

	fmt.Println(x, x2, x3, a, b, c)
}
