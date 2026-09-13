package arrays

/*
Longest Subarray with Sum K
- This is the best in case of sliding windows as it is determinstic, when start increase value decreases and end increase value increases
*/

func FindLongestSum(nums []int, k int) int {
	sum, start, end, res := 0, 0, 0, 0

	for end < len(nums) {
		sum += nums[end]

		for sum > k {
			sum -= nums[start]
			start++
		}

		if sum == k {
			res = max(res, end-start+1)
		}
		end++
	}

	return res
}
