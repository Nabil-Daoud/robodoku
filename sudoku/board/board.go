package board

import (
	"fmt"
	"io/ioutil"

	"github.com/Nabil-Daoud/robodoku/sudoku/spot"
)

// Board is a structure that holds an N x N Sudoku board.
type Board struct {
	N     int
	Spots [9][9]spot.Spot
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// PopulateBoard takes a filename and populates the Spots
func (board *Board) PopulateBoard(filename string) {
	dat, err := ioutil.ReadFile(filename)
	check(err)

	// fmt.Printf("The data is read in as a %T and is of length %v\n", dat, len(dat))
	// lines := []string{strings.Split(string(dat), "\n")}

	// The code bewlow can be used to determine the dimensions of the board.
	// for i := 0; i < len(dat); i++ {
	// 	fmt.Printf("i: %v\n", i)
	// 	if dat[i] == 10 {
	// 		fmt.Println(">>>> In the if block.")
	// 		board.N = i
	// 		fmt.Printf("Board size, N set to %v\n", i)
	// 		fmt.Printf("N = %v\n", i)
	// 		break
	// 	}
	// 	fmt.Println("Loop.")
	// }
	// fmt.Println("<<<< made it out of the for loop. phew")
	// fmt.Println("the board is ", board.N, " by ", string(board.N), ".")

	for i, item := range dat {
		// fmt.Print(string(item))
		// fmt.Print(" | ", i, ": ", string(item))
		if item != uint8(10) {
			// if it's not a blank, solve it; otherwise set all possible to true.
			if item != uint8(32) {
				board.Spots[i/10][i%10].SolveSpot(item - 48)
			} else {
				board.Spots[i/10][i%10].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
			}
		}
	}

	// fmt.Println("\n<<<<<<<<<<<<<<<< length of dat = ", len(dat))
	// fmt.Println("")
	// for _, thing := range dat {
	// 	fmt.Print(thing, " . ")
	// 	if thing == 10 {
	// 		fmt.Println("")
	// 	}
	// }
	// fmt.Println("")

}

// PrintBoard prints the solved values on the Sodoku board.
func (board *Board) PrintBoard() {
	for i, row := range board.Spots {
		for j, spot := range row {
			if spot.Solved {
				fmt.Printf("%v ", spot.Value)
			} else {
				fmt.Print("  ")
			}

			if j%3 == 2 && j != 8 {
				fmt.Print("|")
			}
		}
		fmt.Print("\n")
		if i%3 == 2 && i != 8 {
			fmt.Println("------|------|------")
		}
	}
}

// UpdatePossible takes the board and updates the possible values based on rows,
// columns, and squares. It returns true if a possible square was updated.
func (board *Board) UpdatePossible() bool {
	return board.updatePossibleRow() || board.updatePossibleCol() || board.updatePossibleSqr()
	// return board.updatePossibleCol() || board.updatePossibleSqr()

	// updated := board.updatePossibleRow()
	// updated = updated || board.updatePossibleCol()
	// updated = updated || board.updatePossibleSqr()
	// return updated
}

// updatePossibleRow() loops through the rows and updates the Possible array
// based on solved spots in the same row.
func (board *Board) updatePossibleRow() bool {
	updated := false

	// Loop through the rows
	for i, row := range board.Spots { // i
		// Loop through the spots in the row to populate the values the row has
		rowHas := [9]bool{} // false, false, false, false, false, false, false, false, false}
		var updatedPossible [9]bool

		for _, spot := range row {
			if spot.Solved {
				rowHas[spot.Value-1] = true
			}
		}

		// Loop through the spots in the row again and update the possible values
		// based on what the row has.
		for j, spot := range row { // k
			if !spot.Solved {
				updatedPossible = spot.Possible

				for k, possible := range spot.Possible {
					if possible && rowHas[k] {
						updatedPossible[k] = false
						updated = true
					}
				}
				board.Spots[i][j].UpdatePossible(updatedPossible)
			}
		}
	}
	return updated
}

// updatePossibleCol() loops through the columns and updates the Possible array
// based on solved spots in the same column.
func (board *Board) updatePossibleCol() bool {
	updated := false
	// Loop through the columns
	for i := range board.Spots {
		// Loop through the spots in the column to populate the values the column has
		colHas := [9]bool{}
		var updatedPossible [9]bool

		for j := range board.Spots {
			if board.Spots[j][i].Solved {
				colHas[board.Spots[j][i].Value-1] = true
			}
		}
		// Loop through the spots in the column again and update their possible
		// values based on what the column has.
		for j := range board.Spots {
			if !board.Spots[j][i].Solved {
				updatedPossible = board.Spots[j][i].Possible

				for k, possible := range board.Spots[j][i].Possible {
					if possible && colHas[k] {
						updatedPossible[k] = false
						updated = true
					}
					board.Spots[j][i].UpdatePossible(updatedPossible)
				}
			}
		}
	}
	return updated
}

// updatePossibleSqr() loops through the squares and updates the Possible array
// based on solved spots in the same square.
func (board *Board) updatePossibleSqr() bool {
	updated := false
	var updatedPossible [9]bool

	// loop through the squares
	for sqrRowIndex := 0; sqrRowIndex < 3; sqrRowIndex++ {
		for sqrColIndex := 0; sqrColIndex < 3; sqrColIndex++ {
			sqrHas := board.getSqrHas(sqrRowIndex, sqrColIndex)

			// loop through the spots in the square and update the possible values
			// based on what the column has.
			for spotRowInSqr := 0; spotRowInSqr < 3; spotRowInSqr++ {
				for spotColInSqr := 0; spotColInSqr < 3; spotColInSqr++ {
					spotRow, spotCol := 3*sqrRowIndex+spotRowInSqr, 3*sqrColIndex+spotColInSqr
					// fmt.Printf("spotRow = %v, spotCol = %v.\nsqrRowIndex = %v, sqrColIndex = %v\nspotRowInSqr = %v, spotColInSqr = %v", spotRow, spotCol, sqrRowIndex, sqrColIndex, spotRowInSqr, spotColInSqr)
					if !board.Spots[spotRow][spotCol].Solved {
						updatedPossible = board.Spots[spotRow][spotCol].Possible

						for k, possible := range board.Spots[spotRow][spotCol].Possible {
							if possible && sqrHas[k] {
								updatedPossible[k] = false
								updated = true
							}
							board.Spots[spotRow][spotCol].UpdatePossible(updatedPossible)
						}
					}
				}
			}
		}
	}
	return updated
}

func (board *Board) getSqrHas(sqrRowIndex, sqrColIndex int) [9]bool {
	sqrHas := [9]bool{}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			spotRow := 3*sqrRowIndex + i
			spotCol := 3*sqrColIndex + j
			if board.Spots[spotRow][spotCol].Solved {
				sqrHas[board.Spots[spotRow][spotCol].Value-1] = true
			}
		}
	}
	return sqrHas
}
