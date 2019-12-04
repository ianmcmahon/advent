package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func lines() ([]int64, error) {
	f, err := os.Open("input")
	if err != nil {
		return nil, err
	}

	out := []int64{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		change := scanner.Text()
		s, err := strconv.ParseInt(change, 10, 32)
		if err != nil {
			return nil, err
		}
		out = append(out, s)
	}

	return out, nil
}

func main() {
	freq := int64(0)
	seen := map[int64]bool{}

	changes, err := lines()
	if err != nil {
		panic(err)
	}

	for {
		for _, d := range changes {
			seen[freq] = true
			freq += d
			if _, ok := seen[freq]; ok {
				fmt.Printf("seen %d twice\n", freq)
				os.Exit(0)
			}
		}
	}

}
