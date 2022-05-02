package main

import "fmt"

//session 51
func main() {
	const (
		Sunday float64 = iota + 1.1
		Monday
		Tuesday
		Wednesday int = 99 + iota
	)
	fmt.Println(Sunday, Monday, Tuesday, Wednesday)
	//1.1 2.1 3.1 102
}
