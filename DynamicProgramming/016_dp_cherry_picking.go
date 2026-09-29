package dynamicprogramming

/*
Cherry Picking

You are given a rows x cols matrix grid representing a field of cherries where grid[i][j] represents the number of cherries that you can collect from the (i, j) cell.

You have two robots that can collect cherries for you:

Robot #1 is located at the top-left corner (0, 0), and
Robot #2 is located at the top-right corner (0, cols - 1).
Return the maximum number of cherries collection using both robots by following the rules below:

From a cell (i, j), robots can move to cell (i + 1, j - 1), (i + 1, j), or (i + 1, j + 1).
When any robot passes through a cell, It picks up all cherries, and the cell becomes an empty cell.
When both robots stay in the same cell, only one takes the cherries.
Both robots cannot move outside of the grid at any moment.
Both robots should reach the bottom row in grid.

Explanation:
- There are 2 robots, either one can move one can stay but if you think what will it do by staging there
- So let's consider we can move two robots at once.
- If we observe the constraints it can move in 3 directions but towards down, by this we can consider that robots move towards down, we can consider at the same moment both stay on the same row
- If they both stay on the same row at the same time, there are only two cases, either they fall on same cell or different cell
- If they fall on same cell then we add once or else if they fall on different cell then they can collect both
- They have too many combination one can move other than move right or some other directions so we need to cumulate all and take max out of it
*/
import "math"

func recursion016(row int, col1 int, col2 int, m int, n int, grid [][]int, cache *[][][]int) int {
	if col1 < 0 || col1 >= n || col2 < 0 || col2 >= n {
		return 0
	}

	if row == m-1 {
		if col1 == col2 {
			return grid[row][col1]
		} else {
			return grid[row][col1] + grid[row][col2]
		}
	}

	if (*cache)[row][col1][col2] != -1 {
		return (*cache)[row][col1][col2]
	}

	curr := 0

	if col1 == col2 {
		curr = grid[row][col1]
	} else {
		curr = grid[row][col1] + grid[row][col2]
	}

	maxNext := math.MinInt

	for d1 := -1; d1 <= 1; d1++ {
		for d2 := -1; d2 <= 1; d2++ {
			maxNext = max(maxNext, recursion016(row+1, col1+d1, col2+d2, m, n, grid, cache))
		}
	}

	(*cache)[row][col1][col2] = curr + maxNext

	return curr + maxNext
}

func cherryPickup(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	cache := make([][][]int, m)

	for i := range m {
		cache[i] = make([][]int, n)
		for j := range n {
			cache[i][j] = make([]int, n)
			for k := range n {
				cache[i][j][k] = -1
			}
		}
	}

	return recursion016(0, 0, n-1, m, n, grid, &cache)
}
