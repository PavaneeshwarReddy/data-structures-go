package graphs

import "math"

/*
Find the shortest path from single V to all the other nodes
- We can do this by using simple BFS
- Consider vertex as 0

- Standard approach is like picking individual nodes and processing them in shortest path algos
*/

func ShortestPathBFS(n int, adj map[int][]int) []int {
	dist := []int{}
	for range n {
		dist = append(dist, math.MaxInt)
	}

	queue := []int{0}
	dist[0] = 0

	for len(queue) > 0 {
		curr := queue[0]
		queue := queue[1:]

		for _, next := range adj[curr] {
			if dist[next] > dist[curr]+1 {
				dist[next] = dist[curr] + 1
				queue = append(queue, next)
			}
		}
	}

	return dist

}
