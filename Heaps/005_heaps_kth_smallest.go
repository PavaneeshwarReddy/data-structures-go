package heaps

/*
Kth smallest in an array
- Add elements in the min heap and pop k-1 times and return the peek
*/

func FindKthSmallest(arr []int, k int) int {
	minHeap := MinHeap{}

	for _, val := range arr {
		minHeap.Insert(val)
	}

	for k > 1 {
		minHeap.Pop()
		k -= 1
	}

	return minHeap.Peek()
}
