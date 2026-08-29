package heaps

/*
Kth largest element in array
- Using MaxHeap just pop k-1 times and element at the top is the kth largest element
*/

func FindKthLargest(arr []int, k int) int {
	maxHeap := MaxHeap{}
	for _, val := range arr {
		maxHeap.Insert(val)
	}

	for k > 1 {
		maxHeap.Pop()
		k -= 1
	}

	return maxHeap.Peek()
}
