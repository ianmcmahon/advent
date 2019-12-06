package day6

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

type checksumCase struct {
	body   string
	orbits int
}

var checksumCases = []checksumCase{
	{"COM", 0},
	{"B", 1},
	{"D", 3},
	{"L", 7},
}

var example = `
COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
`

func TestExample(t *testing.T) {
	spaceMap := ParseMap(bytes.NewBufferString(example))
	for _, tc := range checksumCases {
		if c := spaceMap.computeChecksum(tc.body); c != tc.orbits {
			t.Errorf("%s orbits expected %d got %d", tc.body, tc.orbits, c)
		}
	}

	if c := spaceMap.Checksum(); c != 42 {
		t.Errorf("The answer to the ultimate question of space, maps, and orbits, is NOT %d", c)
	}
}

func TestPart1(t *testing.T) {
	file, err := os.Open("puzzledata.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	spaceMap := ParseMap(file)
	fmt.Printf("Checksum: %d\n", spaceMap.Checksum())
}

var example2 = `
COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN
`

func TestExample2(t *testing.T) {
	spaceMap := ParseMap(bytes.NewBufferString(example2))
	if v := spaceMap.TransfersRequired("YOU", "SAN"); v != 4 {
		t.Errorf("expected 4, got %d", v)
	}
}

func TestPart2(t *testing.T) {
	file, err := os.Open("puzzledata.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	spaceMap := ParseMap(file)
	fmt.Printf("Transfers Required: %d\n", spaceMap.TransfersRequired("YOU", "SAN"))
}
