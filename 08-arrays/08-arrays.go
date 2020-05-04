package main

import "fmt"

func main() {

	// creating array
	fmt.Println("Creating and array")
	var x [5]float64
	fmt.Println(x) //  initialized with zero values
	x[4] = 50.1    // accessing elements
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	// printing array's length
	fmt.Println("\n\n\nArray's length")
	fmt.Println(len(x))

	// iterating over an array
	fmt.Println("\n\n\nIterating over an array")
	x = [5]float64{10, 9, 8, 7, 6} // type inference in arrays
	var total float64 = 0
	for i, value := range x {
		fmt.Println("Adding the value number", i+1)
		total += value
	}
	fmt.Println(total / float64(len(x)))

	// ignoring index while iterating
	fmt.Println("\n\n\nIgnoring index while iterating")
	total = 0
	for _, value := range x {
		total += value
	}
	fmt.Println(total / float64(len(x))) // Output is 8

	// questions
	// x = [6]float64{1, 2, 3, 4, 5, 6}
}
