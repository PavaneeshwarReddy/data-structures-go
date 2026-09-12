package graphs

/*
Number of Enclaves
- We just need to traverse entire boundary
- We just need to call dfs
- Left over land is the answer
*/

func dfs013(sr int, sc int, grid [][]int, vis [][]bool) {
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

		if i < m && i >= 0 && j < n && j >= 0 && grid[i][j] == 1 && !vis[i][j] {
			dfs013(i, j, grid, vis)
		}
	}
}

func numEnclaves(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		if !vis[i][0] && grid[i][0] == 1 {
			dfs013(i, 0, grid, vis)
		}
	}

	for i := 0; i < m; i++ {
		if !vis[i][n-1] && grid[i][n-1] == 1 {
			dfs013(i, n-1, grid, vis)
		}
	}

	for j := 0; j < n; j++ {
		if !vis[0][j] && grid[0][j] == 1 {
			dfs013(0, j, grid, vis)
		}
	}

	for j := 0; j < n; j++ {
		if !vis[m-1][j] && grid[m-1][j] == 1 {
			dfs013(m-1, j, grid, vis)
		}
	}
	res := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 && !vis[i][j] {
				res++
			}
		}
	}

	return res

}
