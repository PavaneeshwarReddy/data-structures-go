package dynamicprogramming

/*
Whether a array can be divided into 2 sum equal subsequences
- Base Condition: If we reach our target then we return true or else false
- Explore Paths: Either choose or not choose an element in the sequence
- Cache: key will be idx, curr as this is what our previous computations compute, sometimes using map can be overhelming, if predictable values are minimal use 2 D array, [n][total] this is predictable so we can use this
*/

func recursion011(idx int, curr int, n int, total int, nums []int, cache *[][]int) bool {
	if curr == total/2 {
		return true
	}

	if idx == n || curr > total/2 {
		return false
	}

	if (*cache)[idx][curr] != 0 {
		if (*cache)[idx][curr] == 1 {
			return true
		}
		return false
	}

	notChoose := recursion011(idx+1, curr, n, total, nums, cache)
	choose := recursion011(idx+1, curr+nums[idx], n, total, nums, cache)

	if choose || notChoose {
		(*cache)[idx][curr] = 1
	} else {
		(*cache)[idx][curr] = -1
	}

	return choose || notChoose
}

func canPartition(nums []int) bool {
	total := 0
	for _, val := range nums {
		total += val
	}
	if total%2 != 0 {
		return false
	}
	cache := make([][]int, len(nums)+1)
	for i := 0; i <= len(nums); i++ {
		cache[i] = make([]int, total+1)
		for j := 0; j <= total; j++ {
			cache[i][j] = 0
		}
	}
	return recursion011(0, 0, len(nums), total, nums, &cache)
}
