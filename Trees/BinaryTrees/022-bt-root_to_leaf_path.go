package binarytrees

import "fmt"

/*
Root to leaf path
*/

func rootToLeafRecursive(root *Node, key int, path *[]*Node) bool {
	if root == nil {
		return false
	}
	*path = append(*path, root)
	if root.left == nil && root.right == nil && root.value == key {
		return true
	}
	leftCheck := rootToLeafRecursive(root.left, key, path)
	rightCheck := rootToLeafRecursive(root.right, key, path)

	isFound := leftCheck || rightCheck

	if !(isFound) {
		*path = (*path)[0 : len(*path)-1]
	}
	return isFound
}

func (bt *BinaryTree) RootToLeafPath(key int) {
	path := []*Node{}
	rootToLeafRecursive(bt.Root, key, &path)
	for _, v := range path {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
