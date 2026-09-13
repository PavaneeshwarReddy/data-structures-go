package binarysearchtrees

/*
BST
- In BST we need not search entire tree for results
- We can move some constraints that BST has
*/

func searchBST(root *Node, val int) *Node {
	if root == nil {
		return root
	}

	if root.Val == val {
		return root
	}

	if root.Val > val {
		return searchBST(root.Left, val)
	}
	return searchBST(root.Right, val)

}
