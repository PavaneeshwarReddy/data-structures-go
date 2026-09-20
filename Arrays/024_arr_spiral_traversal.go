package arrays

/*
Spiral Traversal

- It's simple when you left to right, you are same top but moving from left -> right
- When you move to the bottom, you are moving from top to bottom and right will be decrements
- When move bottom right to bottom left that means you are from right -> left and bottom should move up
- When left bottom to left top, you are moving from bottom to top in the left
*/

func spiralOrder(matrix [][]int) []int {
	left, top := 0, 0
	right := len(matrix[0]) - 1
	bottom := len(matrix) - 1

	res := []int{}

	for left <= right && top <= bottom {
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		top++

		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		right--

		if top <= bottom {
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}
	}

	return res
}
