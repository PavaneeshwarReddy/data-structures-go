package bitmanipulation

/*
Given an integer array nums where every element appears three times except for one, which appears exactly once. Find the single element and return it.
You must implement a solution with a linear runtime complexity and use only constant extra space.

Explanation:
- There many different ways we can explore this
Method-1: Check the count using map and solve using count if equal to 1 then return that result
Method-2: Pick each bit and count how many bits are set if equals multiple of 3 then we can leave or else we can add it to the result
Method-3: Using Buckets -this has to be known earlier to get this idea
		- Ones keep only elements which repeated one time
		- Twos keep only elements which repeated two times
		- Three times doesn't exists any where in Ones or Twos
		- Ones = ( Ones ^ num ) & ~( Twos ), this completement can be implemented directly by ^num
		  Twos = ( Twos ^ num ) & ~( Ones )
*/

func singleNumber(nums []int) int {
	ones := 0
	twos := 0

	for _, val := range nums {
		ones = (ones ^ val) & (^(twos))
		twos = (twos ^ val) & (^(ones))
	}

	return ones
}
