package SudokuSolve

func IsValid(n int, row int, column int, sudoku *[][]int) bool {
	for _, c := range (*sudoku)[row] {
		if c == n {
			return false
		}
	}

	for _, row := range *sudoku {
		if row[column] == n {
			return false
		}
	}

	rowStart := int(row/3) * 3
	columnStart := int(column/3) * 3

	for i := rowStart; i < rowStart+3; i++ {
		for j := columnStart; j < columnStart+3; j++ {
			if (*sudoku)[i][j] == n {
				return false
			}
		}
	}

	return true
}

func Solve(sudoku *[][]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if (*sudoku)[i][j] == 0 {
				for n := 1; n < 10; n++ {
					if IsValid(n, i, j, sudoku) {
						(*sudoku)[i][j] = n
						if Solve(sudoku) == true {
							return true
						}
						(*sudoku)[i][j] = 0
					}
				}
				return false
			}
		}
	}

	return true
}
