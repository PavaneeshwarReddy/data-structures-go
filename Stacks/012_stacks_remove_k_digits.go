package stacks

/*
Remove k digits, You can remove k digits to make the number smaller
- Core idea is will my current result effect if i have the information of earlier results

- If you want to take this then you need to see whether you have greatest element in before so that you ignore it and pick this

- Loop through all elements, push element if empty,  if stack is not empty pop until curr element is greater than top and k > 0

- Sometimes in this case 1234 you cannot pop you just keep on inserting to avoid that issue after all the loop if k > 0. pop elements from back

- Sometimes after popping all you may be left with 002 or 00 then you need to pop leading zeroes
*/

type Stack012 struct {
	Elements []byte
}

func (s *Stack012) Push(key byte) {
	s.Elements = append(s.Elements, key)
}

func (s *Stack012) Pop() byte {
	front := s.Elements[len(s.Elements)-1]
	s.Elements = s.Elements[:len(s.Elements)-1]
	return front
}

func (s *Stack012) Empty() bool {
	return len(s.Elements) == 0
}

func (s *Stack012) Top() byte {
	return s.Elements[len(s.Elements)-1]
}

func removeKdigits(num string, k int) string {
	stack := Stack012{}
	n := len(num)

	if k == 0 {
		return num
	}

	if k == n {
		return "0"
	}

	for i := 0; i < n; i++ {

		for k > 0 && !stack.Empty() && stack.Top() > num[i] {
			stack.Pop()
			k--
		}

		stack.Push(num[i])
	}

	for k > 0 {
		stack.Pop()
		k--
	}

	i := 0
	for i < len(stack.Elements) {
		if stack.Elements[i] != '0' {
			break
		}
		i++
	}

	if i == len(stack.Elements) {
		return "0"
	}

	return string(stack.Elements[i:])
}
