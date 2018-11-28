package main

import (
	// "os"

	"fmt"
	"io/ioutil"
	// "path/filepath"
	"github.com/Nabil-Daoud/robodoku/solver"
	"github.com/Nabil-Daoud/robodoku/sudoku/board"
)

// puzzle_path = ARGV[0]
// puzzle_text = File.read(puzzle_path)
// solver = Solver.new(puzzle_text)
// solver.solve()

func main() {
	var myBoard board.Board
	dat := ReadFile("/Users/nabil/go/src/github.com/Nabil-Daoud/robodoku/testdata/easy.txt")
	myBoard = board.Build(dat)

	myBoard.PrintBoard()

	fmt.Println()

	solvedBoard := solver.Phase1(myBoard)
	fmt.Println("Board after Phase1:")
	solvedBoard.PrintBoard()

	// fmt.Println("four_by_four.txt >>>>>>>>>>>>>>")
	// dat = sudoku.ReadFile("/Users/nabil/go/src/github.com/Nabil-Daoud/robodoku/testdata/four_by_four_trivial.txt")
	// fmt.Printf("%v\n", len(dat))
	// fmt.Println("two_by_two.txt >>>>>>>>>>>>>>")
	// dat = sudoku.ReadFile("/Users/nabil/go/src/github.com/Nabil-Daoud/robodoku/testdata/two_by_two_trivial.txt")
	// fmt.Printf("%v\n", len(dat))
	// fmt.Println("medium.txt >>>>>>>>>>>>>>")
	// dat = sudoku.ReadFile("/Users/nabil/go/src/github.com/Nabil-Daoud/robodoku/testdata/medium.txt")
	// fmt.Printf("%v\n", len(dat))
}

// ReadFile takes a board from a file and returns it as an array of uint8.
func ReadFile(filename string) []uint8 {
	dat, err := ioutil.ReadFile(filename)
	check(err)
	return dat
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
