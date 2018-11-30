package solver

import (
	"testing"

	"github.com/Nabil-Daoud/robodoku/sudoku/board"
	"github.com/stretchr/testify/assert"
)

func TestGetRowHas(t *testing.T) {
	assert := assert.New(t)
	testBoard := board.Board{Spots: [][]uint8{{0, 0, 1, 2}, {0, 0, 3, 4}, {0, 1, 2, 3}, {1, 2, 3, 4}}}
	got := getRowHas(testBoard, 0)
	assert.Equal([]bool{true, true, false, false}, got, "getRowHas did not work.")
}

func TestGetColHas(t *testing.T) {
	assert := assert.New(t)
	testBoard := board.Board{Spots: [][]uint8{{0, 0, 1, 2}, {0, 0, 3, 4}, {0, 1, 2, 3}, {1, 2, 0, 4}}}
	got := getColHas(testBoard, 0)
	assert.Equal([]bool{true, false, false, false}, got, "getColHas did not work.")
}

func TestGetSqrHas(t *testing.T) {
	assert := assert.New(t)

	testBoard := board.Board{Spots: [][]uint8{{0, 1}, {2, 1}}}
	got := getSqrHas(testBoard, 0, 0)
	assert.Equal([]bool{false, false}, got, "getSqrHas did not work for 2x2.")

	testBoard = board.Board{Spots: [][]uint8{{0, 0, 1, 2}, {0, 0, 3, 4}, {0, 1, 2, 3}, {1, 2, 3, 4}}}
	got = getSqrHas(testBoard, 0, 3)
	assert.Equal([]bool{true, true, true, true}, got, "getSqrHas did not work.")
}

func TestSpotPossible(t *testing.T) {
	assert := assert.New(t)
	testBoard := board.Board{Spots: [][]uint8{{0, 0, 1, 2}, {0, 0, 3, 4}, {0, 1, 2, 3}, {1, 2, 0, 0}}}
	got := spotPossible(3, 2, testBoard)
	assert.Equal([]bool{false, false, false, true}, got, "spotPossible did not work.")
}

func TestSinglePossible(t *testing.T) {
	assert := assert.New(t)
	testBoard := board.Board{Spots: [][]uint8{{0, 0, 1, 2}, {0, 0, 3, 4}, {0, 1, 2, 3}, {1, 2, 0, 0}}}
	gotBool, gotValue := singlePossible(3, 2, testBoard)
	assert.Equal(true, gotBool, "singlePossible should be true.")
	assert.Equal(uint8(4), gotValue, "singlePossible value should be 4.")
}

func TestSolveSinglePossible(t *testing.T) {
	assert := assert.New(t)
	testBoard := board.Board{Spots: [][]uint8{{0, 0, 1, 2}, {0, 0, 3, 4}, {0, 1, 2, 3}, {0, 2, 0, 0}}}
	gotBool, gotBoard := solveSinglePossible(testBoard)
	assert.Equal(true, gotBool, "solveSinglePossible should be true (on first run).")
	expectedBoard := board.Board{Spots: [][]uint8{{0, 0, 1, 2}, {0, 0, 3, 4}, {4, 1, 2, 3}, {3, 2, 4, 1}}}
	assert.Equal(expectedBoard, gotBoard, "Board solved incorrectly.")

	gotBool, gotBoard = solveSinglePossible(gotBoard)
	assert.Equal(false, gotBool, "solveSinglePossible should be false (on second run).")
	expectedBoard = board.Board{Spots: [][]uint8{{0, 0, 1, 2}, {0, 0, 3, 4}, {4, 1, 2, 3}, {3, 2, 4, 1}}}
	assert.Equal(expectedBoard, gotBoard, "Board solved incorrectly.")
}

