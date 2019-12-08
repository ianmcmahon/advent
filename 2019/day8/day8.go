package day8

import (
	"fmt"
	"io"
)

const layerWidth = 25
const layerHeight = 6

func GetImageLayers(r io.Reader) (out [][]byte, err error) {
	out = make([][]byte, 0)

	for {
		layer := make([]byte, layerWidth*layerHeight)
		var n int
		n, err = r.Read(layer)
		if err != nil {
			return
		}
		if n != layerWidth*layerHeight {
			fmt.Printf("WARN: eof without a full layer, only %d bytes\n", n)
		} else {
			out = append(out, layer)
		}
	}

	return
}

func Checksum(layers [][]byte) int {
	min := layerWidth * layerHeight
	best := -1
	for i, layer := range layers {
		z := CountZeroes(layer)
		if z < min {
			min = z
			best = i
		}
	}
	fmt.Printf("Best layer was layer %d, had %d zeroes\n", best, min)
	fmt.Printf("%v\n", layers[best])

	ones, twos := 0, 0

	for _, c := range layers[best] {
		if c == '1' {
			ones++
		}
		if c == '2' {
			twos++
		}
	}

	return ones * twos
}

func CombineLayers(layers [][]byte) []byte {
	var out = make([]byte, layerHeight*layerWidth)
	copy(out, layers[0]) // top layer goes on first

	for _, layer := range layers {
		// iterate each pixel, if the out map has a transparent, apply the current layer
		for i, c := range layer {
			if out[i] == '2' {
				out[i] = c
			}
		}
	}

	return out
}

func DisplayImage(img []byte) {
	for y := 0; y < layerHeight; y++ {
		for x := 0; x < layerWidth; x++ {
			if img[y*layerWidth+x] != '0' {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func CountZeroes(layer []byte) int {
	zeroes := 0
	for y := 0; y < layerHeight; y++ {
		for x := 0; x < layerWidth; x++ {
			if layer[y*layerWidth+x] == '0' {
				zeroes++
			}
		}
	}
	return zeroes
}
