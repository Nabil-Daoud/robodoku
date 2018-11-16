package main

import (
	// "os"
	"fmt"
	// "path/filepath"
	//
	// "github.com/Nabil-Daoud/robodoku/solver"
	"github.com/Nabil-Daoud/robodoku/sudoku/board"
)

// puzzle_path = ARGV[0]
// puzzle_text = File.read(puzzle_path)
// solver = Solver.new(puzzle_text)
// solver.solve()

func main() {
	var myBoard board.Board

	fmt.Println("easy.txt >>>>>>>>>>>>>>")
	myBoard.PopulateBoard("/Users/nabil/go/src/github.com/Nabil-Daoud/robodoku/puzzles/easy.txt")
	fmt.Println("Populated, using PrintBoard()")
	myBoard.PrintBoard()

	myBoard.UpdatePossible()

}
