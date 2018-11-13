package board

import (
	"fmt"

	"github.com/Nabil-Daoud/robodoku/sudoku/spot"
)

// Board is a structure that holds an N x N Sudoku board.
type Board struct {
	N     int
	spots [][]spot.Spot
}

// PopulateBoard takes a board and a filename and populates the spots
func PopulateBoard(board Board, filename string) {

	board.spots[0][0] = spot.NewSpot(9, '1')

}

// PrintBoard prints the solved values on the Sodoku board.
func PrintBoard() {
	fmt.Println('a')
}
