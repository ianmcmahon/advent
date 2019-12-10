package day10

import (
	"fmt"
	"image"
	"math"
	"sort"
	"strings"
)

type Asteroid struct {
	image.Point
}

func ParseMap(input string) []*Asteroid {
	asteroids := []*Asteroid{}
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				asteroids = append(asteroids, &Asteroid{image.Pt(x, y)})
			}
		}
	}
	return asteroids
}

type Peer struct {
	*Asteroid
	r float64
	θ float64
}

type byTheta []*Peer

func (b byTheta) Len() int {
	return len(b)
}

func (b byTheta) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b byTheta) Less(i, j int) bool {
	return b[i].θ < b[j].θ || (b[i].θ == b[j].θ && b[i].r < b[j].r)
}

func (a *Asteroid) Blast(input []*Asteroid) []*Peer {
	peers := a.Neighbors(input)
	sort.Sort(byTheta(peers))

	blasted := make([]*Peer, 0)

	fmt.Printf("Starting with %d peers\n", len(peers))

	iterCount := 0
	for i := 0; len(peers) > 0; {
		if i == len(peers) {
			i = 0
			iterCount = 0
		}
		fmt.Printf("%s to %s is r %.2f θ %.2f ", a, peers[i], peers[i].r, peers[i].θ)
		if iterCount != 0 && len(blasted) > 0 && len(peers) > 1 && peers[i].θ == blasted[len(blasted)-1].θ {
			i++
			fmt.Printf("skipping!\n")
		} else {
			iterCount++
			blasted = append(blasted, peers[i])
			peers = append(peers[0:i], peers[i+1:]...)
			fmt.Printf("blasting!! #%d len peers: %d  len blasted:  %d\n", len(blasted), len(peers), len(blasted))
		}
	}

	fmt.Printf("blasted %d\n", len(blasted))

	return blasted
}

func (a *Asteroid) Visible(input []*Asteroid) int {
	peers := a.Neighbors(input)
	sort.Sort(byTheta(peers))
	y := []*Peer{}
	for i, n := range peers {
		if i == 0 || peers[i-1].θ != n.θ {
			y = append(y, n)
		}
	}
	/*
		for _, p := range y {
			fmt.Printf("%v r %.2f θ %.2f\n", p, p.r, p.θ)
		}
	*/

	return len(y)
}

func (a *Asteroid) Neighbors(peers []*Asteroid) []*Peer {
	neighbors := []*Peer{}
	for _, p := range peers {
		l := image.Rectangle{a.Point, p.Point}
		θ := math.Atan2(1*float64(l.Dx()), -1*float64(l.Dy()))
		if θ < 0 {
			θ += 2 * math.Pi
		}
		r := math.Sqrt(math.Pow(float64(l.Dx()), 2) + math.Pow(float64(l.Dy()), 2))
		if r == 0 {
			continue
		}
		//fmt.Printf("from a (%d, %d) to p (%d, %d) is r %.2f θ %.2f\n", a.X, a.Y, p.X, p.Y, r, θ)
		neighbors = append(neighbors, &Peer{p, r, θ})
	}
	return neighbors
}
