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

func run2(ch chan int) {
	for i := 0; i < 1; i++ {
		fmt.Println("Run2", i)
	}
	defer func(aa int) {
		ch <- aa
	}(555)
	return
}

func main() {
	ch := make(chan int, 1) //make channel
	ch <- 1                 //assign value
	msg := <-ch             // get value chanel to msg
	fmt.Println(msg)
	fmt.Println("step1")
	run2(ch)
	fmt.Println(<-ch) // for undeadlock
	// return
	ch <- 2
	fmt.Println("step2")
	fmt.Println(<-ch) // for undeadlock
	time.Sleep(1 * time.Second)
}
