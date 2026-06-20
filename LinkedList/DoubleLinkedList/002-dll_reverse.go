package linkedlist

/*
Reverse
*/

func (dll *DoubleLinkedList) Reverse() {
	currPtr := dll.head
	for currPtr != nil {
		nextNode := currPtr.next
		currPtr.next = currPtr.prev
		currPtr.prev = nextNode
		currPtr = nextNode
	}
	dll.head, dll.tail = dll.tail, dll.head
}
