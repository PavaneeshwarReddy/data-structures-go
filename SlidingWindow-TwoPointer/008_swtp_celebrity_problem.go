package slidingwindowtwopointer

/*
Celebrity Problem
- Celebrity knows no one but everyone knows celebrity

- if top points at 0 and bottom at n-1, check if some one knows some one
	matrix[top][bottom] == 1 then that means top is not a celebrity move forward
	matrix[bottom][top] == 1 then that means bottom is not a celebrity move backward
- These stop at a point where top == bottom
	- At this point check if all are 0 then return top or bottom or else -1
*/

func FindCelebrity(matrix [][]int) int {
	m := len(matrix)
	n := len(matrix[0])
	top := 0
	bottom := m - 1

	for bottom > top {

		if matrix[top][bottom] == 1 {
			top++
		} else if matrix[bottom][top] == 1 {
			bottom--
		} else {
			top++
			bottom--
		}
	}

	for j := 0; j < n; j++ {
		if matrix[top][j] != 0 {
			return -1
		}
	}
	return top
}