func TestOnlyPossibleInRow(t *testing.T) {
	assert := assert.New(t)

	testBoard := board.Board{Spots: [][]uint8{{0, 0, 0, 0}, {1, 2, 3, 4}, {2, 3, 4, 1}, {4, 1, 2, 3}}}
	gotBool, gotValue, gotColIndex := onlyPossibleInRow(0, testBoard)
	assert.Equal(true, gotBool, "onlyPossibleInRow should have returned isThere = true.")
	assert.Equal(1, int(gotValue), "onlyPossibleInRow should have returned value = 1")
	assert.Equal(2, gotColIndex, "onlyPossibleInRow should have returned colIndex = 0")

	testBoard = board.Board{Spots: [][]uint8{{0, 0, 0, 0}, {0, 2, 3, 4}, {2, 3, 4, 1}, {4, 1, 2, 3}}}
	gotBool, gotValue, gotColIndex = onlyPossibleInRow(0, testBoard)
	assert.Equal(true, gotBool, "onlyPossibleInRow should have returned isThere = true.")
	assert.Equal(2, int(gotValue), "onlyPossibleInRow should have returned value = 2")
	assert.Equal(3, gotColIndex, "onlyPossibleInRow should have returned colIndex = 0")

	testBoard = board.Board{Spots: [][]uint8{{0, 0, 0, 0}, {0, 2, 3, 4}, {0, 3, 4, 1}, {4, 1, 0, 3}}}
	gotBool, gotValue, gotColIndex = onlyPossibleInRow(0, testBoard)
	assert.Equal(true, gotBool, "onlyPossibleInRow should have returned isThere = true.")
	assert.Equal(3, int(gotValue), "onlyPossibleInRow should have returned value = 3")
	assert.Equal(0, gotColIndex, "onlyPossibleInRow should have returned colIndex = 0")

	testBoard = board.Board{Spots: [][]uint8{{0, 0, 0, 0}, {0, 2, 3, 4}, {0, 0, 4, 1}, {4, 1, 0, 3}}}
	gotBool, gotValue, gotColIndex = onlyPossibleInRow(0, testBoard)
	assert.Equal(true, gotBool, "onlyPossibleInRow should have returned isThere = true.")
	assert.Equal(4, int(gotValue), "onlyPossibleInRow should have returned value = 3")
	assert.Equal(1, gotColIndex, "onlyPossibleInRow should have returned colIndex = 0")

	testBoard = board.Board{Spots: [][]uint8{{0, 0, 0, 0}, {0, 2, 3, 4}, {0, 0, 4, 1}, {0, 1, 0, 3}}}
	gotBool, gotValue, gotColIndex = onlyPossibleInRow(0, testBoard)
	assert.Equal(false, gotBool, "onlyPossibleInRow should have returned isThere = true.")
	assert.Equal(0, int(gotValue), "onlyPossibleInRow should have returned value = 3")
	assert.Equal(0, gotColIndex, "onlyPossibleInRow should have returned colIndex = 0")

}

func TestOnlyPossibleInCol(t *testing.T) {
	assert := assert.New(t)

	testBoard := board.Board{Spots: [][]uint8{{0, 1, 2, 4}, {0, 2, 3, 1}, {0, 3, 4, 2}, {0, 4, 1, 3}}}
	gotBool, gotValue, gotColIndex := onlyPossibleInCol(0, testBoard)
	assert.Equal(true, gotBool, "onlyPossibleInCol should have returned isThere = true.")
	assert.Equal(1, int(gotValue), "onlyPossibleInCol should have returned value = 1")
	assert.Equal(2, gotColIndex, "onlyPossibleInCol should have returned colIndex = 0")

	testBoard = board.Board{Spots: [][]uint8{{0, 0, 2, 4}, {0, 2, 3, 1}, {0, 3, 4, 2}, {0, 4, 1, 3}}}
	gotBool, gotValue, gotColIndex = onlyPossibleInCol(0, testBoard)
	assert.Equal(true, gotBool, "onlyPossibleInCol should have returned isThere = true.")
	assert.Equal(2, int(gotValue), "onlyPossibleInCol should have returned value = 2")
	assert.Equal(3, gotColIndex, "onlyPossibleInCol should have returned colIndex = 0")

	testBoard = board.Board{Spots: [][]uint8{{0, 0, 0, 0}, {0, 2, 3, 4}, {0, 3, 4, 1}, {4, 1, 0, 3}}}
	gotBool, gotValue, gotColIndex = onlyPossibleInCol(0, testBoard)
	assert.Equal(true, gotBool, "onlyPossibleInCol should have returned isThere = true.")
	assert.Equal(3, int(gotValue), "onlyPossibleInCol should have returned value = 3")
	assert.Equal(0, gotColIndex, "onlyPossibleInCol should have returned colIndex = 0")

	testBoard = board.Board{Spots: [][]uint8{{0, 0, 0, 4}, {0, 2, 0, 1}, {0, 3, 4, 0}, {0, 4, 1, 3}}}
	gotBool, gotValue, gotColIndex = onlyPossibleInCol(0, testBoard)
	assert.Equal(true, gotBool, "onlyPossibleInCol should have returned isThere = true.")
	assert.Equal(4, int(gotValue), "onlyPossibleInCol should have returned value = 4")
	assert.Equal(1, gotColIndex, "onlyPossibleInCol should have returned colIndex = 1")

	testBoard = board.Board{Spots: [][]uint8{{0, 0, 0, 0}, {0, 2, 3, 4}, {0, 0, 4, 1}, {0, 1, 0, 3}}}
	gotBool, gotValue, gotColIndex = onlyPossibleInCol(0, testBoard)
	assert.Equal(false, gotBool, "onlyPossibleInCol should have returned isThere = false.")
	assert.Equal(0, int(gotValue), "onlyPossibleInCol should have returned value = 0")
	assert.Equal(0, gotColIndex, "onlyPossibleInCol should have returned colIndex = 0")

}

