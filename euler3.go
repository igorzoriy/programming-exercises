package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n < 4 {
		return false
	}
	max := int(math.Ceil(float64(n / 2)))
	for i := 2; i <= max; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	n := 600851475143
	factor := int(math.Sqrt(float64(n))) + 1

	for factor > 0 {
		if n%factor == 0 && isPrime(factor) {
			break
		}
		factor -= 1
	}

	fmt.Println(factor)
}
