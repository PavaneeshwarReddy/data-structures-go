package binarytrees

/*
Problem: Count total no of nodes
link:
Idea:
- Simply perform any traversal and just add 1 to every recusive call
*/

func countNodes(root *Node) int {
	if root == nil {
		return 0
	}
	return countNodes(root.left) + countNodes(root.right) + 1
}

func (bt *BinaryTree) CountNodes() int {
	return countNodes(bt.Root)
}
