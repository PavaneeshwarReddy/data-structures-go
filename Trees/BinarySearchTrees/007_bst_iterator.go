package binarysearchtrees

/*
BST Iterator
- This can give you what's the next value is
- We have put all left nodes into the stack
- when we are popping out we need to push all to the just right and to the left all ( I mean right child and it's left )
- As we know the inorder traversal gives us in the sorting order which is main logic behind iterator
*/

type BSTIterator struct {
	Stack []*Node
}

func Constructor(root *Node) BSTIterator {
	curr := root
	stack := []*Node{}
	for curr != nil {
		stack = append(stack, curr)
		curr = curr.Left
	}
	return BSTIterator{Stack: stack}
}

func (this *BSTIterator) Next() int {
	top := this.Stack[len(this.Stack)-1]
	this.Stack = this.Stack[:len(this.Stack)-1]

	curr := top.Right
	for curr != nil {
		this.Stack = append(this.Stack, curr)
		curr = curr.Left
	}

	return top.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.Stack) > 0
}
