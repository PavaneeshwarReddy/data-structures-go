package arrays

/*
Print the maximum subarray sum
- Similar to kadens but here we need to know where we are taking and stopping subarray
- The array starts at sum == 0 and ends where curr max exceeds max we stored
*/

import "math"

func maxSubArrayPrint(nums []int) []int {
	resStart, resEnd := -1, -1
	start := -1
	sum := 0
	maxSum := math.MinInt

	for idx, val := range nums {

		if sum == 0 {
			start = idx

		}

		sum += val

		if sum > maxSum {
			maxSum = sum
			resStart = start
			resEnd = idx
		}

		if sum < 0 {
			sum = 0
		}
	}

	return []int{resStart, resEnd}
}
