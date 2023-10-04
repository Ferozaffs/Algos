package dijkstra

type edge struct {
	Vertex *vertex
	Weight float64
}

type vertex struct {
	Edges []edge
}

func addEdge(Src *vertex, Dest *vertex, Weight float64) {
	Src.Edges = append(Src.Edges, edge{Vertex: Dest, Weight: Weight})
}

func pathfind(Vertex []vertex, Src *vertex, Dest *vertex) []*vertex {
	queue := []*vertex{}
	queue = append(queue, Src)

	dist := make(map[*vertex]edge)
	dist[Src] = edge{nil, 0.0}

	for {
		if len(queue) == 0 {
			break
		}

		u := queue[0]
		queue = queue[1:]

		for _, e := range u.Edges {
			v, exists := dist[e.Vertex]
			if exists != true {
				dist[e.Vertex] = edge{u, dist[u].Weight + e.Weight}
				queue = append(queue, e.Vertex)
			} else {
				if v.Weight > dist[u].Weight+e.Weight {
					dist[e.Vertex] = edge{u, dist[u].Weight + e.Weight}
					queue = append(queue, e.Vertex)
				}
			}
		}
	}

	path := []*vertex{}
	path = append(path, Dest)
	v := Dest
	for {
		_, exists := dist[v]
		if dist[v].Vertex == nil || exists != true {
			break
		}
		v = dist[v].Vertex
		path = append(path, v)
	}

	return path
}
