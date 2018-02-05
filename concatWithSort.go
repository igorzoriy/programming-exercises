package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func bubbleSort(slice []int) []int {
	for {
		sortCounter := 0
		for i := 1; i < len(slice); i++ {
			if slice[i-1] > slice[i] {
				val := slice[i]
				slice[i] = slice[i-1]
				slice[i-1] = val
				sortCounter++
			}
		}
		if sortCounter == 0 {
			return slice
		}
	}
}

func concatWithSort(x, y []int) []int {
	return bubbleSort(append(x, y...))
}

func main() {
	printSlice(concatWithSort(make([]int, 0), make([]int, 0)))
	printSlice(concatWithSort([]int{1, 2, 3}, make([]int, 0)))
	printSlice(concatWithSort(make([]int, 0), []int{1, 2, 3}))
	printSlice(concatWithSort([]int{1, 2, 3}, []int{1, 2, 3}))
	printSlice(concatWithSort([]int{1, 2, 5}, []int{3, 4}))
	printSlice(concatWithSort([]int{1, 2, 5}, []int{3, 7, 8, 9}))
	printSlice(concatWithSort([]int{10, 11, 15, 22}, []int{3, 4, 20}))
}
