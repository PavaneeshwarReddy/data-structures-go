package linkedlist

/*
Seggregate odd and even nodes in a single linked list
1. First group odd nodes
2. First node is considered as odd node

*/

func (sll *SingleLinkedList) SeggregateOddEvenNodes() {
	currPtr := sll.head
	var evenHead, oddHead, evenCurrPtr, oddCurrPtr *Node
	c := 1

	for currPtr != nil {
		if c%2 == 0 {
			if evenHead == nil {
				evenHead = currPtr
				evenCurrPtr = evenHead
			} else {
				evenCurrPtr.next = currPtr
				evenCurrPtr = evenCurrPtr.next
			}
		} else {
			if oddHead == nil {
				oddHead = currPtr
				oddCurrPtr = oddHead
			} else {
				oddCurrPtr.next = currPtr
				oddCurrPtr = oddCurrPtr.next
			}
		}

		nextPtr := currPtr.next
		currPtr.next = nil
		currPtr = nextPtr
		c += 1
	}

	if oddHead != nil {
		oddCurrPtr.next = evenHead
	}

}
