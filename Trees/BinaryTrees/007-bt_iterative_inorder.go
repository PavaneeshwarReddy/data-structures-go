package binarytrees

import "fmt"

/*
Iterative Inorder Traversal
- First we will traverse entire left and then when you reach end you will pop
- If still it's nil you pop again and traverse right to check whether its there are not
- We keep on repeating it.

- Always remember first we need to traverse to left and then right by popping the current node
*/

func (bt *BinaryTree) InorderIterative() {
	currNode := bt.Root
	nodes := []*Node{}

	for true {
		if currNode != nil {
			nodes = append(nodes, currNode)
			currNode = currNode.left
		} else {
			if len(nodes) == 0 {
				break
			}
			topNode := nodes[len(nodes)-1]
			nodes = nodes[0 : len(nodes)-1]
			fmt.Print(topNode.value, " ")
			currNode = topNode.right
		}
	}
	fmt.Println()
}
