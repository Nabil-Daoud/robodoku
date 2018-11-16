package spot

// N is the dimention of the board and the number of possible values
// var N int

// Spot is an individual spot on a board. A.k.a cell or square.
type Spot struct {
	Solved   bool
	Value    uint8
	Possible [9]bool
}

// SolveSpot sets the Solved to true, Value to input value, and the values of
// the Possible slice appropriately.
func (spot *Spot) SolveSpot(value uint8) {
	spot.Solved = true
	spot.Value = value
	for i := range spot.Possible {
		if i == int(value-1) {
			spot.Possible[i] = true
		} else {
			spot.Possible[i] = false
		}
	}
}

// UpdatePossible updates the array of possibe values the spot can hold.
func (spot *Spot) UpdatePossible(possible [9]bool) {
	spot.Possible = possible
}

// NewSpot is the constructor for a spot
// func NewSpot(n int, Value uint8) Spot {
// 	newSpot := Spot{}
// 	newSpot.Value = Value
// 	if Value == ' ' {
// 		newSpot.Possible = []bool{true}
// 	} else {
// 		newSpot.Possible = []bool{false}
// 		newSpot.Possible[int(Value)-1] = true
// 	}
// 	return newSpot
// }
