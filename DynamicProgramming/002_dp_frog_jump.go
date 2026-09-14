package dynamicprogramming

import "math"

/*
Frog Jump
- Base Conditions: When you reach n or greater than n return
- Paths To Explore: Either we take 1 or 2, if we take 1 calculate the enery diff for the prev and current, we need to maintain a boundary because we can evaluate effort if it crosses array
- Cache: It will be the idx
*/

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func recursion002(idx int, n int, effort []int, cache *map[int]int) int {
	if idx >= n-1 {
		return 0
	}

	if val, ok := (*cache)[idx]; ok {
		return val
	}

	one := abs(effort[idx]-effort[idx+1]) + recursion002(idx+1, n, effort, cache)

	two := math.MaxInt

	if idx+2 < n {
		two = abs(effort[idx]-effort[idx+2]) + recursion002(idx+2, n, effort, cache)
	}

	(*cache)[idx] = min(one, two)

	return min(one, two)

}

func FrogJump(n int, effort []int) int {
	cache := make(map[int]int)
	return recursion002(0, n, effort, &cache)
}
