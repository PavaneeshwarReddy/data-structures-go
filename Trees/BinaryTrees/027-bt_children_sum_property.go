package binarytrees

/*
Problem: Children Sum Property
Link:
Idea:
	- Ensure that parent are never smaller than the children
	- make the sum, if sum is greater than make the parent node value as sum
	- if sum is smaller, we can pass current parent node value to both
	- We can make a post order traversal
	- what value we can finally place at the current root after traversing entire bottom tree of that node
*/

func sumProperty(root *Node) {
	if root == nil {
		return
	}

	child := 0

	if root.left != nil {
		child += root.left.value
	}
	if root.right != nil {
		child += root.right.value
	}

	if child >= root.value {
		root.value = child
	} else {
		if root.left != nil {
			root.left.value = root.value
		}
		if root.right != nil {
			root.right.value = root.value
		}
	}

	sumProperty(root.left)
	sumProperty(root.right)

	total := 0

	if root.left != nil {
		total += root.left.value
	}
	if root.right != nil {
		total += root.right.value
	}

	if root.left != nil || root.right != nil {
		root.value = total
	}

}

func (bt *BinaryTree) ChildrenSumProperty() {
	sumProperty(bt.Root)
}
