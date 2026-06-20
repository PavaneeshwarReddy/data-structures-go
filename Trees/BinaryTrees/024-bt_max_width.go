package binarytrees

import "slices"

/*
Max Width
*/

func rowColMp(root *Node, row int, col int, mp *map[int][]int) {
	if root == nil {
		return
	}

	rowColMp(root.left, row+1, 2*col, mp)
	rowColMp(root.right, row+1, 2*col+1, mp)

	value, ok := (*mp)[row]
	if !ok {
		(*mp)[row] = []int{col}
	} else {
		(*mp)[row] = append(value, col)
	}
}

func (bt *BinaryTree) MaxWidth() int {
	mp := make(map[int][]int)
	rowColMp(bt.Root, 0, 0, &mp)
	maxWidth := 0

	for _, value := range mp {
		slices.Sort(value)
		maxWidth = max(maxWidth, value[len(value)-1]-value[0])
	}

	return maxWidth

}
