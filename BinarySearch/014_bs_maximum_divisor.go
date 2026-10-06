package binarysearch

/*
Given an array of integers nums and an integer threshold, we will choose a positive integer divisor, divide all the array by it, and sum the division's result. Find the smallest divisor such that the result mentioned above is less than or equal to threshold.
Each result of the division is rounded to the nearest integer greater than or equal to that element. (For example: 7/3 = 3 and 10/2 = 5).
The test cases are generated so that there will be an answer.

Explanation:
- If we see again it seem similar to binary search
- There is a threshold value only range something in between
- We should choose our low and high
	- low = 1, smallest can be as small as 1 only
	- high = will be max, why because it's the largest divisor we can pick so that everything in the array gets divided
- We can move the pointer based on this, if we are unable to meet the threshold we increment the divisor
*/

func calculateThreshold(nums []int, divisor int) int {
	s := 0
	for _, val := range nums {
		s += (val + divisor - 1) / divisor
	}
	return s
}

func smallestDivisor(nums []int, threshold int) int {
	low := 1
	high := func(nums []int) int {
		maxVal := -1
		for _, val := range nums {
			maxVal = max(maxVal, val)
		}
		return maxVal
	}(nums)

	for low <= high {
		mid := (low + high) / 2
		currThreshold := calculateThreshold(nums, mid)
		if currThreshold > threshold {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	return low
}
