package binarysearch

/*
Given an integer array nums and an integer k, split nums into k non-empty subarrays such that the largest sum of any subarray is minimized.
Return the minimized largest sum of the split.
A subarray is a contiguous part of the array.

Explanation:
- This doesnt look like a binary search problem, but ya it is
- This says that we can choose only continous elements not random pick
- Why this became a binary search is search space is fixed and continous, i mean it can go between some value x to y
- Let's say it's binary search then we need to have low and high,
	- low = max of all, let's consider i hold minimum the that means no one can pick more that this value, then that means no one pick anything
	- high = this will be obvious sum of all
- Let's say we exceed our total splits then that means we need to increase the capacity we hold in our hand
*/

func countSplits(nums []int, maxHold int) int {
	splits := 1
	val := 0

	for _, v := range nums {
		if val+v <= maxHold {
			val += v
		} else {
			splits++
			val = v
		}
	}

	return splits

}

func splitArray(nums []int, k int) int {
	low := func(nums []int) int {
		maxVal := -1
		for _, val := range nums {
			maxVal = max(maxVal, val)
		}
		return maxVal
	}(nums)

	high := func(nums []int) int {
		s := 0
		for _, val := range nums {
			s += val
		}
		return s
	}(nums)

	for low <= high {
		mid := (low + high) / 2
		splits := countSplits(nums, mid)
		if splits > k {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}
