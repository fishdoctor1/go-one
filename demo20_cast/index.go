package main

import (
	"fmt"
	"math"
	"reflect"
)

func showInfo(s shape) {
	t := reflect.TypeOf(s).Name()
	switch t {
	case "rectangle":
		r := s.(rectangle) //cast สวมบทบาทเป็น rectangle
		fmt.Printf("Rect wid: %f, heig %f\n", r.width, r.height)
		break
	case "circle":
		c := s.(circle) //cast สวมบทบาทเป็น rectangle
		fmt.Printf("Circle radisu: %f\n", c.radius)
		break
	}
}

func main() {
	r := rectangle{width: 10, height: 10}
	c := circle{radius: 10}
	fmt.Printf("Rectangle: %f\n", getArea(r))
	fmt.Printf("Circle: %f\n", getArea(c))

	showInfo(r)
	showInfo(c)
}

type shape interface {
	area() float64
}

func getArea(s shape) float64 {
	return s.area()
}

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
