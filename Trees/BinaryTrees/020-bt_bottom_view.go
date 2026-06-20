package binarytrees

import "fmt"

/*
Bottom View
*/

func leafNodeTraversal(root *Node) {
	if root == nil {
		return
	}
	if root.left == nil && root.right == nil {
		fmt.Print(root.value, " ")
	}
	leafNodeTraversal(root.left)
	leafNodeTraversal(root.right)
}

func (bt *BinaryTree) BottomView() {
	leafNodeTraversal(bt.Root)
}
