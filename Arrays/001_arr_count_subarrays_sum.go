package arrays

/*
Count total no of subarrays with sum k
- Prefix sum is a different case, it says how many sub arrays are possible which can provide me required sum
- If require sum is k and current sum is prefSum then difference is present in any of the prev array or not
*/

func subarraySum(nums []int, k int) int {
	pfSum := make(map[int]int)
	sum := 0
	res := 0

	pfSum[0] = 1

	for i := range len(nums) {
		sum += nums[i]

		if val, ok := pfSum[sum-k]; ok {
			res += val
		}

		pfSum[sum] += 1
	}

	return res
}
