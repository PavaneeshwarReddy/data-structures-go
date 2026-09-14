package dynamicprogramming

/*
Target Sum
- Base Condition: If we reach our target then we return 1 return 0
- Explore Paths: Choose every element with alternative signs +ve or -ve
- Cache: key will be idx, curr as this is what our previous computations compute
*/

import "strconv"

func recursion014(idx int, curr int, n int, target int, nums []int, cache *map[string]int) int {
	if idx == n {
		if curr == target {
			return 1
		}
		return 0
	}

	key := strconv.Itoa(idx) + "a" + strconv.Itoa(curr)
	if val, ok := (*cache)[key]; ok {
		return val
	}

	add := recursion014(idx+1, curr+nums[idx], n, target, nums, cache)
	sub := recursion014(idx+1, curr-nums[idx], n, target, nums, cache)

	(*cache)[key] = add + sub

	return add + sub
}

func findTargetSumWays(nums []int, target int) int {
	cache := make(map[string]int)
	return recursion014(0, 0, len(nums), target, nums, &cache)
}
