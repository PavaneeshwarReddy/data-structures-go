package bitmanipulation

/*
Check whether number is odd or not
- If last bit is 1 then only it can be odd or else it cannot be odd at any moment
*/

func isOdd(n int) bool {
	if (n & 1) == 1 {
		return true
	}
	return false
}
