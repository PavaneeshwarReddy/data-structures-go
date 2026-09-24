package binarytrees

/*
Diameter
- If we think about max depth, it gives either the max left or max right
- But for diameter to be max we need to add these 2 and search for a max result
*/

func diameterRecursive(root *Node, result *int) int {
	if root == nil {
		return 0
	}

	leftDepth := diameterRecursive(root.left, result)
	rightDepth := diameterRecursive(root.right, result)

	*result = max(leftDepth+rightDepth, *result)

	return max(leftDepth, rightDepth) + 1
}

func (bt *BinaryTree) Diameter() int {
	result := 0
	diameterRecursive(bt.Root, &result)
	return result
}
