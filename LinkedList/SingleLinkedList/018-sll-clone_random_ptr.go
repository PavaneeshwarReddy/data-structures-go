package linkedlist


func findPosFromStart(head *Node, search *Node) int {
	idx := 0
	currPtr := head
	for currPtr != search {
		currPtr = currPtr.next
		idx += 1
	}
	return idx
}

func getRandomPtrNode(head *Node, idx int) *Node {
	currPtr := head
	for idx > 0 {
		currPtr = currPtr.next
		idx -= 1
	}
	return currPtr
}

func copyRandomList(head *Node) *Node {

	dummyHead := &Node{value: -1, next: nil, random: nil}
	dummyPtr := dummyHead
	currPtr := head

	for currPtr != nil {
		dummyPtr.next = &Node{value: currPtr.value, next: nil, random: nil}
		dummyPtr = dummyPtr.next
		currPtr = currPtr.next
	}

	dummyPtr = dummyHead.next
	currPtr = head

	for currPtr != nil {
		if currPtr.random != nil {
			idx := findPosFromStart(head, currPtr.random)
			randomPtrNode := getRandomPtrNode(dummyHead.next, idx)
			dummyPtr.random = randomPtrNode
		}

		currPtr = currPtr.next
		dummyPtr = dummyPtr.next

	}

	return dummyHead.next

}
