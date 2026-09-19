package binarysearch

/*

Lower Occ:
- If you found then you need to move towards left
- Only store when mid value is same as target

Upper Occ:
- If you found then you need to move towards right
- Only store when mid value is same as target

*/

func lowerOcc(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	res := -1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			right = mid - 1
			if nums[mid] == target {
				res = mid
			}
		} else {
			left = mid + 1
		}
	}
	return res
}

func upperOcc(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	res := -1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > target {
			right = mid - 1

		} else {
			left = mid + 1
			if nums[mid] == target {
				res = mid
			}
		}
	}
	return res
}

func searchRange(nums []int, target int) []int {
	return []int{lowerOcc(nums, target), upperOcc(nums, target)}
}
