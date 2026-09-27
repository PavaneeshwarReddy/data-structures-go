package bitmanipulation

/*
Given an integer array nums, in which exactly two elements appear only once and all the other elements appear exactly twice. Find the two elements that appear only once. You can return the answer in any order.
You must write an algorithm that runs in linear runtime complexity and uses only constant extra space.

Explanation:
	- First we will xor of all numbers in the array
	- We are left with a ^ b which we want to find out, but we cannot return them seperately because they are xored
	- If a bit is 1 at some point we can for sure say that one of bit in a or b is 1 and else is 0 at that position
	- So now we iterate again and put numbers in 2 buckets, if current is 1 for xor and current bit is 1 for curr num then we put in 1s bucket

The main idea is if it has different bits they will be in different buckets for sure if not they will be in same bucket.

-Element -> represent in Element in binary, complement it and add 1
*/

func singleNumber2(nums []int) []int {
	cXor := 0
	for _, val := range nums {
		cXor ^= val
	}

	diffBit := cXor & -cXor // we want to extract the right most bit set bit that is point where they both differ
	a, b := 0, 0

	for _, num := range nums {
		if diffBit&num != 0 {
			a ^= num
		} else {
			b ^= num
		}
	}

	return []int{a, b}
}
