package binarytrees

/*

Check 2 Trees are similar

*/

func (bt *BinaryTree) CheckSimilarity(root1 *Node, root2 *Node) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if (root1 == nil && root2 != nil) || (root1 != nil && root2 == nil) || (root1.value != root2.value) {
		return false
	}

	return bt.CheckSimilarity(root1.left, root2.left) && bt.CheckSimilarity(root1.right, root2.right)
}
