package graphs

/*
Course Schedule - 2
- We need to print the valid path
- DAG Cyclic Check + Topo Sort
*/

func dfs019(v int, adj map[int][]int, vis []bool, st *[]int) {
	vis[v] = true
	if nodes, ok := adj[v]; ok {
		for _, node := range nodes {
			if !vis[node] {
				dfs019(node, adj, vis, st)
			}
		}
	}
	*st = append(*st, v)
}

func findOrder(numCourses int, prerequisites [][]int) []int {
	adj := constructDirectedAdjacencyList(prerequisites)
	vis := make([]bool, numCourses)
	inPath := make([]bool, numCourses)

	for i := 0; i < numCourses; i++ {
		if !vis[i] {
			if DetectCycle(i, adj, vis, inPath) {
				return []int{}
			}
		}

	}

	st := []int{}
	newVis := make([]bool, numCourses)

	for i := 0; i < numCourses; i++ {
		if !newVis[i] {
			dfs019(i, adj, newVis, &st)
		}
	}

	return st
}
