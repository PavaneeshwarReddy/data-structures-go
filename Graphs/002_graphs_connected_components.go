package graphs

/*
Connected Components
- Total individual linked components in the graph has to calculated
- Maintain a visited array and visit each value where it has nodes, if visited leave or else increment

Input: V=4, edges=[[0,1],[1,2]]
Output: 2
Explanation: Vertices {0,1,2} forms the first component and vertex 3 forms the second component.
Input:V = 7, edges = [[0, 1], [1, 2], [2, 3], [4, 5]]
Output: 3
Explanation: The edges [0, 1], [1, 2], [2, 3] form a connected component with vertices {0, 1, 2, 3}
The edge [4, 5] forms another connected component with vertices {4, 5}.
Therefore, the graph has 3 connected components: {0, 1, 2, 3}, {4, 5}, and the isolated vertices {6}.

NOTE:
Same can be applied for both directed and undirected, in some cases directed can have more components then undirected like this

A -> B <- C ( DIRECTED )
This means  2 different components

A - B - C ( UNDIRECTED )
This means 1 component
*/

func dfs(node int, adj map[int][]int, vis []bool) {
	if vis[node] {
		return
	}

	vis[node] = true

	if nodes, ok := adj[node]; ok {
		for _, newNode := range nodes {
			dfs(newNode, adj, vis)
		}
	}
}

func FindConnectedComponents(n int, edges [][]int) int {
	adj := constructDirectedAdjacencyList(edges)
	vis := make([]bool, n)
	result := 0
	for node := range adj {
		if !vis[node] {
			dfs(node, adj, vis)
			result++
		}
	}
	return result
}
