package main

import (
	"fmt"
	"time"
)

func run1() {
	for i := 0; i < 100; i++ {
		fmt.Println("Run1", i)
	}
}

func run2() {
	for i := 0; i < 100; i++ {
		fmt.Println("Run2", i)
	}
}

func main() {
	//go เหมือนแบ่งออกไปทำอีก เทรด
	go run1()
	go run2()

	time.Sleep(5 * time.Second)
}
