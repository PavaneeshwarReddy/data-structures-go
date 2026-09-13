package arrays

/*
Rotate a Sorted Array
- Rotate last k elements
- Rotate first left over elements n - k - 1
- Rotate Whole
*/

func rotate1(nums []int, start int, end int) {
	for start <= end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

func rotate(nums []int, k int) {
	k = k % len(nums)

	rotate1(nums, len(nums)-k, len(nums)-1)
	rotate1(nums, 0, len(nums)-k-1)
	rotate1(nums, 0, len(nums)-1)
}
