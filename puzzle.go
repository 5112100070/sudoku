package main

import "fmt"

type Puzzle struct {
	Board [][]int
}

func (p Puzzle) IsValidSudoku() bool {
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			if !p.isDigitValid(row, col, p.Board[row][col]) {
				return false
			}
		}
	}

	return true
}

func (p Puzzle) PrintBoard() {
	for row := 0; row < N; row++ {
		for col := 0; col < N; col++ {
			fmt.Printf(" %d ", p.Board[row][col])
		}
		fmt.Println("")
	}
}

func (p Puzzle) isDigitValid(row, col, digit int) bool {
	// note: integer division 'rounds down' by ignoring all decimal places
	startRow := row / 3 * 3
	startCol := col / 3 * 3

	var isFirstDataH bool
	var isFirstDataV bool
	var isFirstData3x3 bool
	var firstDataPosH, firstDataPosV, firstDataPos3x3 string

	for i := 0; i < N; i++ {
		// check the corresponding row and column
		if (p.Board[row][i] == digit ||
			p.Board[i][col] == digit ||
			// check the corresponding 3x3 section
			p.Board[startRow+i/3][startCol+i%3] == digit) && (isFirstDataH && isFirstDataV && isFirstData3x3) {

			if isFirstDataH {
				fmt.Printf("Wrong data on position %d,%d and %s\n", row, i, firstDataPosH)
			}

			if isFirstDataV {
				fmt.Printf("Wrong data on position %d,%d and %s\n", i, col, firstDataPosV)
			}

			if isFirstData3x3 {
				fmt.Printf("Wrong data on position %d,%d and %s", startRow+i/3, startCol+i%3, firstDataPos3x3)
			}

			return false
		}

		if p.Board[row][i] == digit && !isFirstDataH {
			firstDataPosH = fmt.Sprintf("%d,%d", row, i)
			isFirstDataH = true
		}

		if p.Board[i][col] == digit && !isFirstDataV {
			firstDataPosV = fmt.Sprintf("%d,%d", i, col)
			isFirstDataV = true
		}

		if p.Board[startRow+i/3][startCol+i%3] == digit && !isFirstData3x3 {
			firstDataPos3x3 = fmt.Sprintf("%d,%d", startRow+i/3, startCol+i%3)
			isFirstData3x3 = true
		}
	}

	return true
}
