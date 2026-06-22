package sortings

/*
Recursive Bubble Sort
	- It is almost similar to iterative approach
	- Outer loop is replaced with recursive
*/

func recursiveBS(arr *[]int, end int) {
	if end == 0 {
		return
	}

	for i := 0; i < end-1; i++ {
		if (*arr)[i] > (*arr)[i+1] {
			(*arr)[i], (*arr)[i+1] = (*arr)[i+1], (*arr)[i]
		}
	}

	recursiveBS(arr, end-1)

}

func RecursiveBubbleSort(arr []int) []int {
	recursiveBS(&arr, len(arr))
	return arr
}
