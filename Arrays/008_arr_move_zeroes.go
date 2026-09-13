package arrays

/*
Move Zeroes to the end
- j, pointer to place the next non zero element, it goes as prev index of current i or else same
*/

func moveZeroes(nums []int) {
	j := 0
	for i := range nums {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j += 1
		}
	}
}