func TestOnlyPossibleInSqr(t *testing.T) {
	assert := assert.New(t)

	testBoard := board.Board{Spots: [][]uint8{{0, 0, 1, 0}, {0, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}}}
	gotBool, gotValue, gotRowIndex, gotColIndex := onlyPossibleInSqr(0, 0, testBoard)
	assert.Equal(true, gotBool, "onlyPossibleInSqr should have returned isThere = true.")
	assert.Equal(1, int(gotValue), "onlyPossibleInSqr should have returned value = 1")
	assert.Equal(1, gotRowIndex, "onlyPossibleInSqr should have returned rowIndex = 1")
	assert.Equal(0, gotColIndex, "onlyPossibleInSqr should have returned colIndex = 0")

	testBoard = board.Board{Spots: [][]uint8{{0, 0, 0, 0}, {0, 0, 1, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}}}
	gotBool, gotValue, gotRowIndex, gotColIndex = onlyPossibleInSqr(0, 0, testBoard)
	assert.Equal(true, gotBool, "onlyPossibleInSqr should have returned isThere = true.")
	assert.Equal(1, int(gotValue), "onlyPossibleInSqr should have returned value = 1")
	assert.Equal(0, gotRowIndex, "onlyPossibleInSqr should have returned rowIndex = 0")
	assert.Equal(0, gotColIndex, "onlyPossibleInSqr should have returned colIndex = 0")

	testBoard = board.Board{Spots: [][]uint8{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 1, 0, 0}, {0, 0, 0, 0}}}
	gotBool, gotValue, gotRowIndex, gotColIndex = onlyPossibleInSqr(0, 0, testBoard)
	assert.Equal(false, gotBool, "isThere")
	assert.Equal(0, int(gotValue), "value")
	assert.Equal(0, gotRowIndex, "rowIndex")
	assert.Equal(0, gotColIndex, "colIndex")
}

func Test(t *testing.T) {

	testBoard, _ := board.Build([]uint8{32, 32, 32, 50, 54, 32, 55, 32, 49, 10, 54, 56, 32, 32, 55, 32, 32, 57, 32, 10, 49, 57, 32, 32, 32, 52, 53, 32, 32, 10, 56, 50, 32, 49, 32, 32, 32, 52, 32, 10, 32, 32, 52, 54, 32, 50, 57, 32, 32, 10, 32, 53, 32, 32, 32, 51, 32, 50, 56, 10, 32, 32, 57, 51, 32, 32, 32, 55, 52, 10, 32, 52, 32, 32, 53, 32, 32, 51, 54, 10, 55, 32, 51, 32, 49, 56, 32, 32, 32, 10})
	Phase1(testBoard)
	Phase2(testBoard)

	solveOnlyPossibleRow(testBoard)
	solveOnlyPossibleCol(testBoard)
	// var tests = []struct {
	// 	value uint8
	// 	want  int
	// }{
	// 	{1, 1},
	// 	{2, 2},
	// }
	// var got int
	// for _, c := range tests {
	// 	got = int(c.value)
	// 	if got != c.want {
	// 		t.Errorf("Cell(%v) == %v, want %v", c.value, got, c.want)
	// 	}
	// }
}
