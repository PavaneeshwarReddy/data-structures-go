package graphs

/*
Flood Fill
- Simple scenarion where we have to use BFS inorder to get the related cells filled
*/

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	stack := [][2]int{}
	startColor := image[sr][sc]

	stack = append(stack, [2]int{sr, sc})
	m := len(image)
	n := len(image[0])
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}

	for len(stack) > 0 {
		newNodes := [][2]int{}

		for _, node := range stack {
			i := node[0]
			j := node[1]

			image[i][j] = color

			if i+1 < m && image[i+1][j] == startColor && !vis[i+1][j] {
				image[i+1][j] = color
				vis[i+1][j] = true
				newNodes = append(newNodes, [2]int{i + 1, j})
			}

			if i-1 >= 0 && image[i-1][j] == startColor && !vis[i-1][j] {
				image[i-1][j] = color
				vis[i-1][j] = true
				newNodes = append(newNodes, [2]int{i - 1, j})
			}

			if j+1 < n && image[i][j+1] == startColor && !vis[i][j+1] {
				image[i][j+1] = color
				vis[i][j+1] = true
				newNodes = append(newNodes, [2]int{i, j + 1})
			}

			if j-1 >= 0 && image[i][j-1] == startColor && !vis[i][j-1] {
				image[i][j-1] = color
				vis[i][j-1] = true
				newNodes = append(newNodes, [2]int{i, j - 1})
			}

		}

		stack = newNodes
	}

	return image

}
