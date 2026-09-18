package stacks

/*
Asteroid collisions
- Each asteroid has a sign and weight
- Sign positive ->, sign negative <-
- If greater negative collides then postive goes away, if same weight then both get destroyed

- Push ppsitive elements to the stack, if negative occurs check whether top is positive and check if greater -ve pop the top, if same pop and return alive false and don't include

*/

type Stack010 struct {
	Elements []int
}

func (s *Stack010) Push(key int) {
	s.Elements = append(s.Elements, key)
}

func (s *Stack010) Pop() int {
	front := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return front
}

func (s *Stack010) Empty() bool {
	return len(s.Elements) == 0
}

func (s *Stack010) Top() int {
	return s.Elements[len(s.Elements)-1]
}

func signCheck(x int, y int) bool {
	return x < 0 && y < 0
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func asteroidCollision(nums []int) []int {
	stack := Stack010{}
	n := len(nums)
	for i := 0; i < n; i++ {
		curr := nums[i]

		if curr > 0 {
			stack.Push(curr)
			continue
		}

		alive := true

		for alive && !stack.Empty() && stack.Top() > 0 {

			top := stack.Top()

			if top == abs(nums[i]) {
				stack.Pop()
				alive = false
			} else if top < abs(nums[i]) {
				stack.Pop()
			} else if top > abs(nums[i]) {
				alive = false
			}

		}

		if alive {
			stack.Push(nums[i])
		}
	}

	return stack.Elements
}
