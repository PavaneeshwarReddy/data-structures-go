package dynamicprogramming

/*
Unique Paths With Obstacles
- Base Condition: If we reach to end cell then we can return 1 or else if it crosses then we can return 0, if current cell has a obstacle then return 0
- Explore Paths: Given visit down and right all
- Cache: key can be either we convert it to a string or maintain a new type called Pair and use that as a key, but for now we use string as key

*/

import "strconv"

func recursion007(sr int, sc int, m int, n int, grid [][]int, cache *map[string]int) int {

	if sr >= m || sc >= n || grid[sr][sc] == 1 {
		return 0
	}
	if sr == m-1 && sc == n-1 {
		return 1
	}

	key := strconv.Itoa(sr) + "a" + strconv.Itoa(sc)

	if val, ok := (*cache)[key]; ok {
		return val
	}

	down := recursion007(sr+1, sc, m, n, grid, cache)
	right := recursion007(sr, sc+1, m, n, grid, cache)

	(*cache)[key] = down + right

	return down + right

}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	cache := make(map[string]int)
	return recursion007(0, 0, len(obstacleGrid), len(obstacleGrid[0]), obstacleGrid, &cache)
}
