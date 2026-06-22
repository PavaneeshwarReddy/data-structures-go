package sortings

/*
Quick Sort
	- Select a pivot element any random
	- Rearrange elements in such a way that, smaller elements on left and greater on right of pivot
	- Recursively apply the same for all left and right sub arrays
*/

func parition(arr *[]int, low int, high int) int {
	paritionKey := (*arr)[high]

	i := low - 1

	for j := low; j < high; j++ {
		if (*arr)[j] <= paritionKey {
			i++
			(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
		}
	}

	(*arr)[i+1], (*arr)[high] = (*arr)[high], (*arr)[i+1]

	return i + 1
}

func quickSort(arr *[]int, low int, high int) {
	if low < high {
		paritionIdx := parition(arr, low, high)

		quickSort(arr, low, paritionIdx-1)
		quickSort(arr, paritionIdx+1, high)
	}
}

func QuickSort(arr []int) []int {
	quickSort(&arr, 0, len(arr)-1)
	return arr
}
