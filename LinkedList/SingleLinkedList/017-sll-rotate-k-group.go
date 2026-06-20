package linkedlist

func reverseList(head *Node) *Node {
	var prevPtr *Node
	currPtr := head

	for currPtr != nil {
		nextPtr := currPtr.next
		currPtr.next = prevPtr
		prevPtr = currPtr
		currPtr = nextPtr
	}

	return prevPtr
}

func isValidKGroup(head *Node, k int) (ok bool, lastNode *Node) {
	lastNode = head
	k -= 1
	for lastNode != nil && k > 0 {
		lastNode = lastNode.next
		k -= 1
	}
	if k == 0 && lastNode != nil {
		return true, lastNode
	}
	return false, nil
}

func reverseKGroup(head *Node, k int) *Node {
	ok, lastNode := isValidKGroup(head, k)

	if !ok {
		return head
	}

	nextList := lastNode.next
	lastNode.next = nil

	newHead := reverseList(head)
	head.next = reverseKGroup(nextList, k)

	return newHead
}
