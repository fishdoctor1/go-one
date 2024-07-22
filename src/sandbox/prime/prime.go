package prime

import (
	"fmt"
	"math"
)

var notPrimes [1000000]bool

func init() {
	fmt.Println("init")
	sqrtLen := int(math.Ceil(math.Sqrt(float64(len(notPrimes)))))
	for i := 2; i < sqrtLen; i++ {
		if notPrimes[i] {
			//if true continue
			//default false at var notPrimes
			continue
		}
		for j := i; j < len(notPrimes); j += i {
			notPrimes[j] = true
		}
	}
	fmt.Println("init done")
}

func IsPrime(x int) bool {
	return !notPrimes[x]
}
