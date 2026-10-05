package dynamicprogramming

/*
You are given an integer array prices where prices[i] is the price of a given stock on the ith day, and an integer k.
Find the maximum profit you can achieve. You may complete at most k transactions: i.e. you may buy at most k times and sell at most k times.
Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).

Explanation:
- It is very similar to previous problem, but just instead of having capacity as 2 we have some dynamic value
*/
func recursion020(idx int, buy int, cap int, n int, prices []int, cache *[][2][]int) int {

	if cap == 0 || idx == n {
		return 0
	}

	if (*cache)[idx][buy][cap] != -1 {
		return (*cache)[idx][buy][cap]
	}

	profit := 0

	if buy == 1 {
		take := -prices[idx] + recursion020(idx+1, 0, cap, n, prices, cache)
		notTake := recursion020(idx+1, 1, cap, n, prices, cache)

		profit = max(take, notTake)
	} else {
		take := prices[idx] + recursion020(idx+1, 1, cap-1, n, prices, cache)
		notTake := recursion020(idx+1, 0, cap, n, prices, cache)

		profit = max(take, notTake)
	}

	(*cache)[idx][buy][cap] = profit

	return profit

}

func maxProfit20(cap int, prices []int) int {
	n := len(prices)
	cache := make([][2][]int, n)
	for i := range n {
		for j := range 2 {
			cache[i][j] = make([]int, cap+1)
			for k := range cap + 1 {
				cache[i][j][k] = -1
			}
		}
	}
	return recursion020(0, 1, cap, n, prices, &cache)
}
