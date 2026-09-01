package heaps

/*
Find kth largest frequently occuring elements in array
*/

type Element struct {
	Num   int
	Count int
}

type KthMaxHeap struct {
	nodes []Element
}

func (mh *KthMaxHeap) Insert(key Element) {
	mh.nodes = append(mh.nodes, key)

	idx := len(mh.nodes) - 1

	for idx > 0 {
		parentIdx := (idx - 1) / 2

		if parentIdx >= 0 && mh.nodes[parentIdx].Count < mh.nodes[idx].Count {
			mh.nodes[parentIdx], mh.nodes[idx] = mh.nodes[idx], mh.nodes[parentIdx]
		}
		idx = parentIdx
	}
}

func (mh *KthMaxHeap) Pop() Element {
	peekElement := mh.nodes[0]
	mh.nodes[0] = mh.nodes[len(mh.nodes)-1]
	mh.nodes = mh.nodes[:len(mh.nodes)-1]

	idx := 0

	for true {
		leftChild := 2*idx + 1
		rightChild := 2*idx + 2
		largest := idx

		if leftChild < len(mh.nodes) && mh.nodes[leftChild].Count > mh.nodes[largest].Count {
			largest = leftChild
		}

		if rightChild < len(mh.nodes) && mh.nodes[rightChild].Count > mh.nodes[largest].Count {
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

func (mh *KthMaxHeap) Peek() Element {
	return mh.nodes[0]
}

func (mh *KthMaxHeap) Length() int {
	return len(mh.nodes)
}

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, val := range nums {
		count, ok := freqMap[val]
		if !ok {
			count = 0
		}
		count += 1
		freqMap[val] = count
	}

	maxHeap := KthMaxHeap{}

	for k, v := range freqMap {
		maxHeap.Insert(Element{Num: k, Count: v})
	}

	result := []int{}
	for k > 0 {
		result = append(result, maxHeap.Pop().Num)
		k -= 1
	}

	return result
}
