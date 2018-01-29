package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n < 4 {
		return false
	}

	for i := 2; i <= n/2; i++ {
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
		factor--
	}

	fmt.Println(factor)
}
