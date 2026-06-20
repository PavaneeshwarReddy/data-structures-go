package linkedlist

/*
Detect Loop
*/

func (sll *SingleLinkedList) DetectLoop() bool {
	slow := sll.head
	fast := sll.head

	for fast != nil && fast.next.next != nil {
		if fast == slow {
			return true
		}
		fast = fast.next.next
		slow = slow.next
	}

	return false
}
