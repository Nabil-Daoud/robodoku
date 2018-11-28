package board

import "fmt"

// Board is a structure that holds an N x N Sudoku board.
type Board struct {
	Spots [][]uint8
}

// Build takes a slice of uint8, creates and populates known board values
func Build(dat []uint8) Board {
	var (
		board Board
		n     int
	)

	if len(dat) == 90 { // standard 9x9 board
		n = 9
		board = populate(n, dat, newBoard(n))
	} else if len(dat) == 20 { // 4x4 board
		n = 4
		board = populate(n, dat, newBoard(n))
	} else if len(dat) == 6 { // 2x2 board
		n = 2
		board = populate(n, dat, newBoard(n))
	} else {
		fmt.Printf("Error: board data not a recognized size, %v.\n", len(dat))
	}
	return board
}

// newBoard takes the dimension, n, and returns an nxn Board.
func newBoard(n int) Board {
	b := make([][]uint8, n)
	for i := range b {
		b[i] = make([]uint8, n)
	}
	return Board{b}
}

// populate takes the board size, n; board data, dat; and a Board.
// It returns a board with the values filled in.
func populate(n int, dat []uint8, board Board) Board {
	const (
		blank uint8 = 32
		eol   uint8 = 10
	)
	for i, item := range dat {
		// if the value is not eol and not blank set the value
		if item != eol && item != blank {
			// TODO confirm that the value is in the range 0 < x <= len(board.Spots)
			board.Spots[i/(n+1)][i%(n+1)] = (item - 48)
		}
	}
	return board
}

