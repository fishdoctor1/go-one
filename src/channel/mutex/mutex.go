package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const N = 1000
	add := 0
	sub := 0
	mu := sync.Mutex{}

	for i := 0; i < N; i++ {
		add++
		go func() {
			defer func() {
				mu.Lock() //ใช้แก้ปัญหาจาก for range photos ได้คล้ายกับการใช้ waitgroup
				sub--
				mu.Unlock()
			}()
		}()
	}

	for {
		h, m, s := time.Now().Clock()
		fmt.Printf("%d-%d-%d %d %d\r", h, m, s, add, sub)

		if add == N && sub == -N {
			fmt.Println("Success")
			break
		}
	}

	fmt.Println("Main Done", add, sub)
}
