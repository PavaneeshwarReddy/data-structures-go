package linkedlist

func middleList(head *Node) *Node {
	slow := head
	fast := head.next

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	return slow
}

func mergeList(leftList *Node, rightList *Node) *Node {
	dummyHead := &Node{value: -1, next: nil}
	dummyPtr := dummyHead

	for leftList != nil && rightList != nil {
		if leftList.value > rightList.value {
			dummyPtr.next = rightList
			rightList = rightList.next
		} else {
			dummyPtr.next = leftList
			leftList = leftList.next
		}
		dummyPtr = dummyPtr.next
	}

	for leftList != nil {
		dummyPtr.next = leftList
		leftList = leftList.next
		dummyPtr = dummyPtr.next
	}

	for rightList != nil {
		dummyPtr.next = rightList
		rightList = rightList.next
		dummyPtr = dummyPtr.next
	}

	return dummyHead.next

}

func splitList(head *Node) *Node {
	if head == nil || head.next == nil {
		return head
	}

	mid := middleList(head)

	leftList := head
	rightList := mid.next
	mid.next = nil

	return mergeList(splitList(leftList), splitList(rightList))

}

func (sll *SingleLinkedList) SortList() {
	sll.head = splitList(sll.head)
}
