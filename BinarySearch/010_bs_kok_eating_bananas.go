package binarysearch

/*
Koko Eating Bananas

- A monkey can pick with any speed k, and it can only pick one pile at a time. If k is larger than pile it picks it eats everything or k once and then later pile
- It has to complete in H hours

- Here by this k can vary from 1->max(piles), we should find min of k such that it will be completed in H hours

*/

func isValid(k int, piles []int, h int) bool {
	if k == 0 {
		return false
	}
	currH := 0
	for _, val := range piles {
		currH += val / k
		if val%k != 0 {
			currH += 1
		}

		if currH > h {
			return false
		}
	}

	return true
}

func minEatingSpeed(piles []int, h int) int {
	left := 0
	right := 0
	res := 0

	for _, val := range piles {
		right = max(right, val)
	}

	for left <= right {
		mid := (left + right) / 2
		if isValid(mid, piles, h) {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return res
}
