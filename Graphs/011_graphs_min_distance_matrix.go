package graphs

/*
01 Matrix
- We have to find the min distance of 0 cell for each cell
- If you see if we go to every cell and do a BFS it will take more time,
- This is an inverse pattern in graphs, instead of going each cell, we do BFS from every 0 cell and increment their adjancent cells value

*/

func updateMatrix(mat [][]int) [][]int {
	stack := [][2]int{}
	m := len(mat)
	n := len(mat[0])

	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)

		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				stack = append(stack, [2]int{i, j})
			} else {
				res[i][j] = -1
			}
		}
	}

	dirs := [][2]int{
		{0, 1},
		{1, 0},
		{-1, 0},
		{0, -1},
	}

	for len(stack) > 0 {
		newNodes := [][2]int{}

		for _, node := range stack {
			for _, dir := range dirs {
				i := dir[0] + node[0]
				j := dir[1] + node[1]
				if i < m && i >= 0 && j < n && j >= 0 && res[i][j] == -1 {
					res[i][j] = res[node[0]][node[1]] + 1
					newNodes = append(newNodes, [2]int{i, j})
				}
			}
		}

		stack = newNodes
	}

	return res
}
