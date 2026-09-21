package math

import (
	"fmt"
	"math"
)

/*
Reverse integer
-power(2,31) <=x <= power(2,31)-1
if reverse crosses the 32 bit range then return 0

1. Consider we will reverse the entire integer but not the last digit
2. If last digit is greater than max last digit then that means we cannot reverse

XXXXXXXX(left over) X ( last digit )
first we check whether left over part is greater than maxInt left over or not
*/
func reverse(num int) int {
	maxInt := math.MaxInt32
	minInt := math.MinInt32

	res := 0
	for num != 0 {
		lastDigit := num % 10
		num = num / 10
		if (res > (maxInt / 10)) || ((res == (maxInt / 10)) && lastDigit >= maxInt%10) {
			return 0
		}
		if (res < (minInt / 10)) || (res == (minInt/10) && lastDigit <= minInt%10) {
			return 0
		}

		res = (res * 10) + lastDigit

		fmt.Println(res)

	}

	return res
}
