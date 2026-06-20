package binarytrees
import "fmt"

/*
Iterative Inorder Traversal
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
			nodes = nodes[0:len(nodes)-1]
			fmt.Print(topNode.value, " ")
			currNode = topNode.right
		}
	}
	fmt.Println()
}

