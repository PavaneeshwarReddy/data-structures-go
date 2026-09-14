package dynamicprogramming

import "math"

/*
Frog Jump With K steps
- Base Conditions: When frog reaches n-1 the last step return
- Paths to Explore: For every idx we need to check whether idx + j < n and then we need to pass to the recursion
- Cache: It will be a idx
*/

func abs003(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func recursion003(idx int, k int, n int, effort []int, cache *map[int]int) int {
	if idx == n-1 {
		return 0
	}

	if val, ok := (*cache)[idx]; ok {
		return val
	}

	res := math.MaxInt
	for i := 1; i <= k; i++ {
		if idx+i < n {
			res = min(res, abs003(effort[idx]-effort[idx+i])+recursion003(idx+i, k, n, effort, cache))
		}
	}

	(*cache)[idx] = res
	return res
}

func FrogJumpK(k int, n int, effort []int) int {
	cache := make(map[int]int)
	return recursion003(0, k, n, effort, &cache)
}
