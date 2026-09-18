package stacks

import "errors"

/*
Stacks
- Push: Which pushes element at the top
- Top: If there are any elements it returns the peek or else error
- Pop: Removes and returns the top most element
- Length: Returns length of the stack

ErrStackEmpty: If stack is empty an error is thrown
*/

var ErrStackEmpty = errors.New("Stack is empty")

type Stack struct {
	Elements []int
}

func (s *Stack) Insert(val int) {
	s.Elements = append(s.Elements, val)
}

func (s *Stack) Top() (int, error) {
	if len(s.Elements) == 0 {
		return -1, ErrStackEmpty
	}
	return s.Elements[len(s.Elements)-1], nil
}

func (s *Stack) Pop() (int, error) {
	l := len(s.Elements)
	if l == 0 {
		return -1, ErrStackEmpty
	}
	topElement := s.Elements[l-1]
	s.Elements = s.Elements[:l-1]
	return topElement, nil
}
