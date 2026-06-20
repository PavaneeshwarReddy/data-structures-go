package binarytrees
import "fmt"

/*
Postorder traversal
order: left -> right -> root
*/


func (bt *BinaryTree) PostorderTraversal(root *Node) {
	if root == nil {
		return
	}

	bt.PostorderTraversal(root.left)
	bt.PostorderTraversal(root.right)
	fmt.Print(root.value, " ")
}
