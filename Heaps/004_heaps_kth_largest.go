package heaps

/*
Kth largest element in array
- Using MaxHeap just pop k-1 times and element at the top is the kth largest element

- You can either choose minHeap with capacity k and maintain that
- You can choose maxHeap with n capacity and pop the values out
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
