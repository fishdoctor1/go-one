package main

import (
	"fmt"
	"math/rand"
	"time"
)

func running(t chan<- string) {
	// rand.Seed(time.Now().UnixNano())
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	t <- "1"
}
func main() {
	t1 := make(chan string)
	t2 := make(chan string)
	t3 := make(chan string)

	go running(t1)
	go running(t2)
	go running(t3)

	select {
	case <-t1:
		fmt.Println("t1")

	case <-t2:
		fmt.Println("t2")

	case <-t3:
		fmt.Println("t3")
	}
}
