package dynamicprogramming

import "strconv"

/*
Partition Subset Sum
- Base Condition: If we reach our target then we return true or else false
- Explore Paths: Either choose or not choose an element in the sequence
- Cache: key will be idx, curr as this is what our previous computations compute
*/

func recursion010(idx int, curr int, n int, target int, nums []int, cache *map[string]bool) bool {
	if curr == target {
		return true
	}
	if idx == n {
		return false
	}

	key := strconv.Itoa(idx) + "a" + strconv.Itoa(curr)
	if val, ok := (*cache)[key]; ok {
		return val
	}

	choose := recursion010(idx+1, curr+nums[idx], n, target, nums, cache)
	noChoose := recursion010(idx+1, curr, n, target, nums, cache)

	(*cache)[key] = choose || noChoose

	return choose || noChoose

}

func PartitionSubset(nums []int, target int) bool {
	cache := make(map[string]bool)
	return recursion010(0, 0, len(nums), target, nums, &cache)
}
