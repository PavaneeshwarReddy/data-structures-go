package linkedlist

/*
Find Y intersection Node
Approach-1: Using Memory
Approach-2: Using a different Algo
*/

func (sll *SingleLinkedList) YIntersectionMemory(head1 *Node, head2 *Node) *Node {
	set1 := make(map[*Node]bool)
	set2 := make(map[*Node]bool)

	for head1 != nil && head2 != nil {
		set1[head1] = true
		set2[head2] = true

		if _, ok := set1[head2]; ok {
			return head2
		}

		if _, ok := set1[head1]; ok {
			return head1
		}
	}

	for head1 != nil {
		if _, ok := set1[head2]; ok {
			return head2
		}
	}

	for head2 != nil {
		if _, ok := set1[head1]; ok {
			return head1
		}
	}

	return nil

}

func (sll *SingleLinkedList) YIntersectionAlgo(head1 *Node, head2 *Node) *Node {
	head1Ptr := head1
	head2Ptr := head2

	for head1Ptr != nil || head2Ptr != nil {
		if head1Ptr == nil && head2Ptr != nil {
			head1Ptr = head2
		}

		if head2Ptr == nil && head1Ptr != nil {
			head2Ptr = head1
		}

		if head1Ptr == head2Ptr {
			return head1Ptr
		}

		head1Ptr = head1Ptr.next
		head2Ptr = head2Ptr.next

	}

	return nil
}
