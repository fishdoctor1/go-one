package main

//session 80 anonymous function concept
import "fmt"

func createF1() func() int {
	var x = 0
	return func() int {
		x++
		return x
	}
}

func createF2(list []int) []func() {
	result := []func(){}

	for _, v := range list {
		fmt.Println("inside createF:", v)
		result = append(result, func() {
			fmt.Println(v)
		})
	}
	return result
}

func createF3(list []int) []func() {
	result := []func(){}

	for _, v := range list {
		captureValue := v
		fmt.Println("inside createF:", v)
		result = append(result, func() {
			fmt.Println(captureValue)
		})
	}
	return result
}

func main() {
	first()
	fmt.Println("==========================")
	second()
	//func second() value หลังจาก append จะได้ value สุดท้ายเสมอ
	fmt.Println("==========================")
	third()
	//แก้ปัญหา v() ที่ ได้ value สุดท้ายเสมอด้วย captureValue
}

func first() {
	f := createF1()
	fmt.Println(f())
	fmt.Println(f())
	f2 := createF1()
	fmt.Println(f2())
	fmt.Println(f2())
}

func second() {
	fList2 := []func(){}
	fList2 = createF2([]int{108, 12, 30, 4, 5})
	for _, v := range fList2 {
		v()
	}
}

func third() {
	fList3 := []func(){}
	fList3 = createF3([]int{108, 12, 30, 4, 5})
	for _, v := range fList3 {
		v()
	}
}
