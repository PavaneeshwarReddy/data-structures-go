package dynamicprogramming

/*
Unique Paths in a Grid
- Base Condition: If we reach to end cell then we can return 1 or else if it crosses then we can return 0
- Explore Paths: Given visit down and right all
- Cache: key can be either we convert it to a string or maintain a new type called Pair and use that as a key, but for now we use string as key

*/

import "strconv"

func recursion006(sr int, sc int, m int, n int, cache *map[string]int) int {
	if sr == m-1 && sc == n-1 {
		return 1
	}

	if sr >= m || sc >= n {
		return 0
	}

	key := strconv.Itoa(sr) + "a" + strconv.Itoa(sc)

	if val, ok := (*cache)[key]; ok {
		return val
	}

	down := recursion006(sr+1, sc, m, n, cache)
	right := recursion006(sr, sc+1, m, n, cache)

	(*cache)[key] = down + right
	return down + right
}

func uniquePaths(m int, n int) int {
	cache := make(map[string]int)
	return recursion006(0, 0, m, n, &cache)
}
