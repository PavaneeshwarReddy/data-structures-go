package arrays

/*
Subarray Sum Max
- If sum goes negative don't carry the sum.

- You need no choose any element here, if it goes negative we can ignore
*/

import "math"

func maxSubArray(nums []int) int {
	sum := 0
	maxSum := math.MinInt

	for _, val := range nums {
		sum += val

		if sum > maxSum {
			maxSum = sum
		}

		if sum < 0 {
			sum = 0
		}
	}

	return maxSum
}
