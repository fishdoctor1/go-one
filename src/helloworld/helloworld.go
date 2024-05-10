package main

import (
	"fmt"
	"go-one/src/myFmt"
	"log"
)

func main() {
	fmt.Println("hello")
	myFmt.Hi("ss")
	fmt.Println(stringutil.Reverse("HI"))
	log.Println("hhhh") // use log.go in vendor
}
