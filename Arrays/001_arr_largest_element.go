package arrays

import "math"

/*
Largest Element in an array
*/

func FindMaxArray(nums []int) int {
	res := math.MinInt
	for _, val := range nums {
		res = max(res, val)
	}

	return res
}
