package arrays

/*
Max Consecutive Ones
*/

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	res := 0

	for _, val := range nums {
		if val == 1 {
			count += 1
		}
		res = max(res, count)
		if val == 0 {
			count = 0
		}
	}

	return res
}
