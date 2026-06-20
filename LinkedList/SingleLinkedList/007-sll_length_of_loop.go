package linkedlist

/*
Length of the loop
*/

func (sll *SingleLinkedList) LengthOfTheLoop() int {
	slow := sll.head
	fast := sll.head
	length := 0

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
		if slow == fast {
			temp := sll.head
			for temp != slow {
				temp = temp.next
				slow = slow.next
			}
			
			tempPtr  := slow.next
			length = 1
			for tempPtr != slow {
				length+=1
				tempPtr  =tempPtr.next
				if slow == tempPtr {
					return length
				}
			}

		}
	}
	return 0
}
