package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	done := make(chan bool)
	go train(done)
	trainStatus := <-done
	fmt.Println("Done status:", trainStatus)
}

func train(ch chan bool) {
	// t := time.Now()
	// rand.Seed(t.Nanoseconds)
	x := rand.Intn(1000)
	fmt.Println("train", x, " millisecond")
	time.Sleep(time.Duration(x) * time.Millisecond)
	ch <- true
}

//Result-----------------------------------
// train 81  millisecond
// Done status: true
