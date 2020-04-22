package main

import "fmt"

func main() {
	grades := []float64{10, 9, 8, 7, 6}
	fmt.Println("Average:", average(grades))

	minimum, maximum, average := testStatistics(grades)
	fmt.Println("Min:", minimum, "Max:", maximum, "Average:", average)

	fmt.Println("Average:", variadicAverage(10, 9, 8, 7, 6))
	fmt.Println("Average:", variadicAverage(grades...))
}

func average(grades []float64) float64 {
	var total float64 = 0
	for _, value := range grades {
		total += value
	}
	return total / float64(len(grades))
}

func testStatistics(grades []float64) (min float64, max float64, average float64) {
	min = grades[0]
	max = grades[0]
	total := 0.0

	for _, value := range grades {
		if min > value {
			min = value
		}
		if max < value {
			max = value
		}
		total += value
	}
	average = total / float64(len(grades))
	return
}

func variadicAverage(grades ...float64) float64 {
	total := 0.0
	for _, value := range grades {
		total += value
	}
	return total / float64(len(grades))
}
