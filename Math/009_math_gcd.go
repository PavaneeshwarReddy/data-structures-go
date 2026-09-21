package math

/*
GCD
-If you taken an example who is divisor who is divident and remainder you get this
*/

func GCD(a int, b int) int {

	if a == 0 {
		return b
	}

	return GCD(b%a, a)
}
