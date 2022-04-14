package main

import "fmt"

func main() {
	msg := "some"
	var msgPointer *string = &msg
	fmt.Println(msg)         //some
	fmt.Println(*msgPointer) //some

	changeMessage(&msg, "new1")
	fmt.Println(msg)          //new1
	fmt.Println("&msg", &msg) //addr1
	changeMessage(msgPointer, "new2")
	fmt.Println(msg)          //new2
	fmt.Println("&msg", &msg) //addr1

	changeMessage2(&msg, "new3")
	fmt.Println(msg)          //new3
	fmt.Println("&msg", &msg) //addr1
}

func changeMessage(aPointer *string, newM string) {
	*aPointer = newM
	fmt.Println("change1", *aPointer)
	fmt.Println("&apointer", &aPointer)
}

func changeMessage2(aPointer *string, newM string) {
	*aPointer = newM
	fmt.Println("change2", *aPointer)
	fmt.Println("&apointer", &aPointer)
	// fmt.Println(*aPointer) //cannot access value
	// หัวฟังชัน รับ value pointer รับค่าต้องเป็น value pointer
	// &aPointer เปลี่ยนไปเสมอ แต่ &msg ยังคงตำแหน่งเดิม

	//some time access value like *(&aPointer)

}
