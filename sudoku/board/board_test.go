package board

import (
	"fmt"
	"testing"
)

func TestPopulateBoard(t *testing.T) {
	var wantBoard Board
	var gotBoard Board

	wantBoard.PopulateBoard("/Users/nabil/go/src/github.com/Nabil-Daoud/robodoku/puzzles/easy.txt")

	wantBoard.spots[0][0].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
	wantBoard.spots[0][1].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
	wantBoard.spots[0][2].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
	wantBoard.spots[0][3].Solved = true
	wantBoard.spots[0][3].Value = 2
	wantBoard.spots[0][3].Possible = [9]bool{false, true}
	wantBoard.spots[0][4].Solved = true
	wantBoard.spots[0][4].Value = 6
	wantBoard.spots[0][4].Possible = [9]bool{false, false, false, false, false, true}
	wantBoard.spots[0][5].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
	wantBoard.spots[0][6].Solved = true
	wantBoard.spots[0][6].Value = 7
	wantBoard.spots[0][6].Possible = [9]bool{false, false, false, false, false, false, true}
	wantBoard.spots[0][7].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
	wantBoard.spots[0][8].Solved = true
	wantBoard.spots[0][8].Value = 1
	wantBoard.spots[0][8].Possible = [9]bool{true}

	wantBoard.spots[1][0].Solved = true
	wantBoard.spots[1][0].Value = 6
	wantBoard.spots[1][0].Possible = [9]bool{false, false, false, false, false, true}
	wantBoard.spots[1][1].Solved = true
	wantBoard.spots[1][1].Value = 8
	wantBoard.spots[1][1].Possible = [9]bool{false, false, false, false, false, false, false, true}
	wantBoard.spots[1][2].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
	wantBoard.spots[1][3].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
	wantBoard.spots[1][4].Solved = true
	wantBoard.spots[1][4].Value = 7
	wantBoard.spots[1][4].Possible = [9]bool{false, false, false, false, false, false, true}
	wantBoard.spots[1][5].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
	wantBoard.spots[1][6].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
	wantBoard.spots[1][7].Solved = true
	wantBoard.spots[1][7].Value = 9
	wantBoard.spots[1][7].Possible = [9]bool{false, false, false, false, false, false, false, false, true}
	wantBoard.spots[1][8].Possible = [9]bool{true, true, true, true, true, true, true, true, true}

	wantBoard.spots[2][0].Solved = true
	wantBoard.spots[2][0].Value = 1
	wantBoard.spots[2][0].Possible = [9]bool{true}
	wantBoard.spots[2][1].Solved = true
	wantBoard.spots[2][1].Value = 9
	wantBoard.spots[2][1].Possible = [9]bool{false, false, false, false, false, false, false, false, true}
	wantBoard.spots[2][2].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
	wantBoard.spots[2][3].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
	wantBoard.spots[2][4].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
	wantBoard.spots[2][5].Solved = true
	wantBoard.spots[2][5].Value = 4
	wantBoard.spots[2][5].Possible = [9]bool{false, false, false, true}
	wantBoard.spots[2][6].Solved = true
	wantBoard.spots[2][6].Value = 5
	wantBoard.spots[2][6].Possible = [9]bool{false, false, false, false, true}
	wantBoard.spots[2][7].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
	wantBoard.spots[2][8].Possible = [9]bool{true, true, true, true, true, true, true, true, true}

	gotBoard.PopulateBoard("/Users/nabil/go/src/github.com/Nabil-Daoud/robodoku/puzzles/easy.txt")

	for i, row := range wantBoard.spots {
		for j, wantSpot := range row {
			if wantSpot != gotBoard.spots[i][j] {
				t.Errorf("Did not load the spot at %v, %v correctly", i, j)
			}
			if wantSpot.Solved != gotBoard.spots[i][j].Solved {
				t.Errorf("Did not load the spot.Solved at %v, %v correctly. Wanted %t, got %t", i, j, wantSpot.Solved, gotBoard.spots[i][j].Solved)
			}
			if wantSpot.Value != gotBoard.spots[i][j].Value {
				t.Errorf("Did not load the spot.Value at %v, %v correctly. Wanted %v, got %v", i, j, wantSpot.Value, gotBoard.spots[i][j].Value)
			}
			if wantSpot.Possible != gotBoard.spots[i][j].Possible {
				t.Errorf("Did not load the spot.Possible at %v, %v correctly. \nWanted %t, \n   got %t", i, j, wantSpot.Possible, gotBoard.spots[i][j].Possible)
			}
		}
	}
	if wantBoard != gotBoard {
		t.Error("Did not load the board correctly")
	}
}

func TestUpdatePossibleRow(t *testing.T) {
	var wantBoard Board
	var gotBoard Board
	var didUpdate, didNotUpdate bool

	wantBoard.PopulateBoard("/Users/nabil/go/src/github.com/Nabil-Daoud/robodoku/puzzles/easy.txt")
	gotBoard.PopulateBoard("/Users/nabil/go/src/github.com/Nabil-Daoud/robodoku/puzzles/easy.txt")

	fmt.Printf("Spot[0][0].Possible before: %t\n", gotBoard.spots[0][0].Possible)
	fmt.Printf("board.N before: %v\n", gotBoard.N)
	didUpdate = gotBoard.updatePossibleRow()
	fmt.Printf("Spot[0][0].Possible after: %t\n", gotBoard.spots[0][0].Possible)
	fmt.Printf("board.N after: %v\n", gotBoard.N)

	if !didUpdate {
		t.Error("UpdatePossibleRow should have returned true on first call.")
	}

	didNotUpdate = gotBoard.updatePossibleRow()

	if didNotUpdate {
		t.Error("UpdatePossibleRow should have returned false on second call.")
	}

	wantBoard.spots[0][0].Possible = [9]bool{false, false, true, true, true, false, false, true, true}
	wantBoard.spots[0][1].Possible = [9]bool{false, false, true, true, true, false, false, true, true}
	wantBoard.spots[0][2].Possible = [9]bool{false, false, true, true, true, false, false, true, true}
	wantBoard.spots[0][5].Possible = [9]bool{false, false, true, true, true, false, false, true, true}
	wantBoard.spots[0][7].Possible = [9]bool{false, false, true, true, true, false, false, true, true}

	for i, row := range wantBoard.spots {
		for j, wantSpot := range row {
			if wantSpot != gotBoard.spots[i][j] {
				t.Errorf("Did not update the spot at %v, %v correctly", i, j)
			}
			if wantSpot.Solved != gotBoard.spots[i][j].Solved {
				t.Errorf("Did not update the spot.Solved at %v, %v correctly. Wanted %t, got %t", i, j, wantSpot.Solved, gotBoard.spots[i][j].Solved)
			}
			if wantSpot.Value != gotBoard.spots[i][j].Value {
				t.Errorf("Did not update the spot.Value at %v, %v correctly. Wanted %v, got %v", i, j, wantSpot.Value, gotBoard.spots[i][j].Value)
			}
			if wantSpot.Possible != gotBoard.spots[i][j].Possible {
				t.Errorf("Did not update the spot.Possible at %v, %v correctly. \nWanted %t, \n   got %t", i, j, wantSpot.Possible, gotBoard.spots[i][j].Possible)
			}
		}
	}
	if wantBoard != gotBoard {
		t.Error("Did not update the board correctly")
	}
}

func Test(t *testing.T) {
	var tests = []struct {
		value rune
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
