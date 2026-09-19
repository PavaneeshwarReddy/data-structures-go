package binarysearch

/*
Floor:
- The value which is just smaller than current
- You are trying to approach mid but from minimum side

Ceil:
- Value which is just greater than current
- You are tying to approach mid but from max side

Equality is valid both conditions

*/

func floor(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	res := len(nums)

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] <= target {
			res = nums[mid]
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return res
}

func ceil(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	res := len(nums)

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			res = nums[mid]
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return res
}
