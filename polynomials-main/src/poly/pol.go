package poly

import (
    "fmt"
    "strconv"
    "sort"
)

var (
    e0 = "\u2070"
    e1 = "\u00b9"
    e2 = "\u00b2"
    e3 = "\u00b3"
    e4 = "\u2074"
    e5 = "\u2075"
    e6 = "\u2076"
    e7 = "\u2077"
    e8 = "\u2078"
    e9 = "\u2079"
    ep = "\u207a"
    em = "\u207b"
)

type Term struct {
    Value int
    Grade int
}

type Pol struct{
    pol map[int]int
}

func (p Pol) String() string {
    if len(p.pol) == 0 { //i.e. p-p
        return fmt.Sprintf("0")
    }
    if _, ok := p.pol[0]; ok && len(p.pol) == 1 && p.pol[0] == 0 { //i.e. GetPol(poly.Term{0, 0})
        return fmt.Sprintf("0")
    }
    if _, ok := p.pol[0]; ok && p.pol[0] == 0 {
        delete(p.pol, 0)
    }
    var keys []int
    for j := range p.pol {
        keys = append(keys, j)
    }
    sort.Sort(sort.Reverse(sort.IntSlice(keys)))

    ret := ""
    for i, k := range keys {
        if p.pol[k] < 0 {
            ret += "\b"
        }
        if p.pol[k] != 1 || (p.pol[k] == 1 && k == 0){
            ret += strconv.Itoa(p.pol[k])
        }
        if k > 0 {
            ret += "x"
        }
        if k > 1 {
            for _, c := range strconv.Itoa(k) {
                switch c {
                case '0':
                    ret += e0
                case '1':
                    ret += e1
                case '2':
                    ret += e2
                case '3':
                    ret += e3
                case '4':
                    ret += e4
                case '5':
                    ret += e5
                case '6':
                    ret += e6
                case '7':
                    ret += e7
                case '8':
                    ret += e8
                case '9':
                    ret += e9
                }
            }
        }
        if i != len(keys)-1 {
            ret +=  " +"
        }
    }
    return fmt.Sprintf(ret)
}

func GetPol(terms ...Term) Pol {
    var pol map[int]int
    pol = make(map[int]int)
    for _, term := range terms {
        if term.Value != 0 || (term.Grade == 0 && term.Value == 0)  {  // omit 0 coeficient term, excetp poly.Term{0, 0}
            pol[term.Grade] = term.Value
        }
    }
    return Pol{pol}
}

func (p1 Pol) Plus(p2 Pol) Pol {
    var p3 map[int]int
    p3 = make(map[int]int)
    for k, v := range p1.pol {
        if v + p2.pol[k] != 0 {  // omit 0 coeficient term
            p3[k] = v + p2.pol[k]
        }
    }
    for k, v := range p2.pol {
        if p1.pol[k] + v != 0 {  // omit 0 coeficient term
            p3[k] = p1.pol[k] + v

        }
    }
    return Pol{p3}
}

func (p1 Pol) Plus2(p2 Pol) Pol {
    var x []Term
    for n := range p1.Plus2a(p2) {
        x = append(x, n)
    }
    return GetPol(x...)
}

func (p1 Pol) Plus2a(p2 Pol) <-chan Term {
    out := make(chan Term)
    go func(){
        for k, v := range p1.pol {
            if v + p2.pol[k] != 0 {  // omit 0 coeficient term
                out <- Term{v + p2.pol[k], k}
                fmt.Printf("%v\n",Term{v + p2.pol[k], k})
            }
        }
        for k, v := range p2.pol {
            if p1.pol[k] + v != 0 {  // omit 0 coeficient term
                out <- Term{p1.pol[k] + v, k}
                fmt.Printf("%v\n",Term{v + p2.pol[k], k})
            }
        }
        close(out)
    }()
    return out
}

func (p1 Pol) Minus(p2 Pol) Pol {
    var p3 map[int]int
    p3 = make(map[int]int)

    for k, v := range p1.pol {
        if v - p2.pol[k] != 0 {  // omit 0 coeficient term
            p3[k] = v - p2.pol[k]
        }
    }

    for k, v := range p2.pol {
        if p1.pol[k] - v != 0 {  // omit 0 coeficient term
            p3[k] = p1.pol[k] - v
        }
    }
    return Pol{p3}
}

func (p1 Pol) Times(p2 Pol) Pol {
    var p3 map[int]int
    p3 = make(map[int]int)

    for k, v := range p1.pol {
        for k2, v2 := range p2.pol {
            p3[k+k2] = p3[k+k2] + (v * v2)
            if v3, ok := p3[k+k2]; ok && v3==0 {
                delete(p3, k+k2)
            }
        }
    }
    return Pol{p3}
}

func (p1 Pol) Compose(p2 Pol) Pol {
    var p3 map[int]int
    p3 = make(map[int]int)
    p3[0]=0
    r := Pol{p3}
    var keys []int
    for j := range p1.pol {
        keys = append(keys, j)
    }
    sort.Sort(sort.Reverse(sort.IntSlice(keys)))
    for _, k := range keys {
        var t map[int]int
        t = make(map[int]int)
        t[0]=p1.pol[k]
        r = Pol{t}.Plus(p2.Times(r))
    }
    return r
}

func (p1 Pol) Evaluate(x int) int {
    r := 0
    var keys []int
    for j := range p1.pol {
        keys = append(keys, j)
    }
    sort.Sort(sort.Reverse(sort.IntSlice(keys)))
    if len(keys) > 0 {
        for i := keys[0]; i >= 0; i-- {
            r = p1.pol[i] + (x * r)
        }
    }
    return r
}

func (p1 Pol) Differentiate() Pol {
    var p3 map[int]int
    p3 = make(map[int]int)

    for k, v := range p1.pol {
        if k != 0 {
            p3[k-1]=v*k
        }
    }
    return Pol{p3}
}

func (p1 Pol) NumericIntegration(min int, max int, blocks int) int {

	var c chan int = make(chan int)
	delta := (max - min) / blocks
	var total int

	for i := 1; i <= blocks; i++ {
		go func(i int) {
			temp := p1.Evaluate(min + delta*i)
			c <- temp
		}(i)
	}

	for i := 1; i <= blocks; i++ {
		temp := <-c
		total += temp * delta
	}

	return total
}
