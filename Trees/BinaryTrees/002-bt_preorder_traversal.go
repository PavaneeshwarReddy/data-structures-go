package binarytrees

import "fmt"

/*
Preorder traversal
Order: root ->  left -> right
*/

func (bt *BinaryTree) PreorderTraversal(root *Node) {
	if root == nil {
		return
	}

	fmt.Print(root.value, " ")
	bt.PreorderTraversal(root.left)
	bt.PreorderTraversal(root.right)
}
