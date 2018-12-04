package main

import (
	"fmt"
	"os"

	"github.com/Nabil-Daoud/robodoku/cmd/solve"
	"github.com/Nabil-Daoud/robodoku/sudoku"
	"github.com/Nabil-Daoud/robodoku/sudoku/board"
	"github.com/gocraft/health"
	cli "github.com/jawher/mow.cli"
)

// (gocraft/health) stream - allows for app logging and metrics
var stream = health.NewStream()

// (jawher/mow.cli) Global app option available to any of the commands
var filename *string
var filenameSetByUser bool

func main() {
<<<<<<< HEAD
	var myBoard board.Board

	// fmt.Println("easy.txt >>>>>>>>>>>>>>")
	// dat := ReadFile("/Users/nabil/go/src/github.com/Nabil-Daoud/robodoku/testdata/easy.txt")
	// fmt.Printf("\ndat is of type %T\n", dat)
	// fmt.Printf("%v\n", len(dat))
	myBoard.PopulateBoard("/Users/nabil/go/src/github.com/Nabil-Daoud/robodoku/testdata/easy.txt")
	myBoard.PrintBoard()

	fmt.Println()

	solvedBoard := solver.Phase1(myBoard)
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
=======
	var err error
>>>>>>> functional

	// Log to stdout (can also use WriteSink to write to a log file, Syslog, etc.)
	stream.AddSink(&health.WriterSink{os.Stdout})
	setup(err)

	app := cli.App("robodoku", "Sudoku player and solver")
	// (jawher/mow.cli) Define our top-level global option(s)
	filename = app.StringOpt("f file", "easy.txt", "puzzle file name")

	// Declare the first command, which is invocable with "robodoku solve"
	app.Command("solve", "solve the sudoku puzzle", func(cmd *cli.Cmd) {
		filepath := "/Users/nabil/go/src/github.com/Nabil-Daoud/robodoku/testdata/" + *filename
		// Run this function when the solve command is invoked
		solve.Action(fmt.Sprint(filepath))
	})

	// Declare the second command, which is invocable with "robodoku play"
	app.Command("play", "play the sudoku puzzle", func(cmd *cli.Cmd) {
		// Run this function when the play command is invoked
		cmd.Action = func() {
			filepath := "/Users/nabil/go/src/github.com/Nabil-Daoud/robodoku/testdata/" + *filename

			var myBoard board.Board
			dat := sudoku.ReadFile(filepath)
			myBoard, err = board.Build(dat)
			myBoard.PrintBoard()
			fmt.Println("This game mode is not currently implemented.")
		}
	})

	app.Run(os.Args)
}

func setup(err error) {
	// Logging and instrumentation happens within the context of a job.
	job := stream.NewJob("Robodoku Setup")
	var currentDir string
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
}
