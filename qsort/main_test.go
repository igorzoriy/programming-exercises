package main

import (
	"reflect"
	"testing"
)

type testpair struct {
	input  []int
	output []int
}

var testCases = []testpair{
	{[]int{}, []int{}},
	{[]int{42}, []int{42}},
	{[]int{3, 1}, []int{1, 3}},
	{[]int{1, 1, 2, 2}, []int{1, 1, 2, 2}},
	{[]int{2, 5, 7, 4, 8, 1}, []int{1, 2, 4, 5, 7, 8}},
	{
		[]int{-20, 7, 102, -55, 0, 12, 8},
		[]int{-55, -20, 0, 7, 8, 12, 102},
	},
}

func TestDomino(t *testing.T) {
	for _, testCase := range testCases {
		result := qsort(testCase.input)
		if !reflect.DeepEqual(result, testCase.output) {
			t.Error(
				"For", testCase.input,
				"expected", testCase.output,
				"got", result,
			)
		}
	}
}
