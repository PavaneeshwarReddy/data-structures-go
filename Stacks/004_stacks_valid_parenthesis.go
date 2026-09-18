package stacks

/*
Valid Parenthesis
- First we need to create a stack
- If the incoming elements matches opp then pop the top or else push the element
*/
type Stack004 struct {
	Elements []byte
}

func (s *Stack004) Push(key byte) {
	s.Elements = append(s.Elements, key)
}

func (s *Stack004) Pop() byte {
	front := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return front
}

func (s *Stack004) Empty() bool {
	return len(s.Elements) == 0
}

func (s *Stack004) Top() byte {
	return s.Elements[len(s.Elements)-1]
}

func isValid(s string) bool {
	stack := Stack004{}
	mp := map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	for i := len(s) - 1; i >= 0; i-- {
		if stack.Empty() {
			stack.Push(s[i])
		} else {
			if mp[stack.Top()] == s[i] {
				stack.Pop()
			} else {
				stack.Push(s[i])
			}
		}
	}

	return stack.Empty()
}
