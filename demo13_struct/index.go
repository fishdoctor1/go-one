package main

import "fmt"

func main() {
	var p1 product
	p1.name = "ard"
	p1.price = 100
	p1.stock = 100
	show(p1)
}

type product struct {
	name  string
	price int
	stock int
}

func show(p product) {
	fmt.Println(p)
}
