package main

import (
	"fmt"
)

func main() {
	fmt.Println("Begin")
	//Explicit
	var tmp1 int = 0
	var tmp2 string = "Hello"
	var tmp3 bool = true
	//Implicit
	tmp5 := 5

	fmt.Println(tmp1)
	fmt.Println(tmp2)
	fmt.Println(tmp3)
	fmt.Println(tmp5)
}
