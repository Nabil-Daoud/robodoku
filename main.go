package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Nabil-Daoud/robodoku/solver"
	"github.com/Nabil-Daoud/robodoku/sudoku/board"
	"github.com/gocraft/health"
)

// puzzle_path = ARGV[0]
// puzzle_text = File.read(puzzle_path)
// solver = Solver.new(puzzle_text)
// solver.solve()

var stream = health.NewStream()

func main() {
	var err error
	var currentDir string
	// Log to stdout (can also use WriteSink to write to a log file, Syslog, etc.)
	stream.AddSink(&health.WriterSink{os.Stdout})

	// Logging and instrumentation happens within the context of a job.
	job := stream.NewJob("Robodoku Setup")
	currentDir, err = os.Getwd()

	if err != nil {
		job.EventErr("os.Getwd", err)
	}

	if err == nil {
		job.Complete(health.Success)
	} else {
		job.Complete(health.Error)
	}

	fmt.Printf("The current directory is %s\n", currentDir)

	var myBoard board.Board
	dat := ReadFile("/Users/nabil/go/src/github.com/Nabil-Daoud/robodoku/testdata/medium1.txt")
	myBoard, err = board.Build(dat)

	myBoard.PrintBoard()

	fmt.Println()

	solvedBoard := solver.Phase1(myBoard)
	fmt.Println("Board after Phase1:")
	solvedBoard.PrintBoard()

	solvedBoard = solver.Phase2(solvedBoard)
	fmt.Println("\nBoard after Phase2:")
	solvedBoard.PrintBoard()

	solvedBoard = solver.Phase1(solvedBoard)
	fmt.Println("\nBoard after Phase1 again:")
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
