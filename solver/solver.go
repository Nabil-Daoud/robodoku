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
	} else if boardLength == 2 { // this case is degenerate
		return sqrHas
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
			if value == uint8(0) { // if unsolved
				hasSinglePossible, val := singlePossible(rowIndex, colIndex, board)
				if hasSinglePossible {
					didSolve = true
					board.Spots[rowIndex][colIndex] = val
					// fmt.Println("SolveSinglePossible got one!")
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
	return board
}

func onlyPossibleInRow(rowIndex int, board board.Board) (isThere bool, value uint8, colIndex int) {
	var numCandidate int

	// loop through all the values
	for i, hasValue := range getRowHas(board, rowIndex) {
		// if the row is missing a value
		if !hasValue {
			// loop through the spots in the row
			for j := range board.Spots {
				// if the spot can be the value from the first loop
				if spotPossible(rowIndex, j, board)[i] {
					// increment the number of candidates
					numCandidate++
					value = uint8(i + 1)
					colIndex = j
				} else {
				}
			}
			if numCandidate == 1 {
				isThere = true
				return
			}
			numCandidate = 0
		}
	}
	return false, 0, 0
}

func solveOnlyPossibleRow(board board.Board) (bool, board.Board) {
	var didSolve bool

	for rowIndex := range board.Spots {
		isThere, value, colIndex := onlyPossibleInRow(rowIndex, board)
		if isThere {
			// fmt.Println("On board:")
			// board.PrintBoard()
			// fmt.Printf("Spot[%v][%v] is only in row that can be %v\n\n", rowIndex, colIndex, value)

			didSolve = true
			board.Spots[rowIndex][colIndex] = value

		}
	}
	return didSolve, board
}

func onlyPossibleInCol(colIndex int, board board.Board) (isThere bool, value uint8, rowIndex int) {
	var numCandidate int

	// loop through all the values
	for i, hasValue := range getColHas(board, colIndex) {
		// if the column is missing a value
		if !hasValue {
			// loop through the spots in the column
			for j := range board.Spots {
				// if the spot can be the value from the first loop
				if spotPossible(j, colIndex, board)[i] {
					// increment the number of candidates
					numCandidate++
					value = uint8(i + 1)
					rowIndex = j
				}
			}
			if numCandidate == 1 {
				isThere = true
				return
			}
			numCandidate = 0
		}
	}
	return false, 0, 0
}

func solveOnlyPossibleCol(board board.Board) (bool, board.Board) {
	var didSolve bool

	for colIndex := range board.Spots {
		isThere, value, rowIndex := onlyPossibleInCol(colIndex, board)
		if isThere {
			// fmt.Println("On board:")
			// board.PrintBoard()
			// fmt.Printf("Spot[%v][%v] is only in column that can be %v\n\n", rowIndex, colIndex, value)

			didSolve = true
			board.Spots[rowIndex][colIndex] = value
		}
	}
	return didSolve, board
}

func onlyPossibleInSqr(sqrRowIndex int, sqrColIndex int, board board.Board) (isThere bool, value uint8, rowIndex int, colIndex int) {
	var numCandidate int
	boardLength := len(board.Spots)
	var squareLength int

	if boardLength == 9 {
		squareLength = 3
	} else if boardLength == 4 {
		squareLength = 2
	} else if boardLength == 2 { // this case is degenerate
		return false, 0, 0, 0
	}

	// loop through all the values
	for i, hasValue := range getSqrHas(board, sqrRowIndex*squareLength, sqrColIndex*squareLength) {
		// fmt.Printf("Checking value %v\n", i+1)

		// if the square is missing a value
		if !hasValue {
			// fmt.Printf("The [%v],[%v] square does not have %v\n", sqrRowIndex, sqrColIndex, i+1)
			// loop through the spots in the square
			for j := 0; j < squareLength; j++ {
				for k := 0; k < squareLength; k++ {
					// if the spot can be the value from the first loop
					if spotPossible(sqrRowIndex*squareLength+j, sqrColIndex*squareLength+k, board)[i] {
						// fmt.Printf("Spot at [%v][%v] can be %v\n", sqrRowIndex*squareLength+j, sqrColIndex*squareLength+k, i+1)

						// increment the number of candidates
						numCandidate++
						value = uint8(i + 1)
						rowIndex = j
					}
				}
			}
			if numCandidate == 1 {
				isThere = true
				return
			}
			numCandidate = 0
		}
	}
	return false, 0, 0, 0
}

func solveOnlyPossibleSqr(board board.Board) (bool, board.Board) {
	var didSolve bool
	var numSquares int
	boardLength := len(board.Spots)

	if boardLength == 9 {
		numSquares = 3
	} else if boardLength == 4 {
		numSquares = 2
	} else if boardLength == 2 { // this case is degenerate
		return false, board
	}

	for sqrRowIndex := 0; sqrRowIndex < numSquares; sqrRowIndex++ {
		for sqrColIndex := 0; sqrColIndex < numSquares; sqrColIndex++ {
			isThere, value, rowIndex, colIndex := onlyPossibleInSqr(sqrRowIndex, sqrColIndex, board)
			if isThere {
				// fmt.Println("On board:")
				// board.PrintBoard()
				// fmt.Printf("Spot[%v][%v] is only in square that can be %v\n\n", rowIndex, colIndex, value)

				didSolve = true
				board.Spots[rowIndex][colIndex] = value
			}
		}
	}
	return didSolve, board
}

func solveOnlyPossible(board board.Board) (didSolve bool, solvedBoard board.Board) {
	var didSolveRow, didSolveCol, didSolveSqr bool

	solvedBoard = board

	didSolveRow, solvedBoard = solveOnlyPossibleRow(solvedBoard)
	didSolveCol, solvedBoard = solveOnlyPossibleCol(solvedBoard)
	didSolveSqr, solvedBoard = solveOnlyPossibleSqr(solvedBoard)

	didSolve = didSolveRow || didSolveCol || didSolveSqr
	// fmt.Println("solveOnlyPossible got one!")
	return

	// for rowIndex, row := range board.Spots {
	// 	for colIndex, value := range row {
	// 		if value == uint8(0) { // if unsolved
	// 			hasSinglePossible, val := singlePossible(rowIndex, colIndex, board)
	// 			if hasSinglePossible {
	// 				didSolve = true
	// 				board.Spots[rowIndex][colIndex] = val
	// 			}
	// 		}
	// 	}
	// }

}

// Phase2 goes through the rows, columns and squares (peer groups) and solve
// if a spot is the only one of its peers that can be a given value
func Phase2(board board.Board) board.Board {
	var didSolve, done bool
	// count := 0
	for !done {
		// count++
		// fmt.Printf("Lap count: %v\n", count)
		didSolve, board = solveOnlyPossible(board)
		done = !didSolve
	}
	return board
}
