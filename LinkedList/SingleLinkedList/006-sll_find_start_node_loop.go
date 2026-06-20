package linkedlist

/*
Detect starting node of the loop
*/


func (sll *SingleLinkedList)DetechLoopNode() *Node {
	slow := sll.head
	fast := sll.head

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			temp := sll.head
			for temp != slow {
				temp = temp.next
				slow = slow.next
			}
			return temp
		}
	}
	return nil
}
