package stacks

/*
Trapping Rain Water
- Consider you intialize leftMax and rightMax
- If leftMax is smaller thatn rightMax
	- That means for current cell we have a support for right, leftMax - curr will give us the trapped rain water
- or else vice versa
*/

func trap(nums []int) int {

	left := 0
	right := len(nums) - 1
	res := 0

	leftMax, rightMax := nums[left], nums[right]
	left++
	right--

	for left <= right {

		if leftMax <= rightMax {
			if leftMax <= nums[left] {
				leftMax = nums[left]
			} else {
				res += leftMax - nums[left]
			}
			left++
		} else {
			if rightMax <= nums[right] {
				rightMax = nums[right]
			} else {
				res += rightMax - nums[right]
			}
			right--
		}

	}

	return res
}
