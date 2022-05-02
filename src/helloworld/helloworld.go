package main

import (
	"fmt"
	"github.com/golang/example/stringutil"
	"myFmt"
	"log"
)

func main() {
	fmt.Println("hello")
	myFmt.Hi("ss")
	fmt.Println(stringutil.Reverse("HI"))
	log.Println("hhhh")// use log.go in vendor
}
