package domino

import "testing"

type testpair struct {
	input  string
	output int
}

var testCases = []testpair{
	{"", 0},
	{"1-1", 1},
	{"1-2,1-2", 1},
	{"3-2,2-1,1-4,4-4,5-4,4-2,2-1", 4},
	{"5-5,5-5,4-4,5-5,5-5,5-5,5-5,5-5,5-5,5-5", 7},
}

func TestDomino(t *testing.T) {
	for _, testCase := range testCases {
		result := domino(testCase.input)
		if result != testCase.output {
			t.Error(
				"For", testCase.input,
				"expected", testCase.output,
				"got", result,
			)
		}
	}
}
