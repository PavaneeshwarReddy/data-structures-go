package dynamicprogramming

/*
A ninja has planned a n-day training schedule. Each day he has to perform one of three activities - running, stealth training, or fighting practice. The same activity cannot be done on two consecutive days and the ninja earns a specific number of merit points, based on the activity and the given day.
Given a n x 3-sized matrix, where matrix[i][0], matrix[i][1], and matrix[i][2], represent the merit points associated with running, stealth and fighting practice, on the (i+1)th day respectively. Return the maximum possible merit points that the ninja can earn.

Explanation:
- By seeing we are able to understand that, every day he can either choose of 3 tasks, but which is not equal to the task he choose
- That is why I kept another task which is 4 which means he hasn't choosen any tasks, which invaidates condition when looped until 3
*/

func recursion023(day int, choose int, m int, n int, matrix [][]int, cache *[][]int) int {
	if day == m {
		return 0
	}

	if (*cache)[day][choose] != -1 {
		return (*cache)[day][choose]
	}

	maxEarn := 0

	for i := range 3 {
		if choose != i {
			take := matrix[day][i] + recursion023(day+1, i, m, n, matrix, cache)
			notTake := recursion023(day+1, choose, m, n, matrix, cache)

			maxEarn = max(maxEarn, max(take, notTake))
		}
	}

	(*cache)[day][choose] = maxEarn

	return maxEarn
}

func ninjaTraining(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	cache := make([][]int, m)
	for i := range m {
		cache[i] = make([]int, 4)
		for j := range 4 {
			cache[i][j] = -1
		}
	}
	return recursion023(0, 3, m, n, matrix, &cache)
}
