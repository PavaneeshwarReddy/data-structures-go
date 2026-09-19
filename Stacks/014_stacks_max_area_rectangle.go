package stacks

/*
Max area of rectangles formed by 1's in a matrix

- If you consider it is same as area under histogram but for every row, first we need to calculate tower heights that is possible at each row
- Then apply same logic for every row and take max of all
*/
type Stack014 struct {
	Elements []int
}

func (s *Stack014) Push(key int) {
	s.Elements = append(s.Elements, key)
}

func (s *Stack014) Pop() int {
	front := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return front
}

func (s *Stack014) Empty() bool {
	return len(s.Elements) == 0
}

func (s *Stack014) Top() int {
	return s.Elements[len(s.Elements)-1]
}

func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	heights := make([][]int, m)
	for i := 0; i < m; i++ {
		heights[i] = make([]int, n)
	}
	res := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				if matrix[i][j] == '1' {
					heights[i][j] = 1
				} else {
					heights[i][j] = 0
				}
			} else {
				if matrix[i][j] == '1' {
					heights[i][j] = heights[i-1][j] + 1
				} else {
					heights[i][j] = 0
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		stack := Stack014{}

		prevSmaller := make([]int, n)
		nextSmaller := make([]int, n)

		for j := 0; j < n; j++ {
			if stack.Empty() {
				prevSmaller[j] = -1
			} else {
				for !stack.Empty() && heights[i][stack.Top()] >= heights[i][j] {
					stack.Pop()
				}
				if !stack.Empty() {
					prevSmaller[j] = stack.Top()
				} else {
					prevSmaller[j] = -1
				}
			}
			stack.Push(j)
		}

		stack = Stack014{}

		for j := n - 1; j >= 0; j-- {
			if stack.Empty() {
				nextSmaller[j] = n
			} else {
				for !stack.Empty() && heights[i][stack.Top()] > heights[i][j] {
					stack.Pop()
				}
				if !stack.Empty() {
					nextSmaller[j] = stack.Top()
				} else {
					nextSmaller[j] = n
				}
			}
			stack.Push(j)
		}

		for j := 0; j < n; j++ {
			res = max(heights[i][j]*(nextSmaller[j]-prevSmaller[j]-1), res)
		}
	}

	return res

}
