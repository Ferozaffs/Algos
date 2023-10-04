package dijkstra

import (
	"slices"
	"testing"
)

func TestDijkstraSimple(t *testing.T) {
	vertices := make([]vertex, 4)

	addEdge(&vertices[0], &vertices[1], 1.5)
	addEdge(&vertices[0], &vertices[2], 2.0)
	addEdge(&vertices[1], &vertices[3], 0.5)
	addEdge(&vertices[2], &vertices[3], 1.0)

	path := pathfind(vertices, &vertices[0], &vertices[3])
	if len(path) != 3 {
		t.Fatalf("Path is not correct length: %d expected 3", len(path))
	}

	correctPath := []*vertex{&vertices[0], &vertices[1], &vertices[3]}
	for _, p := range path {
		if slices.Contains(correctPath, p) == false {
			t.Fatalf("Path is invalid")
		}
	}
}

func TestDijkstraComplex(t *testing.T) {
	vertices := make([]vertex, 10)

	addEdge(&vertices[0], &vertices[1], 1.5)
	addEdge(&vertices[0], &vertices[2], 2.0)
	addEdge(&vertices[1], &vertices[3], 0.5)
	addEdge(&vertices[1], &vertices[4], 1.0)
	addEdge(&vertices[2], &vertices[5], 2.0)
	addEdge(&vertices[2], &vertices[6], 1.5)
	addEdge(&vertices[3], &vertices[7], 0.5)
	addEdge(&vertices[4], &vertices[8], 2.0)
	addEdge(&vertices[5], &vertices[7], 0.5)
	addEdge(&vertices[6], &vertices[8], 1.5)

	path := pathfind(vertices, &vertices[0], &vertices[8])
	if len(path) != 4 {
		t.Fatalf("Path is not correct length: %d expected 4", len(path))
	}

	correctPath := []*vertex{&vertices[0], &vertices[1], &vertices[4], &vertices[8]}
	for _, p := range path {
		if slices.Contains(correctPath, p) == false {
			t.Fatalf("Path is invalid")
		}
	}
}
