package sudoku

import "io/ioutil"

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
