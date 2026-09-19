package binarysearch

/*
Search in Rotate Sorted array, this can contain duplicates

- In the prev code we can consider that every value can be greater or smaller
- Here left mid and right can become equal and we can't make a move condition
- To avoid that if they are equal then we can simply incroement left and decrement right
- Prev approach we can directly say which array is sorted and which side is not, but due to duplicate we can't tell if they all are equal
*/

func searchSorted2(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
			continue
		}

		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
