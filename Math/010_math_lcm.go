package math

/*
LCM
- For some value of x = a*z = b*k
- Take the smallest number of both and greater and at some point both are having same value
*/

func LCM(a int, b int) int {
	g := max(a, b)
	s := min(a, b)

	for i := g; i <= a*b; i += g {
		if i%s == 0 {
			return i
		}
	}

	return -1
}

/*
Relation between and LCM and GCD
GCD(a, b) * LCM(b,a) = a*b
*/
