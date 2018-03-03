package main

import "testing"

type testpair struct {
	input  int
	output bool
}

var testCases = []testpair{
	{0, false},
	{1, false},
	{2, true},
	{3, false},
	{4, true},
	{5, false},
	{6, false},
	{7, false},
	{8, true},
	{9, false},
	{10, false},
	{16, true},
	{20, false},
	{24, false},
	{32, true},
	{50, false},
	{64, true},
	{100, false},
	{500, false},
	{512, true},
	{2048, true},
	{3072, false},
}

func TestDomino(t *testing.T) {
	for _, testCase := range testCases {
		result := degreeOfTwo(testCase.input)
		if result != testCase.output {
			t.Error(
				"For", testCase.input,
				"expected", testCase.output,
				"got", result,
			)
		}
	}
}
