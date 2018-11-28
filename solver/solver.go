package solver

import (
	"github.com/Nabil-Daoud/robodoku/sudoku/board"
)

var myBoard board.Board

// Phase1 solves unsolved spots that have only one possible value
func Phase1(board board.Board) board.Board {
	for board.UpdatePossible() {
		board.SolveSinglePossible()
	}
	return board
}

// Phase2 goes through the rows, columns and squares looking for and solving
// squares that have the only possilbe for any given value
func Phase2(board board.Board) board.Board {

	return board
}

func rowNeeds(rowIndex int, board board.Board) [9]bool {
	rowNeeds := [9]bool{true, true, true, true, true, true, true, true, true}
	for colIndex := range board.Spots {
		if board.Spots[rowIndex][colIndex].Solved {
			rowNeeds[board.Spots[rowIndex][colIndex].Value-1] = false
		}
	}
	return rowNeeds
}
