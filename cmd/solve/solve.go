package solve

import (
	"fmt"

	"github.com/Nabil-Daoud/robodoku/solver"
	"github.com/Nabil-Daoud/robodoku/sudoku"
	"github.com/Nabil-Daoud/robodoku/sudoku/board"
)

// Action solves the Sudoku puzzle
func Action(filepath string) {
	// var err error
	fmt.Printf("the file path read as %v\n", filepath)

	var myBoard board.Board
	dat := sudoku.ReadFile(filepath)
	myBoard, _ = board.Build(dat)

	fmt.Println("\nInitial Board:")
	myBoard.PrintBoard()

	fmt.Println()

	solvedBoard, didChange := solver.Phase1(myBoard)
	fmt.Printf("Board after Phase1: (changed = %t)\n", didChange)
	solvedBoard.PrintBoard()

	solvedBoard, didChange = solver.Phase2(solvedBoard)
	fmt.Printf("\nBoard after Phase2: (changed = %t)\n", didChange)
	solvedBoard.PrintBoard()

	solvedBoard, didChange = solver.Phase1(solvedBoard)
	fmt.Printf("\nBoard after Phase1 again: (changed = %t)\n", didChange)
	solvedBoard.PrintBoard()
}
