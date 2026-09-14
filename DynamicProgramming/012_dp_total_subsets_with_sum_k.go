package dynamicprogramming

import "strconv"

/*
Count total subsets with sum equals K
Whether a array can be divided into 2 sum equal subsequences
- Base Condition: If we reach our target then we return true or else false
- Explore Paths: Either choose or not choose an element in the sequence
- Cache: key will be idx, curr as this is what our previous computations compute
*/

func recursion012(idx int, curr int, n int, target int, nums []int, cache *map[string]int) int {

	if curr == target {
		return 1
	}

	if idx == n {
		return 0
	}

	key := strconv.Itoa(idx) + "a" + strconv.Itoa(curr)
	if val, ok := (*cache)[key]; ok {
		return val
	}

	choose := recursion012(idx+1, curr+nums[idx], n, target, nums, cache)
	notChoose := recursion012(idx+1, curr, n, target, nums, cache)

	(*cache)[key] = choose + notChoose

	return choose + notChoose

}

func TotalSubsetsWithSumK(nums []int, target int) int {
	cache := make(map[string]int)
	return recursion012(0, 0, len(nums), target, nums, &cache)
}
