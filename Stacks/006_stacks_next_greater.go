package stacks

/*
For every element find the next greater element on the right
- If we see there is some contraint that towards right
- Then we move from the right, if there are elements in the stack then
we remove until we found a greater element then current one or else -1
- Atlast we insert that element

*/

type Stack006 struct {
	Elements []int
}

func (s *Stack006) Push(key int) {
	s.Elements = append(s.Elements, key)
}

func (s *Stack006) Pop() int {
	front := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return front
}

func (s *Stack006) Empty() bool {
	return len(s.Elements) == 0
}

func (s *Stack006) Top() int {
	return s.Elements[len(s.Elements)-1]
}

func nextGreaterElements(nums []int) []int {
	stack := Stack006{}
	n := len(nums)
	results := make([]int, n)
	for i := range n {
		results[i] = -1
	}

	for i := n - 1; i >= 0; i-- {
		if !stack.Empty() {
			for !stack.Empty() && stack.Top() < nums[i] {
				stack.Pop()
			}
			if !stack.Empty() {
				results[i] = stack.Top()
			}
		}
		stack.Push(nums[i])
	}

	return results

}
