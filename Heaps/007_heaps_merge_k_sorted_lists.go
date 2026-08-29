package heaps

/*
Merge K Sorted Linked Lists

Approach - 1:
- Collect elements from all linked lists, put them in an array sort them
- Create a new linked lists for each sorted element and return

Approach - 2:
- Similar to merge sort
- Take 2 pairs process merging between them

Approach - 3: ( Custom MinHeap implementation )
- Consider a min Heap ( int, LinkedList )
- Keep popping values and mapping them without using an extra space
*/

type Node struct {
	Prev *Node
	Next *Node
	Val  int
}

type CustomMinHeap struct {
	nodes []*Node
}

func (cmh *CustomMinHeap) Insert(node *Node) {
	cmh.nodes = append(cmh.nodes, node)

	idx := len(cmh.nodes) - 1

	for idx > 0 {
		parentIdx := (idx - 1) / 2

		if parentIdx >= 0 && cmh.nodes[parentIdx].Val > cmh.nodes[idx].Val {
			cmh.nodes[parentIdx].Val, cmh.nodes[idx].Val = cmh.nodes[idx].Val, cmh.nodes[parentIdx].Val
		}

		idx = parentIdx
	}
}

func (cmh *CustomMinHeap) Pop() *Node {
	peekNode := cmh.nodes[0]
	cmh.nodes[0] = cmh.nodes[len(cmh.nodes)-1]
	cmh.nodes = cmh.nodes[:len(cmh.nodes)-1]

	idx := 0

	for true {
		leftChild := 2*idx + 1
		rightChild := 2*idx + 2
		smallest := idx

		if leftChild < len(cmh.nodes) && cmh.nodes[leftChild].Val < cmh.nodes[smallest].Val {
			smallest = leftChild
		}

		if rightChild < len(cmh.nodes) && cmh.nodes[rightChild].Val < cmh.nodes[smallest].Val {
			smallest = rightChild
		}

		if smallest != idx {
			cmh.nodes[smallest].Val, cmh.nodes[idx].Val = cmh.nodes[idx].Val, cmh.nodes[smallest].Val
			idx = smallest
		} else {
			break
		}
	}
	return peekNode
}

func (cmh *CustomMinHeap) Length() int {
	return len(cmh.nodes)
}

func MergeKSortedists(lists []*Node) *Node {
	customMinHeap := CustomMinHeap{}

	for _, val := range lists {
		temp := val
		for temp != nil {
			nxtTemp := temp.Next
			temp.Next = nil
			customMinHeap.Insert(temp)
			temp = nxtTemp
		}
	}

	result := Node{Val: -1}
	temp := &result

	for customMinHeap.Length() > 0 {
		temp.Next = customMinHeap.Pop()
		temp = temp.Next
	}

	return result.Next
}
