package dynamicprogramming

/*
You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock at any time. However, you can sell and buy the stock multiple times on the same day, ensuring you never hold more than one share of the stock.
Find and return the maximum profit you can achieve.

Explanation:
- You can buy and sell any number of times
- At every instance you get this chance
  - If you have bought earlier then you either sell or leave
  - If you haven't bough earler then you can either buy or leave

- There are around 4 combinations which you can do
- Just choose the max profit one
- Once you buy something then that means it your profit became -currVal + recursion018
*/
func recursion018(idx int, buy int, n int, prices []int, cache *[][2]int) int {

	if idx >= n {
		return 0
	}

	if (*cache)[idx][buy] != -1 {
		return (*cache)[idx][buy]
	}

	take, notTake := 0, 0

	if buy == 1 {
		take = -prices[idx] + recursion018(idx+1, 0, n, prices, cache)
		notTake = recursion018(idx+1, 1, n, prices, cache)
	} else {
		take = prices[idx] + recursion018(idx+1, 1, n, prices, cache)
		notTake = recursion018(idx+1, 0, n, prices, cache)
	}

	(*cache)[idx][buy] = max(take, notTake)

	return max(take, notTake)
}

func maxProfit(prices []int) int {
	cache := make([][2]int, len(prices))
	for i := 0; i < len(prices); i++ {
		cache[i][0] = -1
		cache[i][1] = -1
	}
	return recursion018(0, 1, len(prices), prices, &cache)
}
