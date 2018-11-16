package spot

import (
	"testing"
)

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

func TestUpdatePossible(t *testing.T) {
	var tests = []struct {
		possible [9]bool
		spot     Spot
	}{
		{[9]bool{}, Spot{false, ' ', [9]bool{true, true, true, true, true, true, true, true, true}}},
		{[9]bool{true}, Spot{false, ' ', [9]bool{true, true, true, true, true, true, true, true, true}}},
		{[9]bool{false, false, true, true, true}, Spot{false, ' ', [9]bool{true, true, true, true, true, true, true, true, true}}},
	}

	for _, c := range tests {
		c.spot.UpdatePossible(c.possible)

		if c.possible != c.spot.Possible {
			t.Errorf("Did not update possible array correctly.\nwanted: %t\ngot:    %t", c.possible, c.spot.Possible)
		}
	}
}

func TestSinglePossible(t *testing.T) {
	var tests = []struct {
		spot      Spot
		wantBool  bool
		wantValue int
	}{
		{Spot{false, ' ', [9]bool{true}}, true, 1},
		{Spot{false, ' ', [9]bool{false, true}}, true, 2},
		{Spot{false, ' ', [9]bool{false, false, true}}, true, 3},
		{Spot{false, ' ', [9]bool{false, false, false, true}}, true, 4},
		{Spot{false, ' ', [9]bool{false, false, false, false, true}}, true, 5},
		{Spot{false, ' ', [9]bool{false, false, false, false, false, true}}, true, 6},
		{Spot{false, ' ', [9]bool{false, false, false, false, false, false, true}}, true, 7},
		{Spot{false, ' ', [9]bool{false, false, false, false, false, false, false, true}}, true, 8},
		{Spot{false, ' ', [9]bool{false, false, false, false, false, false, false, false, true}}, true, 9},

		{Spot{false, ' ', [9]bool{true, true}}, false, 0},
		{Spot{false, ' ', [9]bool{true, false, true}}, false, 0},
		{Spot{false, ' ', [9]bool{true, true, true, true, true}}, false, 0},
		{Spot{false, ' ', [9]bool{}}, false, 0},
	}

	var (
		gotBool  bool
		gotValue int
	)

	for _, c := range tests {
		gotBool, gotValue = c.spot.SinglePossible()
		if gotBool != c.wantBool {
			t.Errorf("SinglePossible for %t is %t, want %t", c.spot.Possible, gotBool, c.wantBool)
		}
		if gotValue != c.wantValue {
			t.Errorf("SinglePossible for %t value is %v, want %v", c.spot.Possible, gotValue, c.wantValue)
		}
	}
}

// func Test(t *testing.T) {
// 	var tests = []struct {
// 		value uint8
// 		want  int
// 	}{
// 		{1, 1},
// 		{2, 2},
// 	}
// 	var got int
// 	for _, c := range tests {
// 		got = int(c.value)
// 		if got != c.want {
// 			t.Errorf("Cell(%v) == %v, want %v", c.value, got, c.want)
// 		}
// 	}
// }
