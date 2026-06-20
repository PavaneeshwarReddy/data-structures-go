package binarytrees

import "fmt"

/*
Postorder traversal using 2 stacks iterative traversal
*/

func (bt *BinaryTree) PostordrerIterativeTwoStacks() {
	firstStack := []*Node{bt.Root}
	secondStack := []*Node{}

	for len(firstStack) > 0 {
		topNode := firstStack[len(firstStack)-1]
		firstStack = firstStack[0 : len(firstStack)-1]
		secondStack = append(secondStack, topNode)
		if topNode.left != nil {
			firstStack = append(firstStack, topNode.left)
		}
		if topNode.right != nil {
			firstStack = append(firstStack, topNode.right)
		}
	}
	for len(secondStack) > 0 {
		fmt.Print(secondStack[len(secondStack)-1].value, " ")
		secondStack = secondStack[0 : len(secondStack)-1]
	}
	fmt.Println()
}
