package arrays

/*
Next Permutation
- Always we need to keep common prefix and try to rearrange right digits whether we can be able to make it bigger then current
- If not decrease the prefix and move on the

- To make the next permutation posible at every index starting from right, there should be a value which is greater than current idx on the right
- num[i] < num[i+1] will be our break point where we can stop, which means until that it continously increasing , with just increasing we cannot form next permutation
- find the smallest value which is just greater than current idx and swap those, but after swapping those we need to make sure that the current number which we chose is smaller, to make that possible we need sort the right part
*/

import "sort"

func findJustMax(start int, end int, key int, nums []int) int {
	res := start
	for i := end; i > start; i-- {
		if nums[i] > key {
			res = i
			break
		}
	}
	return res
}

func nextPermutation(nums []int) {
	n := len(nums)
	for idx := n - 2; idx >= 0; idx-- {
		if nums[idx] < nums[idx+1] {
			smallIdx := findJustMax(idx+1, n-1, nums[idx], nums)
			nums[smallIdx], nums[idx] = nums[idx], nums[smallIdx]

			suffix := nums[idx+1:]

			sort.Slice(suffix, func(i, j int) bool {
				return suffix[i] < suffix[j]
			})

			return
		}
	}

	i := 0
	j := n - 1

	for i <= j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
