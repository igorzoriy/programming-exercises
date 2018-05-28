package main

import "fmt"

func stairsStep(n, prevLevel int) int {
	if n < 1 {
		return 0
	} else if n < 3 {
		if n < prevLevel {
			return 1
		} else {
			return 0
		}
	}

	count := 0
	if n < prevLevel {
		count = 1
	}

	for i := 1; i < n; i++ {
		baseLevel := n - i
		if baseLevel >= prevLevel {
			continue
		}
		count += stairsStep(i, baseLevel)
	}

	return count
}

func stairs(n int) int {
	return stairsStep(n, n+1)
}

func main() {
	fmt.Println("1 - ", stairs(1))
	fmt.Println("2 - ", stairs(2))
	fmt.Println("3 - ", stairs(3))
	fmt.Println("4 - ", stairs(4))
	fmt.Println("5 - ", stairs(5))
	fmt.Println("6 - ", stairs(6))
	fmt.Println("7 - ", stairs(7))
	fmt.Println("8 - ", stairs(8))
	fmt.Println("9 - ", stairs(9))
	fmt.Println("10 - ", stairs(10))
	fmt.Println("11 - ", stairs(11))
	fmt.Println("12 - ", stairs(12))
}
