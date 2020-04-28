package main

import "fmt"

func main() {
	var total float64 = 87
	var numberOfCourses int = 10
	var average float64 = total / float64(numberOfCourses)
	fmt.Println("Average is:", average)
}
