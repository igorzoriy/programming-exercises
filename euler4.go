package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(n int) bool {
	testString := strconv.Itoa(n)
	length := len(testString)

	if length == 1 {
		return false
	}

	for i := 0; i < length/2; i++ {
		if testString[i] != testString[length-1-i] {
			return false
		}
	}

	return true
}

func main() {
	max := 999
	result := 0

	for i := max; i > 0; i-- {
		for j := max; j > 0; j-- {
			r := i * j
			if isPalindrome(r) && r > result {
				result = r
			}
		}
	}

	fmt.Println(result)
}
