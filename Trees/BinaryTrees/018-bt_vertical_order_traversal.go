package binarytrees

import "slices"

/*
Vertical Order Traversal
*/

type NodeInfo struct {
	row   int
	value int
	col   int
}

func verticalRecursive(root *Node, info *[]NodeInfo, row int, col int) {
	if root == nil {
		return
	}
	verticalRecursive(root.left, info, row+1, col-1)
	*info = append(*info, NodeInfo{row: row, value: root.value, col: col})
	verticalRecursive(root.right, info, row+1, col+1)
}

func (bt *BinaryTree) VerticalOrderTraversal() [][]int {
	info := []NodeInfo{}
	result := [][]int{}

	verticalRecursive(bt.Root, &info, 0, 0)

	slices.SortFunc(info, func(a, b NodeInfo) int {
		if a.col != b.col {
			return a.col - b.col
		}
		if a.row != b.row {
			return a.row - b.row
		}
		return a.value - b.value
	})

	currCol := info[0].col
	tempResult := []int{}

	for _, v := range info {
		if currCol != v.col {
			if len(tempResult) > 0 {
				result = append(result, tempResult)
			}
			currCol = v.col
			tempResult = []int{v.value}
		} else {
			tempResult = append(tempResult, v.value)
		}
	}

	if len(tempResult) > 0 {
		result = append(result, tempResult)
	}

	return result
}
