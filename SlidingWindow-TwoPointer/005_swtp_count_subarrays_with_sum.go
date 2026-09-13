package slidingwindowtwopointer

/*
Count Subarrays with required sum ( Binary Array )
- You can either go with Prefix Sum or sliding window and 2 pointer technique
- Count all valid that can become a valid array for <=goal and <=goal-1 then subtract these we will get what we require for goal

- Why this works, this is predictable, value only increases or decreases this is the logic when we can apply Sliding Window
- Use this when question says exactly or count something all possible outcomes like this or else use normal sliding window for longest, shortest, etc
*/

func countPoss(nums []int, k int) int {
	if k < 0 {
		return 0
	}
	count, sum, start, end := 0, 0, 0, 0

	for end < len(nums) {
		sum += nums[end]

		for sum > k {
			sum -= nums[start]
			start++
		}

		count += end - start + 1
		end += 1
	}

	return count
}

func numSubarraysWithSum(nums []int, k int) int {
	return countPoss(nums, k) - countPoss(nums, k-1)
}
