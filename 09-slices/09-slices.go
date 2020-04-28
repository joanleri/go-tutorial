package main

import "fmt"

func main() {
	var x []float64
	x = make([]float64, 5)
	x[0] = 1
	fmt.Println(x)
	x[1] = 10
	fmt.Println(x)

	// selecting underlying array capacity
	var x2 []float64 = make([]float64, 5, 10)
	fmt.Println(len(x2))
	fmt.Println(cap(x2))

	// creating slices from an array
	arr := [5]float64{1, 2, 3, 4, 5}
	x3 := arr[0:5]
	x4 := arr[1:3]
	fmt.Println(x3)
	fmt.Println(x4)
	x4[0] = -10
	fmt.Println(x3)
	fmt.Println(arr)

	// appending
	x5 := append(x3, 6, 7)
	fmt.Println(x5)
	fmt.Println(arr)

	_ = append(x4, 100, 101)
	fmt.Println(arr)
	fmt.Println(x5)
}
