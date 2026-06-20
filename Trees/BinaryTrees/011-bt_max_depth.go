package binarytrees

import (
	"math"
)

/*
Max Depth
*/

func (bt *BinaryTree) MaxDepth(root *Node, depth int) int {
	if root == nil {
		return depth
	}
	return int(math.Max(float64(bt.MaxDepth(root.left, depth+1)), float64(bt.MaxDepth(root.right, depth+1))))
}
