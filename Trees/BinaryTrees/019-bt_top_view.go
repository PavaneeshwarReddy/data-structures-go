package binarytrees

import "fmt"

/*
Top View
- You just need to traverse to the left entirely and to the right entirely
*/

func leftTreeTraversal(root *Node) {
	if root == nil {
		return
	}

	fmt.Print(root.value, " ")
	leftTreeTraversal(root.left)
}

func rightTreeTraversal(root *Node) {
	if root == nil {
		return
	}

	fmt.Print(root.value, " ")
	rightTreeTraversal(root.right)
}

func (bt *BinaryTree) TopView() {
	leftTreeTraversal(bt.Root.left)
	fmt.Print(bt.Root.value, " ")
	rightTreeTraversal(bt.Root.right)
}
