package arrays

/*
Two Sum
You can simply use dictionary to track the individual elements
*/

func twoSum(nums []int, target int) []int {
	cache := make(map[int]int)

	for idx, val := range nums {
		if tIdx, ok := cache[target-val]; ok {
			return []int{tIdx, idx}
		}
		cache[val] = idx
	}

	return []int{-1, -1}
}
