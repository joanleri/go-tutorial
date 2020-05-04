package main

import "fmt"

func main() {
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Println("a:", a, "b:", b)

	cPtr := new(int)
	fmt.Println("c:", *cPtr)
	*cPtr = 30
	fmt.Println("c:", *cPtr)
}

func swap(aPtr *int, bPtr *int) {
	temp := *aPtr
	*aPtr = *bPtr
	*bPtr = temp
}
