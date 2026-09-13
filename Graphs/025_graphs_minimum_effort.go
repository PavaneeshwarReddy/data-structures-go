package graphs

/*
Minimum Effort Path
- There are two things, max abs difference will become the effort and that should be minimum

*/

import "math"

type GraphNode025 struct {
	Sr     int
	Sc     int
	Height int
}

type MinHeapCustom025 struct {
	Elements []GraphNode025
}

func (mh *MinHeapCustom025) Insert(key GraphNode025) {
	mh.Elements = append(mh.Elements, key)
	idx := len(mh.Elements) - 1

	for idx > 0 {
		parentIdx := (idx - 1) / 2
		if parentIdx >= 0 && mh.Elements[parentIdx].Height > mh.Elements[idx].Height {
			mh.Elements[idx], mh.Elements[parentIdx] = mh.Elements[parentIdx], mh.Elements[idx]
		}
		idx = parentIdx
	}
}

func (mh *MinHeapCustom025) Peek() GraphNode025 {
	return mh.Elements[0]
}

func (mh *MinHeapCustom025) Length() int {
	return len(mh.Elements)
}

func (mh *MinHeapCustom025) Pop() GraphNode025 {
	top := mh.Elements[0]
	mh.Elements[0] = mh.Elements[len(mh.Elements)-1]
	mh.Elements = mh.Elements[:len(mh.Elements)-1]

	idx := 0

	for true {
		leftChild := 2*idx + 1
		rightCHild := 2*idx + 2
		smallest := idx

		if leftChild < len(mh.Elements) && mh.Elements[leftChild].Height < mh.Elements[smallest].Height {
			smallest = leftChild
		}

		if rightCHild < len(mh.Elements) && mh.Elements[rightCHild].Height < mh.Elements[smallest].Height {
			smallest = rightCHild
		}

		if smallest != idx {
			mh.Elements[smallest], mh.Elements[idx] = mh.Elements[idx], mh.Elements[smallest]
			idx = smallest
		} else {
			break
		}
	}

	return top

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minimumEffortPath(heights [][]int) int {
	m := len(heights)
	n := len(heights[0])

	dist := make([][]int, m)
	for i := range m {
		dist[i] = make([]int, n)
		for j := range n {
			dist[i][j] = math.MaxInt
		}
	}

	minHeap := MinHeapCustom025{}
	minHeap.Insert(GraphNode025{Sr: 0, Sc: 0, Height: 0})
	dist[0][0] = 0

	for minHeap.Length() > 0 {
		top := minHeap.Pop()
		dirs := [][2]int{
			{1, 0},
			{0, 1},
			{-1, 0},
			{0, -1},
		}

		if top.Sr == m-1 && top.Sc == n-1 {
			return top.Height
		}

		for _, dir := range dirs {
			i := dir[0] + top.Sr
			j := dir[1] + top.Sc
			if i < m && i >= 0 && j < n && j >= 0 {
				edgeEffort := abs(heights[top.Sr][top.Sc] - heights[i][j])
				newEffort := max(edgeEffort, top.Height)
				if dist[i][j] > newEffort {
					dist[i][j] = newEffort
					minHeap.Insert(GraphNode025{Sr: i, Sc: j, Height: newEffort})
				}
			}
		}
	}

	return dist[m-1][n-1]

}
