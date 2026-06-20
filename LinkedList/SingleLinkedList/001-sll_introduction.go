package linkedlist

import "fmt"

/*
This introduces to SingleLinkedList.
1. InsertValue
2. DeleteValue
3. FindValue
4. Length
5. PrintList
*/

type LinkedList interface {
	InsertValue(value int)
	DeleteValue(value int)
	FindValue(value int)
	Length()
	PrintList()
	MiddleLength()
}

type Node struct {
	value  int
	next   *Node
	random *Node
}

type SingleLinkedList struct {
	head *Node
	tail *Node
}

func (sll *SingleLinkedList) InsertValue(value int) {
	new_node := Node{value: value, next: nil}

	if sll.head == nil {
		sll.head = &new_node
		sll.tail = sll.head
	} else {
		sll.tail.next = &new_node
		sll.tail = sll.tail.next
	}
}

func (sll *SingleLinkedList) DeleteValue(value int) {
	var prev_node *Node
	current_node := sll.head

	for current_node != nil {
		if current_node.value == value {
			break
		}
		prev_node = current_node
		current_node = current_node.next
	}

	if current_node == sll.head && current_node == sll.tail {
		sll.head = nil
		sll.tail = nil
	} else if current_node == sll.head {
		sll.head = current_node.next
		current_node.next = nil
	} else if current_node == sll.tail {
		prev_node.next = current_node.next
		sll.tail = prev_node
	} else {
		prev_node.next = current_node.next
	}
}

func (sll *SingleLinkedList) FindValue(value int) bool {
	current_node := sll.head

	for current_node != nil {
		// (*current_node).value ==  current_node.value
		if current_node.value == value {
			return true
		}
		current_node = current_node.next
	}

	return false
}

func (sll *SingleLinkedList) Length() int {
	current_node := sll.head
	len := 0

	for current_node != nil {
		len += 1
		current_node = current_node.next
	}

	return len
}

func (sll *SingleLinkedList) PrintList() {
	current_node := sll.head

	for current_node != nil {
		fmt.Print(current_node.value, " ")
		current_node = current_node.next
	}
	fmt.Println()
}
