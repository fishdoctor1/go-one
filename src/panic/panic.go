package main

import "fmt"

func main() {
	a(4)

}

func a(a int) {
	defer func() {
		fmt.Println("defer A")
	}()
	b(a)
}

func b(a int) {
	defer func() {
		fmt.Println("defer B")
	}()
	c(a)

}

func c(a int) {
	defer func() {
		fmt.Println("defer C")
	}()
	if a == 4 {
		panic("panic in c")
	}
	fmt.Println("ccccccccc")
}
