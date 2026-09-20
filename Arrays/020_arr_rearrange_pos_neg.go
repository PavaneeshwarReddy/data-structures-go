package arrays

/*
Rearrange negatives and positives into alternating fasion but pos as start
*/

func rearrangeArray(nums []int) []int {
	res := make([]int, len(nums))
	pos := 0
	neg := 1

	for _, val := range nums {
		if val > 0 {
			res[pos] = val
			pos += 2
		} else {
			res[neg] = val
			neg += 2
		}
	}

	return res

}
