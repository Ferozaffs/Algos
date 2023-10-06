package DFS

type Coordinate struct {
	x int
	y int
}

func CheckCoordinate(c Coordinate, nodes [][]int, visited [][]bool, path *[]Coordinate) bool {
	directions := [4]Coordinate{
		{x: 1, y: 0},
		{x: 0, y: 1},
		{x: -1, y: 0},
		{x: 0, y: -1},
	}

	if c.x < 0 || c.y < 0 || c.x >= len(nodes) || c.y >= len(nodes[0]) {
		return false
	} else if nodes[c.x][c.y] == 1 {
		return false
	} else if visited[c.x][c.y] == true {
		return false
	} else if nodes[c.x][c.y] == 6 {
		return true
	}

	visited[c.x][c.y] = true

	for _, d := range directions {
		newCoord := c
		newCoord.x += d.x
		newCoord.y += d.y

		if CheckCoordinate(newCoord, nodes, visited, path) {
			*path = append(*path, newCoord)
			return true
		}
	}

	return false
}

func Solve(nodes [][]int) []Coordinate {
	path := []Coordinate{}
	visited := [][]bool{}

	start := Coordinate{}

	for x, row := range nodes {
		visited = append(visited, []bool{})
		for y, val := range row {
			if val == 3 {
				start.x = x
				start.y = y
			}
			visited[x] = append(visited[x], false)
		}
	}

	CheckCoordinate(start, nodes, visited, &path)
	path = append(path, start)

	return path
}
