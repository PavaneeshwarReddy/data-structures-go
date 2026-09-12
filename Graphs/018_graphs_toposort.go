package graphs

/*
TopoSort
- When there is no cycle exisits inside the graph
- There is a order that can be written in which
- a , b, c , d, e, f  if there exisits some order where a -> f in such a way a always appear before f
- This is only possible in DAG
*/

func dfs018(v int, adj map[int][]int, vis []bool, st []int) {
	vis[v] = true
	if nodes, ok := adj[v]; ok {
		for _, node := range nodes {
			if !vis[node] {
				dfs018(node, adj, vis, st)
			}
		}
	}
	st = append(st, v)
}

func TopoSort(n int, adj map[int][]int) []int {
	st := []int{}
	vis := make([]bool, n)
	for key := range adj {
		if !vis[key] {
			dfs018(key, adj, vis, st)
		}
	}
	i := 0
	j := n - 1

	// reverse
	for i <= j {
		st[i] = st[j]
		i++
		j--
	}
	return st
}
