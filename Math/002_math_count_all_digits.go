package math

/*
Count all digits of a number
*/

func CountDigits(num int) int {
	c := 0
	for num > 0 {
		num = num / 10
		c += 1
	}

	return num
}
