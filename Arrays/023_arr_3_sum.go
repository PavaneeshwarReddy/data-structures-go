package arrays

/*
3 Sum - Sorted Array

First approach:
Brute force 3 inner loops

Second approach: Using hashmap
target - (num[i] + num[j]) is orrcured in our previous iterator or not

Third approach: 2 pointers
i is constant
j (i) and k (points to end ) can vary, if sum if greater then you can move k else j

*/

func ThreeSum(nums []int, target int) bool {
	for i := range nums {
		j := i + 1
		k := len(nums) - 1

		for j <= k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == target {
				return true
			}
			if sum > target {
				k--
			} else if sum < target {
				j++
			}
		}
	}

	return false
}
