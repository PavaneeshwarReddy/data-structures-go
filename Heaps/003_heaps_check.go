package heaps

/*
Check if an array is min heap or not
- All the parent idices should be less than their left and right child
*/

func CheckMinHeap(arr []int) bool {
	for idx, val := range arr {
		leftChild := 2*idx + 1
		rightChild := 2*idx + 2

		if leftChild < len(arr) && val > arr[leftChild] {
			return false
		}

		if rightChild < len(arr) && val > arr[rightChild] {
			return false
		}
	}
	return true
}
