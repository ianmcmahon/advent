package day6

import (
	"bufio"
	"io"
	"strings"
)

type SpaceMap map[string]string

func ParseMap(input io.Reader) SpaceMap {
	spaceMap := make(map[string]string, 0)

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Split(line, ")")
		if len(fields) != 2 {
			continue
		}
		//fmt.Printf("%s - %s orbits %s\n", line, fields[1], fields[0])
		spaceMap[fields[1]] = fields[0]
	}

	return spaceMap
}

func (spaceMap SpaceMap) Checksum() int {
	// iterate over all the keys in the map,
	// this is every body that orbits something
	// for each body, recursively count the orbits
	var checksum int

	for body, _ := range spaceMap {
		checksum += spaceMap.computeChecksum(body)
	}

	return checksum
}

func (spaceMap SpaceMap) computeChecksum(body string) int {
	if body == "COM" {
		return 0
	}
	return 1 + spaceMap.computeChecksum(spaceMap[body])
}

func (spaceMap SpaceMap) pathFromCOM(body string) []string {
	if body == "COM" {
		return []string{"COM"}
	}
	return append(spaceMap.pathFromCOM(spaceMap[body]), body)
}

func longest(a, b []string) []string {
	if len(a) > len(b) {
		return a
	}
	return b
}

func reverse(ss []string) {
	last := len(ss) - 1
	for i := 0; i < len(ss)/2; i++ {
		ss[i], ss[last-i] = ss[last-i], ss[i]
	}
}

func (spaceMap SpaceMap) TransfersRequired(A, B string) int {
	// get the paths from COM to the body we are orbiting
	// dereferencing here to chop "YOU" and "SAN" out of the path
	pathA := spaceMap.pathFromCOM(spaceMap[A])
	pathB := spaceMap.pathFromCOM(spaceMap[B])

	// both paths start with COM
	// pull all common nodes off both paths
	var lastNode string
	for pathA[0] == pathB[0] {
		lastNode = pathA[0]
		pathA = pathA[1:]
		pathB = pathB[1:]
	}
	// now they've diverged, we stored the common ancestor
	// reverse pathA and stick them together with common node in between

	reverse(pathA)
	pathA = append(pathA, lastNode)
	pathA = append(pathA, pathB...)

	// N nodes connected by N-1 edges, the transfers
	return len(pathA) - 1
}
