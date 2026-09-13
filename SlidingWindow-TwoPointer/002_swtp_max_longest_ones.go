package slidingwindowtwopointer

/*
Max Consecutive Ones
- These always have a condition where we can either increase or decrease size of windows
- In this problem it is total flips you can perform
- Peform the flip if it's possible, I mean if if it's a 0 flip it
- If flip count is more than allowed then try to increment start so that the condition matches, I mean shrinking the window
- When it satisfies calculate the range
*/

func longestOnes(nums []int, k int) int {
	start, end, maxOnes, flipped := 0, 0, 0, 0

	for end < len(nums) {
		if nums[end] == 0 {
			flipped++
		}

		for flipped > k {
			if nums[start] == 0 {
				flipped--
			}
			start++
		}

		if flipped <= k {
			maxOnes = max(maxOnes, end-start+1)
		}

		end++

	}

	return maxOnes
}
