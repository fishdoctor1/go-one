package main

import (
	"fmt"
)

func routine1(c chan int, countTo int) {
	for i := 0; i < countTo; i++ {
		c <- i
	}
	close(c) //ปิด channel

}

func main() {
	ch := make(chan int, 10) //make channel
	go routine1(ch, 10)

	/*for true {
		value, ok := <-ch
		fmt.Println("chan", ok)
		fmt.Println("l:", value)
		if !ok {
			fmt.Println("close")
			break
		}
	}*/
	// result
	// 	chan true
	// l: 0
	// chan true
	// l: 1
	// chan true
	// l: 2
	// chan true
	// l: 3
	// chan true
	// l: 4
	// chan true
	// l: 5
	// chan true
	// l: 6
	// chan true
	// l: 7
	// chan true
	// l: 8
	// chan true
	// l: 9
	// chan false
	// l: 0
	// close
	for value := range ch {
		fmt.Println(value)
	}
	// result0
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// if not close then deadlock
}
