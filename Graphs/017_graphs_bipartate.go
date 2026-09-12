package graphs

/*
Bipartate Graph
- Where in which if we divide graph into two sets then every node in set A should have an edge with other node in set B
- Basically are we able to color the graph in 2 colors, so that adjacent node always contain different colors
*/

func dfs017(v int, color int, graph [][]int, colors []int) bool {

	colors[v] = color

	for _, node := range graph[v] {
		if colors[node] != 0 && colors[node] == color {
			return false
		}
		if colors[node] == 0 {
			if color == -1 {
				if !dfs017(node, 1, graph, colors) {
					return false
				}
			} else {
				if !dfs017(node, -1, graph, colors) {
					return false
				}
			}
		}
	}

	return true

}

func isBipartite(graph [][]int) bool {
	colors := make([]int, len(graph))
	result := true
	for idx, _ := range graph {
		if colors[idx] == 0 {
			result = result && dfs017(idx, -1, graph, colors)
		}
	}
	return result
}
