package bitmanipulation

/*
Given two integers L and R. Find the XOR of the elements in the range [L , R].

Explanation:
- It's all about seeing patterns
- Write all the patterns from N=1 to N=10, you will be finding these results
- N = 1, 5, 9 ,.... = 1
- N = 2, 4, 6, 8,.... = 3, 4, 7, 8

- When divided by 4
	- remainder: 1, then 1
	- remainder: 2, then N+1
	- remainder: 3, then 0
	- remainder: 0, then N
*/

func Xor(n int) int {
	switch n % 4 {
	case 1:
		return 1
	case 2:
		return n + 1
	case 3:
		return 0
	}

	return n
}

func findRangeXOR(l, r int) int {
	xorUntilL := Xor(l)
	xorUntilR := Xor(r)

	return xorUntilR ^ xorUntilL ^ l // when we xor until R even L cancels out, so we need to include that l also

}
