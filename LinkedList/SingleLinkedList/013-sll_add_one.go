package linkedlist

/*
Add one
*/

func (sll *SingleLinkedList) AddOne() {

	sll.Reverse()

	carryForward := 1
	currPtr := sll.head

	for currPtr != nil {
		value := currPtr.value + carryForward
		currPtr.value = value % 10
		carryForward = value / 10
		currPtr = currPtr.next
	}

	if carryForward != 0 {
		sll.tail.next = &Node{value: carryForward, next: nil}
		sll.tail = sll.tail.next
	}

	sll.Reverse()

}
