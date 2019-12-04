package day4

import (
	"fmt"
	"strconv"
)

/*
It is a six-digit number.
The value is within the range given in your puzzle input.
Two adjacent digits are the same (like 22 in 122345).
Going from left to right, the digits never decrease; they only ever increase or stay the same (like 111123 or 135679).
Other than the range rule, the following are true:

111111 meets these criteria (double 11, never decreases).
223450 does not meet these criteria (decreasing pair of digits 50).
123789 does not meet these criteria (no double).
How many different passwords within the range given in your puzzle input meet these criteria?

Your puzzle input is 245318-765747.
.
--- Part Two ---
An Elf just remembered one more important detail: the two adjacent matching digits are not part of a larger group of matching digits.

Given this additional criterion, but still ignoring the range rule, the following are now true:

112233 meets these criteria because the digits never decrease and all repeated digits are exactly two digits long.
123444 no longer meets the criteria (the repeated 44 is part of a larger group of 444).
111122 meets the criteria (even though 1 is repeated more than twice, it still contains a double 22).
How many different passwords within the range given in your puzzle input meet all of the criteria?
*/

var inputRangeMin = 245318
var inputRangeMax = 765747

func Check(v int) bool {
	digits := []byte(fmt.Sprintf("%d", v))
	return CheckIncreasing(digits) && CheckDouble(digits)
}

func CheckDouble(digits []byte) bool {
	m := map[byte]int{}
	for _, c := range digits {
		if _, ok := m[c]; !ok {
			m[c] = 0
		}
		m[c] = m[c] + 1
	}
	for _, d := range m {
		if d == 2 {
			return true
		}
	}
	return false
}

func CheckIncreasing(digits []byte) bool {
	max := 0
	for _, c := range digits {
		d, _ := strconv.Atoi(fmt.Sprintf("%c", c))
		if d < max {
			return false
		}
		max = d
	}
	return true
}
