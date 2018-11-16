package board

import (
	"testing"
)

func TestPopulateBoard(t *testing.T) {
	var wantBoard Board
	var gotBoard Board

	// Only hand enter the first three columns and use this line to get
	// rows 4 through 9 into the board.
	wantBoard.PopulateBoard("/Users/nabil/go/src/github.com/Nabil-Daoud/robodoku/puzzles/easy.txt")

	wantBoard.Spots[0][0].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
	wantBoard.Spots[0][1].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
	wantBoard.Spots[0][2].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
	wantBoard.Spots[0][3].Solved = true
	wantBoard.Spots[0][3].Value = 2
	wantBoard.Spots[0][3].Possible = [9]bool{false, true}
	wantBoard.Spots[0][4].Solved = true
	wantBoard.Spots[0][4].Value = 6
	wantBoard.Spots[0][4].Possible = [9]bool{false, false, false, false, false, true}
	wantBoard.Spots[0][5].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
	wantBoard.Spots[0][6].Solved = true
	wantBoard.Spots[0][6].Value = 7
	wantBoard.Spots[0][6].Possible = [9]bool{false, false, false, false, false, false, true}
	wantBoard.Spots[0][7].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
	wantBoard.Spots[0][8].Solved = true
	wantBoard.Spots[0][8].Value = 1
	wantBoard.Spots[0][8].Possible = [9]bool{true}

	wantBoard.Spots[1][0].Solved = true
	wantBoard.Spots[1][0].Value = 6
	wantBoard.Spots[1][0].Possible = [9]bool{false, false, false, false, false, true}
	wantBoard.Spots[1][1].Solved = true
	wantBoard.Spots[1][1].Value = 8
	wantBoard.Spots[1][1].Possible = [9]bool{false, false, false, false, false, false, false, true}
	wantBoard.Spots[1][2].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
	wantBoard.Spots[1][3].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
	wantBoard.Spots[1][4].Solved = true
	wantBoard.Spots[1][4].Value = 7
	wantBoard.Spots[1][4].Possible = [9]bool{false, false, false, false, false, false, true}
	wantBoard.Spots[1][5].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
	wantBoard.Spots[1][6].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
	wantBoard.Spots[1][7].Solved = true
	wantBoard.Spots[1][7].Value = 9
	wantBoard.Spots[1][7].Possible = [9]bool{false, false, false, false, false, false, false, false, true}
	wantBoard.Spots[1][8].Possible = [9]bool{true, true, true, true, true, true, true, true, true}

	wantBoard.Spots[2][0].Solved = true
	wantBoard.Spots[2][0].Value = 1
	wantBoard.Spots[2][0].Possible = [9]bool{true}
	wantBoard.Spots[2][1].Solved = true
	wantBoard.Spots[2][1].Value = 9
	wantBoard.Spots[2][1].Possible = [9]bool{false, false, false, false, false, false, false, false, true}
	wantBoard.Spots[2][2].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
	wantBoard.Spots[2][3].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
	wantBoard.Spots[2][4].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
	wantBoard.Spots[2][5].Solved = true
	wantBoard.Spots[2][5].Value = 4
	wantBoard.Spots[2][5].Possible = [9]bool{false, false, false, true}
	wantBoard.Spots[2][6].Solved = true
	wantBoard.Spots[2][6].Value = 5
	wantBoard.Spots[2][6].Possible = [9]bool{false, false, false, false, true}
	wantBoard.Spots[2][7].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
	wantBoard.Spots[2][8].Possible = [9]bool{true, true, true, true, true, true, true, true, true}

	gotBoard.PopulateBoard("/Users/nabil/go/src/github.com/Nabil-Daoud/robodoku/puzzles/easy.txt")

	for i, row := range wantBoard.Spots {
		for j, wantSpot := range row {
			if wantSpot != gotBoard.Spots[i][j] {
				t.Errorf("Did not load the spot at %v, %v correctly", i, j)
			}
			if wantSpot.Solved != gotBoard.Spots[i][j].Solved {
				t.Errorf("Did not load the spot.Solved at %v, %v correctly. Wanted %t, got %t", i, j, wantSpot.Solved, gotBoard.Spots[i][j].Solved)
			}
			if wantSpot.Value != gotBoard.Spots[i][j].Value {
				t.Errorf("Did not load the spot.Value at %v, %v correctly. Wanted %v, got %v", i, j, wantSpot.Value, gotBoard.Spots[i][j].Value)
			}
			if wantSpot.Possible != gotBoard.Spots[i][j].Possible {
				t.Errorf("Did not load the spot.Possible at %v, %v correctly. \nWanted %t, \n   got %t", i, j, wantSpot.Possible, gotBoard.Spots[i][j].Possible)
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

	// fmt.Printf("Spot[0][0].Possible before: %t\n", gotBoard.Spots[0][0].Possible)
	// fmt.Printf("board.N before: %v\n", gotBoard.N)
	didUpdate = gotBoard.updatePossibleRow()
	// fmt.Printf("Spot[0][0].Possible after: %t\n", gotBoard.Spots[0][0].Possible)
	// fmt.Printf("board.N after: %v\n", gotBoard.N)

	if !didUpdate {
		t.Error("UpdatePossibleRow should have returned true on first call.")
	}

	didNotUpdate = gotBoard.updatePossibleRow()

	if didNotUpdate {
		t.Error("UpdatePossibleRow should have returned false on second call.")
	}

	wantBoard.updatePossibleRow()

	wantBoard.Spots[0][0].Possible = [9]bool{false, false, true, true, true, false, false, true, true}
	wantBoard.Spots[0][1].Possible = [9]bool{false, false, true, true, true, false, false, true, true}
	wantBoard.Spots[0][2].Possible = [9]bool{false, false, true, true, true, false, false, true, true}
	wantBoard.Spots[0][3].Possible = [9]bool{false, true}
	wantBoard.Spots[0][4].Possible = [9]bool{false, false, false, false, false, true}
	wantBoard.Spots[0][5].Possible = [9]bool{false, false, true, true, true, false, false, true, true}
	wantBoard.Spots[0][6].Possible = [9]bool{false, false, false, false, false, false, true}
	wantBoard.Spots[0][7].Possible = [9]bool{false, false, true, true, true, false, false, true, true}
	wantBoard.Spots[0][8].Possible = [9]bool{true}

	wantBoard.Spots[1][0].Possible = [9]bool{false, false, false, false, false, true}
	wantBoard.Spots[1][1].Possible = [9]bool{false, false, false, false, false, false, false, true}
	wantBoard.Spots[1][2].Possible = [9]bool{true, true, true, true, true}
	wantBoard.Spots[1][3].Possible = [9]bool{true, true, true, true, true}
	wantBoard.Spots[1][4].Possible = [9]bool{false, false, false, false, false, false, true}
	wantBoard.Spots[1][5].Possible = [9]bool{true, true, true, true, true}
	wantBoard.Spots[1][6].Possible = [9]bool{true, true, true, true, true}
	wantBoard.Spots[1][7].Possible = [9]bool{false, false, false, false, false, false, false, false, true}
	wantBoard.Spots[1][8].Possible = [9]bool{true, true, true, true, true}

	wantBoard.Spots[2][0].Possible = [9]bool{true}
	wantBoard.Spots[2][1].Possible = [9]bool{false, false, false, false, false, false, false, false, true}
	wantBoard.Spots[2][2].Possible = [9]bool{false, true, true, false, false, true, true, true}
	wantBoard.Spots[2][3].Possible = [9]bool{false, true, true, false, false, true, true, true}
	wantBoard.Spots[2][4].Possible = [9]bool{false, true, true, false, false, true, true, true}
	wantBoard.Spots[2][5].Possible = [9]bool{false, false, false, true}
	wantBoard.Spots[2][6].Possible = [9]bool{false, false, false, false, true}
	wantBoard.Spots[2][7].Possible = [9]bool{false, true, true, false, false, true, true, true}
	wantBoard.Spots[2][8].Possible = [9]bool{false, true, true, false, false, true, true, true}

	for i, row := range wantBoard.Spots {
		for j, wantSpot := range row {
			if wantSpot != gotBoard.Spots[i][j] {
				t.Errorf("Did not update the spot at %v, %v correctly", i, j)
			}
			if wantSpot.Solved != gotBoard.Spots[i][j].Solved {
				t.Errorf("Did not update the spot.Solved at %v, %v correctly. Wanted %t, got %t", i, j, wantSpot.Solved, gotBoard.Spots[i][j].Solved)
			}
			if wantSpot.Value != gotBoard.Spots[i][j].Value {
				t.Errorf("Did not update the spot.Value at %v, %v correctly. Wanted %v, got %v", i, j, wantSpot.Value, gotBoard.Spots[i][j].Value)
			}
			if wantSpot.Possible != gotBoard.Spots[i][j].Possible {
				t.Errorf("Did not update the spot.Possible at %v, %v correctly. \nWanted %t, \n   got %t", i, j, wantSpot.Possible, gotBoard.Spots[i][j].Possible)
			}
		}
	}
	if wantBoard != gotBoard {
		t.Error("Did not update the board correctly")
	}
}

func TestUpdatePossibleCol(t *testing.T) {
	var wantBoard Board
	var gotBoard Board
	var didUpdate, didNotUpdate bool

	wantBoard.PopulateBoard("/Users/nabil/go/src/github.com/Nabil-Daoud/robodoku/puzzles/easy.txt")
	gotBoard.PopulateBoard("/Users/nabil/go/src/github.com/Nabil-Daoud/robodoku/puzzles/easy.txt")

	didUpdate = gotBoard.updatePossibleCol()
	if !didUpdate {
		t.Error("UpdatePossibleCol should have returned true on first call.")
	}

	didNotUpdate = gotBoard.updatePossibleCol()
	if didNotUpdate {
		t.Error("UpdatePossibleCol should have returned false on second call.")
	}

	wantBoard.updatePossibleCol()
	wantBoard.Spots[0][0].Possible = [9]bool{false, true, true, true, true, false, false, false, true}
	wantBoard.Spots[1][0].Possible = [9]bool{false, false, false, false, false, true, false, false, false} // 6
	wantBoard.Spots[2][0].Possible = [9]bool{true, false, false, false, false, false, false, false, false} // 1
	wantBoard.Spots[3][0].Possible = [9]bool{false, false, false, false, false, false, false, true, false} // 8
	wantBoard.Spots[4][0].Possible = [9]bool{false, true, true, true, true, false, false, false, true}
	wantBoard.Spots[5][0].Possible = [9]bool{false, true, true, true, true, false, false, false, true}
	wantBoard.Spots[6][0].Possible = [9]bool{false, true, true, true, true, false, false, false, true}
	wantBoard.Spots[7][0].Possible = [9]bool{false, true, true, true, true, false, false, false, true}
	wantBoard.Spots[8][0].Possible = [9]bool{false, false, false, false, false, false, true, false, false} // 7

	wantBoard.Spots[0][1].Possible = [9]bool{true, false, true, false, false, true, true, false, false}
	wantBoard.Spots[1][1].Possible = [9]bool{false, false, false, false, false, false, false, true}
	wantBoard.Spots[2][1].Possible = [9]bool{false, false, false, false, false, false, false, false, true}
	wantBoard.Spots[3][1].Possible = [9]bool{false, true}
	wantBoard.Spots[4][1].Possible = [9]bool{true, false, true, false, false, true, true, false, false}
	wantBoard.Spots[5][1].Possible = [9]bool{false, false, false, false, true}
	wantBoard.Spots[6][1].Possible = [9]bool{true, false, true, false, false, true, true, false, false}
	wantBoard.Spots[7][1].Possible = [9]bool{false, false, false, true}
	wantBoard.Spots[8][1].Possible = [9]bool{true, false, true, false, false, true, true, false, false}

	wantBoard.Spots[0][2].Possible = [9]bool{true, true, false, false, true, true, true, true}
	wantBoard.Spots[1][2].Possible = [9]bool{true, true, false, false, true, true, true, true}
	wantBoard.Spots[2][2].Possible = [9]bool{true, true, false, false, true, true, true, true}
	wantBoard.Spots[3][2].Possible = [9]bool{true, true, false, false, true, true, true, true}
	wantBoard.Spots[4][2].Possible = [9]bool{false, false, false, true}
	wantBoard.Spots[5][2].Possible = [9]bool{true, true, false, false, true, true, true, true}
	wantBoard.Spots[6][2].Possible = [9]bool{false, false, false, false, false, false, false, false, true}
	wantBoard.Spots[7][2].Possible = [9]bool{true, true, false, false, true, true, true, true}
	wantBoard.Spots[8][2].Possible = [9]bool{false, false, true}

	for i, row := range wantBoard.Spots {
		for j, wantSpot := range row {
			if wantSpot != gotBoard.Spots[i][j] {
				t.Errorf("Did not update the spot at %v, %v correctly", i, j)
			}
			if wantSpot.Solved != gotBoard.Spots[i][j].Solved {
				t.Errorf("Did not update the spot.Solved at %v, %v correctly. Wanted %t, got %t", i, j, wantSpot.Solved, gotBoard.Spots[i][j].Solved)
			}
			if wantSpot.Value != gotBoard.Spots[i][j].Value {
				t.Errorf("Did not update the spot.Value at %v, %v correctly. Wanted %v, got %v", i, j, wantSpot.Value, gotBoard.Spots[i][j].Value)
			}
			if wantSpot.Possible != gotBoard.Spots[i][j].Possible {
				t.Errorf("Did not update the spot.Possible at %v, %v correctly. \nWanted %t, \n   got %t", i, j, wantSpot.Possible, gotBoard.Spots[i][j].Possible)
			}
		}
	}
	if wantBoard != gotBoard {
		t.Error("Did not update the board correctly")
	}
}

func TestSqrHas(t *testing.T) {
	var board Board
	board.PopulateBoard("/Users/nabil/go/src/github.com/Nabil-Daoud/robodoku/puzzles/easy.txt")

	board.getSqrHas(0, 0)
	var tests = []struct {
		sqrRowIndex int
		sqrColIndex int
		want        [9]bool
	}{
		// r, c, 	1,  	2,  		3,		4, 		5, 		6, 			7, 		8, 		9
		{0, 0, [9]bool{true, false, false, false, false, true, false, true, true}},
		{0, 1, [9]bool{false, true, false, true, false, true, true, false, false}},
		{0, 2, [9]bool{true, false, false, false, true, false, true, false, true}},

		{1, 0, [9]bool{false, true, false, true, true, false, false, true, false}},
		{1, 1, [9]bool{true, true, true, false, false, true, false, false, false}},
		{1, 2, [9]bool{false, true, false, true, false, false, false, true, true}},

		{2, 0, [9]bool{false, false, true, true, false, false, true, false, true}},
		{2, 1, [9]bool{true, false, true, false, true, false, false, true, false}},
		{2, 2, [9]bool{false, false, true, true, false, true, true, false, false}},
	}
	var got [9]bool
	for _, c := range tests {
		got = board.getSqrHas(c.sqrRowIndex, c.sqrColIndex)
		if got != c.want {
			t.Errorf("getSqrHas(%v, %v) = %t,\n    wanted %t", c.sqrRowIndex, c.sqrColIndex, got, c.want)
		}
	}
}

func TestUpdatePossibleSqr(t *testing.T) {
	var wantBoard Board
	var gotBoard Board
	var didUpdate, didNotUpdate bool

	wantBoard.PopulateBoard("/Users/nabil/go/src/github.com/Nabil-Daoud/robodoku/puzzles/easy.txt")
	gotBoard.PopulateBoard("/Users/nabil/go/src/github.com/Nabil-Daoud/robodoku/puzzles/easy.txt")

	didUpdate = gotBoard.updatePossibleSqr()
	if !didUpdate {
		t.Error("UpdatePossibleSqr should have returned true on first call.")
	}

	didNotUpdate = gotBoard.updatePossibleSqr()
	if didNotUpdate {
		t.Error("UpdatePossibleSqr should have returned false on second call.")
	}

	wantBoard.updatePossibleSqr()
	wantBoard.Spots[0][0].Possible = [9]bool{false, true, true, true, true, false, true, false, false}
	wantBoard.Spots[0][1].Possible = [9]bool{false, true, true, true, true, false, true, false, false}
	wantBoard.Spots[0][2].Possible = [9]bool{false, true, true, true, true, false, true, false, false}
	wantBoard.Spots[1][0].Possible = [9]bool{false, false, false, false, false, true}               // has 6
	wantBoard.Spots[1][1].Possible = [9]bool{false, false, false, false, false, false, false, true} // has 8
	wantBoard.Spots[1][2].Possible = [9]bool{false, true, true, true, true, false, true, false, false}
	wantBoard.Spots[2][0].Possible = [9]bool{true}                                                         // has 1
	wantBoard.Spots[2][1].Possible = [9]bool{false, false, false, false, false, false, false, false, true} // has 9
	wantBoard.Spots[2][2].Possible = [9]bool{false, true, true, true, true, false, true, false, false}

	wantBoard.Spots[0][3].Possible = [9]bool{false, true}                             // has 2
	wantBoard.Spots[0][4].Possible = [9]bool{false, false, false, false, false, true} // has 6
	wantBoard.Spots[0][5].Possible = [9]bool{true, false, true, false, true, false, false, true, true}
	wantBoard.Spots[1][3].Possible = [9]bool{true, false, true, false, true, false, false, true, true}
	wantBoard.Spots[1][4].Possible = [9]bool{false, false, false, false, false, false, true} // has 7
	wantBoard.Spots[1][5].Possible = [9]bool{true, false, true, false, true, false, false, true, true}
	wantBoard.Spots[2][3].Possible = [9]bool{true, false, true, false, true, false, false, true, true}
	wantBoard.Spots[2][4].Possible = [9]bool{true, false, true, false, true, false, false, true, true}
	wantBoard.Spots[2][5].Possible = [9]bool{false, false, false, true} // has 4

	wantBoard.Spots[0][6].Possible = [9]bool{false, false, false, false, false, false, true} // has 7
	wantBoard.Spots[0][7].Possible = [9]bool{false, true, true, true, false, true, false, true}
	wantBoard.Spots[0][8].Possible = [9]bool{true} // has 1
	wantBoard.Spots[1][6].Possible = [9]bool{false, true, true, true, false, true, false, true}
	wantBoard.Spots[1][7].Possible = [9]bool{false, false, false, false, false, false, false, false, true} // has 9
	wantBoard.Spots[1][8].Possible = [9]bool{false, true, true, true, false, true, false, true}
	wantBoard.Spots[2][6].Possible = [9]bool{false, false, false, false, true} // has 5
	wantBoard.Spots[2][7].Possible = [9]bool{false, true, true, true, false, true, false, true}
	wantBoard.Spots[2][8].Possible = [9]bool{false, true, true, true, false, true, false, true}

	for i, row := range wantBoard.Spots {
		for j, wantSpot := range row {
			if wantSpot != gotBoard.Spots[i][j] {
				t.Errorf("updatePossibleSqr() did not update the spot at %v, %v correctly", i, j)
			}
			if wantSpot.Solved != gotBoard.Spots[i][j].Solved {
				t.Errorf("updatePossibleSqr() did not update the spot.Solved at %v, %v correctly. Wanted %t, got %t", i, j, wantSpot.Solved, gotBoard.Spots[i][j].Solved)
			}
			if wantSpot.Value != gotBoard.Spots[i][j].Value {
				t.Errorf("updatePossibleSqr() did not update the spot.Value at %v, %v correctly. Wanted %v, got %v", i, j, wantSpot.Value, gotBoard.Spots[i][j].Value)
			}
			if wantSpot.Possible != gotBoard.Spots[i][j].Possible {
				t.Errorf("updatePossibleSqr() did not update the spot.Possible at %v, %v correctly. \nWanted %t, \n   got %t", i, j, wantSpot.Possible, gotBoard.Spots[i][j].Possible)
			}
		}
	}
	if wantBoard != gotBoard {
		t.Error("Did not update the board correctly")
	}
}

func TestUpdatePossible(t *testing.T) {
	var wantBoard Board
	var gotBoard Board
	var didUpdate, didNotUpdate bool

	wantBoard.PopulateBoard("/Users/nabil/go/src/github.com/Nabil-Daoud/robodoku/puzzles/easy.txt")
	gotBoard.PopulateBoard("/Users/nabil/go/src/github.com/Nabil-Daoud/robodoku/puzzles/easy.txt")

	didUpdate = gotBoard.UpdatePossible()
	if !didUpdate {
		t.Error("UpdatePossible should have returned true on first call.")
	}

	didNotUpdate = gotBoard.UpdatePossible()
	if didNotUpdate {
		t.Error("UpdatePossible should have returned false on second call.")
	}

	wantBoard.UpdatePossible()
	wantBoard.Spots[0][0].Possible = [9]bool{false, false, true, true, true, false, false, false, false}

	// wantBoard.Spots[0][1].Possible = [9]bool{false, true, true, true, true, false, true, false, false}
	// wantBoard.Spots[0][2].Possible = [9]bool{false, true, true, true, true, false, true, false, false}
	// wantBoard.Spots[1][0].Possible = [9]bool{false, false, false, false, false, true}               // has 6
	// wantBoard.Spots[1][1].Possible = [9]bool{false, false, false, false, false, false, false, true} // has 8
	// wantBoard.Spots[1][2].Possible = [9]bool{false, true, true, true, true, false, true, false, false}
	// wantBoard.Spots[2][0].Possible = [9]bool{true}                                                         // has 1
	// wantBoard.Spots[2][1].Possible = [9]bool{false, false, false, false, false, false, false, false, true} // has 9
	// wantBoard.Spots[2][2].Possible = [9]bool{false, true, true, true, true, false, true, false, false}
	//
	// wantBoard.Spots[0][3].Possible = [9]bool{false, true}                             // has 2
	// wantBoard.Spots[0][4].Possible = [9]bool{false, false, false, false, false, true} // has 6
	// wantBoard.Spots[0][5].Possible = [9]bool{true, false, true, false, true, false, false, true, true}
	// wantBoard.Spots[1][3].Possible = [9]bool{true, false, true, false, true, false, false, true, true}
	// wantBoard.Spots[1][4].Possible = [9]bool{false, false, false, false, false, false, true} // has 7
	// wantBoard.Spots[1][5].Possible = [9]bool{true, false, true, false, true, false, false, true, true}
	// wantBoard.Spots[2][3].Possible = [9]bool{true, false, true, false, true, false, false, true, true}
	// wantBoard.Spots[2][4].Possible = [9]bool{true, false, true, false, true, false, false, true, true}
	// wantBoard.Spots[2][5].Possible = [9]bool{false, false, false, true} // has 4
	//
	// wantBoard.Spots[0][6].Possible = [9]bool{false, false, false, false, false, false, true} // has 7
	// wantBoard.Spots[0][7].Possible = [9]bool{false, true, true, true, false, true, false, true}
	// wantBoard.Spots[0][8].Possible = [9]bool{true} // has 1
	// wantBoard.Spots[1][6].Possible = [9]bool{false, true, true, true, false, true, false, true}
	// wantBoard.Spots[1][7].Possible = [9]bool{false, false, false, false, false, false, false, false, true} // has 9
	// wantBoard.Spots[1][8].Possible = [9]bool{false, true, true, true, false, true, false, true}
	// wantBoard.Spots[2][6].Possible = [9]bool{false, false, false, false, true} // has 5
	// wantBoard.Spots[2][7].Possible = [9]bool{false, true, true, true, false, true, false, true}
	// wantBoard.Spots[2][8].Possible = [9]bool{false, true, true, true, false, true, false, true}

	for i, row := range wantBoard.Spots {
		for j, wantSpot := range row {
			if wantSpot != gotBoard.Spots[i][j] {
				t.Errorf("UpdatePossible() did not update the spot at %v, %v correctly", i, j)
			}
			if wantSpot.Solved != gotBoard.Spots[i][j].Solved {
				t.Errorf("UpdatePossible() did not update the spot.Solved at %v, %v correctly. Wanted %t, got %t", i, j, wantSpot.Solved, gotBoard.Spots[i][j].Solved)
			}
			if wantSpot.Value != gotBoard.Spots[i][j].Value {
				t.Errorf("UpdatePossible() did not update the spot.Value at %v, %v correctly. Wanted %v, got %v", i, j, wantSpot.Value, gotBoard.Spots[i][j].Value)
			}
			if wantSpot.Possible != gotBoard.Spots[i][j].Possible {
				t.Errorf("UpdatePossible() did not update the spot.Possible at %v, %v correctly. \nWanted %t, \n   got %t", i, j, wantSpot.Possible, gotBoard.Spots[i][j].Possible)
			}
		}
	}
	if wantBoard != gotBoard {
		t.Error("Did not update the board correctly")
	}
}

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
