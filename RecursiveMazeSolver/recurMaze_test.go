package recurMaze

import (
	"testing"
)

func TestMazeSimple(t *testing.T) {
	maze := [][]int{
		{1, 1, 1, 3},
		{0, 0, 0, 0},
		{6, 1, 1, 1},
	}

	path := SolveMaze(maze)

	solution := [6]Coordinate{
		{x: 2, y: 0},
		{x: 1, y: 0},
		{x: 1, y: 1},
		{x: 1, y: 2},
		{x: 1, y: 3},
		{x: 0, y: 3},
	}

	for i, c := range solution {
		if c.x != path[i].x || c.y != path[i].y {
			t.Fatalf("Path is invalid, expected %d,%d got %d,%d", c.x, c.y, path[i].x, path[i].y)
		}
	}
}

func TestMazeComplex(t *testing.T) {
	maze := [][]int{
		{1, 1, 1, 1, 1, 1, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 1},
		{0, 1, 0, 1, 0, 3, 0, 1},
		{0, 0, 0, 1, 1, 1, 0, 0},
		{1, 1, 1, 6, 1, 1, 1, 0},
		{0, 0, 0, 0, 1, 0, 0, 0},
		{0, 1, 1, 1, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	path := SolveMaze(maze)

	solution := []Coordinate{
		{x: 4, y: 3},
		{x: 5, y: 3},
		{x: 5, y: 2},
		{x: 5, y: 1},
		{x: 5, y: 0},
		{x: 6, y: 0},
		{x: 7, y: 0},
		{x: 7, y: 1},
		{x: 7, y: 2},
		{x: 7, y: 3},
		{x: 7, y: 4},
		{x: 7, y: 5},
		{x: 7, y: 6},
		{x: 7, y: 7},
		{x: 6, y: 7},
		{x: 5, y: 7},
		{x: 4, y: 7},
		{x: 3, y: 7},
		{x: 3, y: 6},
		{x: 2, y: 6},
		{x: 2, y: 5},
	}

	for i, c := range solution {
		if c.x != path[i].x || c.y != path[i].y {
			t.Fatalf("Path is invalid, expected %d,%d got %d,%d", c.x, c.y, path[i].x, path[i].y)
		}
	}
}
