package day2

import (
	"fmt"
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
		fmt.Printf("trying input: %v\n", tc.input)
		output := process(tc.input)
		fmt.Printf(" got: %v  expected: %v\n", output, tc.expected)
		if !reflect.DeepEqual(output, tc.expected) {
			t.Fail()
		}
	}
}

var input = []int{1, 0, 0, 3, 1, 1, 2, 3, 1, 3, 4, 3, 1, 5, 0, 3, 2, 10, 1, 19, 1, 6, 19, 23, 1, 23, 13, 27, 2, 6, 27, 31, 1, 5, 31, 35, 2, 10, 35, 39, 1, 6, 39, 43, 1, 13, 43, 47, 2, 47, 6, 51, 1, 51, 5, 55, 1, 55, 6, 59, 2, 59, 10, 63, 1, 63, 6, 67, 2, 67, 10, 71, 1, 71, 9, 75, 2, 75, 10, 79, 1, 79, 5, 83, 2, 10, 83, 87, 1, 87, 6, 91, 2, 9, 91, 95, 1, 95, 5, 99, 1, 5, 99, 103, 1, 103, 10, 107, 1, 9, 107, 111, 1, 6, 111, 115, 1, 115, 5, 119, 1, 10, 119, 123, 2, 6, 123, 127, 2, 127, 6, 131, 1, 131, 2, 135, 1, 10, 135, 0, 99, 2, 0, 14, 0}

func TestPart1(t *testing.T) {
	// make changes required by problem:
	// Once you have a working computer, the first step is to restore the gravity assist program (your puzzle input) to the "1202 program alarm" state it had just before the last computer caught fire. To do this, before running the program, replace position 1 with the value 12 and replace position 2 with the value 2. What value is left at position 0 after the program halts?

	part1Input := make([]int, len(input))
	copy(part1Input, input)
	part1Input[1] = 12
	part1Input[2] = 2

	out := process(part1Input)
	fmt.Printf("%v\n", out)
}

func TestPart2(t *testing.T) {
	// The inputs should still be provided to the program by replacing the values at addresses 1 and 2, just like before. In this program, the value placed in address 1 is called the noun, and the value placed in address 2 is called the verb. Each of the two input values will be between 0 and 99, inclusive.

	// Find the input noun and verb that cause the program to produce the output 19690720. What is 100 * noun + verb? (For example, if noun=12 and verb=2, the answer would be 1202.)

	part2Input := make([]int, len(input))

	/*
		// here's the lazy brute force method
		for noun := 0; noun < 100; noun++ {
			for verb := 50; verb < 51; verb++ {
				copy(part2Input, input)
				part2Input[1] = noun
				part2Input[2] = verb
				output := process(part2Input)
				fmt.Printf("(%d, %d) = %d\n", noun, verb, output[0])
			}
		}
	*/

	// ok for an iterative solution, let's start at 50, 50
	// the noun makes large changes, the verb makes small changes
	// so we will effectively binary search for the noun range, then
	// binary search for the verb

	minnoun, minverb := 50, 0
	maxnoun, maxverb := 50, 99

	//expected := 1690720
	expected := 19690720

	// let's prime min/max with the actual value at 50,50
	copy(part2Input, input)
	part2Input[1], part2Input[2] = minnoun, minverb
	output := process(part2Input)
	fmt.Printf("(%d, %d) = %d\n", minnoun, minverb, output[0])
	min, max := output[0], output[0]

	// now, if our output is lower than min, we need to hunt backwards
	// if it's higher than max, we need to hunt forwards
	for {
		//fmt.Printf("currently: (%d, %d) to (%d, %d)  (%d - %d)\n", minnoun, minverb, maxnoun, maxverb, min, max)
		// this is terminating condition
		if maxnoun == minnoun && min <= expected && expected <= max {
			fmt.Printf("found the noun? between (%d,%d) and (%d,%d) (%d - %d)\n", minnoun, minverb, maxnoun, maxverb, min, max)
			break
		}

		if expected > min {
			if expected > max {
				minnoun = maxnoun
				maxnoun = minnoun + (100-minnoun)/2
				copy(part2Input, input)
				part2Input[1], part2Input[2] = maxnoun, maxverb
				output := process(part2Input)
				fmt.Printf("(%d, %d) = %d\n", maxnoun, maxverb, output[0])
				max = output[0]
				if max < expected { // if noun, 99 is still too small, we can move minnoun up to here
					minnoun = maxnoun
				}
			} else {
				minnoun += (maxnoun + 1 - minnoun) / 2
				copy(part2Input, input)
				part2Input[1], part2Input[2] = minnoun, minverb
				output := process(part2Input)
				fmt.Printf("(%d, %d) = %d\n", minnoun, minverb, output[0])
				min = output[0]
				if min > expected {
					maxnoun = minnoun
				}
			}
		}
		if expected < min {
			fmt.Printf("expected < min\n")
			// do the opposite, but I don't think we will need to?
		}
	}

	// at this point, minnoun == maxnoun == solution noun
	// so we need to iterate to find the verb, same concept

	for {
		// this is terminating condition
		if maxverb == minverb && min <= expected && expected <= max {
			fmt.Printf("found the verb? between (%d,%d) and (%d,%d) (%d - %d)\n", minnoun, minverb, maxnoun, maxverb, min, max)
			break
		}

		try := minverb + (maxverb-minverb)/2

		copy(part2Input, input)
		part2Input[1], part2Input[2] = minnoun, try
		output := process(part2Input)
		fmt.Printf("(%d, %d) = %d\n", minnoun, try, output[0])
		if output[0] >= expected {
			max = output[0]
			maxverb = try
		}
		if output[0] <= expected {
			min = output[0]
			minverb = try
		}
	}
}
