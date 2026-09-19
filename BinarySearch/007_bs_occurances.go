package binarysearch

/*
Find total occurances
- Upper Occurance - Lower Occurance + 1
*/

func findOccurances(nums []int, target int) int {
	uO := upperOcc(nums, target)
	lC := lowerOcc(nums, target)

	return uO - lC + 1
}
