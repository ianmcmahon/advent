package day5

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/ianmcmahon/advent/2019/intcode"
)

type testCase struct {
	input          []int
	expectedMem    []int
	userInput      []int
	expectedOutput []int
}

var testCases = []testCase{
	{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}, nil, []int{}},
	{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}, nil, []int{}},
	{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}, nil, []int{}},
	{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}, nil, []int{}},
	{[]int{1002, 4, 3, 4, 33}, []int{1002, 4, 3, 4, 99}, nil, []int{}},
	{[]int{1101, 100, -1, 4, 0}, []int{1101, 100, -1, 4, 99}, nil, []int{}},
	{[]int{3, 0, 4, 0, 99}, []int{1, 0, 4, 0, 99}, []int{1}, []int{1}},
	{ // position mode, 1 EQ 8 == 0
		[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
		[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, 0, 8},
		[]int{1},
		[]int{0},
	}, { // position mode, 8 EQ 8 == 1
		[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
		[]int{3, 9, 8, 9, 10, 9, 4, 9, 99, 1, 8},
		[]int{8},
		[]int{1},
	}, { // position mode, 7 LT 8 == 1
		[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
		[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, 1, 8},
		[]int{7},
		[]int{1},
	}, { // position mode, 8 LT 8 == 0
		[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
		[]int{3, 9, 7, 9, 10, 9, 4, 9, 99, 0, 8},
		[]int{8},
		[]int{0},
	}, { // immediate mode, 1 EQ 8 = 0
		[]int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
		[]int{3, 3, 1108, 0, 8, 3, 4, 3, 99},
		[]int{1},
		[]int{0},
	}, { // immediate mode, 8 EQ 8 = 1
		[]int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
		[]int{3, 3, 1108, 1, 8, 3, 4, 3, 99},
		[]int{8},
		[]int{1},
	}, { // immediate mode, 7 LT 8 = 1
		[]int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
		[]int{3, 3, 1107, 1, 8, 3, 4, 3, 99},
		[]int{7},
		[]int{1},
	}, { // immediate mode, 8 LT 8 = 0
		[]int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
		[]int{3, 3, 1107, 0, 8, 3, 4, 3, 99},
		[]int{8},
		[]int{0},
	}, { // position mode jump test, outputs 0 for 0 input
		[]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
		[]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, 0, 0, 1, 9},
		[]int{0},
		[]int{0},
	}, { // position mode jump test, outputs 1 for nonzero input
		[]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
		[]int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, 100, 1, 1, 9},
		[]int{100},
		[]int{1},
	}, { // immediate mode jump test, outputs 0 for 0 input
		[]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
		[]int{3, 3, 1105, 0, 9, 1101, 0, 0, 12, 4, 12, 99, 0},
		[]int{0},
		[]int{0},
	}, { // immediate mode jump test, outputs 1 for nonzero input
		[]int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
		[]int{3, 3, 1105, 100, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
		[]int{100},
		[]int{1},
	},
}

func TestExamples(t *testing.T) {
	for _, tc := range testCases {
		//fmt.Printf("trying input: %v\n", tc.input)
		memory, output := intcode.Process(tc.input, tc.userInput...)
		if !reflect.DeepEqual(memory, tc.expectedMem) {
			t.Errorf(" got: %v  expected: %v\n", memory, tc.expectedMem)
		}
		if !reflect.DeepEqual(output, tc.expectedOutput) {
			t.Errorf(" got: %v  expected: %v\n", output, tc.expectedOutput)
		}
	}
}

func TestLargerExample(t *testing.T) {
	program := []int{
		3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31,
		1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104,
		999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99,
	}
	/*
		The above example program uses an input instruction to ask for a single number.
		The program will then output 999 if the input value is below 8, output 1000
		if the input value is equal to 8, or output 1001 if the input value is greater than 8.
	*/
	io := map[int][]int{
		1:   []int{999},
		7:   []int{999},
		8:   []int{1000},
		9:   []int{1001},
		100: []int{1001},
	}

	for input, expected := range io {
		_, output := intcode.Process(program, input)
		if !reflect.DeepEqual(output, expected) {
			t.Errorf(" got: %v  expected: %v\n", output, expected)
		}
	}
}

func TestPart1(t *testing.T) {
	input := make([]int, len(puzzleInput))
	copy(input, puzzleInput)

	_, output := intcode.Process(input, 1)
	fmt.Printf("Diagnostic code: %d\n", output[len(output)-1])
}

func TestPart2(t *testing.T) {
	input := make([]int, len(puzzleInput))
	copy(input, puzzleInput)

	_, output := intcode.Process(input, 5)
	fmt.Printf("Diagnostic code: %d\n", output[len(output)-1])
}

var puzzleInput = []int{3, 225, 1, 225, 6, 6, 1100, 1, 238, 225, 104, 0, 1101, 32, 43, 225, 101, 68, 192, 224, 1001, 224, -160, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 2, 224, 1, 223, 224, 223, 1001, 118, 77, 224, 1001, 224, -87, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 6, 224, 1, 223, 224, 223, 1102, 5, 19, 225, 1102, 74, 50, 224, 101, -3700, 224, 224, 4, 224, 1002, 223, 8, 223, 1001, 224, 1, 224, 1, 223, 224, 223, 1102, 89, 18, 225, 1002, 14, 72, 224, 1001, 224, -3096, 224, 4, 224, 102, 8, 223, 223, 101, 5, 224, 224, 1, 223, 224, 223, 1101, 34, 53, 225, 1102, 54, 10, 225, 1, 113, 61, 224, 101, -39, 224, 224, 4, 224, 102, 8, 223, 223, 101, 2, 224, 224, 1, 223, 224, 223, 1101, 31, 61, 224, 101, -92, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 4, 224, 1, 223, 224, 223, 1102, 75, 18, 225, 102, 48, 87, 224, 101, -4272, 224, 224, 4, 224, 102, 8, 223, 223, 1001, 224, 7, 224, 1, 224, 223, 223, 1101, 23, 92, 225, 2, 165, 218, 224, 101, -3675, 224, 224, 4, 224, 1002, 223, 8, 223, 101, 1, 224, 224, 1, 223, 224, 223, 1102, 8, 49, 225, 4, 223, 99, 0, 0, 0, 677, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1105, 0, 99999, 1105, 227, 247, 1105, 1, 99999, 1005, 227, 99999, 1005, 0, 256, 1105, 1, 99999, 1106, 227, 99999, 1106, 0, 265, 1105, 1, 99999, 1006, 0, 99999, 1006, 227, 274, 1105, 1, 99999, 1105, 1, 280, 1105, 1, 99999, 1, 225, 225, 225, 1101, 294, 0, 0, 105, 1, 0, 1105, 1, 99999, 1106, 0, 300, 1105, 1, 99999, 1, 225, 225, 225, 1101, 314, 0, 0, 106, 0, 0, 1105, 1, 99999, 1107, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 329, 1001, 223, 1, 223, 1007, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 344, 1001, 223, 1, 223, 108, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 359, 1001, 223, 1, 223, 7, 226, 226, 224, 1002, 223, 2, 223, 1005, 224, 374, 101, 1, 223, 223, 107, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 389, 1001, 223, 1, 223, 1007, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 404, 1001, 223, 1, 223, 1107, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 419, 1001, 223, 1, 223, 108, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 434, 1001, 223, 1, 223, 1108, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 449, 1001, 223, 1, 223, 1108, 677, 226, 224, 102, 2, 223, 223, 1005, 224, 464, 1001, 223, 1, 223, 107, 226, 226, 224, 102, 2, 223, 223, 1006, 224, 479, 1001, 223, 1, 223, 1008, 226, 226, 224, 102, 2, 223, 223, 1005, 224, 494, 101, 1, 223, 223, 7, 677, 226, 224, 1002, 223, 2, 223, 1005, 224, 509, 101, 1, 223, 223, 8, 226, 677, 224, 1002, 223, 2, 223, 1006, 224, 524, 1001, 223, 1, 223, 1007, 226, 226, 224, 1002, 223, 2, 223, 1006, 224, 539, 101, 1, 223, 223, 1008, 677, 677, 224, 1002, 223, 2, 223, 1006, 224, 554, 101, 1, 223, 223, 1108, 677, 677, 224, 102, 2, 223, 223, 1006, 224, 569, 101, 1, 223, 223, 1107, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 584, 1001, 223, 1, 223, 8, 677, 226, 224, 1002, 223, 2, 223, 1006, 224, 599, 101, 1, 223, 223, 1008, 677, 226, 224, 102, 2, 223, 223, 1006, 224, 614, 1001, 223, 1, 223, 7, 226, 677, 224, 1002, 223, 2, 223, 1005, 224, 629, 101, 1, 223, 223, 107, 226, 677, 224, 102, 2, 223, 223, 1005, 224, 644, 101, 1, 223, 223, 8, 677, 677, 224, 102, 2, 223, 223, 1005, 224, 659, 1001, 223, 1, 223, 108, 677, 677, 224, 1002, 223, 2, 223, 1005, 224, 674, 101, 1, 223, 223, 4, 223, 99, 226}
