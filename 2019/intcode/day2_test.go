package intcode

import (
	"reflect"
	"testing"
)

type testCase struct {
	input    []int
	expected []int
}

var testCases = []testCase{
	{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
	{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
	{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
	{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
}

func TestExamples(t *testing.T) {
	for _, tc := range testCases {
		//fmt.Printf("trying input: %v\n", tc.input)
		output := Process(tc.input)
		//fmt.Printf(" got: %v  expected: %v\n", output, tc.expected)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fail()
		}
	}
}
