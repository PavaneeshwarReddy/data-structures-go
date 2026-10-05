package dynamicprogramming

/*
Count Total No of Squares
Given a m * n matrix of ones and zeros, return how many square submatrices have all ones.

Explanation:
- The core idea is how many squares end at this cell
- If its a zero then we can simply return 0 or else we can take minimum of top, left, diagnol values and check
- For every cell we need to check how many squares end at this position recursively
- We can simply memoize it later
*/

func recursion017(i int, j int, m int, n int, matrix [][]int, cache *[][]int) int {
	if i < 0 || j < 0 {
		return 0
	}

	if (*cache)[i][j] != -1 {
		return (*cache)[i][j]
	}

	if matrix[i][j] == 0 {
		(*cache)[i][j] = 0
		return 0
	}

	left := recursion017(i-1, j, m, n, matrix, cache)
	right := recursion017(i, j-1, m, n, matrix, cache)
	diagnol := recursion017(i-1, j-1, m, n, matrix, cache)

	(*cache)[i][j] = 1 + min(left, right, diagnol)

	return (*cache)[i][j]

}

func countSquares(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	ans := 0
	cache := make([][]int, m)
	for i := 0; i < m; i++ {
		cache[i] = make([]int, n)
		for j := 0; j < n; j++ {
			cache[i][j] = -1
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			ans += recursion017(i, j, m, n, matrix, &cache)
		}
	}
	return ans
}
