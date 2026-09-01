package heaps

/*
What does it mean by MinHeap restricted to size K
and if something crosses the length we pop first element
That means every time we insert we are making sure that smallest element is being removed
We are left with largest k numbers in the array
*/

type KthLargest struct {
	heap MinHeap
	k    int
}

func Constructor1(k int, nums []int) KthLargest {
	maxHeap := MinHeap{}
	for _, val := range nums {
		maxHeap.Insert(val)
	}
	for maxHeap.Length() > k {
		maxHeap.Pop()
	}
	return KthLargest{heap: maxHeap, k: k}
}

func (this *KthLargest) Add(val int) int {
	this.heap.Insert(val)
	if this.heap.Length() > this.k {
		this.heap.Pop()
	}
	return this.heap.Peek()
}
