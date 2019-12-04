package day4

import (
	"fmt"
	"testing"
)

var testCases = map[int]bool{
	122345: true,
	111111: false,
	223450: false,
	123789: false,
	112233: true,
	123444: false,
	111122: true,
}

func TestCheck(t *testing.T) {
	for key, expected := range testCases {
		if got := Check(key); got != expected {
			t.Errorf("%d: expected %v got %v\n", key, expected, got)
		}
	}
}

func TestPart1(t *testing.T) {
	count := 0
	for i := inputRangeMin; i <= inputRangeMax; i++ {
		if Check(i) {
			count++
			fmt.Printf("%d\n", i)
		}
	}
	fmt.Printf("I counted %d\n", count)
}
