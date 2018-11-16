package solver

import (
	"fmt"

	"github.com/Nabil-Daoud/robodoku/sudoku/board"
)

var myBoard board.Board

// Phase1 solves unsolved spots that have only one possible value
func Phase1(board *board.Board) {
	fmt.Println("I'm in Phase 1.")

	myBoard.PopulateBoard("/Users/nabil/go/src/github.com/Nabil-Daoud/robodoku/puzzles/easy.txt")
	for board.UpdatePossible() {
		board.SolveSinglePossible()
	}
}
