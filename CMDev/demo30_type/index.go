package main

import "fmt"

type KG float64
type LB float64

func (lb LB) toKG() KG {
	return KG(lb / 2.204)
}
func main() {

	k := KG(1)
	l := LB(1)
	result := k + 3
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println(result)        // Implicitly
	fmt.Printf("%T\n", result) // Implicitly type = KG
	// fmt.Println(a == b) //cannot compare a == b (mismatched types KG and LB)
	lb := LB(2.204)
	fmt.Println(l, "LB =", l.toKG(), "KG")
	fmt.Println(k == lb.toKG())
}
