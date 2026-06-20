package linkedlist


/*
Reverse of a single linked list
*/


func (sll *SingleLinkedList) Reverse() {
	var prevPtr *Node
	currPtr := sll.head

	for currPtr != nil {
		nextPtr := currPtr.next
		currPtr.next = prevPtr
		prevPtr = currPtr
		currPtr = nextPtr
	}

	sll.tail = sll.head
	sll.head = prevPtr
}
