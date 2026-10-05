package dynamicprogramming

/*
You are given an array prices where prices[i] is the price of a given stock on the ith day.
Find the maximum profit you can achieve. You may complete as many transactions as you like (i.e., buy one and sell one share of the stock multiple times) with the following restrictions:
After you sell your stock, you cannot buy stock on the next day (i.e., cooldown one day).
Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).

Explanation:
- this is also similar to buy and sell any no of times of a stock
- A new addition is that if you are willing to sell today, then we can simply ignore the nxt day
*/

func recursion022(idx int, buy int, n int, prices []int, cache *[][2]int) int {
	if idx >= n {
		return 0
	}

	if (*cache)[idx][buy] != -1 {
		return (*cache)[idx][buy]
	}

	profit := 0

	if buy == 1 {
		take := -prices[idx] + recursion022(idx+1, 0, n, prices, cache)
		notTake := recursion022(idx+1, 1, n, prices, cache)

		profit = max(take, notTake)
	} else {
		take := prices[idx] + recursion022(idx+2, 1, n, prices, cache)
		notTake := recursion022(idx+1, 0, n, prices, cache)

		profit = max(take, notTake)
	}

	(*cache)[idx][buy] = profit

	return profit
}

func maxProfit022(prices []int) int {
	n := len(prices)
	cache := make([][2]int, n)
	for i := range n {
		for j := range 2 {
			cache[i][j] = -1
		}
	}
	return recursion022(0, 1, len(prices), prices, &cache)
}
