package arrays

/*
Check if Array is Sorted or not, rotated array
- This gives a hint that, if the array elements I mean, nums[i] > nums[i+1] this can happen only one time
- If this happens more that one time then the array is not sorted
- We need to check the last rotate indices also
*/

func check(nums []int) bool {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			count += 1
		}
	}
	if nums[len(nums)-1] > nums[0] {
		count += 1
	}
	if count > 1 {
		return false
	}
	return true
}
