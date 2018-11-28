package board

import (
	"fmt"
	"io/ioutil"

	"github.com/Nabil-Daoud/robodoku/sudoku/spot"
)

// Board is a structure that holds an N x N Sudoku board.
type Board struct {
	// N     int
	Spots [9][9]spot.Spot
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Populate takes a slice of uint8 and populates the Spots on the bBoard
func (board *Board) Populate(dat []uint8) {
	for i, item := range dat {
		if item != uint8(10) {
			// if it's not a blank, solve it; else set all possible to true.
			if item != uint8(32) {
				board.Spots[i/10][i%10].SolveSpot(item - 48)
			} else {
				board.Spots[i/10][i%10].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
			}
		}
	}
}

// PopulateBoard takes a filename and populates the Spots
func (board *Board) PopulateBoard(filename string) {
	dat, err := ioutil.ReadFile(filename)
	check(err)

	for i, item := range dat {
		if item != uint8(10) {
			// if it's not a blank, solve it; else set all possible to true.
			if item != uint8(32) {
				board.Spots[i/10][i%10].SolveSpot(item - 48)
			} else {
				board.Spots[i/10][i%10].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
			}
		}
	}
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
	didUpdateRow := board.updatePossibleRow()
	didUpdateCol := board.updatePossibleCol()
	didUpdateSqr := board.updatePossibleSqr()
	return didUpdateRow || didUpdateCol || didUpdateSqr
}

// updatePossibleRow() loops through the rows and updates the Possible array
// based on solved spots in the same row.
func (board *Board) updatePossibleRow() bool {
	updated := false

	// Loop through the rows
	for rowIndex, row := range board.Spots {
		rowHas := board.getRowHas(row)
		var updatedPossible [9]bool

		// Loop through the spots in the row again and update the possible values
		// based on what the row has.
		for colIndex, spot := range row {
			if !spot.Solved {
				updatedPossible = spot.Possible

				for valueIndex, possible := range spot.Possible {
					if possible && rowHas[valueIndex] {
						updatedPossible[valueIndex] = false
						updated = true
					}
				}
				board.Spots[rowIndex][colIndex].UpdatePossible(updatedPossible)
			}
		}
	}
	return updated
}

func (board Board) getRowHas(row [9]spot.Spot) [9]bool {
	rowHas := [9]bool{}
	// Loop through the spots in the row to populate the values the row has
	for _, spot := range row {
		if spot.Solved {
			rowHas[spot.Value-1] = true
		}
	}

	return rowHas
}

// updatePossibleCol() loops through the columns and updates the Possible array
// based on solved spots in the same column.
func (board *Board) updatePossibleCol() bool {
	updated := false

	// Loop through the columns, i
	for colIndex := range board.Spots {
		// Loop through the spots in the column to populate the values the column has
		colHas := board.getColHas(colIndex)
		var updatedPossible [9]bool

		// Loop through the spots in the column again and update their possible
		// values based on what the column has.
		for rowIndex := range board.Spots {
			if !board.Spots[rowIndex][colIndex].Solved {
				updatedPossible = board.Spots[rowIndex][colIndex].Possible

				for valueIndex, possible := range board.Spots[rowIndex][colIndex].Possible {
					if possible && colHas[valueIndex] {
						updatedPossible[valueIndex] = false
						updated = true
					}
					board.Spots[rowIndex][colIndex].UpdatePossible(updatedPossible)
				}
			}
		}
	}
	return updated
}

func (board Board) getColHas(colIndex int) [9]bool {
	colHas := [9]bool{}

	for rowIndex := range board.Spots {
		if board.Spots[rowIndex][colIndex].Solved {
			colHas[board.Spots[rowIndex][colIndex].Value-1] = true
		}
	}

	return colHas
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

						for valueIndex, possible := range board.Spots[spotRow][spotCol].Possible {
							if possible && sqrHas[valueIndex] {
								updatedPossible[valueIndex] = false
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

// SolveSinglePossible looks for and solves unsolved spots
// that have a single possible value
func (board *Board) SolveSinglePossible() {
	for rowIndex, row := range board.Spots {
		for colIndex, spot := range row {
			if !spot.Solved {
				hasSinglePossible, value := spot.SinglePossible()
				if hasSinglePossible {
					// fmt.Println("On board: ")
					// board.PrintBoard()
					fmt.Printf("Spot(%v, %v) can only be %v.\n\n", rowIndex, colIndex, value)
					board.Spots[rowIndex][colIndex].SolveSpot(uint8(value))
				}
			}
		}
	}
}
