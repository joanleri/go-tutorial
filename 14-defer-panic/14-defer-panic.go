package main

import "fmt"

func main() {
	deferingFunction()

	grades := []float64{}
	panicAverage(grades)
}

func deferingFunction() {
	defer fmt.Println("This will be printed until the end...")
	fmt.Println("This will be printed first...")
	fmt.Println("This will be printed second...")
	fmt.Println("This will be printed third...")
}

func panicAverage(grades []float64) float64 {
	defer func() {
		str := recover()
		if str != nil {
			fmt.Println(str)
		}
	}()

	if len(grades) == 0 {
		panic("LENGTH OF SLICE CANNOT BE ZERO")
	}

	total := 0.0
	for _, value := range grades {
		total += value
	}

	return total / float64(len(grades))
}
