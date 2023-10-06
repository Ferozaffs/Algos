package Dijkstra

import (
	"slices"
	"testing"
)

func TestDijkstraSimple(t *testing.T) {
	vertices := make([]Vertex, 4)

	AddEdge(&vertices[0], &vertices[1], 15)
	AddEdge(&vertices[0], &vertices[2], 20)
	AddEdge(&vertices[1], &vertices[3], 5)
	AddEdge(&vertices[2], &vertices[3], 10)

	path := pathfind(vertices, &vertices[0], &vertices[3])
	if len(path) != 3 {
		t.Fatalf("Path is not correct length: %d expected 3", len(path))
	}

	correctPath := []*Vertex{&vertices[0], &vertices[1], &vertices[3]}
	for _, p := range path {
		if slices.Contains(correctPath, p) == false {
			t.Fatalf("Path is invalid")
		}
	}
}

func TestDijkstraComplex(t *testing.T) {
	vertices := make([]Vertex, 10)

	AddEdge(&vertices[0], &vertices[1], 15)
	AddEdge(&vertices[0], &vertices[2], 20)
	AddEdge(&vertices[1], &vertices[3], 5)
	AddEdge(&vertices[1], &vertices[4], 10)
	AddEdge(&vertices[2], &vertices[5], 20)
	AddEdge(&vertices[2], &vertices[6], 15)
	AddEdge(&vertices[3], &vertices[7], 5)
	AddEdge(&vertices[4], &vertices[8], 20)
	AddEdge(&vertices[5], &vertices[7], 5)
	AddEdge(&vertices[6], &vertices[8], 15)

	path := pathfind(vertices, &vertices[0], &vertices[8])
	if len(path) != 4 {
		t.Fatalf("Path is not correct length: %d expected 4", len(path))
	}

	correctPath := []*Vertex{&vertices[0], &vertices[1], &vertices[4], &vertices[8]}
	for _, p := range path {
		if slices.Contains(correctPath, p) == false {
			t.Fatalf("Path is invalid")
		}
	}
}
