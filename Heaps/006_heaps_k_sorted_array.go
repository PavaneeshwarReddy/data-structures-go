package heaps

/*
Sort K Sorted Array
- Every element in the array is at max at a distance of k from it's sorted position
- Meaning it can be either in between i-k to i+k inclusive

Why to use minHeap ?
- If we take a minHeap of first k elements then the first element can only be in that range (0 -> k )
- If we move the window slowly by popping out and inserting in then we are able to sort the array
- Normal sort O(nlogn) whereas this O(nlogk)
*/

func SortKSortedArr(arr []int, k int) []int {
	result := []int{}

	minHeap := MinHeap{}

	for i := 0; i <= k; i++ {
		minHeap.Insert(arr[i])
	}

	for i := k + 1; i < len(arr); i++ {
		result = append(result, minHeap.Pop())
		minHeap.Insert(arr[i])
	}

	for minHeap.Length() > 0 {
		result = append(result, minHeap.Pop())
	}

	return result

}
