package board

import (
	"fmt"
	"io/ioutil"

	"github.com/Nabil-Daoud/robodoku/sudoku/spot"
)

// Board is a structure that holds an N x N Sudoku board.
type Board struct {
	N     int
	spots [9][9]spot.Spot
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// PopulateBoard takes a filename and populates the spots
func (board *Board) PopulateBoard(filename string) {
	dat, err := ioutil.ReadFile(filename)
	check(err)

	// fmt.Printf("The data is read in as a %T and is of length %v\n", dat, len(dat))
	// lines := []string{strings.Split(string(dat), "\n")}

	// The code bewlow can be used to determine the dimensions of the board.
	// for i := 0; i < len(dat); i++ {
	// 	fmt.Printf("i: %v\n", i)
	// 	if dat[i] == 10 {
	// 		fmt.Println(">>>> In the if block.")
	// 		board.N = i
	// 		fmt.Printf("Board size, N set to %v\n", i)
	// 		fmt.Printf("N = %v\n", i)
	// 		break
	// 	}
	// 	fmt.Println("Loop.")
	// }
	// fmt.Println("<<<< made it out of the for loop. phew")
	// fmt.Println("the board is ", board.N, " by ", string(board.N), ".")

	for i, item := range dat {
		// fmt.Print(string(item))
		// fmt.Print(" | ", i, ": ", string(item))
		if item != uint8(10) {
			// if it's not a blank, solve it; otherwise set all possible to true.
			if item != uint8(32) {
				board.spots[i/10][i%10].SolveSpot(item - 48)
			} else {
				board.spots[i/10][i%10].Possible = [9]bool{true, true, true, true, true, true, true, true, true}
			}
		}
	}

	// fmt.Println("\n<<<<<<<<<<<<<<<< length of dat = ", len(dat))
	// fmt.Println("")
	// for _, thing := range dat {
	// 	fmt.Print(thing, " . ")
	// 	if thing == 10 {
	// 		fmt.Println("")
	// 	}
	// }
	// fmt.Println("")

}

// PrintBoard prints the solved values on the Sodoku board.
func (board *Board) PrintBoard() {
	for i, row := range board.spots {
		for j, spot := range row {
			if spot.Solved {
				fmt.Printf("%v ", spot.Value)
			} else {
				fmt.Print("  ")
			}

			if j%3 == 2 && j != 8 {
				fmt.Print("|")
			}
		}
		fmt.Print("\n")
		if i%3 == 2 && i != 8 {
			fmt.Println("------|------|------")
		}
	}
}

// UpdatePossible takes the board and updates the possible values based on rows,
// columns, and squares. It returns true if a possible square was updated.
func (board *Board) UpdatePossible() bool {
	return board.updatePossibleRow() || board.updatePossibleCol() // || board.updatePossibleSqr()
}

// updatePossibleRow() loops through the rows and updates the Possible array
// based on solved spots in the same row.
func (board *Board) updatePossibleRow() bool {
	updated := false

	board.N = 8

	// Loop through the rows
	for _, row := range board.spots { // i
		// Loop through the spots in the row to populate the values the row has
		rowHas := [9]bool{} // false, false, false, false, false, false, false, false, false}
		for _, spot := range row {
			if spot.Solved {
				rowHas[spot.Value-1] = true
				// fmt.Printf("Row %v has value %v. ", i, spot.Value)
			}
		}
		// fmt.Println("rowHas: ", rowHas)

		// Loop through the spots in the row again and update the possible values
		// based on what the column has.
		for _, spot := range row { // k
			if !spot.Solved {
				fmt.Printf("Possible before: %t\n", spot.Possible)
				for j, possible := range spot.Possible {
					if possible && rowHas[j] {
						// fmt.Printf("Updating spot at [%v, %v], can't be %v\n", i, k, j+1)
						// fmt.Printf("Row %v, spot %v has %v. ", i, k, j+1)
						spot.Possible[j] = false
						updated = true
					}
				}
				fmt.Printf("Possible after: %t\n\n", spot.Possible)
			}
		}
	}
	return updated
}

// updatePossibleCol() loops through the columns and updates the Possible array
// based on solved spots in the same column.
func (board *Board) updatePossibleCol() bool {
	updated := false
	// Loop through the columns
	for i := range board.spots {
		// Loop through the spots in the column to populate the values the column has
		var colHas = [9]bool{}
		for j := range board.spots {
			if board.spots[j][i].Solved {
				colHas[board.spots[j][i].Value-1] = true
			}
		}
		// Loop through the spots in the column again and update their possible
		// values based on what the column has.
		for j := range board.spots {
			if !board.spots[j][i].Solved {
				for k, possible := range board.spots[j][i].Possible {
					if possible && colHas[k] {
						board.spots[j][i].Possible[k] = false
						updated = true
					}
				}
			}
		}
	}
	return updated
}
