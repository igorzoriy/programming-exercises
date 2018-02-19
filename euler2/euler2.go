package main

import "fmt"

func main() {
	sum := 0
	fibonacci := []int{0, 1}

	for i := 2; ; i++ {
		current := fibonacci[i-2] + fibonacci[i-1]

		if current > 4000000 {
			break
		}

		if current%2 == 0 {
			sum += current
		}
		fibonacci = append(fibonacci, current)
	}
	fmt.Println(sum)
}
