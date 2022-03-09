package albart

import (
	"math/rand"
	"testing"

	"github.com/fogleman/gg"
)

func Test_LaneAnchor(t *testing.T) {
	rand.Seed(8675309)
	corners := [4]gg.Point{
		gg.Point{0, 1000},
		gg.Point{1000, 1000},
		gg.Point{1000, 800},
		gg.Point{0, 800},
	}
	p := LaneAnchor(corners)

	if p.X < 0 || p.X > 1000 {
		t.Logf("expected x anchor to be within range 0-1000 for north lane, got %.2f", p.X)
		t.Fail()
	}

	if p.Y < 800 || p.Y > 1000 {
		t.Logf("expected y anchor to be within range 800-1000 for north lane, got %.2f", p.Y)
		t.Fail()
	}

}

func Test_LaneCorners(t *testing.T) {
	rand.Seed(8675309)

	t.Run("north", func(t *testing.T) {
		t.Skip()
	})

}
