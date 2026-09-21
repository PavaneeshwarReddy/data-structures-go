package math

import "math"

/*
Largest Digit
*/

func LargestDigit(num int) int {
	res := math.MinInt
	for res > 0 {
		res = max(res, res%10)
		res = res / 10
	}

	return res
}
