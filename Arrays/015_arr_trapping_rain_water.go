package arrays

/*
Trapping rain water
- Heights are given and we need to store the amount of water stored between heights
- If you consider at current height min of left Max and right Max is the answer - diff how much it can hold between these 2 maximas

- Done using prefix some
*/

func trap(nums []int) int {
	n := len(nums)
	leftMaxArr := make([]int, n)
	rightMaxArr := make([]int, n)

	leftMaxArr[0] = nums[0]
	rightMaxArr[n-1] = nums[n-1]

	for i := 1; i < n; i++ {
		leftMaxArr[i] = max(leftMaxArr[i-1], nums[i])
		rightMaxArr[n-1-i] = max(rightMaxArr[n-i], nums[n-i-1])
	}

	result := 0

	for i := 0; i < n; i++ {
		diff := min(leftMaxArr[i], rightMaxArr[i]) - nums[i]
		if diff >= 0 {
			result += diff
		}
	}

	return result
}
