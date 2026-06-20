package linkedlist


/*
Middle Node of the Single Linked List
Approaches:

Second Middle for even nodes, exact Middle for odd nodes
1. Using Length
2. Using Tortoise and Hare
*/


func (sll *SingleLinkedList) MiddleLength() int {
	l := 0
	tempPtr := sll.head

	for tempPtr != nil {
		l+=1
		tempPtr=tempPtr.next
	}

	tempPtr = sll.head
	l = l/2	

	for l>0 {
		tempPtr = tempPtr.next
		l-=1
	}

	return tempPtr.value
}




func (sll *SingleLinkedList) MiddleTortoiseHare() int {
	slow := sll.head
	fast := sll.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow.value
}
