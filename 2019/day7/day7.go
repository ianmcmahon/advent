package day7

import (
	"github.com/ianmcmahon/advent/2019/intcode"
)

var part1phases = []int{0, 1, 2, 3, 4}
var part2phases = []int{5, 6, 7, 8, 9}

func Signal(phases []int) int {
	signal := 0

	for _, p := range phases {
		_, output := intcode.Process(program(), p, signal)
		signal = output[0]
	}

	return signal
}

func MaxSignal() int {
	var max = 0
	permutations := permutePhases(part1phases)
	for _, phases := range permutations {
		if signal := Signal(phases); signal > max {
			max = signal
		}
	}
	return max
}

func MaxFeedbackSignal() int {
	var max = 0
	permutations := permutePhases(part2phases)
	for _, phases := range permutations {
		signal, err := AmplifiersHeadless(phases)
		if err != nil {
			panic(err)
		}
		if signal > max {
			max = signal
		}
	}
	return max
}

func permutePhases(phases []int) (permuts [][]int) {
	var xs = make([]int, len(phases))
	copy(xs, phases)

	var rc func([]int, int)
	rc = func(a []int, k int) {
		if k == len(a) {
			permuts = append(permuts, append([]int{}, a...))
		} else {
			for i := k; i < len(xs); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}

var program = func() []int {
	var p = make([]int, len(puzzleInput))
	copy(p, puzzleInput)
	return p
}
