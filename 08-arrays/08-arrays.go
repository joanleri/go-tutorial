package main

import "fmt"

func main() {

	// creating array
	var x [5]float64
	x[4] = 50.1
	fmt.Println(x)

	// printing array's length
	fmt.Println(len(x))

	// for loop in arrays
	x = [5]float64{10, 9, 8, 7, 6}
	var total float64 = 0
	for i, value := range x {
		fmt.Println("Adding the value number", i+1)
		total += value
	}
	fmt.Println(total / float64(len(x)))

	// ignoring index
	total = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x))) // Output is 8
}
