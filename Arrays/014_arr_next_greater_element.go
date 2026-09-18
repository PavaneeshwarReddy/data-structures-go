package arrays

/*
Next Greater Element
For element in nums1 find nums1[i] == nums[j] and then after
find a greater element to the right side of it, if you found return or else -1

*/

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	result := []int{}

	for _, val := range nums1 {
		foundIdx := -1
		for idx2 := range nums2 {
			if nums2[idx2] == val {
				foundIdx = idx2
				break
			}
		}
		nxtGreater := -1
		if foundIdx != -1 {
			for i := foundIdx + 1; i < len(nums2); i++ {
				if val < nums2[i] {
					nxtGreater = nums2[i]
					break
				}
			}
		}

		result = append(result, nxtGreater)
	}

	return result
}
