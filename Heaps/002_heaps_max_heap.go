package heaps

/*
Heap ( Priority Queue ):
- Complete Binary Tree - All levels are completely filled except last level
- Elements are processed with some priority instead of simply popping out elements

MaxHeap:
- Parent node value is greater than it's child value
- Insert(): Insert a key at the end, swap until it's parent smaller then the current element O(logN)
- Pop(): Replace root with last value and delete the last value, now heapify from the top O(logN)
- Delete(): Same as pop, but before that we need to find which node we need to delete and replace that with last element O(logN)
- Peek(): return the top element, O(1)
*/

type MaxHeap struct {
	nodes []int
}

func (mh *MaxHeap) Insert(key int) {
	mh.nodes = append(mh.nodes, key)

	idx := len(mh.nodes) - 1

	for idx > 0 {
		parentIdx := (idx - 1) / 2

		if parentIdx >= 0 && mh.nodes[parentIdx] < mh.nodes[idx] {
			mh.nodes[parentIdx], mh.nodes[idx] = mh.nodes[idx], mh.nodes[parentIdx]
		}
		idx = parentIdx
	}
}

func (mh *MaxHeap) Pop() int {
	peekElement := mh.nodes[0]
	mh.nodes[0] = mh.nodes[len(mh.nodes)-1]
	mh.nodes = mh.nodes[:len(mh.nodes)-1]

	idx := 0

	for true {
		leftChild := 2*idx + 1
		rightChild := 2*idx + 2
		largest := idx

		if leftChild < len(mh.nodes) && mh.nodes[leftChild] > mh.nodes[largest] {
			largest = leftChild
		}

		if rightChild < len(mh.nodes) && mh.nodes[rightChild] > mh.nodes[largest] {
			largest = rightChild
		}

		if largest != idx {
			mh.nodes[largest], mh.nodes[idx] = mh.nodes[idx], mh.nodes[largest]
			idx = largest
		} else {
			break
		}

	}

	return peekElement
}

func (mh *MaxHeap) Peek() int {
	return mh.nodes[0]
}

func (mh *MaxHeap) Length() int {
	return len(mh.nodes)
}
