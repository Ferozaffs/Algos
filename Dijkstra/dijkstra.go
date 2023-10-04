package dijkstra

type Edge struct {
	vertex *Vertex
	weight float64
}

type Vertex struct {
	edges []Edge
}

func AddEdge(src *Vertex, dest *Vertex, weight float64) {
	src.edges = append(src.edges, Edge{vertex: dest, weight: weight})
}

func pathfind(vertex []Vertex, src *Vertex, dest *Vertex) []*Vertex {
	queue := []*Vertex{}
	queue = append(queue, src)

	dist := make(map[*Vertex]Edge)
	dist[src] = Edge{nil, 0.0}

	for {
		if len(queue) == 0 {
			break
		}

		u := queue[0]
		queue = queue[1:]

		for _, e := range u.edges {
			v, exists := dist[e.vertex]
			if exists != true {
				dist[e.vertex] = Edge{u, dist[u].weight + e.weight}
				queue = append(queue, e.vertex)
			} else {
				if v.weight > dist[u].weight+e.weight {
					dist[e.vertex] = Edge{u, dist[u].weight + e.weight}
					queue = append(queue, e.vertex)
				}
			}
		}
	}

	path := []*Vertex{}
	path = append(path, dest)
	v := dest
	for {
		_, exists := dist[v]
		if dist[v].vertex == nil || exists != true {
			break
		}
		v = dist[v].vertex
		path = append(path, v)
	}

	return path
}
