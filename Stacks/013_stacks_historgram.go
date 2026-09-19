package stacks

/*
Histogram Area
- Find the max area covered

- For every height if we see how left and right it can spread such that the current is the minimum
- So we can simply nums[i] * width
*/

type Stack013 struct {
	Elements []int
}

func (s *Stack013) Push(key int) {
	s.Elements = append(s.Elements, key)
}

func (s *Stack013) Pop() int {
	front := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return front
}

func (s *Stack013) Empty() bool {
	return len(s.Elements) == 0
}

func (s *Stack013) Top() int {
	return s.Elements[len(s.Elements)-1]
}

func largestRectangleArea(nums []int) int {
	stack := Stack013{}
	n := len(nums)
	prevSmaller := make([]int, n)
	nextSmaller := make([]int, n)

	for i := 0; i < n; i++ {
		if stack.Empty() {
			prevSmaller[i] = -1
		} else {
			for !stack.Empty() && nums[stack.Top()] >= nums[i] {
				stack.Pop()
			}
			if !stack.Empty() {
				prevSmaller[i] = stack.Top()
			} else {
				prevSmaller[i] = -1
			}
		}
		stack.Push(i)
	}

	stack = Stack013{}

	for i := n - 1; i >= 0; i-- {
		if stack.Empty() {
			nextSmaller[i] = n
		} else {
			for !stack.Empty() && nums[stack.Top()] > nums[i] {
				stack.Pop()
			}
			if !stack.Empty() {
				nextSmaller[i] = stack.Top()
			} else {
				nextSmaller[i] = n
			}
		}
		stack.Push(i)
	}

	res := 0

	for i := 0; i < n; i++ {
		res = max(nums[i]*(nextSmaller[i]-prevSmaller[i]-1), res)
	}
	return res
}
