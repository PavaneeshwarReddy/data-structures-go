package bitmanipulation

/*
Count total no of set bits
*/

func hammingWeight(n int) int {
	count := 0
	for n > 0 {
		count += (n & 1)
		n = n >> 1
	}
	return count
}
