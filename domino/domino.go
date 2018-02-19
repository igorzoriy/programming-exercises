package main

import (
	"fmt"
	"strings"
)

func domino(str string) int {
	dominos := strings.Split(str, ",")
	if len(dominos) == 0 {
		return 0
	} else if len(dominos) == 1 {
		return 1
	}

	maxCounter := 0
	counter := 1
	for i := 1; i < len(dominos); i++ {
		prev := dominos[i-1]
		current := dominos[i]
		if prev[2] == current[0] {
			counter++
		} else {
			if counter > maxCounter {
				maxCounter = counter
			}
			counter = 1
		}
	}
	if counter > maxCounter {
		maxCounter = counter
	}

	return maxCounter
}

func main() {
	fmt.Println(domino("1-1"))
	fmt.Println(domino("1-2,1-2"))
	fmt.Println(domino("3-2,2-1,1-4,4-4,5-4,4-2,2-1"))
	fmt.Println(domino("5-5,5-5,4-4,5-5,5-5,5-5,5-5,5-5,5-5,5-5"))
}
