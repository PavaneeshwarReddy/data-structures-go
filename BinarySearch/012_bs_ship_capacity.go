package binarysearch

/*
Capacity To Ship Packages
- Left can be max(weights) and Right can be sum(weights)
*/

func isValid012(weight int, weights []int, days int) bool {
	totalDays := 1
	currWeight := 0
	for _, val := range weights {

		if currWeight+val > weight {
			currWeight = val
			totalDays++
		} else {
			currWeight += val
		}

		if totalDays > days {
			return false
		}
	}

	return true
}

func shipWithinDays(weights []int, days int) int {
	res := 0
	left := func(weights []int) int {
		m := 0
		for _, val := range weights {
			m = max(m, val)
		}
		return m
	}(weights)
	right := func(weights []int) int {
		sum := 0
		for _, val := range weights {
			sum += val
		}
		return sum
	}(weights)

	for left <= right {
		mid := (left + right) / 2
		if isValid012(mid, weights, days) {
			res = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return res

}
