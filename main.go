package main

import (
	"fmt"
)

var N = 9

func main() {
	cases := []Cases{
		Case1, Case2, Case3, Case4, Case5, Case6, Case7,
	}

	for i, c := range cases {
		puzzle := Puzzle{
			Board: c,
		}

		var isValidStr string
		puzzle.PrintBoard()
		if puzzle.IsValidSudoku() {
			isValidStr = "Valid"
		} else {
			isValidStr = "Not Valid"
		}

		fmt.Println(fmt.Sprintf("SUDOKU Case Number: %d is %s", i+1, isValidStr))
	}
}
