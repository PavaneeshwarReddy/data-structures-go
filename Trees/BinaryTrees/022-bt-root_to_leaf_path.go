package binarytrees

import "fmt"

/*
Root to leaf path
- We can do any traversal but pre order is better
- Whenever we see a node we add it to the path
- if that node doesn't make any sense, I mean we are unable to location from the left or right we again pop the addenode
- If we found then we can return
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
