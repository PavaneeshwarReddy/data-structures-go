package sortings

/*
Recursive Insertion Sort
	-
*/

func recursiveIS(arr *[]int, end int) {
	if end > len(*arr)-1 {
		return
	}

	key := (*arr)[end]
	j := end - 1

	for j >= 0 && (*arr)[j] > key {
		(*arr)[j+1] = (*arr)[j]
		j -= 1
	}

	(*arr)[j+1] = key

	recursiveIS(arr, end+1)
}

func RecursiveInsertionSort(arr []int) []int {
	recursiveIS(&arr, 1)
	return arr
}
