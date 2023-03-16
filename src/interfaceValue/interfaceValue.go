package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// inside interface is Type,Value
	var w io.Writer
	fmt.Printf("%T, %v \n", w, w) //nil nil

	w = os.Stdout
	fmt.Printf("%T, %v \n", w, w) //*os.File, address

	w = &bytes.Buffer{}
	fmt.Printf("%T, %v \n", w, w) //*bytes.Buffer,

	w = nil
	fmt.Printf("%T, %v \n", w, w) //nil nil

	var byteBuff *bytes.Buffer
	fmt.Println("byteBuff", byteBuff == nil) //true
	fmt.Println("w1:", w == nil)             //true //Type=nil, Value=nil
	w = byteBuff
	fmt.Println("w2:", w == nil) // false //Type=*ByteBuff, Value=nil
	printBuff()

	//ระลึกไว้ว่าไม่ควรเอา interface ใดๆเข้า condition (interface == nil)
}

func printBuff() {
	fmt.Println("----------------printBuff-----------------")
	var byteBuff *bytes.Buffer
	var w io.Writer
	fmt.Printf("%T %v \n", w, w)
	w = byteBuff
	fmt.Printf("%T %v \n", w, w)
	// var w = byteBuff //ไม่พัง ถ้าประกาศเป็น Type io.Writer ก่อนจะพัง งง  //Type=*bytes.Buffer, Value=<nil>
	if w != nil {
		fmt.Println("w !!!!!!!!!!!!!!!!!!! nil")
		fmt.Println(w.Write(([]byte("hello"))))
	}
	//from head2 commit2
}
