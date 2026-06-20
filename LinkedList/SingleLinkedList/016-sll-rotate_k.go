package linkedlist

func (sll *SingleLinkedList) RotateByK(k int) {
	l := sll.Length()
	if l == 0 {
		return
	}
	k = k % l
	for k > 0 {
		tailPtr := sll.head
		for tailPtr.next != nil && tailPtr.next.next != nil {
			tailPtr = tailPtr.next
		}
		if tailPtr.next != nil {
			tailNode := tailPtr.next
			tailPtr.next = nil
			tailNode.next = sll.head
			sll.head = tailNode
		}
		k -= 1
	}
}
