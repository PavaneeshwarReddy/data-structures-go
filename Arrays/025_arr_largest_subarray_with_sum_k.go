package arrays

/*
Largest subarray with sum 0

- Here the meaning be like, we are searching for equal prefix sum, if this occurs then that means inbetween them all the sum = 0
- We don't include the first because it started their and ended at different so we should not include this
*/

func LargestSubarrayWithSumZero(nums []int) int {
	prefSum := make(map[int]int)
	sum := 0
	mx := 0

	for idx := range nums {
		sum += nums[idx]

		if sum == 0 {
			mx = max(mx, idx+1)
		}

		requiredSum := -sum
		if val, ok := prefSum[requiredSum]; ok {
			mx = max(mx, val-idx)
		}

		prefSum[-sum] = idx
	}

	return mx
}
