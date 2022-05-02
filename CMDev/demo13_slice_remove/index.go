package main

import "fmt"

func main() {
	var number = []int{1, 2, 3, 4, 5, 6, 7, 8}
	showSlice(number)

	//ลบหัว
	number = number[1:len(number)]
	showSlice(number)
	number = number[1:len(number)]
	showSlice(number)
	// ลบหาง
	number = number[0 : len(number)-1]
	showSlice(number)
	number = removeIndex(number, 1)
	showSlice(number)
}

func showSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func removeIndex(s []int, index int) []int {
	// return append(s[:index], s[index+1:]...) //... = spread //len=8 cap=8 slice=[1 3 4 5 6 7 8]
	// return append(s[:index], s[index+1:][0]) //... = spread
	return append(s[:index], s[4]) //... = spread //len=8 cap=8 slice=[3 7]
}
