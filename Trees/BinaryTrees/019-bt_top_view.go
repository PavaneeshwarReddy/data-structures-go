package binarytrees

import "fmt"

/*
Top View
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
