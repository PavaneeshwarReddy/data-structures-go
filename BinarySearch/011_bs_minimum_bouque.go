package binarysearch

/*
Minimum Days to make a bouquet
- There are m bouquet and bloomDay will decide at which day flower will boom and we need to select k adjacent flowers to make a single one

- min : 1 day, max : max(bloomDays)

- If we are able to make m bouqets with k adjacent flowers then it is valid
*/

func isValidDay(day int, bloomDay []int, m int, k int) bool {
	bloomed := 0
	i := 0
	cK := 0
	for i < len(bloomDay) {
		if bloomDay[i] <= day {
			cK++
		} else {
			cK = 0
		}
		if cK == k {
			cK = 0
			bloomed++
		}
		i++
	}

	return bloomed >= m
}

func minDays(bloomDay []int, m int, k int) int {
	left := 1
	right := 1

	for _, val := range bloomDay {
		right = max(right, val)
	}

	res := -1
	for left <= right {
		mid := (left + right) / 2

		if isValidDay(mid, bloomDay, m, k) {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return res
}
