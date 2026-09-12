package graphs

/*
Rotten Oranges
- Add all rotten oranges to the stack
- Then process all adjancent 4 positions for a fresh orange
- If found one rotten it and add it to the stack
- At last minutes - 1 because we don't really count the first orange as it is already rottened

*/

func orangesRotting(grid [][]int) int {

	rottenOranges := [][2]int{}
	m := len(grid)
	n := len(grid[0])

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 2 {
				rottenOranges = append(rottenOranges, [2]int{i, j})
			}
		}
	}

	minutes := 0
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}

	for len(rottenOranges) > 0 {
		newRotten := [][2]int{}
		for _, org := range rottenOranges {
			if vis[org[0]][org[1]] {
				continue
			}
			vis[org[0]][org[1]] = true

			if org[0]+1 < m && grid[org[0]+1][org[1]] == 1 && !vis[org[0]+1][org[1]] {
				newRotten = append(newRotten, [2]int{org[0] + 1, org[1]})
				grid[org[0]+1][org[1]] = 2
			}
			if org[0]-1 >= 0 && grid[org[0]-1][org[1]] == 1 && !vis[org[0]-1][org[1]] {
				newRotten = append(newRotten, [2]int{org[0] - 1, org[1]})
				grid[org[0]-1][org[1]] = 2
			}
			if org[1]+1 < n && grid[org[0]][org[1]+1] == 1 && !vis[org[0]][org[1]+1] {
				newRotten = append(newRotten, [2]int{org[0], org[1] + 1})
				grid[org[0]][org[1]+1] = 2
			}
			if org[1]-1 >= 0 && grid[org[0]][org[1]-1] == 1 && !vis[org[0]][org[1]-1] {
				newRotten = append(newRotten, [2]int{org[0], org[1] - 1})
				grid[org[0]][org[1]-1] = 2
			}
		}
		rottenOranges = newRotten
		minutes++
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	if minutes == 0 {
		return 0
	}
	return minutes - 1

}
