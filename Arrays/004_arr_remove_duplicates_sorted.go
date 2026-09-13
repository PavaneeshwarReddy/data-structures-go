package arrays

/*
Remove Duplicates From a sorted array
- The main intution works like a sliding window but not the same
*/

func removeDuplicates(nums []int) int {
	start := 0
	end := 0

	for end < len(nums) {
		nums[start] = nums[end]
		for end < len(nums) && nums[start] == nums[end] {
			end++
		}
		start++
	}

	return start
}
