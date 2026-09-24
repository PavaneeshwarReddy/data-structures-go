package binarytrees

import "math"

/*
Balanced Tree
- If depth at every point is not more than abs 1 then it is called balance
- To do that for every node, we need right max and left max depth and calculate diff to get the final result
- If found more than 1 then we can return false
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
