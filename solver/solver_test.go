package solver

import (
	"testing"

	"github.com/Nabil-Daoud/robodoku/sudoku/board"
)

func Test(t *testing.T) {
	var tests = []struct {
		value uint8
		want  int
	}{
		{1, 1},
		{2, 2},
	}

	var board board.Board
	Phase1(board)

	var got int
	for _, c := range tests {
		got = int(c.value)
		if got != c.want {
			t.Errorf("Cell(%v) == %v, want %v", c.value, got, c.want)
		}
	}
}
