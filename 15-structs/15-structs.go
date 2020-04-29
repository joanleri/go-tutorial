package main

import (
	"fmt"
	"math"
)

func main() {
	var c1 Circle
	c1.x = 0
	c1.y = 0
	c1.r = 10

	c2 := Circle{x: 0, y: 0, r: 10}

	c3 := Circle{0, 0, 10}

	cPtr1 := &Circle{0, 0, 10}

	cPtr2 := new(Circle)
	cPtr2.x = 0
	cPtr2.y = 0
	cPtr2.r = 10

	fmt.Println(c1.r, c2.r, c3.r, cPtr1.r, cPtr2.r)
	fmt.Println("The area of the circle is:", c1.area())

	// passing a pointer to a struct
	c1.doubleRadious()
	fmt.Println("c1 radious is now", c1.r)

	// interfaces
	aCircle := Circle{r: 10, x: 0, y: 0}
	aSquare := Square{l: 10}
	aRectangle := Rectangle{w: 10, h: 20}
	shapes := []Shape{aCircle, aSquare, aRectangle}
	fmt.Println("Total area is:", totalArea(shapes))

	// type assertions
	var aShape Shape = Circle{x: 1, y: 1, r: 33}
	var undelyingCirle Circle = aShape.(Circle)
	fmt.Println("The underlying circle's radious is", undelyingCirle.r)

	probablyACirlce, ok := aShape.(Circle)
	if ok {
		fmt.Println("It is a circle!")
		fmt.Println("Its radius is", probablyACirlce.r)
	} else {
		fmt.Println("It might be a square...")
	}
}

// Circle : abstraction of a circle
// with the coordinates of its center
// and radius
type Circle struct {
	x float64
	y float64
	r float64
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) doubleRadious() {
	c.r = 2 * c.r
}

// Square : a square struct
type Square struct {
	l float64
}

func (s Square) area() float64 {
	return s.l * s.l
}

// Rectangle : a rectangle struct
type Rectangle struct {
	h float64
	w float64
}

func (r Rectangle) area() float64 {
	return r.h * r.w
}

// Shape : a shape interface
type Shape interface {
	area() float64
}

func totalArea(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.area()
	}
	return total
}
