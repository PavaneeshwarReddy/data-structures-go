package main

import (
	sortings "dsa-go/Sortings"
	"fmt"
)

func main() {

	arr := []int{1, 2, -3, 100}
	sortings.QuickSort(arr)

	fmt.Printf("%v", arr)

}
