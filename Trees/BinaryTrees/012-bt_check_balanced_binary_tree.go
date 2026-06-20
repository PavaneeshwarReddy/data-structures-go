package binarytrees

import "math"

/*
Balanced Tree
*/

func balancedTree(root *Node, result *bool) int {
	if root == nil {
		return 0
	}
	leftDepth := balancedTree(root.left, result)
	rightDepth := balancedTree(root.right, result)

	if math.Abs(float64(leftDepth-rightDepth)) > 1 {
		*result = false
	}
	return max(leftDepth, rightDepth) + 1
}

func (bt *BinaryTree) CheckBalancedTree() bool {
	result := true
	balancedTree(bt.Root, &result)
	return result
}
