package main

import "fmt"

func main() {
	fmt.Println(sum(3, 2))
	x, y, title := swapv2(10, 20)
	fmt.Printf("%d %d %s\n", x, y, title)
}

func sum(a int, b int) int {
	return a + b
}

func swap(a int, b int) (int, int) {
	return b, a
}

func swapv2(a int, b int) (x, y int, title string) {
	y = a
	x = b
	title = "sdf"
	return
}
