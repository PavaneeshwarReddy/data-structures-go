package stacks

/*
Sum of Subarray Ranges

- The range of a subarray of nums is the difference between the largest and smallest element in the subarray.

- The means for how many subarrays does this i is smaller and largest then we can subtract

*/

type Stack011 struct {
	Elements []int
}

func (s *Stack011) Push(key int) {
	s.Elements = append(s.Elements, key)
}

func (s *Stack011) Pop() int {
	front := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return front
}

func (s *Stack011) Empty() bool {
	return len(s.Elements) == 0
}

func (s *Stack011) Top() int {
	return s.Elements[len(s.Elements)-1]
}

func subArrayRanges(nums []int) int64 {
	stack := Stack011{}
	n := len(nums)
	prevSmaller := make([]int, n)
	nextSmaller := make([]int, n)

	prevGreater := make([]int, n)
	nextGreater := make([]int, n)

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

	stack = Stack011{}

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

	stack = Stack011{}

	for i := 0; i < n; i++ {
		if stack.Empty() {
			prevGreater[i] = -1
		} else {
			for !stack.Empty() && nums[stack.Top()] <= nums[i] {
				stack.Pop()
			}
			if !stack.Empty() {
				prevGreater[i] = stack.Top()
			} else {
				prevGreater[i] = -1
			}
		}
		stack.Push(i)
	}

	stack = Stack011{}

	for i := n - 1; i >= 0; i-- {
		if stack.Empty() {
			nextGreater[i] = n
		} else {
			for !stack.Empty() && nums[stack.Top()] < nums[i] {
				stack.Pop()
			}
			if !stack.Empty() {
				nextGreater[i] = stack.Top()
			} else {
				nextGreater[i] = n
			}
		}
		stack.Push(i)
	}

	var res int64

	for i := 0; i < n; i++ {
		minRes := int64(nums[i] * ((i - prevSmaller[i]) * (nextSmaller[i] - i)))
		maxRes := int64(nums[i] * ((i - prevGreater[i]) * (nextGreater[i] - i)))

		res += maxRes - minRes
	}

	return res

}
