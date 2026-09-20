package arrays

/*
Majority Element
- We can simply maintain a counter and validate
*/

func majorityElement(nums []int) int {
	currElement := -1
	count := 0

	for _, val := range nums {
		if count == 0 {
			currElement = val
		}

		if currElement == val {
			count++
		} else {
			count--
		}
	}

	return currElement
}
