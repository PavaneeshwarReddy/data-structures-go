package linkedlist

/*
Delete Middle Node
Similar to removing nth node from back
*/

func (sll *SingleLinkedList) DeleteMiddleNode() {

	l := 0
	currPtr := sll.head

	for currPtr != nil {
		l += 1
		currPtr = currPtr.next
	}

	hops := (l / 2) - 1

	currPtr = sll.head

	for hops > 0 {
		currPtr = currPtr.next
		hops -= 1
	}

	if l == 1 {
		sll.head = nil
	} else if currPtr == sll.head {
		if hops == 0 {
			sll.head.next = sll.head.next.next
		} else {
			sll.head = sll.head.next
		}
	} else {
		currPtr.next = currPtr.next.next
	}

}
