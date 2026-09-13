package slidingwindowtwopointer

/*
Count No of Nice Subarrays
- This is similar to finding exactly that means we have to use k, k-1
- K can be decremented directly if it crosses 0 then we try to bring it back to greater than 0

*/

func nofSub(nums []int, k int) int {
	if k < 0 {
		return 0
	}

	count, start, end := 0, 0, 0

	for end < len(nums) {
		if nums[end]%2 == 1 {
			k -= 1
		}

		for k < 0 {
			if nums[start]%2 == 1 {
				k += 1
			}
			start++
		}

		count += end - start + 1
		end++
	}

	return count
}

func numberOfSubarrays(nums []int, k int) int {
	return nofSub(nums, k) - nofSub(nums, k-1)
}
