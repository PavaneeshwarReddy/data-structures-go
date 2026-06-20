package linkedlist

/*
Add two lists
*/

func getValueCarry(val1 int, val2 int, c int) (val int, cf int) {
	value := val1 + val2 + c
	c = value / 10
	value = value % 10
	return value, c
}

func (sll *SingleLinkedList) addTwoNumbers(l1 *Node, l2 *Node) *Node {
	var result *Node
	var resultPtr *Node
	l1Ptr := l1
	l2Ptr := l2
	carryForward := 0
	value := 0

	for l1Ptr != nil && l2Ptr != nil {
		value, carryForward = getValueCarry(l1Ptr.value, l2Ptr.value, carryForward)

		newNode := &Node{value: value, next: nil}

		if result == nil {
			result = newNode
			resultPtr = result
		} else {
			resultPtr.next = newNode
			resultPtr = resultPtr.next
		}
		l1Ptr = l1Ptr.next
		l2Ptr = l2Ptr.next
	}

	for l1Ptr != nil {
		value, carryForward = getValueCarry(l1Ptr.value, 0, carryForward)
		newNode := &Node{value: value, next: nil}
		if result == nil {
			result = newNode
			resultPtr = result
		} else {
			resultPtr.next = newNode
			resultPtr = resultPtr.next
		}
		l1Ptr = l1Ptr.next
	}

	for l2Ptr != nil {
		value, carryForward = getValueCarry(0, l2Ptr.value, carryForward)
		newNode := &Node{value: value, next: nil}
		if result == nil {
			result = newNode
			resultPtr = result
		} else {
			resultPtr.next = newNode
			resultPtr = resultPtr.next
		}
		l2Ptr = l2Ptr.next
	}

	if carryForward != 0 {
		newNode := &Node{value: carryForward, next: nil}
		resultPtr.next = newNode
	}

	return result

}
