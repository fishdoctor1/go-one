package main

import "fmt"

func main() {
	var number = make([]int, 0, 5) //initial _len cap
	number = append(number, 1)
	number = append(number, 2)
	showSlice(number)

	var number2 []int
	number2 = append(number2, 1)
	number2 = append(number2, 2)
	showSlice(number2)
}

func showSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}
