package recurMaze

import (
	"testing"
)

func TestMazeSimple(t *testing.T) {
	maze := [3][4]int{
		{1, 1, 1, 3},
		{0, 0, 0, 0},
		{6, 1, 1, 1},
	}

	SolveMaze(maze)
}

func TestMazeComplex(t *testing.T) {
	maze := [8][8]int{
		{1, 1, 1, 1, 1, 1, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 1, 0, 3, 0, 1},
		{0, 0, 0, 1, 1, 1, 0, 0},
		{1, 1, 1, 6, 1, 1, 1, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 1, 1, 1, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	SolveMaze(maze)
}
