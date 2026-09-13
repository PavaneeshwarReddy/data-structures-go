package arrays

/*
Rotate the Array by One
*/

func RotateArrByOne(nums []int) {
	prevValue := nums[0]
	for i := 1; i < len(nums); i++ {
		tempValue := nums[i]
		nums[i] = prevValue
		prevValue = tempValue
	}

	nums[0] = prevValue

}
