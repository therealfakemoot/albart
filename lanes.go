package albart

import (
	"log"
	"math/rand"

	"github.com/fogleman/gg"
)

type Lane int

const (
	West Lane = iota
	North
	East
	South
)

func LaneAnchor(corners [4]gg.Point) gg.Point {
	xMin, xMax := corners[3].X, corners[1].X
	yMin, yMax := corners[3].Y, corners[1].Y

	x := float64(rand.Intn(int(xMax - xMin)))
	y := float64(rand.Intn(int(yMax - yMin)))
	p := gg.Point{
		X: xMin + x,
		Y: yMin + y,
	}
	log.Printf("%#v", p)
	return p
}

func LaneCorners(lane Lane, x, y, w, h int) [4]gg.Point {
	// x,y are canvas
	// w,h  are the width and height of the elemnt being laned
	// west lane is (0:w, y)
	// north lane is (x, 0:h)
	// east lane is ((x-w):x, y)
	// south lane is (x, (y-h):y)
	if lane == 0 {
		lane = Lane(rand.Intn(3) + 1)
	}
	fx, fy, fw, fh := float64(x), float64(y), float64(w), float64(h)

	// nw, ne, se, sw
	// corners move clockwise starting at top left
	corners := [4]gg.Point{}
	switch lane {
	case West:
		corners[0] = gg.Point{X: 0.0, Y: fy}
		corners[1] = gg.Point{X: fw, Y: fy}
		corners[2] = gg.Point{X: fw, Y: 0.0}
		corners[3] = gg.Point{X: 0.0, Y: 0.0}
	case North:
		corners[0] = gg.Point{X: 0, Y: fy}
		corners[1] = gg.Point{X: fx, Y: fy}
		corners[2] = gg.Point{X: fx, Y: fy - fh}
		corners[3] = gg.Point{X: 0.0, Y: fy - fh}
	case East:
		corners[0] = gg.Point{X: fx - fw, Y: fy}
		corners[1] = gg.Point{X: fx, Y: fy}
		corners[2] = gg.Point{X: fx, Y: 0.0}
		corners[3] = gg.Point{X: fx - fw, Y: 0.0}
	case South:
		corners[0] = gg.Point{X: 0.0, Y: fh}
		corners[1] = gg.Point{X: fx, Y: fh}
		corners[2] = gg.Point{X: fx, Y: fy}
		corners[3] = gg.Point{X: 0.0, Y: 0.0}
	}
	return corners
}
