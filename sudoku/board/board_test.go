package board

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func expectedNewBoard(n int) Board {
	if n == 2 {
		return Board{[][]uint8{{0, 0}, {0, 0}}}
	}
	if n == 4 {
		return Board{[][]uint8{{0, 0, 0, 0}, {0, 0, 0, 0},
			{0, 0, 0, 0}, {0, 0, 0, 0}}}
	}
	if n == 9 {
		return Board{[][]uint8{{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0}}}
	}
	b := make([][]uint8, n)
	for i := range b {
		b[i] = make([]uint8, n)
	}
	return Board{b}
}

func TestNewBoard(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(expectedNewBoard(2), newBoard(2), "2x2 Board did not initialize correctly.")
	assert.Equal(expectedNewBoard(4), newBoard(4), "4x4 Board did not initialize correctly.")
	assert.Equal(expectedNewBoard(9), newBoard(9), "9x9 Board did not initialize correctly.")
}

// func expectedBuildBoard(values [][]uint8) Board {
//
// }

func TestBuild(t *testing.T) {
	assert := assert.New(t)

	var (
		expectedBoard1, gotBoard1 Board
		expectedBoard2, gotBoard2 Board
		expectedBoard3, gotBoard3 Board
		expectedBoard4, gotBoard4 Board
		gotErr                    error
	)

	expectedBoard1, _ = Build([]uint8{32, 32, 32, 50, 54, 32, 55, 32, 49, 10, 54, 56, 32, 32, 55, 32, 32, 57, 32, 10, 49, 57, 32, 32, 32, 52, 53, 32, 32, 10, 56, 50, 32, 49, 32, 32, 32, 52, 32, 10, 32, 32, 52, 54, 32, 50, 57, 32, 32, 10, 32, 53, 32, 32, 32, 51, 32, 50, 56, 10, 32, 32, 57, 51, 32, 32, 32, 55, 52, 10, 32, 52, 32, 32, 53, 32, 32, 51, 54, 10, 55, 32, 51, 32, 49, 56, 32, 32, 32, 10})
	gotBoard1, gotErr = Build([]uint8{32, 32, 32, 50, 54, 32, 55, 32, 49, 10, 54, 56, 32, 32, 55, 32, 32, 57, 32, 10, 49, 57, 32, 32, 32, 52, 53, 32, 32, 10, 56, 50, 32, 49, 32, 32, 32, 52, 32, 10, 32, 32, 52, 54, 32, 50, 57, 32, 32, 10, 32, 53, 32, 32, 32, 51, 32, 50, 56, 10, 32, 32, 57, 51, 32, 32, 32, 55, 52, 10, 32, 52, 32, 32, 53, 32, 32, 51, 54, 10, 55, 32, 51, 32, 49, 56, 32, 32, 32, 10})
	assert.Equal(expectedBoard1, gotBoard1, "9x9 board did not load correctly")
	assert.Equal(nil, gotErr, "Error should be nil for 9x9 board.")

	expectedBoard2 = Board{[][]uint8{{3, 2, 4, 1}, {4, 1, 3, 2}, {1, 4, 2, 3}, {2, 0, 0, 4}}}
	gotBoard2, gotErr = Build([]uint8{51, 50, 52, 49, 10, 52, 49, 51, 50, 10, 49, 52, 50, 51, 10, 50, 32, 32, 52, 10})
	assert.Equal(expectedBoard2, gotBoard2, "4x4 board did not load correctly.")
	assert.Equal(nil, gotErr, "Error should be nil for 4x4 board.")

	expectedBoard3 = Board{[][]uint8{{1, 2}, {2, 0}}}
	gotBoard3, gotErr = Build([]uint8{49, 50, 10, 50, 32, 10})
	assert.Equal(expectedBoard3, gotBoard3, "2x2 board did not load correctly.")
	assert.Equal(nil, gotErr, "Error should be nil for 2x2 board.")

	gotBoard4, gotErr = Build([]uint8{2})
	assert.Equal(expectedBoard4, gotBoard4, "Board should be zero value.")
	assert.Equal(ErrUnrecognizedBoardSize(1), gotErr, "Should have an error for single value board data")
}
