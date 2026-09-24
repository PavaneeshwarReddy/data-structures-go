package binarytrees

import "fmt"

/*
All traversal
- To get the pre order first we need to always traverse first to the left
- To get the inorder after popping we need to move to the right
- For pre order its after all

Preorder : visit count is 1 , if it visited first then we just print the value
Inorder: visit count is 2, then that means it went through the node once while traversing left and comming back
Postorder: visit count is 3, that means the same went on for 2 times we just pop the value out and add it to postorder
*/

type StackNode struct {
	node  *Node
	count int
}

func (bt *BinaryTree) AllTraversal() {
	stack := []StackNode{StackNode{node: bt.Root, count: 1}}
	preorder := []int{}
	inorder := []int{}
	postorder := []int{}

	for len(stack) > 0 {
		topNode := stack[len(stack)-1]
		if topNode.count == 1 {
			preorder = append(preorder, topNode.node.value)
			stack[len(stack)-1].count += 1
			if topNode.node.left != nil {
				stack = append(stack, StackNode{node: topNode.node.left, count: 1})
			}
			topNode.count += 1
		} else if topNode.count == 2 {
			inorder = append(inorder, topNode.node.value)
			stack[len(stack)-1].count += 1
			if topNode.node.right != nil {
				stack = append(stack, StackNode{node: topNode.node.right, count: 1})
			}
			topNode.count += 1
		} else if topNode.count == 3 {
			postorder = append(postorder, topNode.node.value)
			stack = stack[0 : len(stack)-1]
		}
	}

	for _, v := range preorder {
		fmt.Print(v, " ")
	}

	fmt.Println()

	for _, v := range inorder {
		fmt.Print(v, " ")
	}

	fmt.Println()

	for _, v := range postorder {
		fmt.Print(v, " ")
	}

}
