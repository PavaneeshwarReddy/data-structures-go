package binarysearchtrees

/*
Binary Search Trees
- We can take many approaches
- Mainly we can take MaxHeap but we don't need that complex here
*/

func inorder(root *Node, k int, count *int) int {
	if root == nil {
		return -1
	}

	left := inorder(root.Left, k, count)
	if left != -1 {
		return left
	}

	(*count) += 1
	if (*count) == k {
		return root.Val
	}

	return inorder(root.Right, k, count)
}

func kthSmallest(root *Node, k int) int {
	count := 0
	return inorder(root, k, &count)
}
