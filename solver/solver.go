package solver

import (
	"github.com/Nabil-Daoud/robodoku/sudoku/board"
)

func getRowHas(board board.Board, rowIndex int) []bool {
	var rowHas = make([]bool, len(board.Spots))
	for _, value := range board.Spots[rowIndex] {
		if value != uint8(0) {
			rowHas[value-1] = true
		}
	}
	return rowHas
}

func getColHas(board board.Board, colIndex int) []bool {
	var colHas = make([]bool, len(board.Spots))
	for rowIndex := range board.Spots {
		value := board.Spots[rowIndex][colIndex]
		if value != uint8(0) {
			colHas[value-1] = true
		}
	}
	return colHas
}

func getSqrHas(board board.Board, rowIndex int, colIndex int) []bool {
	boardLength := len(board.Spots)
	var squareLength int
	sqrHas := make([]bool, boardLength)

	if boardLength == 9 {
		squareLength = 3
	} else if boardLength == 4 {
		squareLength = 2
	} else if boardLength == 2 {
		squareLength = 2
	}

	for i := 0; i < squareLength; i++ {
		for j := 0; j < squareLength; j++ {
			spotRow := (rowIndex - rowIndex%squareLength) + i
			spotCol := (colIndex - colIndex%squareLength) + j
			value := board.Spots[spotRow][spotCol]
			if value != uint8(0) {
				sqrHas[value-1] = true
			}
		}
	}
	return sqrHas
}

func spotPossible(rowIndex int, colIndex int, board board.Board) []bool {
	rowHas := getRowHas(board, rowIndex)
	colHas := getColHas(board, colIndex)
	sqrHas := getSqrHas(board, rowIndex, colIndex)
	possibleVal := make([]bool, len(board.Spots))

	for i := range possibleVal {
		possibleVal[i] = !rowHas[i] && !colHas[i] && !sqrHas[i]
	}
	return possibleVal
}

func singlePossible(rowIndex int, colIndex int, board board.Board) (bool, uint8) {
	var numPossible, value int
	possibleVal := spotPossible(rowIndex, colIndex, board)

	for i := range possibleVal {
		if possibleVal[i] {
			numPossible++
			value = i + 1
		}
	}

	if numPossible == 1 {
		return true, uint8(value)
	}
	return false, 0
}

func solveSinglePossible(board board.Board) (bool, board.Board) {
	var didSolve bool

	for rowIndex, row := range board.Spots {
		for colIndex, value := range row {
			if value == uint8(0) {
				hasSinglePossible, val := singlePossible(rowIndex, colIndex, board)
				if hasSinglePossible {
					didSolve = true
					board.Spots[rowIndex][colIndex] = val
				}
			}
		}
	}
	return didSolve, board
}

// Phase1 solves unsolved spots that have only one possible value based on
// other values in its row, column, and square
func Phase1(board board.Board) board.Board {
	var didSolve, done bool

	// count := 0
	for !done {
		// count++
		// fmt.Printf("Lap count: %v\n", count)
		didSolve, board = solveSinglePossible(board)
		done = !didSolve
	}
	// for board.UpdatePossible() {
	// 	board.SolveSinglePossible()
	// }
	return board
}

// Phase2 goes through the rows, columns and squares looking for and solving
// squares that have the only possilbe for any given value
// func Phase2(board board.Board) board.Board {
// 	return board
// }
