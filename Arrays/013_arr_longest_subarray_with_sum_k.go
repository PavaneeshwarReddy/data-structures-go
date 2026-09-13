package arrays

/*
Longest Subarray with Sum k
- Here whenever a question is related to longest subarray but it doesn't say positive then we cannot use sliding window
- We have to go with prefixSum

- Here in prefixSum we store min indices and we iterate over them, if more min is found then we go with that
*/

func FindLongestSubarray(nums []int, k int) int {
	prefSum := make(map[int]int)
	sum := 0
	res := 0

	for i := range nums {
		sum += nums[i]

		if sum == k {
			res = max(res, i+1)
		}

		leftOverSum := sum - k

		if val, ok := prefSum[leftOverSum]; ok {
			res = max(res, i-val+1)
		}

		if _, ok := prefSum[sum]; !ok {
			prefSum[sum] = i
		}
	}

	return res
}
