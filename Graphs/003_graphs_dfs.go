package graphs

/*
Depth First Search
- If works like, when you supply a start node, it starts traversing from there and goes in depth until it reaches edge
- If you want to traverse entire graph then you need to loop on all nodes by placing starting node as required
*/

func DFS(start int, adj map[int][]int, vis []bool, res []int) {
	if vis[start] {
		return
	}

	vis[start] = true
	res = append(res, start)

	if nodes, ok := adj[start]; ok {
		for _, newNode := range nodes {
			DFS(newNode, adj, vis, res)
		}
	}
}
