package binarytrees

import "fmt"

/*
All traversal
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
