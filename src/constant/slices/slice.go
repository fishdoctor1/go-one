package main

import "fmt"

func main() {
	x := [10]int{}
	a := x[0:5]
	// b := x[5:7]
	for i := range a {
		a[i] = i * i
	}
	// a = append(a, b...)
	fmt.Println("x:", x)
	a = append(a, 71, 72, 73)
	fmt.Println("x:", x)
	fmt.Println("a:", a)
	a = append(a, 74, 75, 76)
	fmt.Println("x:", x)
	fmt.Println("a:", a)
	fmt.Println(len(a), cap(a))
}