// PrintBoard prints the solved values on the Sodoku board.
// TODO: Make this function return a string for any size board
func (board *Board) PrintBoard() {
	for i, row := range board.Spots {
		for j, spot := range row {
			if spot != uint8(0) {
				fmt.Printf("%v ", spot)
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
// func (board *Board) UpdatePossible() bool {
// 	didUpdateRow := board.updatePossibleRow()
// 	didUpdateCol := board.updatePossibleCol()
// 	didUpdateSqr := board.updatePossibleSqr()
// 	return didUpdateRow || didUpdateCol || didUpdateSqr
// }

// updatePossibleRow() loops through the rows and updates the Possible array
// based on solved spots in the same row.
// func (board *Board) updatePossibleRow() bool {
// 	updated := false
//
// 	// Loop through the rows
// 	for rowIndex, row := range board.Spots {
// 		rowHas := board.getRowHas(row)
// 		var updatedPossible [9]bool
//
// 		// Loop through the spots in the row again and update the possible values
// 		// based on what the row has.
// 		for colIndex, spot := range row {
// 			if !spot.Solved {
// 				updatedPossible = spot.Possible
//
// 				for valueIndex, possible := range spot.Possible {
// 					if possible && rowHas[valueIndex] {
// 						updatedPossible[valueIndex] = false
// 						updated = true
// 					}
// 				}
// 				board.Spots[rowIndex][colIndex].UpdatePossible(updatedPossible)
// 			}
// 		}
// 	}
// 	return updated
// }
//
// func (board Board) getRowHas(row [9]spot.Spot) [9]bool {
// 	rowHas := [9]bool{}
// 	// Loop through the spots in the row to populate the values the row has
// 	for _, spot := range row {
// 		if spot.Solved {
// 			rowHas[spot.Value-1] = true
// 		}
// 	}
//
// 	return rowHas
// }

// updatePossibleCol() loops through the columns and updates the Possible array
// based on solved spots in the same column.
// func (board *Board) updatePossibleCol() bool {
// 	updated := false
//
// 	// Loop through the columns, i
// 	for colIndex := range board.Spots {
// 		// Loop through the spots in the column to populate the values the column has
// 		colHas := board.getColHas(colIndex)
// 		var updatedPossible [9]bool
//
// 		// Loop through the spots in the column again and update their possible
// 		// values based on what the column has.
// 		for rowIndex := range board.Spots {
// 			if !board.Spots[rowIndex][colIndex].Solved {
// 				updatedPossible = board.Spots[rowIndex][colIndex].Possible
//
// 				for valueIndex, possible := range board.Spots[rowIndex][colIndex].Possible {
// 					if possible && colHas[valueIndex] {
// 						updatedPossible[valueIndex] = false
// 						updated = true
// 					}
// 					board.Spots[rowIndex][colIndex].UpdatePossible(updatedPossible)
// 				}
// 			}
// 		}
// 	}
// 	return updated
// }
//
// func (board Board) getColHas(colIndex int) [9]bool {
// 	colHas := [9]bool{}
//
// 	for rowIndex := range board.Spots {
// 		if board.Spots[rowIndex][colIndex].Solved {
// 			colHas[board.Spots[rowIndex][colIndex].Value-1] = true
// 		}
// 	}
//
// 	return colHas
// }

// updatePossibleSqr() loops through the squares and updates the Possible array
// based on solved spots in the same square.
// func (board *Board) updatePossibleSqr() bool {
// 	updated := false
// 	var updatedPossible [9]bool
//
// 	// loop through the squares
// 	for sqrRowIndex := 0; sqrRowIndex < 3; sqrRowIndex++ {
// 		for sqrColIndex := 0; sqrColIndex < 3; sqrColIndex++ {
// 			sqrHas := board.getSqrHas(sqrRowIndex, sqrColIndex)
//
// 			// loop through the spots in the square and update the possible values
// 			// based on what the column has.
// 			for spotRowInSqr := 0; spotRowInSqr < 3; spotRowInSqr++ {
// 				for spotColInSqr := 0; spotColInSqr < 3; spotColInSqr++ {
// 					spotRow, spotCol := 3*sqrRowIndex+spotRowInSqr, 3*sqrColIndex+spotColInSqr
// 					// fmt.Printf("spotRow = %v, spotCol = %v.\nsqrRowIndex = %v, sqrColIndex = %v\nspotRowInSqr = %v, spotColInSqr = %v", spotRow, spotCol, sqrRowIndex, sqrColIndex, spotRowInSqr, spotColInSqr)
// 					if !board.Spots[spotRow][spotCol].Solved {
// 						updatedPossible = board.Spots[spotRow][spotCol].Possible
//
// 						for valueIndex, possible := range board.Spots[spotRow][spotCol].Possible {
// 							if possible && sqrHas[valueIndex] {
// 								updatedPossible[valueIndex] = false
// 								updated = true
// 							}
// 							board.Spots[spotRow][spotCol].UpdatePossible(updatedPossible)
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// 	return updated
// }
//
// func (board *Board) getSqrHas(sqrRowIndex, sqrColIndex int) [9]bool {
// 	sqrHas := [9]bool{}
// 	for i := 0; i < 3; i++ {
// 		for j := 0; j < 3; j++ {
// 			spotRow := 3*sqrRowIndex + i
// 			spotCol := 3*sqrColIndex + j
// 			if board.Spots[spotRow][spotCol].Solved {
// 				sqrHas[board.Spots[spotRow][spotCol].Value-1] = true
// 			}
// 		}
// 	}
// 	return sqrHas
// }

// SolveSinglePossible looks for and solves unsolved spots
// that have a single possible value
// func (board *Board) SolveSinglePossible() {
// 	for rowIndex, row := range board.Spots {
// 		for colIndex, spot := range row {
// 			if !spot.Solved {
// 				hasSinglePossible, value := spot.SinglePossible()
// 				if hasSinglePossible {
// 					// fmt.Println("On board: ")
// 					// board.PrintBoard()
// 					fmt.Printf("Spot(%v, %v) can only be %v.\n\n", rowIndex, colIndex, value)
// 					board.Spots[rowIndex][colIndex].SolveSpot(uint8(value))
// 				}
// 			}
// 		}
// 	}
// }
