package linkedlist

/*
Palindrome
*/


func (sll *SingleLinkedList) CheckPalindrome() bool {
	 slow := sll.head
	 fast := sll.head

	 for fast.next != nil && fast.next.next != nil {
		 slow = slow.next
		 fast = fast.next
	 }

	 head1 := sll.head
	 head2 := slow.next
	 slow.next=nil

     var prevPtr *Node
	 currPtr := head2

	 for currPtr != nil {
		 nextPtr := currPtr.next
		 currPtr.next = prevPtr
		 prevPtr = currPtr
		 currPtr = nextPtr
	 }

	 head2 = prevPtr

	 for head1 != nil && head2 != nil {
		 if head1.value != head2.value {
			 return false
		 }
		 head1 = head1.next
		 head2 = head2.next
	 }
	 return true
}



