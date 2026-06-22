package sortings

/*
Selection Sort
	- Find the smallest element in the unsorted array and swap with the current idx
*/

func findSmallestIdx(arr []int, start int) int {
	smallIdx := start
	for i := start; i < len(arr); i++ {
		if arr[smallIdx] > arr[i] {
			smallIdx = i
		}
	}
	return smallIdx
}

func SelectionSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		smallIdx := findSmallestIdx(arr, i)
		arr[i], arr[smallIdx] = arr[smallIdx], arr[i]
	}
	return arr
}
