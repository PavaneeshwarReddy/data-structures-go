package dynamicprogramming

/*
Triangle Path Sum
- Base Condition: If we cross the boundaries we can return 0, else return cell value
- Explore Paths: Given visit down and right all
- Cache: key can be either we convert it to a string or maintain a new type called Pair and use that as a key, but for now we use string as key
*/

import "strconv"

func recursion009(sr int, sc int, tri [][]int, cache *map[string]int) int {
	if sc >= len(tri) || sr >= len(tri[sr]) {
		return 0
	}

	key := strconv.Itoa(sr) + "a" + strconv.Itoa(sc)
	if val, ok := (*cache)[key]; ok {
		return val
	}

	if sr == len(tri)-1 {
		return tri[sr][sc]
	}

	adjLeft := recursion009(sr+1, sc, tri, cache) + tri[sr][sc]
	adjRight := recursion009(sr+1, sc+1, tri, cache) + tri[sr][sc]

	(*cache)[key] = min(adjLeft, adjRight)

	return min(adjLeft, adjRight)
}

func minimumTotal(triangle [][]int) int {
	cache := make(map[string]int)
	return recursion009(0, 0, triangle, &cache)
}
