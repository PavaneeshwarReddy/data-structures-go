package dynamicprogramming

/*
Climbing Stairs
- Base Condition: Return when you climb greater or than or equal to n
- Paths To Explore: You need to explore either 1 or 2
- Cache Key: As idx changes that mean if you take either 1 or 2 there are chances you may end up on same stair
*/

func recursion(idx int, n int, cache *map[int]int) int {

	if val, ok := (*cache)[idx]; ok {
		return val
	}

	if idx >= n {
		return 1
	}

	chooseOne := recursion(idx+1, n, cache)
	chooseTwo := recursion(idx+2, n, cache)

	(*cache)[idx] = chooseOne + chooseTwo

	return chooseOne + chooseTwo
}

func climbStairs(n int) int {
	cache := make(map[int]int)
	return recursion(1, n, &cache)
}
