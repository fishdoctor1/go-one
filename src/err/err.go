package main

import (
	"errors"
	"fmt"
	"log"
)

func withdraw(a int) (int, error) {
	if a > 10 {
		return 0, errors.New("insuf a")
	}
	fmt.Errorf("sdsd %v", "db down")
	return a, nil
}
func main() {
	log.SetFlags(1) //format for log.fatalf
	amount, err := withdraw(11)
	if err != nil {
		log.Fatalf("main: %v", err)

		// fmt.Println("main:", err)
		// os.Exit(1)
		// return
	}
	fmt.Println("Please collect your money:", amount)

}
