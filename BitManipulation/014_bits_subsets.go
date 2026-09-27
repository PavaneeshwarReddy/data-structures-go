package bitmanipulation

/*
Given an integer array nums of unique elements, return all possible subsets (the power set).
The solution set must not contain duplicate subsets. Return the solution in any order.

Explanation:
- Let's see how count of subsets varry
	- n = 1, 1
	= n = 2, 4
	- n = 3, 8
- We can see that total subets are n << 1
- Let's see for 4, these are all the combinations incase of len 4
	- 0 0
	- 0 1
	- 1 0
	- 1 1
- By using this logic we will include elements in the result
- It's about choosing for all the subset combinations whether the index can be choosen or not
*/

func subsets(nums []int) [][]int {
	totalSub := 1 << len(nums)

	res := [][]int{}

	for i := range totalSub {
		tempRes := []int{}
		for j := range len(nums) {
			if (i>>j)&1 == 1 {
				tempRes = append(tempRes, nums[j])
			}
		}
		res = append(res, tempRes)
	}
	return res
}
