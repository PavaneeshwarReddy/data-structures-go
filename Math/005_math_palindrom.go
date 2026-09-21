package math

/*
Palindrom
*/

func Palindrom(num int) bool {
	res := 0

	tempNum := num
	for tempNum > 0 {
		res = res*10 + (res % 10)
		tempNum = tempNum / 10
	}

	if res == num {
		return true
	}

	return false
}
