package dynamicprogramming

/*
Minimum Falling Path Sum
- Base Condition: If we cross the boundaries we can return MaxInt, if we reached end then we have to return cell value
- Explore Paths: Given visit down and right all
- Cache: key can be either we convert it to a string or maintain a new type called Pair and use that as a key, but for now we use string as key
*/

import (
	"math"
	"strconv"
)

func recursion008(sr int, sc int, m int, n int, grid [][]int, cache *map[string]int) int {
	if sr >= m || sc >= n {
		return math.MaxInt
	}

	key := strconv.Itoa(sr) + "a" + strconv.Itoa(sc)
	if val, ok := (*cache)[key]; ok {
		return val
	}

	if sr == m-1 && sc == n-1 {
		return grid[sr][sc]
	}

	down := recursion008(sr+1, sc, m, n, grid, cache)
	right := recursion008(sr, sc+1, m, n, grid, cache)

	if down != math.MaxInt {
		down += grid[sr][sc]
	}

	if right != math.MaxInt {
		right += grid[sr][sc]
	}

	(*cache)[key] = min(down, right)

	return min(down, right)
}

func minPathSum(grid [][]int) int {
	cache := make(map[string]int)
	return recursion008(0, 0, len(grid), len(grid[0]), grid, &cache)
}
