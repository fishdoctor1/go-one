package main

import (
	"fmt"
	"time"
)

func routine1(c chan int, payload int) {
	c <- payload
	fmt.Println("s1")
	c <- payload
	fmt.Println("s2")
	c <- payload
	fmt.Println("s3")
}

func main() {
	ch := make(chan int, 5) //make channel
	go routine1(ch, 1)
	//go routine1(ch, 2)
	fmt.Println(<-ch) // for undeadlock if buffer not enough
	fmt.Println(<-ch) // for undeadlock if buffer not enough
	fmt.Println(<-ch) // for undeadlock if buffer not enough
	//if read 4time then deadlock
	time.Sleep(1 * time.Second)
	//--output--------------
	//s1
	//s2
	//s3
	//1
}
