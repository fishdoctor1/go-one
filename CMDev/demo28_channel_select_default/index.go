package main

import (
	"fmt"
	"time"
)

func task(c, quit chan string) {
	for {
		select {
		case c <- "Running...":
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("def")
			time.Sleep(time.Millisecond * 500)
		}
	}
}

func main() {
	c := make(chan string)    //make channel
	quit := make(chan string) //make channel
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}

		//quit <- "end" //ถ้าไม่มี quit จะ deadlock ใช้ default ช่วยได้
	}()
	// time.Sleep(time.Second * 2) //ถ้าไม่มี log ไม่ออก
	task(c, quit) //ถ้าไม่มี log ไม่ออก
	// เลือกใช้อย่างใดอย่างหนึ่ง task or sleep

	// 	def
	// Running...
	// def
	// Running...
	// def
	// Running...
	// def
	// Running...
	// def
	// Running...
	// def
	// Running...
	// def
	// Running...
	// def
	// Running...
	// Running...
	// def
	// def
	// Running...
	// def
	// def
}
