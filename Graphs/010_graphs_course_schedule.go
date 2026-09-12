package graphs

/*
Pre-requisite task
- If we want to perform some task then we need complete some dependent tasks
- n = 5, [[1,4],[2,4],[3,1],[3,2]]
- 4 -> 1, 2 -> 4, 1 -> 3 , 2 -> 3, these are the edges
- Our goal is simple, if somehow the cycle exisits that means a deadlock situation where he can't complete tasks
*/

func detectCycle(curr int, adj map[int][]int, vis []bool, inPath []bool) bool {
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

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := constructDirectedAdjacencyList(prerequisites)
	vis := make([]bool, numCourses)
	inPath := make([]bool, numCourses)

	for i := 0; i < numCourses; i++ {
		if !vis[i] {
			if DetectCycle(i, adj, vis, inPath) {
				return false
			}
		}

	}

	return true
}
