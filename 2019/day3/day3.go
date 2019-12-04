package day3

import (
	"fmt"
	"image"
	"strconv"
	"strings"
)

type segment struct {
	image.Rectangle
	vertical bool
}

func asSegments(input string) ([]segment, error) {
	vectors := strings.Split(input, ",")

	var segments []segment = make([]segment, 0)

	for _, v := range vectors {
		dir, v := v[0], v[1:]
		dist, err := strconv.ParseInt(v, 10, 16)
		if err != nil {
			return nil, err
		}

		vector := image.Point{}
		segment := segment{}

		if len(segments) > 0 {
			lastSeg := segments[len(segments)-1]
			vector = vector.Add(lastSeg.Max)
			segment.Min = lastSeg.Max
		}

		switch dir {
		case 'R':
			vector = vector.Add(image.Pt(int(dist), 0))
		case 'L':
			vector = vector.Add(image.Pt(-1*int(dist), 0))
		case 'U':
			vector = vector.Add(image.Pt(0, int(dist)))
			segment.vertical = true
		case 'D':
			vector = vector.Add(image.Pt(0, -1*int(dist)))
			segment.vertical = true
		default:
			panic(fmt.Errorf("unknown dir: %c\n", dir))
		}

		segment.Max = vector

		segments = append(segments, segment)
	}

	return segments, nil
}

func findIntersections(pathA, pathB []segment) []image.Point {
	crossings := []image.Point{}
	for _, a := range pathA {
		for _, b := range pathB {
			// do we only care about perpendicular lines?
			if a.vertical == b.vertical {
				//fmt.Printf("skipping %v x %v as they are parallel\n", a, b)
				continue
			}

			a.Rectangle = a.Canon()
			b.Rectangle = b.Canon()

			if a.vertical {
				if b.Min.X < a.Min.X && b.Max.X > a.Max.X {
					if a.Min.Y < b.Min.Y && a.Max.Y > b.Max.Y {
						crossings = append(crossings, image.Pt(a.Min.X, b.Min.Y))
					}
				}
			} else {
				if a.Min.X < b.Min.X && a.Max.X > b.Max.X {
					if b.Min.Y < a.Min.Y && b.Max.Y > a.Max.Y {
						crossings = append(crossings, image.Pt(b.Min.X, a.Min.Y))
					}
				}
			}
		}
	}

	return crossings
}

func abs(i int) int {
	if i < 0 {
		return -1 * i
	}
	return i
}

func findMinDistance(points []image.Point) int {
	min := 65535
	for _, pt := range points {
		v := abs(pt.X) + abs(pt.Y)
		if v < min {
			min = v
		}
	}
	return min
}

func stepsToPoint(path []segment, pt image.Point) int {
	steps := 0
	for _, s := range path {
		if s.contains(pt) {
			//fmt.Printf("after %d steps, %v contains our point %v!\n", steps, s, pt)
			finalSegment := segment{Rectangle: image.Rectangle{s.Min, pt}}
			steps += abs(finalSegment.Dx()) + abs(finalSegment.Dy())
			break
		}
		steps += abs(s.Dx()) + abs(s.Dy())
	}
	return steps
}

func (s segment) contains(pt image.Point) bool {
	c := s.Canon()
	if pt.X == c.Min.X && pt.X == c.Max.X {
		return c.Min.Y <= pt.Y && pt.Y <= c.Max.Y
	}
	if pt.Y == c.Min.Y && pt.Y == c.Max.Y {
		return c.Min.X <= pt.X && pt.X <= c.Max.X
	}
	return false
}
