package math

/*
Perfect Number
- If all divisors of a num except itself, sum of them should be equal to the current num
*/

func PerfectNumber(num int) bool {
	res := 0

	for i := 1; i <= (num/2)+1; i++ {
		if (num % i) == 0 {
			res += i
		}
	}

	return res == num
}
