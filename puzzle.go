package main

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

func (p Puzzle) isDigitValid(row, col, digit int) bool {
	// note: integer division 'rounds down' by ignoring all decimal places
	startRow := row / 3 * 3
	startCol := col / 3 * 3

	var isFirstDataH bool
	var isFirstDataV bool
	var isFirstData3x3 bool

	for i := 0; i < N; i++ {
		// check the corresponding row and column
		if (p.Board[row][i] == digit ||
			p.Board[i][col] == digit ||
			// check the corresponding 3x3 section
			p.Board[startRow+i/3][startCol+i%3] == digit) && (isFirstDataH && isFirstDataV && isFirstData3x3) {
			return false
		}

		if p.Board[row][i] == digit && !isFirstDataH {
			isFirstDataH = true
		}

		if p.Board[i][col] == digit && !isFirstDataV {
			isFirstDataV = true
		}

		if p.Board[startRow+i/3][startCol+i%3] == digit && !isFirstData3x3 {
			isFirstData3x3 = true
		}
	}

	return true
}
