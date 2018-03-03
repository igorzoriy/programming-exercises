package main

import "fmt"

func qsort(slice []int) []int {
	length := len(slice)

	if length < 2 {
		return slice
	}

	baseIndex := int(length / 2)
	base := slice[baseIndex]
	left := make([]int, 0, length)
	middle := make([]int, 0, length)
	right := make([]int, 0, length)
	for _, item := range slice {
		if item < base {
			left = append(left, item)
		} else if item == base {
			middle = append(middle, item)
		} else {
			right = append(right, item)
		}
	}

	left = qsort(left)
	right = qsort(right)

	return append(left, append(middle, right...)...)
}

func main() {
	fmt.Println(qsort([]int{6, 7, 2, 4, 1, 0, -10, 6, 15, 42, 11, 7}))
}
