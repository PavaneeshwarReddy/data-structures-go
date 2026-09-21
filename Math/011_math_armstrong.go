package math

/*
Armstrong Number
n = total digits
power(digit1, n) + power(digiti2, n) + .. = num
*/

func CheckArmstrong(num int) bool {
	digits := 0
	tempNum := num

	for tempNum > 0 {
		digits++
		tempNum = num / 10
	}

	tempNum = num
	res := 0
	for tempNum > 0 {
		res = res + (tempNum % 10)
		tempNum = tempNum / 10
	}
	return res == num
}
