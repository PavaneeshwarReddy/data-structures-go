package dynamicprogramming

/*
House Robbery 2, House are placed circular
- Base Condition: If we reaches out of houses then return 0
- Explore Paths: If he had choosen current house then he has to skip adjacent one, if hadn't choose one then he can choose next one
- Cache: Idx can be the key as there are chances that roberrer visits the same house either by choosing or not choosing

Here one thing to notice is,
- We can either choose first one and ignore the last one
- We can either choose last one and ignore the first one
*/

func recursion005(idx int, n int, nums []int, cache *map[int]int) int {
	if idx > n-1 {
		return 0
	}

	if val, ok := (*cache)[idx]; ok {
		return val
	}

	choose := nums[idx] + recursion005(idx+2, n, nums, cache)
	notChoose := recursion005(idx+1, n, nums, cache)

	(*cache)[idx] = max(choose, notChoose)

	return max(choose, notChoose)
}

func rob2(nums []int) int {
	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	cache1 := make(map[int]int)
	case1 := recursion005(0, n-1, nums, &cache1)

	cache2 := make(map[int]int)
	case2 := recursion005(1, n, nums, &cache2)

	return max(case1, case2)
}
