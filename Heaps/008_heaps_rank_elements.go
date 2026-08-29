package heaps

/*
Rank Students  by their Score order
- [20 15 26 2 98 6] -> [4 3 5 1 6 2], the order which they appear when they are sorted
*/

type Student struct {
	Score int
	Idx   int
}

type StudentMinHeap struct {
	nodes []Student
}

func (cmh *StudentMinHeap) Insert(node Student) {
	cmh.nodes = append(cmh.nodes, node)

	idx := len(cmh.nodes) - 1

	for idx > 0 {
		parentIdx := (idx - 1) / 2

		if parentIdx >= 0 && cmh.nodes[parentIdx].Score > cmh.nodes[idx].Score {
			cmh.nodes[parentIdx].Score, cmh.nodes[idx].Score = cmh.nodes[idx].Score, cmh.nodes[parentIdx].Score
		}

		idx = parentIdx
	}
}

func (cmh *StudentMinHeap) Pop() Student {
	peekNode := cmh.nodes[0]
	cmh.nodes[0] = cmh.nodes[len(cmh.nodes)-1]
	cmh.nodes = cmh.nodes[:len(cmh.nodes)-1]

	idx := 0

	for true {
		leftChild := 2*idx + 1
		rightChild := 2*idx + 2
		smallest := idx

		if leftChild < len(cmh.nodes) && cmh.nodes[leftChild].Score < cmh.nodes[smallest].Score {
			smallest = leftChild
		}

		if rightChild < len(cmh.nodes) && cmh.nodes[rightChild].Score < cmh.nodes[smallest].Score {
			smallest = rightChild
		}

		if smallest != idx {
			cmh.nodes[smallest].Score, cmh.nodes[idx].Score = cmh.nodes[idx].Score, cmh.nodes[smallest].Score
			idx = smallest
		} else {
			break
		}
	}
	return peekNode
}

func (cmh *StudentMinHeap) Length() int {
	return len(cmh.nodes)
}

func RankStudentsByScore(scores []int) []int {
	minHeap := StudentMinHeap{}

	for idx, val := range scores {
		minHeap.Insert(Student{Score: val, Idx: idx})
	}
	idx := 1
	for minHeap.Length() > 0 {
		student := minHeap.Pop()
		scores[student.Idx] = idx
		idx++
	}
	return scores
}
