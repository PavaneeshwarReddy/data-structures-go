package binarytrees

/*
Morris Inorder Traversal
	- If there is no Left, that means we can simply append to our result and move right
	- If Left is present, then we can simply find the right most node and point that right to curr node
		- if right pointer of the last node is nil, then point this to curr node and move towards left
		- if right pointer of the last node is pointing to curr node, then break move towards right and include this in result
*/

func (bt *BinaryTree) MorrisInorder() []int {
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
				currNode = currNode.left
			} else {
				prevNode.right = nil
				result = append(result, currNode.value)
				currNode = currNode.right
			}
		}
	}

	return result
}
