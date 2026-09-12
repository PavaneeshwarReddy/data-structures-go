package graphs

/*
Sourrounded Regions
- Concept is if we DFS from a cell and we move adjacent wherever there are 0s and we found some 0 at the border then we should not populate the entire 0s path to X
- Or else we can populate O with X

*/

func dfs012(sr int, sc int, board [][]byte, vis [][]bool, updates *[][2]int) bool {

	m := len(board)
	n := len(board[0])

	sourrounded := true

	if (sr == m-1 || sr == 0) || (sc == n-1 || sc == 0) {
		sourrounded = false
	}

	vis[sr][sc] = true

	dirs := [][2]int{
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
	}

	*updates = append(*updates, [2]int{sr, sc})

	for _, dir := range dirs {
		i := sr + dir[0]
		j := sc + dir[1]
		if i < m && i >= 0 && j < n && j >= 0 && board[i][j] == 'O' && !vis[i][j] {
			if !dfs012(i, j, board, vis, updates) {
				sourrounded = false
			}
		}
	}

	return sourrounded
}

func solve(board [][]byte) {
	m := len(board)
	n := len(board[0])

	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			updates := [][2]int{}
			if board[i][j] == 'O' && !vis[i][j] {
				if dfs012(i, j, board, vis, &updates) {
					for _, cell := range updates {
						board[cell[0]][cell[1]] = 'X'
					}
				}
			}
		}
	}
}
