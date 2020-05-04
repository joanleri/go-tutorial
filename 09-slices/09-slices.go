package main

import "fmt"

func main() {

	// declaring a slice variable
	var x []float64
	fmt.Printf("%T\n", x)

	// creating a slice literal
	x = make([]float64, 5)
	fmt.Printf("%T\n", x)

	// accessing individual elements
	fmt.Println("\n\n\naccessing individual elements...")
	x[0] = 1
	fmt.Println(x)
	x[1] = 10
	fmt.Println(x)

	// selecting the underlying array capacity
	fmt.Println("\n\n\nselecting the underlying array capacity...")
	var x2 []float64 = make([]float64, 5, 10)
	fmt.Println(x2)
	fmt.Println("length:", len(x2))
	fmt.Println("capacity:", cap(x2))

	// creating slices from an array
	fmt.Println("\n\n\ncreating slices from an array...")
	arr := [5]float64{1, 2, 3, 4, 5}
	x3 := arr[0:5]
	x4 := arr[1:3]
	fmt.Println(x3)
	fmt.Println(x4)
	x4[0] = -10
	fmt.Println(x3)
	fmt.Println(arr)
	fmt.Println(x4)

	// appending
	fmt.Println("\n\n\nappending...")
	x5 := append(x3, 6, 7)
	fmt.Println("x5:", x5)
	fmt.Println("arr:", arr)

	_ = append(x4, 100, 101)
	fmt.Println("arr:", arr)
	fmt.Println("x5:", x5)
}
