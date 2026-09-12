package graphs

/*
Detect a cycle in a directed graph
- This is a but confusing incase of directed,
- In case of undirected we if a node diverges and has many paths visiting any path doesn't matter
- But incase of directed it matters, we need to remember this
- A visited node should appear inPath of current traversal if not no issue
*/

func DetectCycle(curr int, adj map[int][]int, vis []bool, inPath []bool) bool {
	vis[curr] = true
	inPath[curr] = true

	if nodes, ok := adj[curr]; ok {
		for _, node := range nodes {
			if !vis[node] {
				if DetectCycle(node, adj, vis, inPath) {
					return true
				}
			} else if inPath[node] {
				return true
			}
		}
	}
	inPath[curr] = false

	return false
}

func DirectedCycle(n int, adj map[int][]int) bool {

	vis := make([]bool, n)
	inPath := make([]bool, n)

	for i := 0; i < n; i++ {
		if !vis[i] {
			if DetectCycle(i, adj, vis, inPath) {
				return true
			}
		}

	}

	return false
}
