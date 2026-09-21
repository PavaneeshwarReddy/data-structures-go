package math

/*
Count total odd digits in a number
*/

func CountOddDigits(num int) int {
	c := 0

	for num > 0 {
		d := num % 10
		if d%2 == 1 {
			c++
		}
		num = num / 10
	}

	return c
}
