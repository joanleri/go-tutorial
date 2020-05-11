package main

import (
    "fmt"
    "poly"
)

func main() {

    p := poly.GetPol(poly.Term{4, 3}, poly.Term{3, 2}, poly.Term{2, 1}, poly.Term{1, 0})
    q := poly.GetPol(poly.Term{3, 2}, poly.Term{5, 0})
    z := poly.GetPol(poly.Term{0, 0})
    last := poly.GetPol(poly.Term{1, 0}, poly.Term{1, 1}, poly.Term{1, 2})

    fmt.Printf("z\t: %v\n",z)
    fmt.Printf("p\t: %v\n",p)
    fmt.Printf("q\t: %v\n",q)
    fmt.Printf("p + q\t: %v\n",p.Plus(q))
    fmt.Printf("p * q\t: %v\n",p.Times(q))
    fmt.Printf("p(q)\t: %v\n",p.Compose(q))
    fmt.Printf("0-p\t: %v\n",z.Minus(p))
    fmt.Printf("p(3)\t: %v\n",p.Evaluate(3))
    fmt.Printf("p'\t: %v\n",p.Differentiate())
    fmt.Printf("p''\t: %v\n",p.Differentiate().Differentiate())
    fmt.Printf("p-p\t: %v\n",p.Minus(p))


	fmt.Printf("last\t: %v\n",last)
	fmt.Printf("last(5)\t: %v\n",last.Evaluate(5))
	fmt.Printf("last.NumericIntegration(-50, 50, 100): %v\n", last.NumericIntegration(-50, 50, 100))

    // fmt.Printf("{1,1}\t: %v\n",poly.GetPol(poly.Term{1, 1}))
    // fmt.Printf("{1,0}\t: %v\n",poly.GetPol(poly.Term{1, 0}))
    // fmt.Printf("{1,1},{1,0}\t: %v\n",poly.GetPol(poly.Term{1, 1},poly.Term{1, 0}))
    // fmt.Printf("{0,0},{1,1}\t: %v\n",poly.GetPol(poly.Term{1, 1},poly.Term{0, 0}))
    // fmt.Printf("{0,0},{0,1},{5,2} | x=5 \t: %v\n",poly.GetPol(poly.Term{0, 0}, poly.Term{0, 1}, poly.Term{5, 2}).Evaluate(5))
    // fmt.Printf("{1,0},{1,2} | x=5 \t: %v\n",poly.GetPol(poly.Term{1, 0}, poly.Term{1, 2}).Evaluate(5))
    // fmt.Println(p.Plus2(q))
}
