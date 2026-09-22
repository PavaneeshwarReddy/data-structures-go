package hash

/*
Find max consecutive subsequence
- If we mark all elements in the num as true in a map
- When we are iterating if we check whether there is any previous element val - 1 that means
  - if yes, this is not the start
  - if no, this is the start

- By doing this we are avoiding unecassary loops for nums
- we iterate on hash as we know that duplicate numbers in the sequence doesn't contribute anything
*/
func longestConsecutive(nums []int) int {
	hash := make(map[int]bool)
	for _, val := range nums {
		hash[val] = true
	}
	maxCount := 0
	for val := range hash {
		if !hash[val-1] {
			curr := val
			currCount := 1
			for hash[curr+1] {
				currCount++
				curr++
			}
			maxCount = max(maxCount, currCount)
		}
	}
	return maxCount
}
