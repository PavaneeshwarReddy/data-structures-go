package binarytrees
import "fmt"

/*
Postorder Traversal Using 1 Stack
*/


func (bt *BinaryTree) PostorderTraversalOneStack() {
	stack := []*Node{}
	currNode := bt.Root

	for len(stack) > 0 || currNode != nil {
		if currNode != nil {
			stack = append(stack, currNode)
			currNode = currNode.left
		} else {
			tempNode := stack[len(stack)-1].right
			if tempNode == nil {
				tempNode = stack[len(stack)-1]
				fmt.Print(tempNode.value, " ")
				stack = stack[0:len(stack)-1]
			
				for len(stack) > 0 && tempNode == stack[len(stack)-1].right {
					tempNode = stack[len(stack)-1]
					fmt.Print(tempNode.value, " ")
					stack = stack[0:len(stack)-1]
				}
			} else {
				currNode = tempNode
			}
		}
	}
}




