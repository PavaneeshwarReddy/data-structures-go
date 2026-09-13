package graphs

import "math"

/*
Dijkstra's Algorithm
- We may think that we can simply use BFS for shortest distance but that doesn't work, that does work only if all the nodes have same weight


- We maintain 3 components mainly
	- Parent Array: Who is the parent of the nodes, this is used to constuct the path, as we move from destination to back back parent to find the last source
	- Dist Array: Which stores the minimum distance
	- MinHeap: Which is picked based on distance
*/

type GraphNode struct {
	Vertex int
	Dist   int
}

type MinHeapCustom struct {
	Elements []GraphNode
}

func (mh *MinHeapCustom) Insert(key GraphNode) {
	mh.Elements = append(mh.Elements, key)
	idx := len(mh.Elements) - 1

	for idx >= 0 {
		parentIdx := (idx - 1) / 2
		if parentIdx >= 0 && mh.Elements[parentIdx].Dist > mh.Elements[idx].Dist {
			mh.Elements[idx], mh.Elements[parentIdx] = mh.Elements[parentIdx], mh.Elements[idx]
		}
		idx = parentIdx
	}
}

func (mh *MinHeapCustom) Peek() GraphNode {
	return mh.Elements[0]
}

func (mh *MinHeapCustom) Length() int {
	return len(mh.Elements)
}

func (mh *MinHeapCustom) Pop() GraphNode {
	top := mh.Elements[0]
	mh.Elements[0] = mh.Elements[len(mh.Elements)-1]
	mh.Elements = mh.Elements[len(mh.Elements)-1:]

	idx := 0

	for true {
		leftChild := 2*idx + 1
		rightCHild := 2*idx + 2
		smallest := idx

		if leftChild < len(mh.Elements) && mh.Elements[leftChild].Dist < mh.Elements[smallest].Dist {
			smallest = leftChild
		}

		if rightCHild < len(mh.Elements) && mh.Elements[rightCHild].Dist < mh.Elements[smallest].Dist {
			smallest = rightCHild
		}

		if smallest != idx {
			mh.Elements[smallest], mh.Elements[idx] = mh.Elements[idx], mh.Elements[smallest]
		} else {
			break
		}
	}

	return top

}

func DijkstraAlgo(start int, end int, n int, adj map[int][][]int) (int, []int) {

	parentArr := make([]int, n)
	distArr := make([]int, n)

	for i := range n {
		parentArr[i] = i
		distArr[i] = math.MaxInt
	}

	minHeap := MinHeapCustom{}
	minHeap.Insert(GraphNode{Vertex: start, Dist: 0})
	distArr[start] = 0

	for minHeap.Length() > 0 {
		currNode := minHeap.Pop()

		for _, node := range adj[currNode.Vertex] {
			if distArr[node[0]] > distArr[currNode.Vertex]+node[1] {
				distArr[node[0]] = distArr[currNode.Vertex] + node[1]
				parentArr[node[0]] = currNode.Vertex
				minHeap.Insert(GraphNode{Vertex: node[0], Dist: distArr[node[0]]})
			}
		}
	}

	shortestDis := distArr[end]

	path := []int{}

	idx := end

	for idx >= 0 {
		path = append(path, parentArr[idx])
		if idx == parentArr[idx] {
			break
		}
		idx = parentArr[idx]
	}

	i := 0
	j := len(path) - 1
	for i <= j {
		path[i], path[j] = path[j], path[i]
		i++
		j--
	}

	return shortestDis, path

}
