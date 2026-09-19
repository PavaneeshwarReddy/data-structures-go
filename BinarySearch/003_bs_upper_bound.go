package binarysearch

/*
Upper bound
- arr[i] > x
*/

func upperBound(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	res := -1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return res
}
