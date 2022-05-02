package main

import "fmt"

func main() {
	var numbers = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println("", numbers["two"]) //2

	var colors = make(map[string]string)
	colors["red"] = "#f00"
	fmt.Println("", colors["red"]) //#f00

	var courses = make(map[string]map[string]int)
	// courses["Android"]["price"] = 200 //call before BAD
	courses["Android"] = make(map[string]int)
	courses["Android"]["price"] = 200 //works
	// courses["Android"] = map[string]int{"price": 200} //or assign

	fmt.Println(courses["Android"]["price"])

	courses["ios"] = make(map[string]int)
	courses["ios"]["price"] = 350 //works
}
