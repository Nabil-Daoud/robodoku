package spot

import (
	"testing"
)

func Test(t *testing.T) {
	var tests = []struct {
		value uint8
		want  int
	}{
		{1, 1},
		{2, 2},
	}
	var got int
	for _, c := range tests {
		got = int(c.value)
		if got != c.want {
			t.Errorf("Cell(%v) == %v, want %v", c.value, got, c.want)
		}
	}
}

func TestSolveSpot(t *testing.T) {
	var tests = []struct {
		value uint8
		spot  Spot
	}{
		{1, Spot{false, ' ', [9]bool{true, true, true, true, true, true, true, true, true}}},
		{2, Spot{false, ' ', [9]bool{true, true, true, true, true, true, true, true, true}}},
		{3, Spot{false, ' ', [9]bool{true, true, true, true, true, true, true, true, true}}},
	}

	for _, c := range tests {
		c.spot.SolveSpot(c.value)
		if c.value != c.spot.Value {
			t.Errorf("spot.Value expected %v, got %v.", c.value, c.spot.Value)
		}
		if !c.spot.Solved {
			t.Error("spot.Solved should be true.")
		}
		if !c.spot.Possible[c.value-1] {
			t.Error("spot.Possible should be true for solution value.")
		}
		for i, d := range c.spot.Possible {
			if i != int(c.value-1) && d {
				t.Error("spot.Possible should be false for all other values.")
			}
		}
	}
}
