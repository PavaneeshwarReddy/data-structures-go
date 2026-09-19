package binarysearch

/*
Search Insert Position
- It's simply lower bound

*/

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	res := len(nums)

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return res
}
