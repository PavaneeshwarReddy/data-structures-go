package binarytrees

/*
Morris Preorder Traversal
	- If there is no Left, that means we can simply append to our result and move right
	- If Left is present, then we can simply find the right most node and point that right to curr node
		- if right pointer of the last node is nil, then point this to curr node and move towards left and include in result
		- if right pointer of the last node is pointing to curr node, then break move towards right
*/

func (bt *BinaryTree) MorrisPreorder() []int {
	result := []int{}
	currNode := bt.Root

	for currNode != nil {
		if currNode.left == nil {
			result = append(result, currNode.value)
			currNode = currNode.right
		} else {
			prevNode := currNode.left

			for prevNode.right != nil && prevNode.right != currNode {
				prevNode = prevNode.right
			}

			if prevNode.right == nil {
				prevNode.right = currNode
				result = append(result, currNode.value)
				currNode = currNode.left
			} else {
				prevNode.right = nil
				currNode = currNode.right
			}
		}
	}

	return result
}
