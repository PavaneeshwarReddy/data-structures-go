package binarytrees

/*
Max Depth
- Whenever you travel towards your left you will increment value by 1 or by right the same
- Return the max depth ever reached and just return max value
*/

func (bt *BinaryTree) MaxDepth(root *Node, depth int) int {
	if root == nil {
		return depth
	}
	return max(bt.MaxDepth(root.left, depth+1), bt.MaxDepth(root.right, depth+1))
}
