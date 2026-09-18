package stacks

/*
Implement Stacks using linked list
Head: Maintains the top element in the stack if you use tail as top then you need to traverse until the end
Emtpy: If there is no head then empty

Push: Just we need to change the head
Pop: We need to remove head and point to next
*/

type Node struct {
	Val  int
	Next *Node
}

type StackList struct {
	Head *Node
}

func (sl *StackList) Push(key int) {
	newNode := Node{
		Val:  key,
		Next: sl.Head,
	}
	sl.Head = &newNode
}

func (sl *StackList) Pop() int {
	front := sl.Head.Val
	sl.Head = sl.Head.Next
	return front

}

func (sl *StackList) Top() int {
	return sl.Head.Val
}

func (sl *StackList) Empty() bool {
	return sl.Head == nil
}
