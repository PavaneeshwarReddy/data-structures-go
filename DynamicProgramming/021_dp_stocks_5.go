package dynamicprogramming

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day, and an integer fee representing a transaction fee.
Find the maximum profit you can achieve. You may complete as many transactions as you like, but you need to pay the transaction fee for each transaction.

Note:
You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).
The transaction fee is only charged once for each stock purchase and sale.

Explanation:
- If we see it seems similar with buy and sell stock any no of times, just with a new addition transaction fees
- This doesn't have any effect on cache
*/
func recursion021(idx int, buy int, fee int, n int, prices []int, cache *[][2]int) int {
	if idx == n {
		return 0
	}

	if (*cache)[idx][buy] != -1 {
		return (*cache)[idx][buy]
	}

	profit := 0

	if buy == 1 {
		take := -prices[idx] - fee + recursion021(idx+1, 0, fee, n, prices, cache)
		notTake := recursion021(idx+1, 1, fee, n, prices, cache)

		profit = max(take, notTake)
	} else {
		take := prices[idx] + recursion021(idx+1, 1, fee, n, prices, cache)
		notTake := recursion021(idx+1, 0, fee, n, prices, cache)

		profit = max(take, notTake)
	}

	(*cache)[idx][buy] = profit

	return profit
}

func maxProfit021(prices []int, fee int) int {
	n := len(prices)
	cache := make([][2]int, n)
	for i := range n {
		for j := range 2 {
			cache[i][j] = -1
		}
	}
	return recursion021(0, 1, fee, len(prices), prices, &cache)
}
