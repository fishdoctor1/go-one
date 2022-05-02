package main

//session 47
import "fmt"

// serach unicode lookup
func main() {
	ex1 := string(0x265e)
	//int to string
	fmt.Println("int to string")
	for i := 0; i < len(ex1); i++ {
		fmt.Printf("%d %x,", i, ex1[i])
	}
	fmt.Println()
	//[]bype to string
	fmt.Println("byte to string")
	ex2 := []byte{0xe2, 0x99, 0x9e}
	ex2String := string(ex2)
	fmt.Println(ex2String)

	//[]rune to string
	fmt.Println("rune to string")
	ex3 := []rune{0x2654}
	fmt.Println(string(ex3))

	//string to byte
	fmt.Println("string to byte")
	ex4 := []byte("hello♔")
	fmt.Println(ex4)

	//string to byte
	fmt.Println("string to rune")
	ex5 := []rune("hello♔")
	fmt.Println(ex5)
}
