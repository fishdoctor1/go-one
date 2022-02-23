package main

import "fmt"

func main() {
	msg := "some"
	var msgPointer *string = &msg
	fmt.Println(msg)         //some
	fmt.Println(*msgPointer) //some

	changeMessage(&msg, "new1")
	fmt.Println(msg)
	changeMessage(msgPointer, "new2")
	fmt.Println(msg)
}

func changeMessage(aPointer *string, newM string) {
	*aPointer = newM
}
