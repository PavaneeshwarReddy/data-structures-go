package graphs

/*
Shortest Path in a Binary Matrix
- Start at 0,0 and end at m-1, n-1
*/

func shortestPathBinaryMatrix(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	queue := [][3]int{}
	vis := make([][]bool, m)
	for i := range m {
		vis[i] = make([]bool, n)
	}
	dirs := [][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	if grid[0][0] == 1 {
		return -1
	}
	queue = append(queue, [3]int{0, 0, 0})

	for len(queue) > 0 {
		newNodes := [][3]int{}
		for _, node := range queue {
			if node[0] == m-1 && node[1] == n-1 {
				return node[2] + 1
			}
			for _, dir := range dirs {
				i := dir[0] + node[0]
				j := dir[1] + node[1]

				if i >= 0 && i < m && j >= 0 && j < n && grid[i][j] == 0 && !vis[i][j] {
					vis[i][j] = true
					newNodes = append(newNodes, [3]int{i, j, node[2] + 1})
				}
			}
		}

		queue = newNodes
	}

	return -1
}
