package heaps

/*
Heap ( Priority Queue ):
- Complete Binary Tree - All levels are completely filled except last level
- Elements are processed with some priority instead of simply popping out elements

MinHeap:
- Parent node value is less than it's child value
- Insert(): Insert a key at the end, swap until it's parent larger then the current element
- Pop(): Replace root with last value and delete the last value, now heapify from the top
- Delete(): Same as pop, but before that we need to find which node we need to delete and replace that with last element
*/

type MinHeap struct {
	nodes []int
}

func (mh *MinHeap) Insert(key int) {
	mh.nodes = append(mh.nodes, key)

	idx := len(mh.nodes) - 1

	for idx > 0 {
		parentIdx := (idx - 1) / 2
		if parentIdx >= 0 && mh.nodes[parentIdx] > mh.nodes[idx] {
			mh.nodes[parentIdx], mh.nodes[idx] = mh.nodes[idx], mh.nodes[parentIdx]
		}
		idx = parentIdx
	}
}

func (mh *MinHeap) Pop() int {
	peekElement := mh.nodes[0]
	mh.nodes[0] = mh.nodes[len(mh.nodes)-1]
	mh.nodes = mh.nodes[:len(mh.nodes)-1]

	idx := 0

	for true {
		leftChild := 2*idx + 1
		rightChild := 2*idx + 2

		smallest := idx

		if leftChild < len(mh.nodes) && mh.nodes[smallest] > mh.nodes[leftChild] {
			smallest = leftChild
		}

		if rightChild < len(mh.nodes) && mh.nodes[smallest] > mh.nodes[rightChild] {
			smallest = rightChild
		}

		if smallest != idx {
			mh.nodes[idx], mh.nodes[smallest] = mh.nodes[smallest], mh.nodes[idx]
			idx = smallest
		} else {
			break
		}
	}

	return peekElement

}

func (mh *MinHeap) Peek() int {
	return mh.nodes[0]
}

func (mh *MinHeap) Length() int {
	return len(mh.nodes)
}
