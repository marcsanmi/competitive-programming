package main

import (
	"fmt"
	"math"
)

const N = 10001

func main() {
	primeNum := 1
	counter := 1
	for counter <= N {
		primeNum++
		if IsPrime(primeNum) {
			counter++
		}
	}
	fmt.Println(primeNum)
}

func IsPrime(value int) bool {
	for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
