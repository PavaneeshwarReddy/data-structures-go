package binarytrees

import "fmt"

/*
Boundary Traversal
*/

func leftTraversal(root *Node, result *[]int) {
	if root == nil || (root.left == nil && root.right == nil) {
		return
	}

	*result = append(*result, root.value)
	if root.left == nil {
		leftTraversal(root.right, result)
	} else {
		leftTraversal(root.left, result)
	}
}

func rightTraversal(root *Node, result *[]int) {
	if root == nil || (root.left == nil && root.right == nil) {
		return
	}

	if root.right == nil {
		rightTraversal(root.left, result)
	} else {
		rightTraversal(root.right, result)
	}
	*result = append(*result, root.value)
}

func inorderTraversal(root *Node, result *[]int) {
	if root == nil {
		return
	}

	inorderTraversal(root.left, result)

	if root.left == nil && root.right == nil {
		*result = append(*result, root.value)
	}

	inorderTraversal(root.right, result)
}

func (bt *BinaryTree) BoundaryTraversal() {
	result := []int{}
	leftTraversal(bt.Root, &result)
	inorderTraversal(bt.Root, &result)
	rightTraversal(bt.Root, &result)

	for _, v := range result {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
