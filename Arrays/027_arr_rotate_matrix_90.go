package arrays

/*
Rotate the matrix by 90 degrees

You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).
You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. DO NOT allocate another 2D matrix and do the rotation.

Explanation:
- Vertical Reverses
- Transpose of i,j with j, i above the upper diagnal
*/

func rotate90(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])

	// vertical reverses
	for j := range n {
		l := 0
		h := m - 1
		for l <= h {
			matrix[l][j], matrix[h][j] = matrix[h][j], matrix[l][j]
			l++
			h--
		}
	}

	// transpose along the upper diagnol
	for i := range m {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

}
