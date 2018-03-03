package main

import "fmt"

func degreeOfTwo(n int) bool {
	if n == 0 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 != 0 {
		return false
	}

	return degreeOfTwo(n / 2)
}

func main() {
	fmt.Println(degreeOfTwo(2048))
}
