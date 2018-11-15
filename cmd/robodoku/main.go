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
	myBoard.PopulateBoard("github.com/Nabil-Daoud/robodoku/puzzles/easy.txt")
	fmt.Println("Populated, using PrintBoard()")
	myBoard.PrintBoard()
	// fmt.Println("easy_solution.txt >>>>>>>>>>>>>>")
	// myBoard.PopulateBoard("github.com/Nabil-Daoud/robodoku/puzzles/easy_solution.txt")
	// fmt.Println("two_by_two_trivial.txt >>>>>>>>>>>>>>")
	// myBoard.PopulateBoard("github.com/Nabil-Daoud/robodoku/puzzles/two_by_two_trivial.txt")
	// fmt.Println("medium.txt >>>>>>>>>>>>>>")
	// myBoard.PopulateBoard("github.com/Nabil-Daoud/robodoku/puzzles/medium.txt")
}
