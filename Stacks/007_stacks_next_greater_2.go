package stacks

/*
Find next greater element but here array is circular
- Still have a base condition that only to the right
- You need to make sure that array is doubled so that it automatically becomes circular.
- i%n will be your index
*/

type Stack007 struct {
	Elements []int
}

func (s *Stack007) Push(key int) {
	s.Elements = append(s.Elements, key)
}

func (s *Stack007) Pop() int {
	front := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return front
}

func (s *Stack007) Empty() bool {
	return len(s.Elements) == 0
}

func (s *Stack007) Top() int {
	return s.Elements[len(s.Elements)-1]
}

func nextGreaterElements2(nums []int) []int {
	stack := Stack007{}
	n := len(nums)
	results := make([]int, n)
	for i := range n {
		results[i] = -1
	}

	newN := 2 * n

	for i := newN - 1; i >= 0; i-- {
		if !stack.Empty() {
			for !stack.Empty() && stack.Top() <= nums[i%n] {
				stack.Pop()
			}
			if !stack.Empty() && i <= n-1 {
				results[i] = stack.Top()
			}
		}
		stack.Push(nums[i%n])
	}

	return results

}
