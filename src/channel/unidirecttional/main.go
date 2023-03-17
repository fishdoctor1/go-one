package main

import "fmt"

func sender(ch chan<- int) {
	ch <- 88
	// close(ch)
}

func f(ch chan int) {
	ch <- 88
	close(ch)
}

func receiver(ch <-chan int, done chan<- bool) {
	fmt.Println("Receive ", <-ch)
	done <- true
}

func main() {
	ch := make(chan int)
	done := make(chan bool)
	go sender(ch) //if not go then deadlock!
	go receiver(ch, done)
	<-done //if comment then print 'Done' before receiver
	fmt.Println("Done")

}
