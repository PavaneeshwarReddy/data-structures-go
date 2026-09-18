package stacks

/*
Sum of Subarrays Minimums
- find the sum of min(b) where ranges over every subarray

- Here we need to think in such a way that
	- arr[i] is part of how many subarrays such that this is mini
	- We can simply perform arr[i] * length

- To find that we can simply use prev smaller element and next smaller element

- How this works, if curr is the smallest on the right then that means it should set to n which is the end

- Return the result as modulo

- first take int64 and perform the operations and then convert into normal
*/

type Stack009 struct {
	Elements []int
}

func (s *Stack009) Push(key int) {
	s.Elements = append(s.Elements, key)
}

func (s *Stack009) Pop() int {
	front := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return front
}

func (s *Stack009) Empty() bool {
	return len(s.Elements) == 0
}

func (s *Stack009) Top() int {
	return s.Elements[len(s.Elements)-1]
}

func sumSubarrayMins(arr []int) int {
	n := len(arr)
	nextSmaller := make([]int, n)
	prevSmaller := make([]int, n)
	stack := Stack009{}

	for i := 0; i < n; i++ {
		if stack.Empty() {
			prevSmaller[i] = -1
		} else {
			for !stack.Empty() && arr[stack.Top()] >= arr[i] {
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

	stack = Stack009{}

	for i := n - 1; i >= 0; i-- {
		if stack.Empty() {
			nextSmaller[i] = n
		} else {
			for !stack.Empty() && arr[stack.Top()] > arr[i] {
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

	const mod int64 = 1_000_000_007

	res := int64(0)

	for i := 0; i < n; i++ {
		left := int64(i - prevSmaller[i])
		right := int64(nextSmaller[i] - i)

		res = (res + int64(arr[i])*left*right) % mod
	}

	return int(res)
}
