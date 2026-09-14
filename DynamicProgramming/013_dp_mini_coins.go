package dynamicprogramming

/*
Minimum Coins
- Base Condition: If we reach our target then we return 0 else we return MaxInt
- Explore Paths: Either choose or not choose an element in the sequence
- Cache: key will be idx, curr as this is what our previous computations compute
*/

import (
	"math"
)

func recursion013(coinIdx int, curr int, target int, coins []int, cache *[]int) int {
	if curr > target {
		return math.MaxInt
	}

	if curr == target {
		return 0
	}

	if (*cache)[curr] != -1 {
		return (*cache)[curr]
	}

	res := math.MaxInt
	for i := 0; i < len(coins); i++ {
		res = min(res, recursion013(i, curr+coins[i], target, coins, cache))
	}

	if res != math.MaxInt {
		res += 1
	}

	(*cache)[curr] = res

	return res
}

func coinChange(coins []int, amount int) int {
	cache := make([]int, amount+1)

	for i := range cache {
		cache[i] = -1
	}

	res := recursion013(0, 0, amount, coins, &cache)
	if res == math.MaxInt {
		return -1
	}
	return res
}
