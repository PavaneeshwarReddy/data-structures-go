package linkedlist

/*
Remove Nth Node from back
*/

func (sll *SingleLinkedList) RemoveNthNodeFromback(n int) {

	l := 0
	currPtr := sll.head

	for currPtr != nil {
		l += 1
		currPtr = currPtr.next
	}

	hops := l - n - 1

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
