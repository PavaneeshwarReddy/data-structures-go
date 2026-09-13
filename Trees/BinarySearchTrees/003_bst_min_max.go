package binarysearchtrees

/*
For Max we can only go to the right
For Min we can only go to the left
*/

func FindMin(root *Node) int {
	if root.Left == nil {
		return root.Val
	}
	return FindMin(root.Left)
}

func FindMax(root *Node) int {
	if root.Right == nil {
		return root.Val
	}
	return FindMax(root.Right)
}
