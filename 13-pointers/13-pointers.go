package main

import "fmt"

func main() {
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Println("a:", a, "b:", b)

	cPtr := new(int)
	*cPtr = 30
	fmt.Println("C:", *cPtr)
}

func swap(aPtr *int, bPtr *int) {
	temp := *aPtr
	*aPtr = *bPtr
	*bPtr = temp
}
