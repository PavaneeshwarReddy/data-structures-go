package graphs

/*
Breadth First Search
- This works as a queue traversing entire level at a single time
*/

func BFS(start int, adj map[int][]int, res []int, vis []bool) {
	if vis[start] {
		return
	}
	stack := []int{}

	stack = append(stack, start)

	for len(stack) > 0 {
		newStack := []int{}
		for _, node := range stack {
			if !vis[node] {
				res = append(res, node)
				if connectedNodes, ok := adj[node]; ok {
					// you can also check whether this is visited and add it to the new stack
					newStack = append(newStack, connectedNodes...)
				}
				vis[node] = true
			}
		}
		stack = newStack
	}
}
