package arrays

/*
Subarrays with xor k
*/

func subarraysWithXorK(nums []int, k int) int {
	res := 0
	prefXor := make(map[int]int)
	curr := 0

	for _, val := range nums {
		curr ^= val

		if curr^k == 0 {
			res++
		}

		xorReq := curr ^ k

		res += prefXor[xorReq]

		prefXor[curr]++
	}
	return res

}
