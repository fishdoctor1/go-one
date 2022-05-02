package main

import (
	"fmt"
	"time"
)

func IntSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func stringSeq() func() string {
	y := 0
	return func() string {
		y++
		// return fmt.Printf("Y = %d", y) //error
		return fmt.Sprintf("Y = %d", y)
	}
}
func main() {
	//ใช้มากในการทำ callback, delegate
	basic()
	lnw()
}

func basic() {
	nextInt := IntSeq()

	fmt.Println(nextInt()) //1
	fmt.Println(nextInt()) //2
	fmt.Println(nextInt()) //3

	newInt := IntSeq()
	fmt.Println(newInt()) //1
	fmt.Println(newInt()) //2

	fmt.Println(nextInt()) //4
	fmt.Println(nextInt()) //5

	newStr := stringSeq()
	fmt.Println(newStr()) //1

	fmt.Println(stringSeq()()) //1
}

func lnw() {
	//งงงงงงงงงงงงงงงงงงงงง
	for i := 0; i < 5; i++ {
		fmt.Println("[1] addr", &i)
		//i := i //if comment this then i=5 all row
		go func() {
			time.Sleep(time.Second * 1)
			fmt.Println("[34] addr", &i)
			fmt.Println(i)
		}()

	}
	time.Sleep(time.Second * 2) //ถ้าไม่มี log ไม่ออก
}
