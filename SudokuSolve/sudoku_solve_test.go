package SudokuSolve

import "testing"

func TestEasySudoku(t *testing.T) {
	sudoku := [][]int{
		{9, 7, 0, 4, 2, 1, 0, 0, 0},
		{1, 0, 2, 0, 0, 6, 8, 7, 9},
		{6, 5, 3, 0, 8, 9, 0, 0, 1},
		{0, 0, 0, 3, 1, 2, 4, 6, 5},
		{5, 0, 0, 9, 7, 8, 1, 0, 0},
		{2, 3, 0, 0, 4, 0, 7, 9, 0},
		{3, 0, 0, 2, 0, 0, 9, 8, 0},
		{0, 2, 6, 8, 0, 7, 5, 0, 3},
		{0, 8, 9, 1, 5, 3, 0, 2, 0},
	}
	solution := [][]int{
		{9, 7, 8, 4, 2, 1, 3, 5, 6},
		{1, 4, 2, 5, 3, 6, 8, 7, 9},
		{6, 5, 3, 7, 8, 9, 2, 4, 1},
		{8, 9, 7, 3, 1, 2, 4, 6, 5},
		{5, 6, 4, 9, 7, 8, 1, 3, 2},
		{2, 3, 1, 6, 4, 5, 7, 9, 8},
		{3, 1, 5, 2, 6, 4, 9, 8, 7},
		{4, 2, 6, 8, 9, 7, 5, 1, 3},
		{7, 8, 9, 1, 5, 3, 6, 2, 4},
	}

	result := Solve(&sudoku)

	if result == false {
		t.Fatalf("Sudoku did not solve!")
	}

	for i, row := range solution {
		for j := range row {
			if solution[i][j] != sudoku[i][j] {
				t.Fatalf("Sudoku mistmatch! Expected %d at %d %d, got %d", solution[i][j], i, j, sudoku[i][j])
			}
		}
	}
}

func TestMultiSolutionSudoku(t *testing.T) {

}

func TestInvalidSudoku(t *testing.T) {

}
