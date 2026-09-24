package binarytrees

import "fmt"

/*
Boundary Traversal
- You move towards left and store results but at the point we reach the end, if we go to left then when left is nil we need to move right
- while the same applies when you are moving from right also
- For leaf nodes we can perform any operation either inorder, post or pre order
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
