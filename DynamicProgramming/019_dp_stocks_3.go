package dynamicprogramming

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.
Find the maximum profit you can achieve. You may complete at most two transactions.
Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).


Explanation:
- We can consider something like capacity which keeps track of our capacity which we are going to use
- Same we can buy and sell in 4 combinations, here capacity only decreases when buy and sell that means when we choose to sell then only capacity decreases
- If we reach the capacity we can return 0
*/

func recursion019(idx int, buy int, cap int, n int, prices []int, cache *[][2][3]int) int {

	if cap == 0 || idx == n {
		return 0
	}

	if (*cache)[idx][buy][cap] != -1 {
		return (*cache)[idx][buy][cap]
	}

	profit := 0

	if buy == 1 {
		take := -prices[idx] + recursion019(idx+1, 0, cap, n, prices, cache)
		notTake := recursion019(idx+1, 1, cap, n, prices, cache)

		profit = max(take, notTake)
	} else {
		take := prices[idx] + recursion019(idx+1, 1, cap-1, n, prices, cache)
		notTake := recursion019(idx+1, 0, cap, n, prices, cache)

		profit = max(take, notTake)
	}

	(*cache)[idx][buy][cap] = profit

	return profit

}

func maxProfit019(prices []int) int {
	n := len(prices)
	cache := make([][2][3]int, n)
	for i := range n {
		for j := range 2 {
			for k := range 3 {
				cache[i][j][k] = -1
			}
		}
	}
	return recursion019(0, 1, 2, n, prices, &cache)
}
