package main

import (
	"reflect"
	"testing"
)

type testpair struct {
	in  int
	out int
}

var testCases = []testpair{
	{0, 0},
	{1, 1},
	{2, 1},
	{3, 2},
	{4, 2},
	{5, 3},
	{6, 4},
	{7, 5},
	{8, 6},
	{9, 8},
	{10, 10},
	{11, 12},
	{12, 15},
}

func TestDomino(t *testing.T) {
	for _, testCase := range testCases {
		result := stairs(testCase.in)
		if !reflect.DeepEqual(result, testCase.out) {
			t.Error(
				"For", testCase.in,
				"expected", testCase.out,
				"got", result,
			)
		}
	}
}
