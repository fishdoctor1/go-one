package main

import (
	"fmt"
)

func task(c, c2, quit chan string) {
	for {
		select {
		case c <- "Running...":
		case c2 <- "Running2...":
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan string)    //make channel
	c2 := make(chan string)   //make channel
	quit := make(chan string) //make channel
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
			fmt.Println(<-c2)
			fmt.Println("hi go func")
		}

		quit <- "end" //ถ้าไม่มี quit จะ deadlock
	}()
	// time.Sleep(time.Second * 2) //ถ้าไม่มี log ไม่ออก
	task(c, c2, quit) //ถ้าไม่มี log ไม่ออก
	// เลือกใช้อย่างใดอย่างหนึ่ง

}
