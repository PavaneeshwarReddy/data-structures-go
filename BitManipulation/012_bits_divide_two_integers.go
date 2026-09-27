package bitmanipulation

/*
Given two integers dividend and divisor, divide two integers without using multiplication, division, and mod operator.
The integer division should truncate toward zero, which means losing its fractional part. For example, 8.345 would be truncated to 8, and -2.7335 would be truncated to -2.
Return the quotient after dividing dividend by divisor.

Note: Assume we are dealing with an environment that could only store integers within the 32-bit signed integer range: [−231, 231 − 1]. For this problem, if the quotient is strictly greater than 231 - 1, then return 231 - 1, and if the quotient is strictly less than -231, then return -231.


Explanation:
Method-1: We can simply take dividend and add it until we cross divisor
Method-2:
	- Consider 22/3 -> 3 * 7 -> 3 * (4+2+1) -> 3 * power(2,2) + 3*power(2,1) + 3*power(2,0)
	- For every iteration we try to remove the max number we can remove from the value
	- Loop-1: 3*2 can be removed, 3*2*2 can be removed, 3*2*2*2 cannot be removed we stop here, we remove 3*2*2 which is 12, here we are able to remove 4, ans = 4
	- Loop-2: we left with 10, we can remove 3*2, we left with 4, here we able to remove 2, 4+2=6
	- Loop-3: we left right 4, we can remove 3*1, we left with 1 we stop, this is the remainder and answer is the quotient, final we are able to remove 1, 6+1=7
*/

import "math"

func divide(dividend int, divisor int) int {
	isNegative := (dividend < 0) != (divisor < 0)

	absDividend := int64(dividend)
	if absDividend < 0 {
		absDividend = -absDividend
	}

	absDivisor := int64(divisor)
	if absDivisor < 0 {
		absDivisor = -absDivisor
	}

	ans := int64(0)

	for absDividend >= absDivisor {
		count := 0
		for absDividend >= (absDivisor << (count + 1)) {
			count++
		}
		ans += (1 << count)
		absDividend -= (absDivisor << count)
	}

	if isNegative {
		ans = -ans
	}

	if ans > math.MaxInt32 {
		return math.MaxInt32
	}
	if ans < math.MinInt32 {
		return math.MinInt32
	}

	return int(ans)

}
