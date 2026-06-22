package sortings

/*
Insertion Sort
	- Consider first idx value is sorted
	- Now compare the value with before sorted array and place at correct idx
*/

func InsertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j -= 1
		}

		arr[j+1] = key
	}
	return arr
}
