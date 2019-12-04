package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type vector struct {
	x, y int
}

type claim struct {
	num  int
	at   vector
	size vector
}

func mustInt(s string) int {
	i, err := strconv.ParseUint(s, 10, 16)
	if err != nil {
		panic(err)
	}
	return int(i)
}

func claims() ([]claim, error) {
	f, err := os.Open("input")
	if err != nil {
		return nil, err
	}

	r := regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)

	out := []claim{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		matches := r.FindStringSubmatch(scanner.Text())

		c := claim{
			num:  mustInt(matches[1]),
			at:   vector{mustInt(matches[2]), mustInt(matches[3])},
			size: vector{mustInt(matches[4]), mustInt(matches[5])},
		}
		out = append(out, c)
	}

	return out, nil
}

func main() {
	fmt.Printf("the answer shall be.... ")
	claimzzz, err := claims()
	if err != nil {
		panic(err)
	}

	clothSize := 1000

	cloth := make([]int, clothSize*clothSize)

	unsullied := map[int]claim{}

	for _, c := range claimzzz {
		unsullied[c.num] = c
		fmt.Printf("claim: %#v\n", c)
		for x := c.at.x; x < c.at.x+c.size.x; x++ {
			for y := c.at.y; y < c.at.y+c.size.y; y++ {
				i := y*clothSize + x
				if cloth[i] > 0 { // we're attempting to claim a square that's been claimed before
					delete(unsullied, cloth[i]) // neither challenger can survive!
					delete(unsullied, c.num)
				}
				cloth[i] = c.num
			}
		}
	}

	fmt.Printf("unsullied: %#v\n", unsullied)
}
