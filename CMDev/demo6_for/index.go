package main

import "fmt"

func main() {
	//fnFor()
	// fnWhile()
	// forEach()
	fnWhileBreak()
}

func fnFor() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Index %d\n", i)
	}
}

func fnWhile() {
	var i = 0
	//or i := 0;
	for i < 10 {
		fmt.Printf("While-Index %d\n", i)
		i++
	}
}

func forEach() {
	courses := []string{"And", "ios", "Re"}
	for index, item := range courses {
		fmt.Printf("%d. %s\n", index, item)
	}

	for _, item := range courses { //_ ใช้เมื่อไม่ต้องการใช้งานตัวแปร index เดี๋ยวมันจะ แดง
		fmt.Printf("%s\n", item)
	}
}

func fnWhileBreak() {
	index := 0
	for true {
		if index > 5 {
			break
		}
		fmt.Printf("While-Index %d\n", index)
		index++
	}
}
