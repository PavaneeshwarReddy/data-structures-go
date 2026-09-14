package dynamicprogramming

/*
House Robery
- Base Condition: If we reaches out of houses then return 0
- Explore Paths: If he had choosen current house then he has to skip adjacent one, if hadn't choose one then he can choose next one
- Cache: Idx can be the key as there are chances that roberrer visits the same house either by choosing or not choosing
*/

func recursion004(idx int, n int, nums []int, cache *map[int]int) int {
	if idx > n-1 {
		return 0
	}

	if val, ok := (*cache)[idx]; ok {
		return val
	}

	choose := nums[idx] + recursion004(idx+2, n, nums, cache)
	notChoose := recursion004(idx+1, n, nums, cache)

	(*cache)[idx] = max(choose, notChoose)

	return max(choose, notChoose)
}

func rob(nums []int) int {
	cache := make(map[int]int)
	return recursion004(0, len(nums), nums, &cache)
}
