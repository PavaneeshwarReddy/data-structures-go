package graphs

/*
Detect a Cycle in Undirected Graph

- If we consider there is always a 2 way edge between every node
- There is a cycle if and only ig the node which we are going to visit is already visited and not same as prev node

*/

func dfs008(curr int, prev int, adj map[int][]int, vis []bool) bool {

	vis[curr] = true

	if nodes, ok := adj[curr]; ok {
		for _, node := range nodes {
			if node == prev {
				continue
			}

			if vis[node] {
				return true
			}

			if dfs008(node, curr, adj, vis) {
				return true
			}
		}
	}

	return false
}
