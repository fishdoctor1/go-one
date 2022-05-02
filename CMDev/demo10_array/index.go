package main

import "fmt"

func main() {
	array := [2]int{1, 2}
	var array2 []int = []int{1, 2}
	var array3 [2]int
	array3[0], array3[1] = 1, 2
	fmt.Println(array[0])
	fmt.Println(array2[0])
	fmt.Printf("3[0] %d\n", array3[0])

	for _, v := range array {
		fmt.Println(v)
	}
	for _, v := range array2 {
		fmt.Println(v)
	}
	for _, v := range array3 {
		fmt.Println(v)
	}
}
