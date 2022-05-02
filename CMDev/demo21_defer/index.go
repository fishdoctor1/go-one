package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("Finish", i) //same onDestroy in angular or settimeout
		//ทำงานตตอนจะจบ func
		//9 8 7 6 5 4 3 2 1 0
	}
	for i := 0; i < 10; i++ {
		fmt.Println("", i)
		//0 1 2 3 4 5 6 7 8 9
	}
}
