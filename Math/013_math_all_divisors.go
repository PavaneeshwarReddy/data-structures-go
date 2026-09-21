package math

/*
All Divisors
*/

func AllDivisors(num int) []int {
	res := []int{}

	for i := 1; i <= (num/2)+1; i++ {
		if (num % i) == 0 {
			res = append(res, i)
		}
	}

	return res
}
