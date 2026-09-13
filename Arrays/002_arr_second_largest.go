package arrays

import "math"

/*
Find Second Largest Element in an Array
*/

func SecondLargest(nums []int) int {
	first := math.MinInt
	second := math.MinInt

	for _, val := range nums {
		if first == math.MinInt {
			first = val
		}
		if first < val {
			second = first
			first = val
		}
	}
	return second
}
