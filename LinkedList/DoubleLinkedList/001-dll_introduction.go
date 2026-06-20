package linkedlist

import "fmt"

/*
InsertValue
PrintList
Length
Search
DeleteValue
*/

type Node struct {
	prev  *Node
	value int
	next  *Node
}

type DoubleLinkedList struct {
	head *Node
	tail *Node
}

type DoubleLinkedListInterface interface {
	InsertValue(val int)
	PrintList()
	Length() int
	DeleteValue()
	Search() *Node
	Reverse()
}

func (dll *DoubleLinkedList) InsertValue(val int) {
	newNode := &Node{value: val}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.prev = dll.tail
		dll.tail.next = newNode
		dll.tail = dll.tail.next
	}
}

func (dll *DoubleLinkedList) PrintList() {
	currPtr := dll.head
	for currPtr != nil {
		fmt.Print(currPtr.value, " ")
		currPtr = currPtr.next
	}
	fmt.Println()
}

func (dll *DoubleLinkedList) Length() int {
	currPtr := dll.head
	l := 0
	for currPtr != nil {
		l += 1
		currPtr = currPtr.next
	}
	return l
}

func (dll *DoubleLinkedList) Search(key int) *Node {
	currPtr := dll.head
	for currPtr != nil {
		if currPtr.value == key {
			return currPtr
		}
		currPtr = currPtr.next
	}
	return nil
}

func (dll *DoubleLinkedList) DeleteValue(key int) {
	searchNode := dll.Search(key)
	if searchNode == nil {
		return
	}
	if searchNode == dll.head && dll.head == dll.tail {
		dll.head = nil
		dll.tail = nil
	} else if searchNode == dll.head {
		nextNode := dll.head.next
		dll.head.next = nil
		nextNode.prev = nil
		dll.head = nextNode
	} else if searchNode == dll.tail {
		prevNode := dll.tail.prev
		dll.tail.prev = nil
		prevNode.next = nil
		dll.tail = prevNode
	} else {
		searchNode.prev.next = searchNode.next
		searchNode.next.prev = searchNode.prev
	}
}
