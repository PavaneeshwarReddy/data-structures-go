package binarysearchtrees

/*
Validate Binary Search Tree
- We start with -infinite, +infinite
- For each node we go below,
	- Left: (min, root.Val)
	- Right: (root.Val, max)
*/

import "math"

func inorder005(root *Node, mi, ma int) bool {
	if root == nil {
		return true
	}

	if root.Val <= mi || root.Val >= ma {
		return false
	}

	return inorder005(root.Left, mi, root.Val) && inorder005(root.Right, root.Val, ma)
}
func isValidBST(root *Node) bool {
	return inorder005(root, math.MinInt, math.MaxInt)
}
