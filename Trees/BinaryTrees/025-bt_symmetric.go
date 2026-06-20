package binarytrees

/*
Symmetric
*/

func symmetric(leftRoot *Node, rightRoot *Node) bool {
	if leftRoot == nil && rightRoot == nil {
		return true
	}

	if (leftRoot == nil && rightRoot != nil) || (leftRoot != nil && rightRoot == nil) || (leftRoot.value != rightRoot.value) {
		return false
	}

	return symmetric(leftRoot.left, rightRoot.right) && symmetric(leftRoot.right, rightRoot.left)
}

func (bt *BinaryTree) Symmetric() bool {
	return symmetric(bt.Root.left, bt.Root.right)
}
