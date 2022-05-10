package main

//session 99 flag interface
import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

// type Value interface {
// 	String() string
// 	Set(string) error
// }

var romanMap = map[string]int{
	"i":   1,
	"ii":  2,
	"iii": 3,
	"iv":  4,
}

type RomanAge struct {
	age string
}

func (r *RomanAge) String() string {
	return strconv.Itoa(romanMap[r.age])
}

func (r *RomanAge) Set(value string) error {
	r.age = value
	return nil
}

func flagRomanAge() *RomanAge {
	romanAge := RomanAge{}
	flag.Var(&romanAge, "age", "help msg flagname")
	return &romanAge
}

func main() {
	// ex1()
	ex2()
}

func ex2() {
	//go run interface.go -name fa -age iii
	var name = flag.String("name", "annonymous", "your name")
	var age = flagRomanAge()
	fmt.Println(*name)
	fmt.Println(age)
}
func ex1() {
	var nameEx1 = flag.String("name", "annonymous", "your name")
	//ถ้ารันแบบ go run interface.go -name fa -age 234
	//output [-name fa -age 234]
	fmt.Println(os.Args[1:])

	fmt.Println(*nameEx1) //annonymous
}
