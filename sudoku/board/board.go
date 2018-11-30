package board

import "fmt"

// Board is a structure that holds an N x N Sudoku board.
type Board struct {
	Spots [][]uint8
}

// ErrUnrecognizedBoardSize is an error for board data of unrecognized lenght
type ErrUnrecognizedBoardSize int

func (e ErrUnrecognizedBoardSize) Error() string {
	return fmt.Sprintf("Unrecognized board data length: %b", int(e))
}

// Build takes a slice of uint8, creates and populates known board values
func Build(dat []uint8) (Board, error) {
	var (
		board Board
		n     int
	)

	if len(dat) == 90 { // standard 9x9 board
		n = 9
	} else if len(dat) == 20 { // 4x4 board
		n = 4
	} else if len(dat) == 6 { // 2x2 board
		n = 2
	} else {
		return board, ErrUnrecognizedBoardSize(len(dat))
	}
	board = populate(n, dat, newBoard(n))
	return board, nil
}

// newBoard takes the dimension, n, and returns an nxn Board.
func newBoard(n int) Board {
	b := make([][]uint8, n)
	for i := range b {
		b[i] = make([]uint8, n)
	}
	return Board{b}
}

// populate takes the board size, n; board data, dat; and a Board.
// It returns a board with the values filled in.
func populate(n int, dat []uint8, board Board) Board {
	const (
		blank uint8 = 32
		eol   uint8 = 10
	)
	for i, item := range dat {
		// if the value is not eol and not blank set the value
		if item != eol && item != blank {
			// TODO confirm that the value is in the range 0 < x <= len(board.Spots)
			board.Spots[i/(n+1)][i%(n+1)] = (item - 48)
		}
	}
	return board
}

// PrintBoard prints the solved values on the Sodoku board.
// TODO: Make this function return a string for any size board
func (board *Board) PrintBoard() {
	for i, row := range board.Spots {
		for j, spot := range row {
			if spot != uint8(0) {
				fmt.Printf("%v ", spot)
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
