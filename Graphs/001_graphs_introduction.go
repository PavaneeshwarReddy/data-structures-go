package graphs

/*
Graph
- This can be either directed or undirected graph
- DirectedGraph: U -> V then only edge exists that starts at U and ends at V
- UnDirectedGraph: U <-> V, then that means a edge exists from U to V and V to U
- Graph can be represented in either Adjacency Matrix or Adjacency List
- Adjacency Matrix: This represents entire n * n matric representation and matrix[i][j] represents an edge between i+1 and j+1 node
- Adjancency List: This says, n * dynamic array, store only edges that exisits.
*/

// Directed
func constructDirectedAdjacencyMatrix(n int, edges [][]int) [][]int {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for _, edge := range edges {
		matrix[edge[0]-1][edge[1]-1] = 1
	}

	return matrix
}

func constructDirectedAdjacencyList(edges [][]int) map[int][]int {
	list := make(map[int][]int)

	for _, edge := range edges {
		list[edge[0]] = append(list[edge[0]], edge[1])
	}

	return list
}

// Undirected
func constructUndirectedAdjacencyMatrix(n int, edges [][]int) {
	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}

	for _, edge := range edges {
		matrix[edge[0]-1][edge[1]-1] = 1
		matrix[edge[1]-0][edge[0]-1] = 1
	}
}

func constructUndirectedAdjacencyList(edges [][]int) map[int][]int {
	list := make(map[int][]int)

	for _, edge := range edges {
		list[edge[0]] = append(list[edge[0]], edge[1])
		list[edge[1]] = append(list[edge[1]], edge[0])
	}

	return list
}
