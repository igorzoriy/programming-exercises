package main

import "fmt"

func divisible(number, limit int) bool {
	for i := 2; i <= limit; i++ {
		if number%i != 0 {
			return false
		}
	}

	return true
}

func main() {
	n := 20
	value := 1

	for {
		if divisible(value, n) {
			break
		}
		value++
	}
	fmt.Println(value)
}
