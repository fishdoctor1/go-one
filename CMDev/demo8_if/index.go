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

	x := 2
	if v := 1; (x < 5) || false {
		fmt.Println("if", v)
	} else {
		fmt.Println("else", v)
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
		fallthrough //ตกลงไปข้างล่างต่อ1เคส
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("def")
		break
	}

	index = 2
	switch {
	case index == 1:
		fmt.Println("0")
	case index == 1:
		fmt.Println("1")
	case index == 2:
		fmt.Println("2")
	default:
		fmt.Println("def")
	}

	index = 2
	switch false {
	case index == 1:
		fmt.Println("0")
	case index == 1:
		fmt.Println("1")
	case index == 2:
		fmt.Println("2")
	default:
		fmt.Println("def")
	}
}
