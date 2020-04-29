package main

import "fmt"

func main() {

	sevenSequence := sequenceGenerator(7)
	fmt.Println(sevenSequence())
	fmt.Println(sevenSequence())
	fmt.Println(sevenSequence())

	grades := []float64{9, 8, 7, 6, 5}
	mapper(grades, func(x float64) float64 { return x + 1 })
	fmt.Println(grades)

	mapper(grades, func(x float64) float64 { return 0.5 * x })
	fmt.Println(grades)

	// mapper(grades, func(x int) int { return 0.5 * int(x) })
	// fmt.Println(grades)
}

func sequenceGenerator(increment int) func() int {
	i := 0
	return func() int {
		i += increment
		return i
	}
}

func mapper(x []float64, f func(float64) float64) {
	for index, value := range x {
		x[index] = f(value)
	}
}
