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
	ch := make(chan int, 1) //make channel
	ch <- 1                 //assign value
	msg := <-ch             // get value chanel to msg
	fmt.Println(msg)
	fmt.Println("step1")
	fmt.Println(<-ch) // for undeadlock

	ch <- 2
	fmt.Println("step2")
	fmt.Println(<-ch) // for undeadlock
	time.Sleep(1 * time.Second)
}
