package graphs

/*
Eventual Safe Nodes
- Anyone who connected to the cycle are never a safe nodes
- Terminal node , a node that doesn't have a outgoing edges
- Safe node, a node which somehow ends at terminal node
*/

func detectCycle020(v int, graph [][]int, vis []bool, inPath []bool, safe *[]bool) bool {
	inPath[v] = true
	vis[v] = true

	isSafe := true
	for _, node := range graph[v] {
		if !vis[node] { // this is if not visited
			isSafe = isSafe && detectCycle020(node, graph, vis, inPath, safe)
		} else if inPath[node] { // if visited but in the path
			isSafe = false
		} else if !(*safe)[node] { // if visited but not in the path
			isSafe = false
		}
	}

	(*safe)[v] = isSafe
	inPath[v] = false

	return isSafe
}

func eventualSafeNodes(graph [][]int) []int {
	safe := make([]bool, len(graph))
	inPath := make([]bool, len(graph))
	vis := make([]bool, len(graph))

	for idx, _ := range graph {
		if !vis[idx] {
			detectCycle020(idx, graph, vis, inPath, &safe)
		}
	}

	result := []int{}

	for idx, val := range safe {
		if val {
			result = append(result, idx)
		}
	}

	return result
}
