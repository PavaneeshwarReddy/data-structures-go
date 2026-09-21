package math

/*
Reverse a number
*/

func ReverseNum(num int) int {
	res := 0

	for num > 0 {
		res = res*10 + (num % 10)
		num = num / 10
	}

	return res
}
