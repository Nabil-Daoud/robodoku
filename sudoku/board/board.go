package board

import (
	"fmt"
	"io/ioutil"

	"github.com/Nabil-Daoud/robodoku/sudoku/spot"
)

// Board is a structure that holds an N x N Sudoku board.
type Board struct {
	N     int
	spots [][]spot.Spot
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// PopulateBoard takes a board and a filename and populates the spots
func PopulateBoard(board *Board, filename string) {
	dat, err := ioutil.ReadFile(filename)
	check(err)
	// lines := []string{strings.Split(string(dat), "\n")}
	for _, line := range dat {
		fmt.Print(string(line))
	}
	fmt.Println("<<<<<<<<<<<<<<<< length of dat = ", len(dat))
	fmt.Println("")
	for _, thing := range dat {
		fmt.Print(thing, " . ")
		if thing == 10 {
			fmt.Println("")
		}
	}
	fmt.Println("")
	// board.spots[0][0] = spot.NewSpot(9, '1')

}

// PrintBoard prints the solved values on the Sodoku board.
func PrintBoard() {
	fmt.Println('a')
}
