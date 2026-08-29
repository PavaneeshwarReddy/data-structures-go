package main

import (
	heaps "dsa-go/Heaps"
	"fmt"
)

func main() {

	arr := []int{6, 5, 3, 2, 8, 10, 9}
	fmt.Println(heaps.SortKSortedArr(arr, 3))

}
