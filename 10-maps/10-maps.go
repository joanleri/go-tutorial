package main

import "fmt"

func main() {

	// creting a map
	var days map[string]int64
	days = make(map[string]int64)
	fmt.Printf("%T\n", days)
	days["Monday"] = 1
	days["Tuesday"] = 2
	days["Wednesday"] = 3
	days["Thursday"] = 4
	days["Friday"] = 5
	days["Saturday"] = 6
	days["Sunday"] = 7

	fmt.Println("length:", len(days)) // is equal to 7

	// deletion
	delete(days, "Monday")
	fmt.Println("new length:", len(days)) // is equal to 6

	// search confirmation
	// common pattern
	if dayNumber, ok := days["Monday"]; ok {
		fmt.Println(dayNumber, ok)
	} else {
		fmt.Println("Monday is not part of the week...")
	}

	if dayNumber, ok := days["Tuesday"]; ok {
		fmt.Println(dayNumber, ok)
	}
}
