package graphs

import "math"

/*
Shortest Path for DAG with weights
 - Topological Sort : This gives you from the source to which order they can occur when you pop the nodes
 - This only works when graph is not cyclic and has a positive weights

 - When u use topo sort it just stays greedy, if gives u the order where start nodes occur before other nodes
 - This enables us to use this algo and then calculate node weights from there
*/

func ShortedPathTopoSort(n int, adj map[int][]int) []int {
	stack := TopoSort(n, adj)
	dist := make([]int, n)
	for i := range n {
		dist[i] = math.MaxInt
	}
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		for _, node := range adj[top] { // when there is a weight then instead of just +1 we need to consider weight of the edge
			if dist[top]+1 < dist[node] {
				dist[node] = dist[node] + 1
			}
		}
	}

	return dist
}
