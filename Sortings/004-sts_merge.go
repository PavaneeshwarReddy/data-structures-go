package sortings

/*
Merge Sort
	- There are mainly 2 operations
	- Split - Which splits the array recursively until it reaches len 1
	- Merge - Which merges two split arrays, by comparing and inserting
*/

func split(arr *[]int, start int, end int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	split(arr, start, mid)
	split(arr, mid+1, end)
	merge(arr, start, mid, mid+1, end)
}

func merge(arr *[]int, leftStart int, leftEnd int, rightStart int, rightEnd int) {
	newArr := []int{}
	i := leftStart
	j := rightStart
	for i <= leftEnd && j <= rightEnd {
		if (*arr)[i] < (*arr)[j] {
			newArr = append(newArr, (*arr)[i])
			i += 1
		} else {
			newArr = append(newArr, (*arr)[j])
			j += 1
		}
	}

	for i <= leftEnd {
		newArr = append(newArr, (*arr)[i])
		i += 1
	}

	for j <= rightEnd {
		newArr = append(newArr, (*arr)[j])
		j += 1
	}

	idx := leftStart

	for _, val := range newArr {
		(*arr)[idx] = val
		idx += 1
	}
}

func MergeSort(arr []int) []int {
	split(&arr, 0, len(arr)-1)
	return arr
}
