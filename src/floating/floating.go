package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Println(x, -x, 1/x, 1/-x, x/x) // 0 -0 +Inf -Inf NaN
	fmt.Println(math.IsNaN(x / x))     //true
	fmt.Println(math.IsInf(-1/x, 0))   //true
	fmt.Println(math.IsInf(1/x, 0))    //true
	fmt.Println(math.IsInf(1/x, 1))    //true
	fmt.Println(math.IsInf(1/x, -1))   //false
}
