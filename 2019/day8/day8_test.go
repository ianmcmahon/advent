package day8

import (
	"fmt"
	"os"
	"testing"
)

func TestPart1(t *testing.T) {
	file, err := os.Open("puzzleinput.txt")
	if err != nil {
		fmt.Printf("%v\n", err)
		t.Fail()
	}

	layers, err := GetImageLayers(file)
	if err != nil {
		t.Error(err)
	}

	checksum := Checksum(layers)
	fmt.Printf("Part1: %d\n", checksum)

	image := CombineLayers(layers)

	DisplayImage(image)
}
