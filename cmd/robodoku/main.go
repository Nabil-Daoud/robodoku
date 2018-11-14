package main

import (
	// "os"
	"fmt"
	// "path/filepath"
	//
	// "robodoku/lib/solver.go"
	"github.com/Nabil-Daoud/robodoku/sudoku/board"
)

// puzzle_path = ARGV[0]
// puzzle_text = File.read(puzzle_path)
// solver = Solver.new(puzzle_text)
// solver.solve()

func main() {
	var myBoard *board.Board
	fmt.Println("easy.txt >>>>>>>>>>>>>>")
	board.PopulateBoard(myBoard, "github.com/Nabil-Daoud/robodoku/puzzles/easy.txt")
	fmt.Println("easy_solution.txt >>>>>>>>>>>>>>")
	board.PopulateBoard(myBoard, "github.com/Nabil-Daoud/robodoku/puzzles/easy_solution.txt")
	fmt.Println("two_by_two_trivial.txt >>>>>>>>>>>>>>")
	board.PopulateBoard(myBoard, "github.com/Nabil-Daoud/robodoku/puzzles/two_by_two_trivial.txt")
	fmt.Println("medium.txt >>>>>>>>>>>>>>")
	board.PopulateBoard(myBoard, "github.com/Nabil-Daoud/robodoku/puzzles/medium.txt")
}
