package main

import "fmt"

func main() {
	var p1 product
	p1.name = "ard"
	p1.price = 100
	p1.stock = 100
	show(p1)
	update(&p1)
	show(p1)

	p1 = p1.setDiscount(1)
	show(p1)
	p1 = p1.clear()
	show(p1)
}

type KG float64
type LB float64
type product struct {
	name  string
	price int
	stock int
}

func compareLB_KB() {
	a := KG(3)
	b := LB(3)
	fmt.Println(a)
	fmt.Println(b)
	// fmt.Println(a == b) //cannot compare a == b (mismatched types KG and LB)
}
func (p product) clear() product {
	//p product บอกว่า เป็นฟังชันของ struct ไหน
	//clear() function name
	//ret  return value
	fmt.Println("clear")
	p.price = 0
	p.stock = 0
	return p
}

func (p product) setDiscount(d int) product {
	fmt.Println("setDisc")
	p.price = p.price - d
	return p
}
func show(p product) {
	fmt.Println(p)

}

func update(p *product) {
	p.price = p.price + 100
	p.stock = 10
}
