package spot

// N is the dimention of the board and the number of possible values
var N int

// Spot is an individual spot on a board. A.k.a cell or square.
type Spot struct {
	Solved   bool
	Value    rune
	Possible []bool
}

// NewSpot is the constructor for a spot
func NewSpot(n int, Value rune) Spot {
	newSpot := Spot{}
	newSpot.Value = Value
	if Value == ' ' {
		newSpot.Possible = []bool{true}
	} else {
		newSpot.Possible = []bool{false}
		newSpot.Possible[int(Value)-1] = true
	}
	return newSpot
}
