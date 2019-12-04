package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func input() ([]string, error) {
	f, err := os.Open("input")
	if err != nil {
		return nil, err
	}
	lines := []string{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {
	lines, err := input()
	if err != nil {
		panic(err)
	}

	threes := 0
	twos := 0

	reverse := []string{}

	for _, l := range lines {
		reverse = append(reverse, Reverse(l))

		isThree := false
		isTwo := false
		hist := map[rune]int{}
		for _, r := range []rune(l) {
			hist[r]++
		}

		for _, v := range hist {
			if v == 3 {
				isThree = true
			}
			if v == 2 {
				isTwo = true
			}
		}
		if isThree {
			threes++
		}
		if isTwo {
			twos++
		}
		fmt.Println()
	}

	fmt.Printf("checksum: %d x %d = %d\n", threes, twos, threes*twos)

	sort.Sort(sort.StringSlice(reverse))

	lastword := ""
	for i, v := range reverse {
		if len(v) != len(lastword) {
			lastword = v
			continue
		}
		differ := charsDiffer(v, lastword)
		if differ == 1 {
			fmt.Printf("%s\n%s\n", Reverse(v), Reverse(reverse[i-1]))
			fmt.Printf("%s differs from above by %d\n", Reverse(v), differ)
		}
		lastword = v
	}
}

func charsDiffer(a, b string) int {
	if len(a) != len(b) {
		return -1
	}

	differ := 0
	for i, _ := range a {
		if a[i] != b[i] {
			differ++
		}
	}
	return differ
}
