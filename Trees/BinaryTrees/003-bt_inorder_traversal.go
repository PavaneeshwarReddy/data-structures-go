package binarytrees
import "fmt"

/*
Inorder traversal
order: left -> root => right
*/


func (bt *BinaryTree) InorderTraversal(root *Node) {
	if root == nil {
		return 
	}

	bt.InorderTraversal(root.left)
	fmt.Print(root.value, " ")
	bt.InorderTraversal(root.right)
}

