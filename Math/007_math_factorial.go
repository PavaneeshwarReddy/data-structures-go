package math

/*
Find factorial
*/

func Factorial(num int) int {
	res := 1

	for num >= 1 {
		res = res * num
		num--
	}
	return res
}
