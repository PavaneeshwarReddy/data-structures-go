package arrays

/*
Search for the number that appears once well other occur twice
*/

func singleNumber(nums []int) int {
	res := 0
	for _, val := range nums {
		res = res ^ val
	}
	return res
}
