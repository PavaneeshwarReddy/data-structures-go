package queues

/*
Queues
Head and Tail are required, we push back using tail and pop front using head
*/

type Node struct {
	Next *Node
	Val  int
}

type QueueList struct {
	Head *Node
	Tail *Node
}

func (ql *QueueList) PushBack(key int) {
	newNode := Node{
		Val: key,
	}
	if ql.Head == nil {
		ql.Head = &newNode
		ql.Tail = &newNode
	} else {
		ql.Tail.Next = &newNode
		ql.Tail = ql.Tail.Next
	}
}

func (ql *QueueList) PopFront() int {
	front := ql.Head.Val
	if ql.Head == ql.Tail {
		ql.Head = nil
		ql.Tail = nil
	} else {
		ql.Head = ql.Head.Next
	}
	return front
}

func (ql *QueueList) Peek() int {
	return ql.Head.Val
}

func (ql *QueueList) Rear() int {
	return ql.Tail.Val
}

func (ql *QueueList) Empty() bool {
	return ql.Head == nil
}
