package graphs

/*
Find total Pronvinces
- This is similar to finding connected components
*/

func dfs005(v int, adj [][]int, vis []bool) {
	if vis[v] {
		return
	}

	vis[v] = true

	for idx, val := range adj[v] {
		if val == 1 {
			dfs005(idx, adj, vis)
		}
	}
}

func findCircleNum(isConnected [][]int) int {
	n := len(isConnected)
	vis := make([]bool, n)
	result := 0

	for i := 0; i < n; i++ {
		if !vis[i] {
			dfs005(i, isConnected, vis)
			result++
		}
	}

	return result

}
