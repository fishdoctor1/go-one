package main

//session 49-50
import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var x = "ทดสอบ"

func main() {
	//ex1()
	convert()
}

func ex1() {
	for i, v := range x {
		fmt.Println(i, v, reflect.TypeOf(v))
		fmt.Printf("%d %c\n", i, v)
	}
	fmt.Println(utf8.RuneCountInString(x)) //5
	for i := 0; i < len(x); {
		r, s := utf8.DecodeRuneInString(x[i:])
		i = i + s
		fmt.Printf("%c-", r)
	} //ท-ด-ส-อ-บ-
}

func stringBuffer() {
	finder := "สอ"
	fmt.Println(bytes.Count([]byte(x), []byte(finder))) //1
	fmt.Println(strings.Count(x, finder))               //1
}

func byteBuffer() {
	buff := bytes.Buffer{}
	//or bufff := strings.Builder{}
	buff.WriteRune('y')
	buff.WriteRune('o')
	fmt.Println(buff.String()) //yo
}

func convert() {
	atoi, _ := strconv.Atoi("123")
	fmt.Println(atoi, reflect.TypeOf(atoi))
	itoa := strconv.Itoa(123)
	fmt.Println(itoa, reflect.TypeOf(itoa))
	fmt.Println(strconv.ParseBool("asdasd"))

	fmt.Println("unicode.IsDigit('a')", unicode.IsDigit('a'))                  // false
	fmt.Println("unicode.IsDigit('1')", unicode.IsDigit('1'))                  // true
	fmt.Println("unicode.IsUpper('a')", unicode.IsUpper('a'))                  // false
	fmt.Println("unicode.IsUpper('A')", unicode.IsUpper('A'))                  //true
	fmt.Println("unicode.IsLetter('A')", unicode.IsLetter('A'))                //true
	fmt.Println("unicode.IsLetter('$')", unicode.IsLetter('$'))                //false
	fmt.Println("unicode.In('a',unicode.Thai)", unicode.In('a', unicode.Thai)) //false
}
