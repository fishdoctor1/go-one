package main

import "fmt"

func main() {
	ch := make(chan int)
	chSquare := make(chan int)
	done := make(chan string)
	go sender(ch, done)
	go square(ch, chSquare, done)
	go receiver(chSquare, done)

	doneStatus := <-done
	fmt.Println("Done status : ", doneStatus)
	doneStatus = <-done
	fmt.Println("Done status : ", doneStatus)
	doneStatus = <-done
	fmt.Println("Done status : ", doneStatus)
}

func sender(ch chan int, done chan string) {
	for i := 0; i <= 5; i++ {
		fmt.Println("Sending value : ", i)
		ch <- i
	}
	close(ch)
	fmt.Println("Sender is about to complete.")
	done <- "Done from Sender"
	fmt.Println("Sender done")
}
func receiver(ch chan int, done chan string) {
	for v := range ch {
		fmt.Println("\tReceive value : ", v)
	}
	done <- "Done from Receive"
	/*
		//way 2 check channel isClose
		for {
			v,ok:= <ch
			if !ok {
				break
			}
		}
	*/
}

func square(ch chan int, chSquare chan int, done chan string) {
	for v := range ch {
		chSquare <- v * v
	}
	close(chSquare)
	done <- "Done from square"
}
