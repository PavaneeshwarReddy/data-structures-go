package binarysearch

/*
Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.
You must not use any built-in exponent function or operator.

Explanation:
- Minimum can be as minimum as 1 but max can go upto x/2
- Later we can move in the space
*/

func mySqrt(x int) int {
	if x <= 1 {
		return x
	}
	low := 0
	high := x / 2

	for low <= high {
		mid := (low + high) / 2
		if mid*mid > x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return high
}
