package solver

import (
	"testing"

	"github.com/Nabil-Daoud/robodoku/sudoku/board"
)

func Test(t *testing.T) {

	var board = board.Build([]uint8{32, 32, 32, 50, 54, 32, 55, 32, 49, 10, 54, 56, 32, 32, 55, 32, 32, 57, 32, 10, 49, 57, 32, 32, 32, 52, 53, 32, 32, 10, 56, 50, 32, 49, 32, 32, 32, 52, 32, 10, 32, 32, 52, 54, 32, 50, 57, 32, 32, 10, 32, 53, 32, 32, 32, 51, 32, 50, 56, 10, 32, 32, 57, 51, 32, 32, 32, 55, 52, 10, 32, 52, 32, 32, 53, 32, 32, 51, 54, 10, 55, 32, 51, 32, 49, 56, 32, 32, 32, 10})
	Phase1(board)
	// fmt.Printf("rowHas %t\n", rowHas(board, 0))
	getRowHas(board, 0)
	getColHas(board, 0)
	getSqrHas(board, 0, 0)
	singlePossible(0, 0, board)

	solveSinglePossible(board)

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
