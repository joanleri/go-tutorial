package main

import "fmt"

func main() {

	// Complex numbers
	c1 := complex(10, 20)
	c2 := 2 + 5i
	fmt.Printf("%T\n", c1)

	// Operations between complex numbers
	fmt.Println(c1 + c2)
	fmt.Println(c1 - c2)
	fmt.Println(c1 * c2)
	fmt.Println(c1 / c2)

	// Extracting real part
	fmt.Println(real(c1 - c2))
	fmt.Printf("%T\n", real(c1-c2))

	// Extracting imaginary part
	fmt.Println(imag(c1 - c2))
	fmt.Printf("%T\n", imag(c1-c2))

	// Modulus for complex numbers
}
