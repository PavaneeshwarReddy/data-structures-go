package binarytrees

/*
Problem: Flatten Binary Tree to Linked List
Link: https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
Idea:
	- Consider this example
			2
		/		\
	  3           4
	- First we will traverse to the right, by keeping track of prev node
	- Once we go to 4, we traverse left, right nothing is there and we will return to the left by storing prevNode as 4
	- After doing with left, when we are accessing 3, we set right of 3 to prevNode which is 4
	- Then we will return 2, here prevNode is 3, so we set it to the right of 2
	- Everywhere we make left pointer as nil
*/

func flatten(root *Node, prevNode **Node) {
	if root == nil {
		return
	}
	flatten(root.right, prevNode)
	flatten(root.left, prevNode)
	root.right = *prevNode
	root.left = nil
	*prevNode = root
}

func (bt *BinaryTree) FlattenTree() {
	var prevNode *Node
	flatten(bt.Root, &prevNode)
}
