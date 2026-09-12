package graphs

/*
Number of islands
- Same like no of connected components but in 4 directions
*/

func dfs014(sr int, sc int, grid [][]byte, vis [][]bool) {
	if vis[sr][sc] {
		return
	}

	m := len(grid)
	n := len(grid[0])

	vis[sr][sc] = true

	dirs := [][2]int{
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
	}

	for _, dir := range dirs {
		i := dir[0] + sr
		j := dir[1] + sc

		if i < m && i >= 0 && j < n && j >= 0 && grid[i][j] == '1' && !vis[i][j] {
			dfs014(i, j, grid, vis)
		}
	}
}

func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])

	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}

	res := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !vis[i][j] && grid[i][j] == '1' {
				dfs014(i, j, grid, vis)
				res++
			}
		}
	}

	return res
}
