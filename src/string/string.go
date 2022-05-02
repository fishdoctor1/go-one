package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	// normalEx()
	runeEx()

}

func normalEx() {
	x := "hi-สวัสดี"
	y := []byte{0x68, 0x69, 0x2d, 0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0xa7, 0xe0, 0xb8, 0xb1, 0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0x94, 0xe0, 0xb8, 0xb5}

	fmt.Println(x, len(x))
	hx := hex.EncodeToString([]byte(x))
	fmt.Println(hx)
	fmt.Println(y)
	for i := 0; i < len(x); i++ {
		// fmt.Println(i)
		fmt.Printf("0x%x,", x[i])
	}
	fmt.Println()
	//ส = e2a (unicodeHEX) = 1110 0010 1010 (bin) = 3byte
	//convert unicode to utf-8 (1110xxxx 10xxxxxx 10xxxxxx) search utf-8 wiki look at enCoding zone
	//ส use 3byte because (e2a) 3 letters
	//ส in utf-8 = 111000001011100010101010
	//ส utf-8 in hexadecimal = e0b8aa
	fmt.Println(string(y[3:6]))          //ส
	fmt.Println(string(0xE2A))           //ส
	fmt.Println("utf-8", "\xe0\xb8\xaa") //ส
}

func runeEx() {
	x := "hi-สวัสดี"
	y := []byte{0x68, 0x69, 0x2d, 0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0xa7, 0xe0, 0xb8, 0xb1, 0xe0, 0xb8, 0xaa, 0xe0, 0xb8, 0x94, 0xe0, 0xb8, 0xb5}
	z := []rune(x)
	fmt.Println(len(x))
	fmt.Println(len(y))
	fmt.Println(len(z))
	fmt.Printf("%q", z[4]) //ว
}
