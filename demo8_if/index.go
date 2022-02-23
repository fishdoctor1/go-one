package main

import "fmt"

func main() {
	someValue := 10
	if someValue == 10 {
		fmt.Printf("==10")
	} else {
		fmt.Printf("!=10")
	}

	if someValue > 10 || someValue < 2 {
		fmt.Printf("someValue > 10 || someVAlue < 2")
	} else {
		fmt.Printf("NOT someValue > 10 || someValue < 20")
	}

	if result := doSomething(); result == "ok" {
		fmt.Printf("ok")
	} else {
		fmt.Println("not")
	}

}

func doSomething() string {
	return "ok"
}

func fnSwitchCase() {
	index := 2
	switch index {
	case 0:
		fmt.Println("0")
		break
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("def")
		break
	}
}
